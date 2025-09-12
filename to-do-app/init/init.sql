CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deadline TIMESTAMP NULL,
    priority INT NOT NULL DEFAULT 1
);

INSERT INTO tasks (title, done, priority) VALUES
('Сделать тестовое задание', false, 3),
('Купить продукты', true, 1),
('Почитать книгу', false, 2);
