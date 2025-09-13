CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deadline TIMESTAMP NULL,
    priority TEXT NOT NULL DEFAULT 'low'
    );

INSERT INTO tasks (title, done, priority) VALUES
('Сделать тестовое задание', false, 'high'),
('Купить продукты', true, 'low'),
('Почитать книгу', false, 'medium');