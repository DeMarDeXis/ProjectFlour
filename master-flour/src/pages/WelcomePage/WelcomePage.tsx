import React from "react";
import {useNavigate} from "react-router-dom";
import './welcome_style.css';
import logo from '../../assets/masterflour.png';

const WelcomePage: React.FC = () => {
    console.log('WelcomePage is rendering');
    const navigation = useNavigate();

    const navigateTGChannel = () => {
        console.log("Navigate to Telegram channel");
    };


    return (
        <div className="welcome-container">
            <header className="welcome-header">
                <div className="header-content">
                    <h1>Мастер Пол</h1>
                    <div className="navbar">
                        <ul>
                            {/*TODO: To learn types of button*/}
                            <li><button type="submit" className="Catalog-But">Каталог</button></li>
                            <li><img src={"https://i.pinimg.com/736x/2a/43/5a/2a435a6985dfc24fa1cda76c5507cd30.jpg"}
                                     alt={"There must be TG"} className={"tg-logo"} onSubmit={navigateTGChannel}/></li>
                        </ul>
                    </div>
                </div>
            </header>
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