import "./ConfirmModal.css";

function ConfirmModal({ show, onClose, onConfirm, task }) {
    if (!show) return null;

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h3>Удалить задачу?</h3>
                <p>Вы уверены, что хотите удалить задачу: <b>{task?.title}</b>?</p>

                <div className="modal-buttons">
                    <button className="cancel" onClick={onClose}>Отмена</button>
                    <button className="confirm" onClick={onConfirm}>Удалить</button>
                </div>
            </div>
        </div>
    );
}

export default ConfirmModal;
