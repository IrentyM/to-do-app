import { useState, useEffect } from 'react';
import './App.css';
import { CreateTask, GetAllTasks, DeleteTask, ToggleTask } from "../wailsjs/go/handler/TaskHandler.js";

function App() {
    const [tasks, setTasks] = useState([]);
    const [title, setTitle] = useState("");
    const [filter, setFilter] = useState("all"); // all | active | done
    const [sortBy, setSortBy] = useState("date"); // date | priority

    useEffect(() => {
        loadTasks();
    }, []);

    async function loadTasks() {
        const data = await GetAllTasks();
        setTasks(data);
    }

    async function addTask() {
        if (title.trim() === "") return;
        await CreateTask({ title, done: false, priority: 1 });
        setTitle("");
        loadTasks();
    }

    async function toggleTask(task) {
        const updated = { ...task, done: !task.done };
        await ToggleTask(updated);
        loadTasks();
    }

    async function deleteTask(task) {
        await DeleteTask(task);
        loadTasks();
    }

    // фильтрация
    const filteredTasks = tasks
        .filter((t) => {
            if (filter === "active") return !t.done;
            if (filter === "done") return t.done;
            return true;
        })
        .sort((a, b) => {
            if (sortBy === "date") {
                return new Date(b.createdAt) - new Date(a.createdAt);
            }
            if (sortBy === "priority") {
                return b.priority - a.priority;
            }
            return 0;
        });

    return (
        <div id="App">
            <h1>📋 To-Do List</h1>

            {/* input */}
            <div className="task-input">
                <input
                    type="text"
                    placeholder="Новая задача..."
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                <button onClick={addTask}>➕ Добавить</button>
            </div>

            {/* фильтры */}
            <div className="filters">
                <div className="filter-group">
                    <button
                        className={filter === "all" ? "active" : ""}
                        onClick={() => setFilter("all")}
                    >
                        Все
                    </button>
                    <button
                        className={filter === "active" ? "active" : ""}
                        onClick={() => setFilter("active")}
                    >
                        Активные
                    </button>
                    <button
                        className={filter === "done" ? "active" : ""}
                        onClick={() => setFilter("done")}
                    >
                        Выполненные
                    </button>
                </div>

                <div className="filter-group">
                    <button
                        className={sortBy === "date" ? "active" : ""}
                        onClick={() => setSortBy("date")}
                    >
                        📅 По дате
                    </button>
                    <button
                        className={sortBy === "priority" ? "active" : ""}
                        onClick={() => setSortBy("priority")}
                    >
                        ⭐ По приоритету
                    </button>
                </div>
            </div>

            {/* список */}
            <ul className="task-list">
                {filteredTasks.map((t) => (
                    <li key={t.id} className={t.done ? "done" : ""}>
                        <div className="task-main">
                            <input
                                type="checkbox"
                                checked={t.done}
                                onChange={() => toggleTask(t)}
                            />
                            <span className="task-title">{t.title}</span>
                        </div>

                        <div className="task-meta">
                            {t.deadline && (
                                <span className="task-deadline">
                                    ⏳ {new Date(t.deadline).toLocaleDateString()}
                                </span>
                            )}
                            <span className="task-priority">⭐ {t.priority}</span>
                        </div>

                        <button className="delete-btn" onClick={() => deleteTask(t)}>
                            🗑
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;
