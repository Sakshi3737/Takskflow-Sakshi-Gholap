TaskFlow — Full Stack Task Management System

1. Overview

TaskFlow is a minimal but production-like task management system where users can:

- Register and log in securely (JWT-based auth)
- Create and manage projects
- Add, update, assign, and track tasks
- Filter tasks by status and assignee

Tech Stack

- Backend: Go (Gin)
- Database: PostgreSQL
- Frontend: HTML, CSS, JavaScript (Vanilla)
- Auth: JWT
- Containerization: Docker

---

2. Architecture Decisions

- Used Gin for lightweight REST APIs
- Used PostgreSQL for relational integrity
- Avoided ORM → used raw SQL for transparency
- JWT authentication for stateless APIs
- Simple frontend to prioritize backend correctness

Tradeoffs

- UI is simple (not React) to save time
- No WebSockets (would add for real-time updates)
- No pagination implemented (easy extension)

---

3. Running Locally

git clone https://github.com/your-name/taskflow
cd taskflow
cp .env.example .env
docker compose up --build

Frontend:
http://localhost:3000

Backend:
http://localhost:8080

---

4. Running Migrations

Migrations run automatically on container start.

---

5. Test Credentials

Email: test@example.com
Password: password123

---

6. API Reference

Auth

POST /auth/register
POST /auth/login

Projects

GET /projects
POST /projects
GET /projects/:id
PATCH /projects/:id
DELETE /projects/:id

Tasks

GET /projects/:id/tasks
POST /projects/:id/tasks
PATCH /tasks/:id
DELETE /tasks/:id

---

7. What I’d Do With More Time

- Add pagination for large datasets
- Implement drag-and-drop Kanban UI fully
- Add role-based permissions
- Add WebSocket-based real-time updates
- Improve UI with React + component library
- Add integration and unit tests

---
