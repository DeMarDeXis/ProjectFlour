import React, { useState } from "react";
import {useNavigate} from "react-router-dom";
import './styles.css';
import logo from '../../assets/masterflour.png'

interface LoginPageProps {
    username: string;
    password: string;
    error: string;
    OnSubmit: (e: React.FormEvent) => void;
    OnUsernameChange: (usernameIn: string) => void;
    OnPasswordChange: (passwordIn: string) => void;
    IsLoading: boolean;
}

interface ErrorToastProps {
    message: string;
    onClose: () => void;
}

const ErrorToast: React.FC<ErrorToastProps> = ({ message, onClose }) => (
    <div className={"error-toast"}>
        <div className={"error-toast-content"}>
            <span className={"error-icon"}>⚠️</span>
            <span className={"error-text"}>{message}</span>
            <button className={"error-close"} onClick={onClose}>x</button>
        </div>
    </div>
)

export const SignInForm: React.FC<LoginPageProps> = ({
    username,
    password,
    error,
    OnSubmit,
    OnUsernameChange,
    OnPasswordChange,
    IsLoading
}) => (
    <div className="login-form">
        <h1>Вход в систему</h1>
        {error && <div className="error-message">{error}</div>}
        <form onSubmit={OnSubmit}>
            <div className="form-group">
                <label htmlFor="username">Системное Имя</label>
                <input
                    id="login"
                    type="text"
                    value={username}
                    onChange={(e) => OnUsernameChange(e.target.value)}
                />
            </div>
            <div className="form-group">
                <label htmlFor="password">Пароль</label>
                <input
                    id="password"
                    type="password"
                    value={password}
                    onChange={(e) => OnPasswordChange(e.target.value)}
                />
            </div>
            <button type="submit" className="login-button" disabled={IsLoading}>
                {IsLoading ? 'Выполняется вход в ситему' : 'Войти'}
            </button>
        </form>
    </div>
)

const LoginPage: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [showErrorToast, setShowErrorToast] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();

    const ERROR_MESSAGES = {
        INVALID_CREDENTIALS: 'Invalid credentials',
        USER_NOT_FOUND: 'User not found',
        SERVER_ERROR: 'Server error',
    };

    const showError = (message:string) => {
        setError(message);
        setShowErrorToast(true);
        setTimeout(() => {
            setShowErrorToast(false);
        }, 5000)
    };

    const handleSignIn = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setShowErrorToast(false);
        console.log('Username:', username);
        console.log('Password:', password);

        if (!username.trim() || !password.trim()) {
            setError('Пожалуйста, заполните все поля.');
            return;
        }

        setIsLoading(true);

        try {
            const requestData = {
                username: username.trim(),
                password: password
            };

            console.log(requestData);

            const response = await fetch('/api/auth/sign-in', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                },
                body: JSON.stringify(requestData),
            });

            if (!response.ok) {
                let errorMessage = 'Ошибка при входе в систему';

                try {
                    const errorData = await response.json();
                    errorMessage = errorData.message || errorMessage;
                } catch {
                    errorMessage = `Ошибка сервера: ${response.status} ${response.statusText}`;
                }

                throw new Error(errorMessage);
            }

            const data = await response.json();
            if (data.token) {
                localStorage.setItem('token', data.token);
            } else {
                throw new Error(ERROR_MESSAGES.INVALID_CREDENTIALS);
            }

            navigate('/home')
        } catch (err) {
            console.error('Error in sign in:', err);

            let errorMessage = "Ошибка при входе в систему";

            if (err instanceof Error) {
                if (err.message === 'Ошибка перхвата токена') {
                    errorMessage = "Ошибка при входе в систему";
                } else {
                    errorMessage = err.message;
                }
            }

            showError(errorMessage);
        } finally {
            setIsLoading(false);
        }
    }

    return (
        <div className="login-container">
            {isLoading && (
                <div className={"overlay"}>
                    <div className="spinner"></div>
                </div>
            )}
            {showErrorToast && (
                <ErrorToast
                    message={error}
                    onClose={() => setShowErrorToast(false)}
                />
            )}
            <SignInForm
                username={username}
                password={password}
                error={error}
                OnSubmit={handleSignIn}
                OnUsernameChange={setUsername}
                OnPasswordChange={setPassword}
                IsLoading={isLoading}
            />
            <div className="login-image">
                <img src={logo} alt="There is must be a logo" className="logo" />
            </div>
        </div>
    );
};

export default LoginPage;