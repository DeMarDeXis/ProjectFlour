import { useState, useEffect } from 'react';
import { useWebSocket } from '../../hooks/useWebSocket.ts';
import './notification.css';

type NotificationType = 'success' | 'error' | 'info' | 'warning';

type Notification = {
    id: string;
    type: NotificationType;
    title: string;
    message: string;
    timestamp: Date;
    isRead: boolean;
};

type NotificationsProps = {
    onClose?: () => void;
};

export const Notifications = ({ onClose }: NotificationsProps) => {
    const [notifications, setNotifications] = useState<Notification[]>([]);
    const [unreadCount, setUnreadCount] = useState(0);

    const getWebSocketUrl = () => {
        const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
        return `${protocol}//${window.location.host}/ws`;
    };

    const wsUrl = getWebSocketUrl();
    // const { connected, sendMessage, disconnect } = useWebSocket(wsUrl, {
    const { connected } = useWebSocket(wsUrl, {
        onOpen: () => {
            console.log('✅ WebSocket connected');
            addNotification({
                type: 'info',
                title: 'WebSocket Подключение',
                message: 'Успешно подключились к серверу',
            });
        },
        onClose: (event?: CloseEvent) => {
            console.log('❌ WebSocket closed', event);
            addNotification({
                type: 'warning',
                title: 'WebSocket Отключение',
                message: event
                    ? `Соединение закрыто (код: ${event.code}, причина: ${event.reason || 'не указана'})`
                    : 'Соединение закрыто',
            });
        },
        onMessage: (message: unknown) => {
            console.log('📨 Received message:', message);
            console.log('📨 Message type:', typeof message);

            const msgText = parseMessage(message);

            addNotification({
                type: 'success',
                title: 'Импорт данных',
                message: msgText,
            });
        },
        onError: (error) => {
            console.error('🔴 WebSocket error:', error);
            addNotification({
                type: 'error',
                title: 'WebSocket Ошибка',
                message: 'Ошибка подключения к серверу',
            });
        },
        reconnect: true,
        reconnectInterval: 3000,
    });

    const parseMessage = (msg: unknown): string => {
        try {
            if (typeof msg === 'string') {
                return msg;
            }

            if (msg && typeof msg === 'object') {
                const obj = msg as Record<string, unknown>;

                if ('message' in obj && typeof obj.message === 'string') {
                    return obj.message;
                }

                if ('data' in obj && obj.data && typeof obj.data === 'object') {
                    const data = obj.data as Record<string, unknown>;
                    if ('message' in data && typeof data.message === 'string') {
                        return data.message;
                    }
                }

                return JSON.stringify(obj);
            }

            return String(msg);
        } catch (error) {
            console.error('Error parsing message:', error);
            return 'Неизвестное сообщение';
        }
    };

    const addNotification = (notification: Omit<Notification, 'id' | 'timestamp' | 'isRead'>) => {
        const newNotification: Notification = {
            id: Math.random().toString(36).substring(2, 9),
            ...notification,
            timestamp: new Date(),
            isRead: false,
        };

        setNotifications((prev) => [newNotification, ...prev]);
        setUnreadCount(prev => prev + 1);

        setTimeout(() => {
            markAsRead(newNotification.id);
        }, 5000);
    };

    const markAsRead = (id: string) => {
        setNotifications(prev =>
            prev.map(n => n.id === id ? { ...n, isRead: true } : n)
        );
        setUnreadCount(prev => Math.max(0, prev - 1));
    };

    const markAllAsRead = () => {
        setNotifications(prev =>
            prev.map(n => ({ ...n, isRead: true }))
        );
        setUnreadCount(0);
    };

    const removeNotification = (id: string) => {
        const wasUnread = notifications.find(n => n.id === id)?.isRead === false;
        setNotifications(prev => prev.filter(n => n.id !== id));
        if (wasUnread) {
            setUnreadCount(prev => Math.max(0, prev - 1));
        }
    };

    const clearAllNotifications = () => {
        setNotifications([]);
        setUnreadCount(0);
    };

    const formatTime = (date: Date): string => {
        const now = new Date();
        const diffMinutes = Math.floor((now.getTime() - date.getTime()) / 60000);

        if (diffMinutes < 1) return 'только что';
        if (diffMinutes < 60) return `${diffMinutes} мин назад`;
        if (diffMinutes < 1440) return `${Math.floor(diffMinutes / 60)} ч назад`;

        return date.toLocaleDateString('ru-RU', {
            day: '2-digit',
            month: '2-digit',
            hour: '2-digit',
            minute: '2-digit'
        });
    };

    // TODO: delete
    useEffect(() => {
        if (notifications.length === 0) {
            setTimeout(() => {
                addNotification({
                    type: 'success',
                    title: 'Успешная операция[TEST]',
                    message: 'Данные успешно импортированы из Excel файла',
                });
                addNotification({
                    type: 'error',
                    title: 'Ошибка сервера[TEST]',
                    message: 'Не удалось сохранить изменения. Попробуйте еще раз.',
                });
                addNotification({
                    type: 'warning',
                    title: 'Внимание[TEST]',
                    message: 'Скоро закончится место на диске. Очистите кэш.',
                });
                addNotification({
                    type: 'info',
                    title: 'Информация[TEST]',
                    message: 'Доступна новая версия приложения. Обновите для получения новых функций.',
                });
            }, 500);
        }
    }, []);

    return (
        <div className="notifications-container">

            <div className="notifications-header">
                <div className="notifications-title">
                    <span className="notifications-title-icon">🔔</span>
                    <span>Уведомления</span>
                    {unreadCount > 0 && (
                        <span className="notifications-unread-count">{unreadCount}</span>
                    )}
                </div>

                <div className="notifications-header-actions">
                    {/* Note STATUS */}
                    <div
                        className={`connection-indicator ${connected ? 'connected' : 'disconnected'}`}
                        title={connected ? 'Подключено к серверу' : 'Отключено от сервера'}
                    >
                        <span className="connection-dot"></span>
                        <span className="connection-text">
                            {connected ? 'Онлайн' : 'Офлайн'}
                        </span>
                    </div>

                    <button
                        className="notifications-mark-all-read"
                        onClick={markAllAsRead}
                        title="Отметить все как прочитанные"
                    >
                        ✅
                    </button>
                    <button
                        className="notifications-clear-all"
                        onClick={clearAllNotifications}
                        title="Очистить все"
                    >
                        🗑️
                    </button>
                    {onClose && (
                        <button
                            className="notifications-close-btn"
                            onClick={onClose}
                            title="Закрыть"
                        >
                            ✕
                        </button>
                    )}
                </div>
            </div>

            <div className="notifications-list">
                {notifications.length === 0 ? (
                    <div className="notifications-empty">
                        <div className="notifications-empty-icon">📭</div>
                        <p>Нет уведомлений</p>
                    </div>
                ) : (
                    notifications.map((notification) => (
                        <div
                            key={notification.id}
                            className={`notification-item ${notification.isRead ? '' : 'unread'} ${notification.type}`}
                            onClick={() => !notification.isRead && markAsRead(notification.id)}
                        >
                            <div className="notification-header">
                                <span className={`notification-type ${notification.type}`}>
                                    {getNotificationTypeEmoji(notification.type)}
                                </span>
                                <span className="notification-time">
                                    {formatTime(notification.timestamp)}
                                </span>
                            </div>
                            <div className="notification-title">{notification.title}</div>
                            <div className="notification-message">{notification.message}</div>
                            <button
                                className="notification-delete-btn"
                                onClick={(e) => {
                                    e.stopPropagation();
                                    removeNotification(notification.id);
                                }}
                                title="Удалить"
                            >
                                ✕
                            </button>
                        </div>
                    ))
                )}
            </div>

            {notifications.length > 0 && (
                <div className="notifications-footer">
                    <button onClick={clearAllNotifications}>
                        Очистить все уведомления
                    </button>
                </div>
            )}
        </div>
    );
};

const getNotificationTypeEmoji = (type: NotificationType): string => {
    switch (type) {
        case 'success': return '✅';
        case 'error': return '❌';
        case 'warning': return '⚠️';
        case 'info': return 'ℹ️';
        default: return 'ℹ️';
    }
};