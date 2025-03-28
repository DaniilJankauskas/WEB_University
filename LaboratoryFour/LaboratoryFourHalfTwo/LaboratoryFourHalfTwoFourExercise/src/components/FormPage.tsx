import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import Cookies from "js-cookie";

const FormPage: React.FC = () => {
    const [formData, setFormData] = useState({ firstName: "", lastName: "", age: "" });
    const navigate = useNavigate();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        // Используем библиотеку js-cookie, для записи в хуки
        // Увидеть что и как можно через devtools (decodeURIComponent(document.cookie)
        Cookies.set("userData", JSON.stringify(formData), { expires: 1 });

        //Переход на страницу отображения
        navigate("/display");
    };

    return (
        <div>
            <h2>Введите ваши данные</h2>
            <form onSubmit={handleSubmit}>
                <input type="text" name="firstName" placeholder="Имя" value={formData.firstName} onChange={handleChange} required />
                <br />
                <input type="text" name="lastName" placeholder="Фамилия" value={formData.lastName} onChange={handleChange} required />
                <br />
                <input type="number" name="age" placeholder="Возраст" value={formData.age} onChange={handleChange} required />
                <br />
                <button type="submit">Сохранить</button>
            </form>
        </div>
    );
};

export default FormPage;