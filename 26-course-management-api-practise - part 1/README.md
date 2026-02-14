# create mod file using this command
-> module github.com/SaddamMohammad1/26-course-management-api-practise

# Install Required Package
-> go get -u github.com/gorilla/mux     <!-- for routing -- >
-> go get github.com/lib/pq         <!-- for postgres driver to connect with posrgres -- >
-> go get github.com/joho/godotenv  <!-- for loading environment variables from .env file -- >

# Project Structure
    26-course-management-api-practise/
    │
    ├── api/
    │   └── main.go
    │
    ├── config/
    │   └── db.go
    │
    ├── models/
    │   ├── course.go
    │   └── author.go
    │
    ├── repository/
    │   ├── course_repository.go
    │   └── author_repository.go
    │
    ├── services/
    │   ├── course_service.go
    │   └── author_service.go
    │
    ├── handlers/
    │   ├── course_handler.go
    │   └── author_handler.go
    │
    ├── router/
    │   └── router.go
    │
    ├── .env
    ├── go.mod

# Command to create the above structure of project to run in terminal
    mkdir api,config,models,repository,services,handlers,router; `
    ni api/main.go, `
    config/db.go, `
    models/course.go, `
    models/author.go, `
    repository/course_repository.go, `
    repository/author_repository.go, `
    services/course_service.go, `
    services/author_service.go, `
    handlers/course_handler.go, `
    handlers/author_handler.go, `
    router/router.go, `
    .env -ItemType File

# Command for create table in postgres database
    Author table:
        CREATE TABLE authors (
            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            fullname VARCHAR(100) NOT NULL,
            website VARCHAR(150)
        );

    Course table:
        CREATE TABLE courses (
            course_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            course_name VARCHAR(150) NOT NULL,
            course_price INTEGER NOT NULL CHECK (course_price > 0),
            author_id INT,
            CONSTRAINT fk_author
                FOREIGN KEY (author_id)
                REFERENCES authors(id)
                ON DELETE SET NULL
        );
