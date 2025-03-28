import { useState } from "react";

const TaskThree: React.FC = () => {
    const [text, setText] = useState("");
    const [wordCount, setWordCount] = useState(0);
    const [charCount, setCharCount] = useState(0);

    const countText = () => {
        const trimmedText = text.trim();
        const charCount = trimmedText.length;
        const wordCount = trimmedText ? trimmedText.split(/\s+/).length : 0;

        setCharCount(charCount);
        setWordCount(wordCount);
    };

    return (
        <div>
            <textarea
                value={text}
                onChange={(e) => setText(e.target.value)}
                placeholder="Введите текст..."
            />

            <br/>

            <button onClick={countText}>Посчитать</button>

            <p>
                <strong>Слов:</strong> {wordCount}
            </p>
            <p>
                <strong>Символов:</strong> {charCount}
            </p>
        </div>
    );
};

export default TaskThree;