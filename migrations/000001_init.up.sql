CREATE TABLE users {
    full_name: VARCHAR(200) NOT NULL,
    phone_number: VARCHAR(100)
}

 migrate -path migrations -database "PATH?sslmode=disable" up
