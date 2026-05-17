# 🛒 SimpleOrder API

[![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat&logo=go)](https://golang.org)
[![Fiber Framework](https://img.shields.io/badge/Fiber-v2-1abc9c?style=flat)](https://gofiber.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![Swagger](https://img.shields.io/badge/Swagger-Docs-85EA2D?style=flat&logo=swagger)](http://swagger.io/)

SimpleOrder is a production-ready, highly scalable RESTful API built for an Order Management System. It is engineered using **Golang** and the **Fiber** web framework, strictly adhering to **Clean Architecture** principles to ensure modularity, testability, and separation of concerns.

## ✨ Features

- **Robust Authentication:** Secure JWT-based authentication and role-based access control (RBAC).
- **Transactional Orders:** Multi-item order placement with atomic database transactions to calculate totals and manage product stock reliably.
- **PDF Generation:** Automated creation of standardized PDF invoices for every order (`go-pdf/fpdf`).
- **Clean Architecture:** Distinct layers for Domain, Repository, Usecase, and Delivery (HTTP).
- **API Documentation:** Auto-generated interactive API documentation using Swagger UI.
- **Containerization:** fully dockerized with a multi-stage Dockerfile and docker-compose.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Fiber (Express inspired web framework)
- **Database:** PostgreSQL
- **ORM:** GORM
- **Security:** bcrypt, JWT (JSON Web Tokens)
- **Documentation:** swaggo/swag

## 📂 Project Structure (Clean Architecture)

```text
internal/
 ├── config/         # Environment variables & configurations
 ├── domain/         # Core business models & custom errors
 ├── repository/     # Data access layer (PostgreSQL / GORM)
 ├── usecase/        # Business logic & application rules
 └── delivery/       # HTTP handlers, routers, and middlewares
cmd/
 └── app/            # Application entrypoint & dependency injection
pkg/                 # Shared utilities (PDF, JWT, Responses)
```

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) 1.22 or higher
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker & Docker Compose](https://www.docker.com/) (Optional)

### 1. Run using Docker (Recommended)
The easiest way to start the project along with the database.

```bash
# 1. Clone the repository
git clone https://github.com/Yulianasetianingrum/simpleorder-api.git
cd simpleorder-api

# 2. Setup Environment Variables
cp .env.example .env

# 3. Start the containers
docker-compose up -d --build
```

### 2. Run Locally (Development)
If you prefer running the application natively.

```bash
# 1. Start PostgreSQL (Make sure to update .env with your DB credentials)
docker-compose up -d db

# 2. Install dependencies
go mod download

# 3. Run the application
make run
# or: go run cmd/app/main.go
```

## 📖 API Documentation

Once the application is running, you can explore the API endpoints using the interactive Swagger UI:

👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

*(Optional: Add a screenshot of your Swagger UI here)*

## 💡 Available Endpoints

### Auth
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Authenticate and get JWT

### Customers & Products
- `GET /api/v1/customers` - List customers (Supports pagination)
- `POST /api/v1/customers` - Add new customer
- `GET /api/v1/products` - List products
- `POST /api/v1/products` - Add new product (Admin Only)

### Orders
- `GET /api/v1/orders` - List all orders
- `POST /api/v1/orders` - Create a multi-item order
- `GET /api/v1/orders/:id/invoice` - Download Order Invoice (PDF)

### Dashboard
- `GET /api/v1/dashboard/stats` - Get summary statistics

## 📜 Postman Collection
A pre-configured Postman collection is included in the root directory (`simpleorder.postman_collection.json`). Simply import it into Postman to start testing the APIs immediately.

---
*Built with ❤️ to demonstrate Clean Architecture in Go.*
