package webSocketHandler

import (
	"ProjectFlour/internal/events"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

const (
	pongWait       = 60 * time.Second
	writeWait      = 10 * time.Second
	pingPeriod     = 30 * time.Second
	maxMessageSize = 5120
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	logger *slog.Logger
	ctx    context.Context
	cancel context.CancelFunc
}

type WebSocketHandler struct {
	log        *slog.Logger
	clients    map[*Client]struct{}
	clientsMUX sync.RWMutex
	eventBus   *events.EventBus
}

func New(log *slog.Logger, eventBus *events.EventBus) *WebSocketHandler {
	return &WebSocketHandler{
		log:      log,
		clients:  make(map[*Client]struct{}),
		eventBus: eventBus,
	}
}

func (wh *WebSocketHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.webSocketHandlers.HandleConnections"

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		wh.log.Error("Error during connection upgrade", "op", op, "error", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:   conn,
		send:   make(chan []byte, 256),
		logger: wh.log,
		ctx:    ctx,
		cancel: cancel,
	}

	wh.registerClient(client)
	wh.log.Info("New client connected", slog.String("client", conn.RemoteAddr().String()))

	subscriptionID := wh.subscribeToEvents(client)
	defer wh.eventBus.Unsubscribe(events.EventFileImported, subscriptionID)

	go client.writePump()
	client.readPump(wh)
}

func (wh *WebSocketHandler) registerClient(c *Client) {
	const op = "internal.handlers.webSocketHandlers.registerClient"
	wh.clientsMUX.Lock()
	defer wh.clientsMUX.Unlock()
	wh.clients[c] = struct{}{}
	wh.log.Debug("New client connected", "op", op)
}

func (wh *WebSocketHandler) unregisterClient(c *Client) {
	wh.clientsMUX.Lock()
	defer wh.clientsMUX.Unlock()
	delete(wh.clients, c)
}

func (c *Client) writePump() {
	//const op = "internal.handlers.webSocketHandlers.writePump"

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
		c.cancel()
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) readPump(wsh *WebSocketHandler) {
	const op = "internal.handlers.webSocketHandlers.readPump"

	defer func() {
		wsh.log.Info("Client disconnected", slog.String("remote", c.conn.RemoteAddr().String()))
		c.cancel()
		_ = c.conn.Close()
		wsh.unregisterClient(c)
	}()

	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		wsh.log.Debug("Received pong", slog.String("client", c.conn.RemoteAddr().String()))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				wsh.log.Error("Client disconnected unexpectedly", "op", op, "error", err)
			} else {
				wsh.log.Error("Error reading message from client", "op", op, "error", err)
			}
			break
		}

		wsh.log.Debug("Message from client", "client", c.conn.RemoteAddr().String(), "msg", string(msg))
	}
}

func (c *Client) close() error {
	select {
	case <-c.ctx.Done():
	default:
		c.cancel()
	}
	_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	return c.conn.Close()
}

func (wh *WebSocketHandler) subscribeToEvents(client *Client) string {
	return wh.eventBus.Subscribe(events.EventFileImported, func(e events.Event) {
		event, ok := e.(events.FileImportedEvent)
		if !ok {
			return
		}
		data, err := json.Marshal(event)
		if err != nil {
			wh.log.Error("Failed to marshal event", "error", err)
			return
		}
		select {
		case client.send <- data:
		default:
			wh.log.Warn("Client send buffer full, closing", "client", client.conn.RemoteAddr())
			_ = client.close()
			wh.registerClient(client)
		}
	})
}

// Broadcast unused for now
//func (wh *WebSocketHandler) Broadcast(message string) {
//	wh.clientsLock.Lock()
//	defer wh.clientsLock.Unlock()
//
//	for conn := range wh.clients {
//		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
//			wh.log.Error("Broadcast failed", "client", conn.RemoteAddr(), "error", err)
//			conn.Close()
//			delete(wh.clients, conn)
//		}
//	}
//}

func (wh *WebSocketHandler) DisconnectAll() {
	wh.clientsMUX.RLock()
	clients := make([]*Client, 0, len(wh.clients))
	for c := range wh.clients {
		clients = append(clients, c)
	}
	wh.clientsMUX.RUnlock()

	var wg sync.WaitGroup

	for _, c := range clients {
		wg.Add(1)
		go func(cli *Client) {
			defer wg.Done()
			_ = cli.close()
			wh.unregisterClient(cli)
		}(c)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		wh.log.Warn("Timeout while waiting for clients to disconnect")
	}
}
