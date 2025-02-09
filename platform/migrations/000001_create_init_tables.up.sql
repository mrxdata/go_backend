CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Europe/Moscow";

CREATE TABLE users (
    id UUID NOT NULL DEFAULT uuid_generate_v4 () PRIMARY KEY,   -- Уникальный идентификатор пользователя
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),         -- Время создания записи
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),         -- Время последнего обновления записи
    status INT NOT NULL DEFAULT 1,                              -- Статус пользователя
    role INT NOT NULL DEFAULT 0,                                -- Роль пользователя
    signed_via INT NOT NULL,                                    -- Начальный Способ авторизации
    tg_user_id BIGINT DEFAULT 0,                             -- Telegram ID пользователя
    tg_first_name VARCHAR(255) DEFAULT NULL,                    -- Имя пользователя в Telegram
    tg_last_name VARCHAR(255) DEFAULT NULL,                     -- Фамилия пользователя в Telegram
    tg_username VARCHAR(255) DEFAULT NULL,                      -- Никнейм пользователя в Telegram
    tg_photo_url TEXT DEFAULT NULL,                             -- Ссылка на фото профиля
    tg_hash TEXT DEFAULT NULL                                   -- Хэш данных Telegram для валидации
);

CREATE INDEX active_users ON users (id) WHERE status = 1;

