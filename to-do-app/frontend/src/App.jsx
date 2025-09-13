import { useState, useEffect } from 'react';
import './App.css';
import ConfirmModal from "./ConfirmModal";
import { CreateTask, GetAllTasks, DeleteTask, ToggleTask } from "../wailsjs/go/handler/TaskHandler.js";

function App() {
    const [tasks, setTasks] = useState([]);
    const [title, setTitle] = useState("");
    const [priority, setPriority] = useState("low");
    const [deadline, setDeadline] = useState("");
    const [filter, setFilter] = useState("all");
    const [sortBy, setSortBy] = useState("date");
    const [taskToDelete, setTaskToDelete] = useState(null);
    const [showModal, setShowModal] = useState(false);

    useEffect(() => {
        loadTasks();
    }, []);

    async function loadTasks() {
        const data = await GetAllTasks();
        setTasks(data);
    }

    async function addTask() {
        if (title.trim() === "") {
            return;
        }

        let deadlineISO = null;
        if (deadline) {
            deadlineISO = new Date(deadline).toISOString();
        }

        await CreateTask({
            title,
            done: false,
            priority,
            deadline: deadlineISO,
        });

        setTitle("");
        setPriority("low");
        setDeadline("");
        loadTasks();
    }


    async function toggleTask(task) {
        const updated = { ...task, done: !task.done };
        await ToggleTask(updated);
        loadTasks();
    }

    function confirmDelete(task) {
        setTaskToDelete(task);
        setShowModal(true);
    }

    async function handleDelete() {
        await DeleteTask(taskToDelete);
        setShowModal(false);
        setTaskToDelete(null);
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
                const order = { high: 3, medium: 2, low: 1 };
                return order[b.priority] - order[a.priority];
            }
            return 0;
        });

    // для отображения
    function priorityLabel(p) {
        if (p === "high") return "🔴 Высокий";
        if (p === "medium") return "🟠 Средний";
        return "🟢 Низкий";
    }

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
                <select value={priority} onChange={(e) => setPriority(e.target.value)}>
                    <option value="low">🟢 Низкий</option>
                    <option value="medium">🟠 Средний</option>
                    <option value="high">🔴 Высокий</option>
                </select>
                <input
                    type="datetime-local"
                    value={deadline}
                    onChange={(e) => setDeadline(e.target.value)}
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
                            <span className="task-priority">{priorityLabel(t.priority)}</span>
                        </div>

                        <button className="delete-btn" onClick={() => confirmDelete(t)}>
                            🗑
                        </button>
                    </li>
                ))}
            </ul>

            <ConfirmModal
                show={showModal}
                task={taskToDelete}
                onClose={() => setShowModal(false)}
                onConfirm={handleDelete}
            />
        </div>
    );
}

export default App;
