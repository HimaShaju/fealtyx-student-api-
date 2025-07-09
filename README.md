# FealtyX Student API (Go)

A simple REST API to perform CRUD operations on a list of students, with AI-generated profile summaries via Ollama.

## Student Attributes

- `id`: Integer
- `name`: String
- `age`: Integer
- `email`: String

## Endpoints

| Method | Endpoint                  | Description                     |
|--------|---------------------------|---------------------------------|
| POST   | `/students`               | Create a new student            |
| GET    | `/students`               | Get all students                |
| GET    | `/students/{id}`          | Get a student by ID             |
| PUT    | `/students/{id}`          | Update a student by ID          |
| DELETE | `/students/{id}`          | Delete a student by ID          |
| GET    | `/students/{id}/summary`  | Get AI-generated summary        |

## ⚙️ Run the App

```bash
go run main.go handlers.go ollama.go
