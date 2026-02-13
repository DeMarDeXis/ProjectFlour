import styled from "styled-components";
// import {useNavigate} from "react-router-dom";

const HeaderContainer = styled.header`
        background-color: var(--second-bg-color);
        padding: 1rem 0;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        position: sticky;
        top: 0;
        z-index: 100;
    `;

const HeaderContent = styled.div`
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 2rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
    
        h1 {
            font-size: 1.8rem;
            margin: 0;
            color: var(--activity-color);
        }
    
    @media (max-width: 768px) {
        flex-direction: column;
        gap: 1rem;
        padding: 0 1rem;
        
        h1 {
            font-size: 1.5rem;            
        }
    }
`;

const NavBar = styled.nav`
    ul {
        list-style: none;
        margin: 0;
        padding: 0;
        display: flex;
        align-items: center;
        gap: 1rem;
    }
    
    li {
        display: flex;
        align-items: center;
    }
`;

const CatalogButton = styled.button`
    background-color: var(--activity-color);
    color: white;
    border: none;
    padding: 0.8rem 1.5rem;
    border-radius: 8px;
    font-weight: 600;
    transition: background-color 0.3s ease;
    
    &:hover {
        background-color: var(--subactivity-hover-color);
    }
`;

const TGLogo = styled.img`
    width: 40px;
    height: 40px;
    border-radius: 50%;
    cursor: pointer;
    transition: transform 0.3s ease;
    
    &:hover {
        transform: scale(1.1);
    }
`;

const HeaderComponent: React.FC = () => {
    const img_src = "https://i.pinimg.com/736x/2a/43/5a/2a435a6985dfc24fa1cda76c5507cd30.jpg";
    // const navigation = useNavigate();

    const navigateTGChannel = () => {
        console.log("Navigate to Telegram channel");
    };

    return (
        <HeaderContainer>
            <HeaderContent>
                <h1>Мастер Пол</h1>
                <NavBar>
                    <ul>
                        <li><CatalogButton type="button" >Каталог</CatalogButton></li>
                        <li><TGLogo src={img_src} alt={"TG logo"} onClick={navigateTGChannel}></TGLogo></li>
                    </ul>
                </NavBar>
            </HeaderContent>
        </HeaderContainer>
    );
};

export default HeaderComponent;