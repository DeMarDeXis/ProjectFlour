import React, { useState } from "react";
import {useNavigate} from "react-router-dom";
import './registr_style.css';
import logo from '../../assets/masterflour.png'

interface RegistrationPageProps {
    name: string;
    username: string;
    password: string;
    error: string;
    OnSubmit: (e: React.FormEvent) => void;
    OnNameChange: (nameIn: string) => void;
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

export const RegistrationForm: React.FC<RegistrationPageProps> = ({
    name,
    username,
    password,
    error,
    OnSubmit,
    OnNameChange,
    OnUsernameChange,
    OnPasswordChange,
    IsLoading
}) => (
    <div className="registration-form">
        <h1>Регистрация в систему</h1>
        {error && <div className="error-message">{error}</div>}
        <form onSubmit={OnSubmit}>
            <div className="form-group">
                <label htmlFor="name">Имя</label>
                <input
                    id="name"
                    type="text"
                    value={name}
                    onChange={(e) => OnNameChange(e.target.value)}
                    disabled={IsLoading}
                />
            </div>
            <div className="form-group">
                <label htmlFor="username">Системный псевдоним</label>
                <input
                    id="username"
                    type="text"
                    value={username}
                    onChange={(e) => OnUsernameChange(e.target.value)}
                    disabled={IsLoading}
                />
            </div>
            <div className="form-group">
                <label htmlFor="password">Пароль</label>
                <input
                    id="password"
                    type="password"
                    value={password}
                    onChange={(e) => OnPasswordChange(e.target.value)}
                    disabled={IsLoading}
                />
            </div>
            <button type="submit" className="registration-button" disabled={IsLoading}>
                {IsLoading ? 'Регистрация...' : 'Зарегистрироваться'}
            </button>
        </form>
    </div>
)

const RegistrationPage: React.FC = () => {
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [showErrorToast, setShowErrorToast] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();

    const showError = (message:string) => {
        setError(message);
        setShowErrorToast(true);
        setTimeout(() => {
            setShowErrorToast(false);
        }, 5000);
    }

    const handlerRegistration = async (e: React.FormEvent)=> {
        e.preventDefault();
        setError('')
        setShowErrorToast(false);

        if (!name.trim() || !username.trim() || !password.trim()) {
            setError('Пожалуйста, заполните все поля.');
            return;
        }

        if (password.length < 8) {
            setError('Пароль должен содержать не менее 8 символов.');
            return;
        }

        setIsLoading(true);

        try {
            const requestData = {
                name: name.trim(),
                username: username.trim(),
                password: password
            };

            console.log(requestData);

            const response = await fetch('/api/auth/sign-up', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json'
                },
                body: JSON.stringify(requestData),
            });

            if (!response.ok) {
                let errorMessage = 'Произошла ошибка при регистрации.';

                try {
                    const errorData = await response.json();
                    errorMessage = errorData.message || errorMessage;
                } catch {
                    errorMessage = `Ошибка сервера: ${response.status} ${response.statusText}`;
                }

                throw new Error(errorMessage);
            }

            navigate('/')
        } catch (err) {
            console.error('Ошибка регистрации:', err);

            let errorMessage = "Произошла ошибка при регистрации.";

            if (err instanceof Error) {
                if (err.message === 'Failed to fetch') {
                    errorMessage = "Не удается подключиться к серверу. Проверьте подключение к интернету и убедитесь, что сервер запущен.";
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
        <div className="registration-container">
            {isLoading && (
                <div className="overlay">
                    <div className="spinner"></div>
                </div>
            )}
            <div className="registration-image">
                <img src={logo} alt="There is must be a logo" className="logo" />
            </div>

            {showErrorToast && (
                <ErrorToast
                    message={error}
                    onClose={() => setShowErrorToast(false)}
                />
            )}
            <RegistrationForm
                name={name}
                username={username}
                password={password}
                error={error}
                OnSubmit={handlerRegistration}
                OnNameChange={setName}
                OnUsernameChange={setUsername}
                OnPasswordChange={setPassword}
                IsLoading={isLoading}
                />
        </div>
    );
};

export default RegistrationPage;