# TaskFlow

A three-tier task management system (React + Go + PostgreSQL) built to demonstrate deploying a Go application to AWS EC2 using a GitHub Actions CI/CD pipeline. On every push to the main branch, the app is automatically built and deployed to the EC2 server.

## What This Project Demonstrates

The focus of this repository is **automated deployment to AWS EC2 via GitHub Actions**. TaskFlow itself is a small, practical sample app — the real subject is the CI/CD pipeline that builds and ships it:

- On every `git push`, GitHub Actions builds the application tier and deploys it to an EC2 instance with no manual steps.
- The app is structured as a realistic three-tier system (presentation, application, data) so the deployment pipeline reflects how a real Go service would ship to production.

## What This App Does

TaskFlow is a simple task manager. It supports full CRUD on tasks:

- Create, view, update, and delete tasks
- Fields: title, description, status, priority, assigned to, due date

## Architecture

```
Presentation (React)  →  Application (Go + Gin)  →  Data (PostgreSQL / Neon)
     frontend/                  backend/                  cloud-hosted
```

- **Presentation tier** — React (Vite) single-page app, consumes the REST API.
- **Application tier** — Go + Gin HTTP API, layered as handler → service → repository.
- **Data tier** — PostgreSQL, hosted on Neon.

## Tech Stack

| Layer      | Technology                          |
|------------|--------------------------------------|
| Frontend   | React (Vite)                         |
| Backend    | Go, Gin, GORM                        |
| Database   | PostgreSQL (Neon cloud)              |
| Deployment | AWS EC2 + GitHub Actions (CI/CD)     |

## Running Locally

### Backend

```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

Requires a `.env` file (see [Environment Variables](#environment-variables)) with `DATABASE_URL` set to a reachable PostgreSQL instance.

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Requires a `.env` file with `VITE_API_URL` pointing at the backend.

## API Documentation

Base path: `/api`

| Method | Endpoint          | Description        |
|--------|-------------------|---------------------|
| GET    | `/api/health`     | Health check        |
| GET    | `/api/tasks`      | List all tasks      |
| POST   | `/api/tasks`      | Create a task        |
| GET    | `/api/tasks/:id`  | Get a single task   |
| PUT    | `/api/tasks/:id`  | Update a task        |
| DELETE | `/api/tasks/:id`  | Delete a task         |

## Environment Variables

| Variable       | Used By  | Description                                | Example                                              |
|----------------|----------|----------------------------------------------|-------------------------------------------------------|
| `DATABASE_URL` | backend  | PostgreSQL connection string                | `postgres://user:pass@host:5432/taskdb?sslmode=require` |
| `PORT`         | backend  | Port the API server listens on               | `8080`                                                 |
| `VITE_API_URL` | frontend | Base URL of the backend API                  | `http://localhost:8080`                                |

## Deployment (AWS EC2 + GitHub Actions CI/CD)

This app is deployed automatically to an AWS EC2 instance via a GitHub Actions CI/CD pipeline on every push.

Step-by-step CI/CD pipeline setup and deployment guide will be added here after deployment.
