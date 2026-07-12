# Task Manager API

A high-performance RESTful API for managing tasks, built as part of the Evolza Training Program. 

## Author
**Thenuja Hansana**

## Tech Stack

<table style="width: 100%; table-layout: fixed; text-align: left;">
  <tr>
    <th style="width: 33.33%; padding: 8px;">Language</th>
    <th style="width: 33.33%; padding: 8px;">Framework</th>
    <th style="width: 33.33%; padding: 8px;">Database</th>
  </tr>
  <tr>
    <td style="padding: 8px;">Go (Golang)</td>
    <td style="padding: 8px;">Fiber v2</td>
    <td style="padding: 8px;">MongoDB</td>
  </tr>
</table>

## Features
- **Clean Architecture:** Separation of concerns using controllers, routes, and models.
- **RESTful Endpoints:** Standard HTTP methods for CRUD operations.
- **NoSQL Database:** MongoDB integration using the official Go driver.
- **Environment Configuration:** Secure `.env` variable loading.

## Project Structure

```text
task_manager_crud/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── controllers/
│   │   └── task_controller.go   # Business logic for tasks
│   ├── database/
│   │   └── database.go          # MongoDB connection setup
│   ├── models/
│   │   └── task.go              # Task struct and schema
│   └── routes/
│       └── task_routes.go       # API route definitions
├── .env                         # Environment variables (ignored in git)
├── go.mod                       # Go module dependencies
└── go.sum                       # Go module checksums