import {BrowserRouter, Navigate, Route, Routes} from 'react-router-dom';
import ProtectedRoute from "./ProtectedRoute.tsx";
import WelcomePage from "../pages/WelcomePage/WelcomePage.tsx";
import RegistrationPage from "../pages/Registration/RegistrationPage.tsx";
import LoginPage from "../pages/LoginPages/LoginPage.tsx";
import HomePage from "../pages/HomePage/HomePage.tsx";

const AppRouter = () => {
    return (
        <BrowserRouter>
            <Routes>
                 {/*Публичные маршруты (без защиты)*/}
                <Route path="/" element={<Navigate to="/welcome" replace/>} />
                <Route path="/login" element={<LoginPage />} />

                {/*<Route path="/login" element={<HomePage />} />*/}

                <Route path="/registration" element={<RegistrationPage />} />
                <Route path="/welcome" element={<WelcomePage />} />

                 {/*Защищенные маршруты*/}
                <Route element={<ProtectedRoute />}>
                    <Route path="/home" element={<HomePage />} />
                </Route>
            </Routes>
        </BrowserRouter>
    );
};


export default AppRouter;