

// Добавить задачу
async function addTask() {
    const input = document.getElementById("taskInput");
    if (input.value.trim() === "") return;
    await AddTask(input.value.trim());
    input.value = "";
    renderTasks();
}

// Отрисовать список задач
async function renderTasks() {
    const tasks = await GetTasks();
    const list = document.getElementById("taskList");
    list.innerHTML = "";
    tasks.forEach(t => {
        const li = document.createElement("li");
        li.className = t.done ? "done" : "";
        li.textContent = t.title;

        const toggleBtn = document.createElement("button");
        toggleBtn.textContent = t.done ? "Вернуть" : "Сделано";
        toggleBtn.onclick = async () => {
            await ToggleTask(t.id);
            renderTasks();
        };

        const delBtn = document.createElement("button");
        delBtn.textContent = "Удалить";
        delBtn.onclick = async () => {
            await DeleteTask(t.id);
            renderTasks();
        };

        li.appendChild(toggleBtn);
        li.appendChild(delBtn);
        list.appendChild(li);
    });
}

document.getElementById("addBtn").onclick = addTask;

renderTasks();
