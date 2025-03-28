import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";

const DisplayPage: React.FC = () => {
    const [userData, setUserData] = useState<{ firstName: string; lastName: string; age: string } | null>(null);

    useEffect(() => {
        const data = Cookies.get("userData"); // Читаем куки
        if (data) {
            setUserData(JSON.parse(data));
        }
    }, []);

    return (
        <div>
            <h2>Данные пользователя</h2>
            {userData ? (
                <div>
                    <p><strong>Имя:</strong> {userData.firstName}</p>
                    <p><strong>Фамилия:</strong> {userData.lastName}</p>
                    <p><strong>Возраст:</strong> {userData.age}</p>
                </div>
            ) : (
                <p>Нет данных. Перейдите на главную страницу и введите данные.</p>
            )}
        </div>
    );
};

export default DisplayPage;