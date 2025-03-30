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

- Go 1.24+
- Docker and Docker Compose
- Make (optional, for using Makefile commands)
- Node.js 16+ (for frontend development)

## Project Structure

```
.
├── backend/            # Go backend application
│   ├── cmd/            # Application entrypoints
│   ├── internal/       # Internal packages
│   ├── migrations/     # Database migrations
│   └── dockers/        # Dockerfile for backend
├── frontend/           # Vue.js frontend application
│   ├── src/            # Frontend source code
│   ├── public/         # Static assets
│   └── package.json    # Frontend dependencies
├── docker-compose.yaml # Docker compose configuration
└── Makefile            # Development commands
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

This project uses Docker containers for development and production environments. This ensures consistent behavior and eliminates the need to install dependencies locally.

### Running the Application

Start all services:
```bash
make init    # First time setup (builds images and runs migrations)
make up      # Start all services
make down    # Stop all services
```

Access points:
- Backend API: http://localhost:8080
- Frontend: http://localhost:3000

### Development Commands

Monitor logs:
```bash
make logs    # View container logs
```

Database operations:
```bash
make migrate-up    # Apply migrations
make migrate-down  # Rollback migrations
```

Cleanup:
```bash
make clean   # Remove temporary files
make rmdb    # Remove database volume
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
