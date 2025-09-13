import { useState, useEffect } from 'react';
import './App.css';
import ConfirmModal from "./ConfirmModal";
import { CreateTask, GetAllTasks, DeleteTask, ToggleTask } from "../wailsjs/go/handler/TaskHandler.js";

function App() {
    const [tasks, setTasks] = useState([]);
    const [title, setTitle] = useState("");
    const [priority, setPriority] = useState("low");
    const [deadline, setDeadline] = useState("");
    const [statusFilter, setStatusFilter] = useState("all"); // all | active | done
    const [deadlineFilter, setDeadlineFilter] = useState("all"); // all | today | week | overdue
    const [sortBy, setSortBy] = useState("date"); // date | priority
    const [taskToDelete, setTaskToDelete] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const [error, setError] = useState("");
    const [theme, setTheme] = useState("light");

    useEffect(() => {
        loadTasks();
    }, []);

    useEffect(() => {
        document.body.setAttribute("data-theme", theme);
    }, [theme]);

    function toggleTheme() {
        setTheme((prev) => (prev === "light" ? "dark" : "light"));
    }

    async function loadTasks() {
        const data = await GetAllTasks();
        setTasks(data);
    }

    async function addTask() {
        if (title.trim() === "") {
            setError("❌ Введите название задачи!");
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
        setError("");
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

    // фильтрация по статусу + deadline
    const now = new Date();
    const startOfToday = new Date(now.getFullYear(), now.getMonth(), now.getDate());
    const endOfToday = new Date(startOfToday);
    endOfToday.setDate(endOfToday.getDate() + 1);

    const endOfWeek = new Date(startOfToday);
    endOfWeek.setDate(endOfWeek.getDate() + 7);

    const filteredTasks = tasks
        .filter((t) => {
            // фильтр по статусу
            if (statusFilter === "active" && t.done) return false;
            if (statusFilter === "done" && !t.done) return false;

            // фильтр по deadline
            if (deadlineFilter === "today") {
                if (!t.deadline) return false;
                const d = new Date(t.deadline);
                return d >= startOfToday && d < endOfToday;
            }
            if (deadlineFilter === "week") {
                if (!t.deadline) return false;
                const d = new Date(t.deadline);
                return d >= startOfToday && d < endOfWeek;
            }
            if (deadlineFilter === "overdue") {
                if (!t.deadline) return false;
                const d = new Date(t.deadline);
                return d < now && !t.done;
            }
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

    // отображение текста приоритета
    function priorityLabel(p) {
        if (p === "high") return "🔴 Высокий";
        if (p === "medium") return "🟠 Средний";
        return "🟢 Низкий";
    }

    return (
        <div id="App">
            <header className="app-header">
                <h1>📋 To-Do List</h1>
                <button className="theme-toggle" onClick={toggleTheme}>
                    {theme === "light" ? "🌙 Тёмная" : "☀️ Светлая"}
                </button>
            </header>

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
            {error && <p className="error-text">{error}</p>}

            {/* фильтры */}
            <div className="filters">

                <select value={statusFilter} onChange={(e) => setStatusFilter(e.target.value)}>
                    <option value="all">Все</option>
                    <option value="active">Активные</option>
                    <option value="done">Выполненные</option>
                </select>

                <select value={deadlineFilter} onChange={(e) => setDeadlineFilter(e.target.value)}>
                    <option value="all">Все по срокам</option>
                    <option value="today">Сегодня</option>
                    <option value="week">На неделю</option>
                    <option value="overdue">Просроченные</option>
                </select>

                <select value={sortBy} onChange={(e) => setSortBy(e.target.value)}>
                    <option value="date">📅 По дате</option>
                    <option value="priority">⭐ По приоритету</option>
                </select>
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
                                    ⏳ {new Date(t.deadline).toLocaleString([], {
                                    year: "numeric",
                                    month: "2-digit",
                                    day: "2-digit",
                                    hour: "2-digit",
                                    minute: "2-digit"
                                })}
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
