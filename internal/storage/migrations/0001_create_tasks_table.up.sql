-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT,
    deadline TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- CREATE INDEX idx_tasks_title ON tasks(title);

-- +goose Down
DROP TABLE tasks;

-- Попробовать применить эту миграцию + написать код для ее применения
-- Выяснить нужно ли устанавливать goose отдельно?
-- id чата e49302e4-020e-4d74-a698-8cd980b5ca4a