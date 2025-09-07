import {useEffect, useRef, useState, useCallback } from 'react';

type Message = (data: unknown) => void;
type OpenHandler = () => void;
type CloseHandler = (ev?: CloseEvent) => void;
type ErrorHandler = (ev?: Event) => void;

interface Options {
    onOpen?: OpenHandler;
    onClose?: CloseHandler;
    onMessage?: Message;
    onError?: ErrorHandler;
    reconnect?: boolean;
    reconnectInterval?: number; //ms
}

export function useWebSocket (url: string, options: Options = {}) {
    const { onOpen, onClose, onMessage, onError, reconnect = true } = options;
    const reconnectInterval = options.reconnectInterval ?? 2000;

    const wsRef = useRef<WebSocket | null>(null)
    const retryRef = useRef(0);
    const timeoutRef = useRef<number | null>(null);
    const [connected, setConnected] = useState(false);

    const connect = useCallback(() => {
        if (wsRef.current && (wsRef.current.readyState === WebSocket.OPEN || wsRef.current.readyState == WebSocket.CONNECTING)) {
            return;
        }

        try {
            console.log('Connecting to WebSocket:', url)
            const socket = new WebSocket(url);
            wsRef.current = socket;

            socket.onopen = () => {
                retryRef.current = 0;
                setConnected(true);
                // onOpen && onOpen();
                onOpen?.();
                console.log('Connected to WebSocket is successful:', url)
            };

            socket.onmessage = (e) => {
                if (!e.data) return;
                try {
                    const parsed = JSON.parse(e.data);
                    // onMessage && onMessage(parsed);
                    onMessage?.(parsed);
                } catch (err) {
                    // onMessage && onMessage(e.data);
                    console.error('Error parsing WebSocket message:', err); // TODO: handle error err
                    onMessage?.(e.data);
                }
            };

            socket.onerror = (ev) => {
                console.warn('Error in WebSocket:', ev);
                // onError && onError(ev);
                onError?.(ev);
            };

            socket.onclose = (ev) => {
                setConnected(false);
                // onClose && onClose(ev);
                onClose?.(ev);
                wsRef.current = null;
                if (reconnect) {
                    retryRef.current++;
                    const delay = Math.min(30000, reconnectInterval * Math.pow(2, retryRef.current - 1))
                    console.log('Reconnecting to WebSocket in', delay, 'ms...');
                    timeoutRef.current = window.setTimeout(connect, delay);
                }
            };
        } catch (err) {
            console.error('Error connecting to WebSocket:', err);
            // onError && onError(err as any);
            // TODO: handle error
            onError?.(err as never)
        }
    }, [url, onOpen, onClose, onMessage, onError, reconnect, reconnectInterval]);

    const disconnect = useCallback(() => {
        if (timeoutRef.current) {
            clearTimeout(timeoutRef.current);
            timeoutRef.current = null;
        }
        const ws = wsRef.current;
        if (!ws) return;
        try {
            ws.close(1000, 'client closing');
        } catch (e) {
            console.error(e)
        } finally {
            wsRef.current = null;
            setConnected(false);
        }
    }, []);

    //TODO: fix any error
    const sendMessage = useCallback((data: any) => {
        const ws = wsRef.current;
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            console.warn('Cannot send message - WebSocket is not connected');
            return false;
        }
        try {
            const payload = typeof data === 'string' ? data : JSON.stringify(data);
            ws.send(payload);
            return true;
        } catch (err) {
            console.error('Error sending message:', err);
            return false;
        }
    }, []);

    useEffect(() => {
        connect();
        return () => {
            disconnect();
        };
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [url]);


    return { connected, sendMessage, disconnect };
}