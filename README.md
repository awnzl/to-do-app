# Todo Application

A full-stack Todo application built with Go backend and Vue.js frontend.

## Features

- Todo list management with CRUD operations
- PostgreSQL database with migrations
- RESTful API using chi router 
- Docker-based development environment
- Database transaction support
- Structured logging and error handling

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Make (optional, for using Makefile commands)
- Node.js 16+ (for frontend development)

## Project Structure

```
.
├── backend/               # Go backend application
│   ├── cmd/              # Application entrypoints
│   ├── internal/         # Internal packages
│   ├── migrations/       # Database migrations
│   └── dockers/         # Dockerfile for backend
├── frontend/             # Vue.js frontend application
│   ├── src/             # Frontend source code
│   ├── public/          # Static assets
│   └── package.json     # Frontend dependencies
├── docker-compose.yml    # Docker compose configuration
└── Makefile             # Development commands
```

## Quick Start

1. Clone the repository:
```bash
git clone <repository-url>
cd todoapp
```

2. Set up environment variables (copy .env.example to .env):
```bash
cp .env.example .env
```

3. Start the application:
```bash
docker-compose up -d
```

Backend API: http://localhost:8080  
Frontend: http://localhost:3000

## Development

### Backend

```bash
cd backend
make run      # Run backend locally
make migrate  # Run database migrations
```

### Frontend

```bash
cd frontend
npm install   # Install dependencies
npm run dev   # Start development server
```

### Running Everything

```bash
docker-compose up -d  # Start all services
```

### Database Migrations

Create a new migration:
```bash
make migrate-create name=your_migration_name
```

Apply migrations:
```bash
make migrate-up
```

Rollback migrations:
```bash
make migrate-down
```

### API Endpoints

Lists:
- `GET    /api/v1/lists`       - Get all lists
- `POST   /api/v1/lists`       - Create list
- `GET    /api/v1/lists/{id}`  - Get single list
- `PUT    /api/v1/lists/{id}`  - Update list
- `DELETE /api/v1/lists/{id}`  - Delete list

Todos:
- `GET    /api/v1/lists/{list_id}/todos`  - Get todos in list
- `POST   /api/v1/lists/{list_id}/todos`  - Create todo
- `GET    /api/v1/todos/{id}`             - Get single todo
- `PUT    /api/v1/todos/{id}`             - Update todo
- `DELETE /api/v1/todos/{id}`             - Delete todo

## About This Project

This is a learning project created to practice:
- Go backend development
- Database transactions
- Docker containerization
- API design
- Development workflows
