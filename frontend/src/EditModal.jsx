import React, { useState, useEffect } from "react";

function EditModal({ show, task, onClose, onSave }) {
    const [title, setTitle] = useState("");
    const [priority, setPriority] = useState("low");
    const [deadline, setDeadline] = useState("");

    useEffect(() => {
        if (task) {
            setTitle(task.title);
            setPriority(task.priority);
            setDeadline(task.deadline ? task.deadline.slice(0, 16) : ""); // формат для datetime-local
        }
    }, [task]);

    if (!show) return null;

    function handleSave() {
        if (title.trim() === "") return;
        const updatedTask = {
            ...task,
            title,
            priority,
            deadline: deadline ? new Date(deadline).toISOString() : null,
        };
        onSave(updatedTask);
    }

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h3>✏️ Редактировать задачу</h3>

                <input
                    type="text"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    placeholder="Название задачи"
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

                <div className="modal-buttons">
                    <button className="cancel" onClick={onClose}>Отмена</button>
                    <button className="confirm" onClick={handleSave}>Сохранить</button>
                </div>
            </div>
        </div>
    );
}

export default EditModal;
