import React from "react";
import {useNavigate} from "react-router-dom";
import HeaderComponent from '../../components/mainHeader/MainHeaderComponent.tsx';
import './welcome_style.css';
import logo from '../../assets/masterflour.png';

const WelcomePage: React.FC = () => {
    const navigation = useNavigate();

    return (
        <div className="welcome-container">
            <HeaderComponent />
            <div className="background-container">
                <div className="logo-container">
                    <img src={logo} alt="There is must be a logo" className="logo" />
                    <div className={"auth-buttons-container"}>
                        <button className="enter-button" type={"submit"}
                                onClick={() => navigation("/login")}>Войти</button>
                        <button className="register-button" type={"submit"}
                                onClick={() => navigation("/registration")}>Зарегистрироваться</button>
                    </div>
                </div>
                <div className={"text-desc-container"}>
                    <h1>Добро пожаловать на сайт компании "Мастер Пол"!</h1>
                    <p>
                        Мы рады приветствовать вас на нашем сайте! Здесь вы найдете всю необходимую информацию о нашей компании,
                        услугах и продуктах, а также сможете связаться с нами в случае необходимости.
                    </p>
                </div>
            </div>
        </div>
    );
};

export default WelcomePage;