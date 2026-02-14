# create mod file using this command
-> github.com/SaddamMohammad1/26-practise-building-apis-format-wise

# Install Required Package
-> go get -u github.com/gorilla/mux     <!-- for routing -- >
-> go get github.com/lib/pq         <!-- for postgres driver to connect with posrgres -- >
-> go get github.com/joho/godotenv  <!-- for loading environment variables from .env file -- >

# Project Structure
    26-practise-building-apis-format-wise/
    │
    ├── api/
    │   └── main.go
    ├── router/
    │   └── router.go
    ├── handlers/
    │   └── course.go
    ├── models/
    │   └── course.go
    ├── services/
    │   └── course_service.go
    ├── repository/
    │   └── course_repo.go
    └── go.mod

# Command to create the above structure of project to run in terminal
    mkdir api, router, config, handlers, models, services, repository

    New-Item .env -ItemType File
    New-Item api/main.go -ItemType File
    New-Item config/db.go -ItemType File
    New-Item router/router.go -ItemType File
    New-Item handlers/course.go -ItemType File
    New-Item models/course.go -ItemType File
    New-Item services/course_service.go -ItemType File
    New-Item repository/course_repo.go -ItemType File
