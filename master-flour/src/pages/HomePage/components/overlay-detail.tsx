import React, { useState } from 'react';
import './OverlayTemplate.css';

type OverlayProps = {
    handleDownloadTemplateFile: (templateName: string) => void;
    operationList: { [key: string]: string };
};

const OverlayTemplate: React.FC<OverlayProps> = ({ handleDownloadTemplateFile, operationList }) => {
    const [isOverlayOpen, setOverlayOpen] = useState(false);

    const toggleOverlay = () => {
        setOverlayOpen(!isOverlayOpen);
    };

    return (
        <div className="overlay-container">
            <button onClick={toggleOverlay} className="button-call-overlay-template-menu">
                Запросить темплейт файла Excel
            </button>

            <div className={`overlay ${isOverlayOpen ? 'open' : ''}`}>
                <div className="overlay-content">
                    <h3>Выберите шаблон для скачивания</h3>
                    <div className="buttons-grid">
                        {Object.keys(operationList).map((key) => (
                            <button
                                key={key}
                                className="import-button"
                                onClick={() => handleDownloadTemplateFile(operationList[key])}
                            >
                                {key}
                            </button>
                        ))}
                    </div>
                    <button onClick={toggleOverlay} className="close-overlay-btn">
                        Закрыть
                    </button>
                </div>
            </div>
        </div>
    );
};

export default OverlayTemplate;
