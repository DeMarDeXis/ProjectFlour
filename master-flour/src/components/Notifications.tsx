import { useState } from 'react';
import { useWebSocket } from '../hooks/useWebSocket';
import './notification.css';

type Notification = {
    id: string;
    type: 'success' | 'error' | 'info' | 'warning';
    title: string;
    message: string;
    timestamp: Date;
};

export const Notifications = () => {
    const [notifications, setNotifications] = useState<Notification[]>([]);
    const getWebSocketUrl = () => {
        const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
        return `${protocol}//${window.location.host}/ws`;
    };

    const wsUrl = getWebSocketUrl();

    console.log('URL passed to useWebSocket:', wsUrl);

    useWebSocket(wsUrl, {
        onOpen: () => {
            console.log('WebSocket connected');
            addNotification({
                type: 'info',
                title: 'WebSocket Подключение',
                message: 'Успешно подключились к серверу',
            });
        },
        onClose: (event?: CloseEvent) => {
            addNotification({
                type: 'warning',
                title: 'WebSocket Ошибка',
                message: event ? `Соединение закрыто (код: ${event.code})` : 'Соединение закрыто',
            });
        },
        onMessage: (message: unknown) => {
            console.log('Received message:', message);

            if (message && typeof message === 'object' && 'message' in message) {
                const msgText =
                    (message as { message?: string; data?: { message: string } }).message ||
                    (message as { data?: { message: string } }).data?.message ||
                    JSON.stringify(message);

                addNotification({
                    type: 'success',
                    title: 'Импорт данных',
                    message: msgText,
                });
            } else {
                addNotification({
                    type: 'error',
                    title: 'Ошибка получения данных',
                    message: 'Неизвестный формат данных от WebSocket',
                });
            }
        },
        onError: (error) => {
            console.log('WebSocket error' + error);
            addNotification({
                type: 'error',
                title: 'WebSocket Ошибка',
                message: 'Ошибка подключения к серверу',
            });
        },
    });

    const addNotification = (notification: Omit<Notification, 'id' | 'timestamp'>) => {
        const newNotification: Notification = {
            id: Math.random().toString(36).substring(2, 9),
            ...notification,
            timestamp: new Date(),
        };

        setNotifications((prev) => [...prev, newNotification]);

        setTimeout(() => {
            setNotifications((prev) => prev.filter((n) => n.id !== newNotification.id));
        }, 5000);
    };

    const removeNotification = (id: string) => {
        setNotifications((prev) => prev.filter((n) => n.id !== id));
    };

    const getNotificationStyle = (type: Notification['type']) => {
        const baseStyle = 'p-4 rounded-lg shadow-md mb-2 flex justify-between items-start';
        switch (type) {
            case 'success':
                return `${baseStyle} bg-green-50 border-l-4 border-green-500`;
            case 'error':
                return `${baseStyle} bg-red-50 border-l-4 border-red-500`;
            case 'warning':
                return `${baseStyle} bg-yellow-50 border-l-4 border-yellow-500`;
            case 'info':
                return `${baseStyle} bg-blue-50 border-l-4 border-blue-500`;
            default:
                return `${baseStyle} bg-gray-50 border-l-4 border-gray-500`;
        }
    };

    const getNotificationIcon = (type: Notification['type']) => {
        switch (type) {
            case 'success': return '✅';
            case 'error': return '❌';
            case 'warning': return '⚠️';
            case 'info': return 'ℹ️';
            default: return '💡';
        }
    };

    return (
        <div className="fixed bottom-4 right-4 w-80 z-50 flex flex-col gap-3">
            {notifications.map((notification) => (
                <div
                    key={notification.id}
                    className={`${getNotificationStyle(notification.type)} notification`}
                >
                    <div className="mr-2 text-xl">{getNotificationIcon(notification.type)}</div>
                    <div className="flex-1">
                        <h4 className="font-medium text-sm">{notification.title}</h4>
                        <p className="text-sm">{notification.message}</p>
                        <p className="text-xs text-gray-500 mt-1">
                            {notification.timestamp.toLocaleTimeString()}
                        </p>
                    </div>
                    <button
                        onClick={() => removeNotification(notification.id)}
                        className="ml-2 text-gray-400 hover:text-gray-600"
                        aria-label="Close notification"
                    >
                        &times;
                    </button>
                </div>
            ))}
        </div>
    );
};