# Course Management API

REST API for managing courses and authors with **pagination**, **filter**, **search**, **login**, and **role-based access** (author vs admin).

---

## Create module

```bash
go mod init github.com/SaddamMohammad1/26-course-management-api-practise
```

## Install required packages

```bash
go get -u github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5
```

---

## Project structure

```
26-course-management-api-practise/
├── api/
│   └── main.go
├── config/
│   └── db.go
├── models/
│   ├── course.go
│   ├── author.go
│   ├── api_response.go
│   ├── error_response.go
│   └── auth.go
├── repository/
│   ├── course_repository.go
│   └── author_repository.go
├── services/
│   ├── course_service.go
│   ├── author_service.go
│   └── errors.go
├── handlers/
│   ├── course_handler.go
│   ├── author_handler.go
│   └── auth_handler.go
├── middleware/
│   └── auth.go
├── router/
│   └── router.go
├── utils/
│   └── response.go
├── .env
├── go.mod
└── README.md
```

---

## Environment variables (.env)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_db_name
JWT_SECRET=your-secret-key-change-in-production
```

---

## Database setup

### 1. Authors table (with auth columns)

```sql
CREATE TABLE authors (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    website VARCHAR(150),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'author'
);
```

### 2. Courses table

```sql
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
```

### 3. If you already have `authors` without auth columns, run:

```sql
ALTER TABLE authors
    ADD COLUMN IF NOT EXISTS email VARCHAR(255) UNIQUE,
    ADD COLUMN IF NOT EXISTS password_hash VARCHAR(255),
    ADD COLUMN IF NOT EXISTS role VARCHAR(20) DEFAULT 'author';

-- Then set NOT NULL and backfill existing rows if needed:
-- UPDATE authors SET email = 'author' || id || '@example.com', password_hash = '', role = 'author' WHERE email IS NULL;
-- ALTER TABLE authors ALTER COLUMN email SET NOT NULL;
-- ALTER TABLE authors ALTER COLUMN password_hash SET NOT NULL;
```

### 4. Create first admin user (after app can hash passwords, use a small script or API later)

You need one admin to add authors. Option A: insert with a pre-generated bcrypt hash:

```sql
-- Password: admin123 (generate your own hash in Go and replace below)
INSERT INTO authors (fullname, website, email, password_hash, role)
VALUES ('Admin', 'https://example.com', 'admin@gmail.com', '$2a$10$iTJkulEg9FTNvi5wttsG4ep2oEto8iQsuUDMc57ofNUhN1RtaVN.S', 'admin');
```

Option B: Temporarily add a signup or use a one-off Go script to hash a password and insert the admin row.

---

## Run the API

```bash
cd api
go run main.go
```

Server runs at `http://localhost:8000`.

---

## API reference

### Authentication

#### POST `/login`

Login as author or admin. Returns JWT and author info.

**Request body:**

```json
{
  "email": "admin@gmail.com",
  "password": "admin123"
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "author": {
      "id": 1,
      "fullname": "Hitesh Choudhary",
      "website": "https://hitesh.ai",
      "email": "author@example.com",
      "role": "author"
    }
  },
  "error": ""
}
```

Use the `token` in subsequent requests:

```
Authorization: Bearer <token>
```

---

### Authors (admin only)

#### POST `/authors`

Create a new author (admin only). New author can then login and create their own courses.

**Headers:** `Authorization: Bearer <admin_jwt>`

**Request body:**

```json
{
  "fullname": "New Author",
  "website": "https://newauthor.com",
  "email": "newauthor@example.com",
  "password": "secure_password"
}
```

**Response:** `201 Created` with author object (no password).

---

### Courses

#### GET `/courses` — List courses (pagination, filter, search)

- **Public:** No `Authorization` → returns all courses (with pagination/filter/search).
- **Author:** With `Authorization` (author role) → returns **only that author’s courses** (same query params apply).

**Query parameters:**

| Parameter   | Type   | Description                          |
|------------|--------|--------------------------------------|
| `page`     | int    | Page number (default: 1)             |
| `limit`    | int    | Items per page (default: 10, max: 100) |
| `search`   | string | Search by course name (partial, case-insensitive) |
| `author_id`| int    | Filter by author ID (ignored when logged-in author) |
| `min_price`| int    | Minimum course price                 |
| `max_price`| int    | Maximum course price                 |

**Examples:**

- `GET /courses?page=1&limit=10`
- `GET /courses?search=go&limit=5`
- `GET /courses?author_id=2&min_price=100&max_price=5000`
- `GET /courses?page=2&limit=20&search=programming`

**Response:**

```json
{
  "success": true,
  "data": {
    "courses": [
      {
        "courseId": 1,
        "courseName": "Go Programming",
        "coursePrice": 4999,
        "author": {
          "id": 1,
          "fullname": "Hitesh Choudhary",
          "website": "https://hitesh.ai"
        }
      }
    ],
    "page": 1,
    "limit": 10,
    "total": 25,
    "totalPage": 3
  },
  "error": ""
}
```

---

#### GET `/courses/{id}` — Get one course

- **Public / Admin:** Can get any course.
- **Author:** Can get **only their own** course; otherwise `403 Forbidden`.

**Response:** Single course object with author.

---

#### POST `/courses` — Create course (author or admin)

- **Author:** Must be logged in; `authorId` in body is **ignored** — course is created for the logged-in author.
- **Admin:** Can set `authorId` in body to create course for any author.

**Headers:** `Authorization: Bearer <token>`

**Request body:**

```json
{
  "courseName": "Go Masterclass",
  "coursePrice": 4999,
  "authorId": 1
}
```

(For author role, `authorId` is overridden by token.)

**Response:** `201 Created` with created course.

---

#### PUT `/courses/{id}` — Update course (author or admin)

- **Author:** Can update **only their own** course.
- **Admin:** Can update any course.

**Headers:** `Authorization: Bearer <token>`

**Request body:** Same as create (`courseName`, `coursePrice`, `authorId`; admin can change `authorId`).

---

#### DELETE `/courses/{id}` — Delete course (author or admin)

- **Author:** Can delete **only their own** course.
- **Admin:** Can delete any course.

**Headers:** `Authorization: Bearer <token>`

**Response:** `200 OK` with message.

---

## Roles summary

| Role   | Login | Add author | List courses     | Create course | Update/Delete course   |
|--------|-------|------------|-------------------|---------------|-------------------------|
| Admin  | Yes   | Yes        | All (with filters)| Any author    | Any course              |
| Author | Yes   | No         | Only own          | Own only      | Own courses only       |
| Public | No    | No         | All (no auth)     | No            | No                      |

---

## Quick test (after creating an admin and an author)

1. **Login as admin:**  
   `POST /login` with admin email/password → get `token`.

2. **Add author:**  
   `POST /authors` with `Authorization: Bearer <admin_token>` and body `fullname`, `website`, `email`, `password`.

3. **Login as author:**  
   `POST /login` with new author email/password → get `token`.

4. **Author creates course:**  
   `POST /courses` with `Authorization: Bearer <author_token>`, body `courseName`, `coursePrice` (no need to send `authorId`).

5. **Author lists only their courses:**  
   `GET /courses` with `Authorization: Bearer <author_token>` → returns only that author’s courses (supports `page`, `limit`, `search`, etc.).

6. **Public list (no auth):**  
   `GET /courses?page=1&limit=10` → returns all courses with pagination.

---

## Commands to create project structure (Windows PowerShell)

```powershell
mkdir api, config, models, repository, services, handlers, router, middleware, utils -Force
ni api/main.go, config/db.go, models/course.go, models/author.go, models/auth.go, models/api_response.go, models/error_response.go, repository/course_repository.go, repository/author_repository.go, services/course_service.go, services/author_service.go, services/errors.go, handlers/course_handler.go, handlers/author_handler.go, handlers/auth_handler.go, middleware/auth.go, router/router.go, utils/response.go, .env -ItemType File
```
