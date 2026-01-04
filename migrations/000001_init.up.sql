CREATE TABLE users {
    full_name: VARCHAR(200) NOT NULL,
    phone_number: VARCHAR(100)
}


-- Создание новой миграции:
-- migrate create -ext sql -dir migrations -seq exampleName
--
-- Где:
--  -ext sql        -> формат файлов (.sql)
--  -dir migrations -> директория для миграций
--  -seq            -> использовать последовательную нумерацию
--  exampleName     -> осмысленное имя миграции (например: create_users_table)

-- Применение миграций:
-- migrate -path migrations -database "PATH?sslmode=disable" up
--
-- PATH — строка подключения к БД (например: postgres://user:pass@localhost:5432/dbname)
-- sslmode=disable используется для локальной разработк