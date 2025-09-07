import { useEffect, useState } from "react";
import { Notifications } from "../../components/Notifications";
import apiFetch from "../../utils/apiFetch.tsx"
import "./contentStyle.css"
import "./home_style.css"

type TypeProduct = {
    ID: number;
    NameType: string;
    Coefficient: number;
};

type TypeMaterial = {
    ID: number;
    NameType: string;
    PercentOfMarriage: number;
};

interface DynamicDataTableProps {
    data: Record<string, unknown>[];
    columns: string[];
}

const DynamicDataTable: React.FC<DynamicDataTableProps> = ({ data, columns }) => {
    return (
        <div style={{ overflowX: "auto" }}>
            <table style={{
                width: "100%",
                borderCollapse: "collapse",
                marginTop: "20px"
            }}>
                <thead>
                <tr>
                    {columns.map(col => (
                        <th key={col} style={{
                            border: "1px solid #ddd",
                            padding: "8px",
                            backgroundColor: "#f2f2f2",
                            textAlign: "left"
                        }}>
                            {col}
                        </th>
                    ))}
                </tr>
                </thead>
                <tbody>
                {data.map((row, index) => (
                    <tr key={index}>
                        {columns.map(col => (
                            <td key={col} style={{
                                border: "1px solid #ddd",
                                padding: "8px"
                            }}>
                                {String(row[col])}
                            </td>
                        ))}
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export const MainContent = () => {
    const [typeProducts, setTypeProducts] = useState<TypeProduct[]>([]);
    const [typeMaterials, setTypeMaterials] = useState<TypeMaterial[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const jwtToken = localStorage.getItem("token")
                const resProducts = await apiFetch('/api/data/company/types_of_products',{
                    method: 'GET',
                    headers: {
                        Authorization: `Bearer ${jwtToken}`,
                    }
                });
                const resMaterials = await fetch('/api/data/company/types_of_materials', {
                    method: 'GET',
                    headers: {
                        Authorization: `Bearer ${jwtToken}`,
                    }
                });

                const dataProducts = await resProducts.json();
                const dataMaterials = await resMaterials.json();

                setTypeProducts(dataProducts);
                setTypeMaterials(dataMaterials);
            } catch (error) {
                console.error('Ошибка при загрузке данных:', error)
            } finally {
                setLoading(false);
            }
        };

        fetchData()
    }, []);

    return (
        <div className="content-section">
            <h2>Главная</h2>
            <p>Добро пожаловать в систему управления проектами Мастер Пол</p>

            <Notifications />

            <div className="dashboard-cards">
                <div className="card card-content">
                    <h3 className="card-title">Типы продукции</h3>
                    {loading ? (
                        <p>Загрузка...</p>
                    ) : !typeProducts ? (
                        <p>Пока ничего нет</p>
                    ) : (
                        <div className="type-list-container">
                            {typeProducts.map((item) => (
                                <div className="type-item" key={item.ID}>
                                    <span className="type-name">{item.NameType}</span>
                                    <span className="type-value">Коэффициент: {item.Coefficient}</span>
                                </div>
                            ))}
                        </div>
                    )}
                </div>

                <div className="card card-content">
                    <h3 className="card-title">Типы материалов</h3>
                    {loading ? (
                        <p>Загрузка...</p>
                    ) : !typeMaterials ? (
                        <p>Пока ничего нет</p>
                    ) : (
                        <div className="type-list-container">
                            {typeMaterials.map((item) => (
                                <div className="type-item" key={item.ID}>
                                    <span className="type-name">{item.NameType}</span>
                                    <span className="type-value">Процент брака: {(item.PercentOfMarriage * 100).toFixed(2)}%</span>
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export const ProjectsContent = () => {
    const getDataButtons = [
        {id: "types_of_products", name: "Типы продукции", icon: "🗂️", api:"/api/data/company/types_of_products"},
        {id: "types_of_materials", name: "Типы материалов", icon: "🗂️", api:"/api/data/company/types_of_materials"},
        {id: "products", name: "Продукция", icon: "🗂️", api:"/api/data/company/products"},
        {id: "partners", name: "Партнеры", icon: "🗂️", api:"/api/data/company/partners"},
        {id: "prods_partners", name: "Продукция партнеров", icon: "🗂️", api:"/api/data/company/product_partners"},
    ]
    
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string  | null>(null);
    const [tableData,  setTableData] = useState<Record<string, unknown>[]>([]);
    const [columns, setColumns] = useState<string[]>([]);

    const fetchDataAndSetTable = async (url: string) => {
        try {
            setLoading(true)
            setError(null)
            const jwtToken = localStorage.getItem("token")
            const response = await fetch(url, {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${jwtToken}`,
                }
            });

            if (!response.ok) {
                throw new Error(`Ошибка при загрузке данных: ${response.status}`);
            }

            const result = await response.json();

            if (Array.isArray(result) && result.length > 0) {
                const keys = Object.keys(result[0]);
                setColumns(keys);
                setTableData(result)
            } else {
                setColumns([])
                setTableData([])
            }
        } catch (err) {
            setError(err instanceof Error ? err.message : "Неизвестная ошибка")
            setTableData([])
            setColumns([])
        } finally {
            setLoading(false)
        }
    };

    return(
        <div className={"content-section"}>
            <h3>Проекты</h3>

            <div className={"project-grid"}>
                {getDataButtons.map((button) => (
                    <button
                        key={button.id}
                        className={"title-card"}
                        onClick={() => fetchDataAndSetTable(button.api)}
                    >
                        <span className={"title-icon"}>{button.icon}</span>
                        <span className={"title-name"}>{button.name}</span>
                    </button>
                ))}
            </div>

            {/* Загрузка / Ошибка / Таблица */}
            {loading && <p>Загрузка данных...</p>}
            {error && <p style={{color: "red"}}>{error}</p>}
            {!loading && !error && tableData.length > 0 && (
                <DynamicDataTable data={tableData} columns={columns} />
            )}
            {!loading && !error && tableData.length === 0 && (
                <p>Нет данных для отображения.</p>
            )}
        </div>
    );
}


export const ScheduleContent = () => (
    <div className={"content-section"}>
        <h2>Календарь</h2>
        <p>Здесь будет типа календарь</p>
    </div>
);

export const ImportContent = () => {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [uploadStatus, setUploadStatus] = useState<string>("");

    type OperationList = { //custom type
        [key: string]: string;
    };

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files.length > 0) {
            setSelectedFile(e.target.files[0]);
            setUploadStatus("");
        }
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            setUploadStatus("Пожалуйста, выберите файл для загрузки.");
            return;
        }

        const operationList: OperationList = {
            "Material_type_import": "types_of_materials",
            "Product_type_import": "types_of_products",
            "Products_import": "products",
            "Partners_import": "partners",
            "Partner_products_import": "product_partners"
        }

        const fileNameWithoutExtension = selectedFile.name.split('.')[0];
        const operationPath = operationList[fileNameWithoutExtension];

        if (!operationPath) {
            setUploadStatus(`Операция для файла "${fileNameWithoutExtension}" не найдена`);
            return;
        }

        const formData = new FormData();
        formData.append("excel_file", selectedFile);

        try {
            const jwtToken = localStorage.getItem("token")
            setUploadStatus("Загрузка...");
            const response = await fetch(`/api/excel/import/${operationPath}`, {
                method: "POST",
                body: formData,
                headers: {
                    Authorization: `Bearer ${jwtToken}`,
                }
            });

            if (!response.ok) {
                throw new Error(`Ошибка при загрузке файла. А именно: ${response.status}`);
            }

            const result = await response.json();
            setUploadStatus(`Файл успешно загружен. Результат: ${result.message}`);
        } catch (error) {
            console.error("Ошибка загрузки файла", error);
            setUploadStatus(`Ошибка: ${error instanceof Error ? error.message : "Неизвестная ошибка"}`);
        }
    };

    return (
        <div className={"content-section"}>
            <h2>Импорт данных</h2>
            <div className={"import-section"}>
                <input type={"file"}
                accept={".xls,.xlsx"}
                onChange={handleFileChange}/>
                <button className={"import-button"}
                disabled={!selectedFile}
                onClick={handleUpload}>Загрузить файл</button>
                {uploadStatus && <p>{uploadStatus}</p>}
            </div>
        </div>
    );
};

export const SettingsContent = () => (
    <div className={"content-section"}>
        <h2>Настройки</h2>
        <div className={"settings-group"}>
            <h3>Общие настройки</h3>
            <label>
            <input type={"checkbox"}/> Уведомления
            </label>
            <label>
            <input type={"checkbox"}/> Темная тема
            </label>
        </div>
    </div>
);
