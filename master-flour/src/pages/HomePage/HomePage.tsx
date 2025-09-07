import {useState} from "react";
import {useNavigate} from "react-router-dom";
import {MainContent, ProjectsContent, ScheduleContent, ImportContent, SettingsContent} from "./contentDesc.tsx";
import logo from '../../assets/masterflour.png';
import './home_style.css';


const HomePage = () => {
    const navigate = useNavigate();
    const [isPanelOpen, setPanelOpen] = useState(true);
    const [activeSection, setActiveSection] = useState("home");


    const handleLogOut = () => {
        localStorage.removeItem('token');
        navigate('/welcome');
    };

    const togglePanel = () => {
        setPanelOpen(!isPanelOpen);
    };

    const handleSectionChange = (sectionID: string) => {
        setActiveSection(sectionID);
    };

    const panelButtons = [
        { id: "main", text: "Главная", icon: "🏠"},
        { id: "projects", text: "Проекты", icon: "📂"},
        { id: "schedule", text: "Календарь", icon: "📅"},
        { id: "import", text: "Импорт данных", icon: "📥"},
        { id: "settings", text: "Настройки", icon: "⚙️"},
        { id: "support", text: "Поддержка", icon: "❓"} //TODO: in progress
    ];

    const renderContent = () => {
        switch (activeSection) {
            case "home":
                return <MainContent/>;
            case "projects":
                return <ProjectsContent/>;
            case "schedule":
                return <ScheduleContent/>;
            case "import":
                return <ImportContent/>;
            case "settings":
                return <SettingsContent/>;
            default:
                return <MainContent/>;
        }
    };

    return (
        <div className={"home-container"}>
            <header className={"home-header"}>
                <div className={"header-left"}>
                    <img src={logo} className={"header-logo"} alt={"There is must be a logo"} />
                    {/*<h1>Мастер Пол</h1>*/}
                    <span className={"app-name"}>Мастер пол</span>
                </div>
                <div className={"header-right"}>
                    <button className={"profile-button"}>
                        <span className={"profile-icon"}>👤</span>
                        <span className={"profile-text"}>Профиль</span>
                    </button>
                    <button className={"logout-button"} onClick={handleLogOut}>Выход</button>
                </div>
            </header>

            <div className={`panel-bar-left ${isPanelOpen ? 'open' : 'closed'}`}>
                <button className={"toggle-panel"} onClick={togglePanel}>
                    {isPanelOpen ? '◄' : '►'}
                </button>

                {panelButtons.map((button) => (
                    <button
                        key={button.id}
                        className={`button ${activeSection === button.id ? 'active' : ''}`}
                        onClick={() => handleSectionChange(button.id)}>
                        <span className={"panel-icon"}>{button.icon}</span>
                        {isPanelOpen && (
                            <span className={"panel-text"}>{button.text}</span>
                        )}
                    </button>
                ))}
            </div>

            <div className="main-content">
                {renderContent()}
            </div>
        </div>
    );
};

export default HomePage;