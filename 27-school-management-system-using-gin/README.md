# create mod file using this command
-> module github.com/SaddamMohammad1/27-school-management-system-using-gin

# Step of creating project
Step 1 → Setup project
Step 2 → DB connection
Step 3 → Basic server run
Step 4 → User model
Step 5 → Register API
Step 6 → Login API
Step 7 → JWT
Step 8 → Auth middleware
Step 9 → RBAC
Step 10 → Student module
Step 11 → Teacher module
Step 12 → Class module
Step 13 → Grade module
Step 14 → Attendance module


# Install Required Package
-> go get github.com/gin-gonic/gin
-> go get github.com/lib/pq
-> go get github.com/joho/godotenv
-> go get github.com/golang-jwt/jwt/v5
-> go get golang.org/x/crypto/bcrypt

# Project Structure
    27-school-management-system-using-gin/
    │
    ├── cmd/
    │   └── main.go
    ├── config/
    │   └── db.go
    ├── models/
    │   ├── school.go
    │   ├── user.go
    │   ├── teacher.go
    │   ├── student.go
    │   ├── class.go
    │   ├── subject.go
    │   ├── grade.go
    │   └── attendance.go
    │
    ├── handlers/
    │   ├── auth_handler.go
    │   ├── user_handler.go
    │   ├── student_handler.go
    │   ├── teacher_handler.go
    │   ├── class_handler.go
    │   ├── subject_handler.go
    │   ├── grade_handler.go
    │   └── attendance_handler.go
    │
    ├── middleware/
    │   ├── auth.go
    │   └── rbac.go
    │
    ├── repository/
    │   ├── user_repo.go
    │   ├── student_repo.go
    │   ├── teacher_repo.go
    │   ├── class_repo.go
    │   ├── grade_repo.go
    │
    ├── services/
    │   ├── auth_service.go
    │   ├── grade_service.go
    │
    ├── utils/
    │   └── jwt.go
    │
    └── .env


# Command for create table in postgres database
    -- Schools
    CREATE TABLE schools (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(150) NOT NULL,
        address TEXT,
        phone VARCHAR(20),
        email VARCHAR(120) UNIQUE,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP
    );

    -- Users
    CREATE TABLE users (
        id BIGSERIAL PRIMARY KEY,
        school_id BIGINT REFERENCES schools(id),
        name VARCHAR(120) NOT NULL,
        email VARCHAR(120) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        role VARCHAR(20) CHECK (role IN ('admin','teacher','student')),
        is_active BOOLEAN DEFAULT TRUE,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP
    );

    -- Teachers
    CREATE TABLE teachers (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT UNIQUE REFERENCES users(id),
        employee_id VARCHAR(50) UNIQUE,
        qualification VARCHAR(150),
        experience_years INT,
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- Classes
    CREATE TABLE classes (
        id BIGSERIAL PRIMARY KEY,
        school_id BIGINT REFERENCES schools(id),
        name VARCHAR(100) NOT NULL,
        section VARCHAR(10),
        teacher_id BIGINT REFERENCES teachers(id),
        academic_year VARCHAR(20),
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- Students
    CREATE TABLE students (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT UNIQUE REFERENCES users(id),
        roll_number VARCHAR(50) NOT NULL,
        class_id BIGINT REFERENCES classes(id),
        admission_date DATE,
        date_of_birth DATE,
        guardian_name VARCHAR(120),
        guardian_phone VARCHAR(20),
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- Subjects
    CREATE TABLE subjects (
        id BIGSERIAL PRIMARY KEY,
        class_id BIGINT REFERENCES classes(id),
        name VARCHAR(100) NOT NULL,
        code VARCHAR(20) UNIQUE,
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- Grades
    CREATE TABLE grades (
        id BIGSERIAL PRIMARY KEY,
        student_id BIGINT REFERENCES students(id),
        subject_id BIGINT REFERENCES subjects(id),
        exam_type VARCHAR(50) NOT NULL,
        marks_obtained NUMERIC(5,2) NOT NULL,
        total_marks NUMERIC(5,2) NOT NULL,
        remarks TEXT,
        created_at TIMESTAMP DEFAULT NOW()
    );

    -- Attendance
    CREATE TABLE attendance (
        id BIGSERIAL PRIMARY KEY,
        student_id BIGINT REFERENCES students(id),
        class_id BIGINT REFERENCES classes(id),
        date DATE NOT NULL,
        status VARCHAR(20) CHECK (status IN ('present','absent','late')),
        marked_by BIGINT REFERENCES teachers(id),
        created_at TIMESTAMP DEFAULT NOW()
    );

