# Task Manager - Evolza Training Program

A full-stack Task Manager application built as part of the Intern Training Program. This project demonstrates clean architecture and separation of concerns using **Golang (Fiber v2)** for the backend REST API, **MongoDB** for persistent storage, and **React (TypeScript + Vite)** for the minimal user interface.

---

## + Features

- **Full CRUD API**: Create, view all, view by ID, update, and delete tasks.
- **Task Status Pipeline**: Transition tasks between `pending`, `in-progress`, and `completed`.
- **Validation**: Server-side request body and field validations.
- **Robust UI**: Interactive frontend with loading indicators, error handling, and responsive components.
- **API Documentation**: Pre-configured Postman Collection with environment variables.

---

## 📁 Repository Structure

```text
task-manager/
├── backend/                             # Golang Fiber REST API
│   ├── cmd/server/main.go               # App entry point
│   ├── config/db.go                     # Database connection configuration
│   ├── controllers/task_controller.go   # Request handlers (HTTP mapping)
│   ├── models/task.go                   # Task schema, status types & validations
│   ├── routes/routes.go                 # API routing registration
│   └── .env                             # Local environment variables
│
├── frontend/                            # React TypeScript UI
│   ├── src/
│   │   ├── components/                  # TaskForm, TaskCard, TaskList
│   │   ├── services/api.ts              # API integration (fetch/axios)
│   │   └── types/task.ts                # Shared TypeScript models
│   ├── App.tsx                          # App state & layout wrapper
│   └── package.json                     # NPM dependencies
│
└── postman/                             # API Collection
    └── Task_Manager.postman_collection.json
```

---

## 🛠️ Prerequisites

Before getting started, ensure you have the following installed on your machine:
1. **Golang** (Go 1.18 or higher) - [Download Go](https://go.dev/dl/)
2. **Node.js** (v16 or higher) - [Download Node.js](https://nodejs.org/)
3. **MongoDB** (Local instance running or MongoDB Atlas Connection URI) - [Download MongoDB](https://www.mongodb.com/try/download/community)
4. **Postman** (For testing endpoints) - [Download Postman](https://www.postman.com/downloads/)

---

## ⚙️ Backend Setup & Execution

### 1. Configure Environment Variables
Inside the `backend/` directory, create a `.env` file:
```env
PORT=4000
MONGO_URI=mongodb://localhost:27017
DB_NAME=task_manager_db
```

### 2. Initialize and Install Dependencies
Navigate to the `backend/` folder and run:
```bash
# Initialize module
go mod init backend

# Install Fiber framework and MongoDB Driver
go get github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/mongo
```

### 3. Start the Server
```bash
go run cmd/server/main.go
```
The server will start running at `http://localhost:4000`. You can test the health-check route at `http://localhost:4000/api/health`.

---

## 💻 Frontend Setup & Execution

### 1. Create React App
Navigate to the `frontend/` folder and run:
```bash
# Create a Vite React TypeScript project
npm create vite@latest . -- --template react-ts

# Install packages
npm install
```

### 2. Start the Frontend Dev Server
```bash
npm run dev
```
The client dashboard will typically load at `http://localhost:5173`.

---

## 🔌 API Documentation

All endpoints receive and return `application/json`.

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| **GET** | `/api/health` | Check backend server status |
| **POST** | `/api/tasks` | Create a new task |
| **GET** | `/api/tasks` | Retrieve all tasks |
| **GET** | `/api/tasks/:id` | Retrieve a task by MongoDB ObjectID |
| **PUT** | `/api/tasks/:id` | Update title, description, or status of a task |
| **DELETE** | `/api/tasks/:id` | Delete task |

---

## 🧪 Testing with Postman

1. Import the Postman Collection located in `postman/Task_Manager.postman_collection.json`.
2. Configure your Environment Variables in Postman:
   - `baseUrl`: `http://localhost:4000`
   - `taskId`: (Populated dynamically or manually from a created task's ObjectID).
