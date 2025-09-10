import { useState, useEffect } from 'react';
import './App.css';
import { AddTask, GetTasks, DeleteTask, ToggleTask } from "../wailsjs/go/main/App.js";

function App() {
    const [tasks, setTasks] = useState([]);
    const [title, setTitle] = useState("");

    // Загружаем задачи при старте
    useEffect(() => {
        loadTasks();
    }, []);

    async function loadTasks() {
        const data = await GetTasks();
        setTasks(data);
    }

    async function addTask() {
        if (title.trim() === "") return;
        await AddTask(title.trim());
        setTitle("");   // очистить поле
        loadTasks();    // обновить список
    }

    async function toggleTask(id) {
        await ToggleTask(id);
        loadTasks();
    }

    async function deleteTask(id) {
        await DeleteTask(id);
        loadTasks();
    }

    return (
        <div id="App">
            <h1>To-Do List</h1>

            <input
                type="text"
                placeholder="Новая задача"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
            />
            <button onClick={addTask}>Добавить</button>

            <ul>
                {tasks.map((t) => (
                    <li key={t.id} className={t.done ? "done" : ""}>
                        {t.title}
                        <button onClick={() => toggleTask(t.id)}>
                            {t.done ? "Вернуть" : "Сделано"}
                        </button>
                        <button onClick={() => deleteTask(t.id)}>Удалить</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;
