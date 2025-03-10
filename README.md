# go-rest-api-mehmet-ogut

# 📚 Building a Simple REST API with Go and Gin-Gonic

## 📌 Project Description

This project aims to develop a basic **REST API** using the **Go programming language** and the **Gin-Gonic framework**.  
The API should interact with a **database (PostgreSQL or SQLite)** and utilize **containerization (Docker)** for deployment.

---

## 🔧 **Project Requirements**

### 🛠 Core Technologies

- **Programming Language:** Go
- **Framework:** Gin-Gonic

---

## 📌 **Functionality: Book Library Management System**

The API will include the following **entities**:

1. **Books**
   - `title`, `author`, `ISBN`, `publication_year`, `description`
2. **Authors**
   - `name`, `biography`, `birth_date`
3. **Reviews**
   - `rating`, `comment`, `date_posted`

### 🔗 **Relationships**

- **One Author → Many Books** (1:N)
- **One Book → Many Reviews** (1:N)
- **Books ↔ Authors** (Bidirectional Relationship)

---

## 🚀 **Required API Endpoints**

### 📚 **Books**

- `GET /api/v1/books` → List all books (pagination)
- `GET /api/v1/books/:id` → Get book details (with author & reviews)
- `POST /api/v1/books` → Create a new book
- `PUT /api/v1/books/:id` → Update a book
- `DELETE /api/v1/books/:id` → Delete a book

### ✍ **Authors**

- `GET /api/v1/authors` → List all authors (with their books)
- `GET /api/v1/authors/:id` → Get author details
- `POST /api/v1/authors` → Create a new author
- `PUT /api/v1/authors/:id` → Update an author
- `DELETE /api/v1/authors/:id` → Delete an author

### ⭐ **Reviews**

- `GET /api/v1/books/:id/reviews` → Get all reviews for a book
- `POST /api/v1/books/:id/reviews` → Add a review to a book
- `PUT /api/v1/reviews/:id` → Update a review
- `DELETE /api/v1/reviews/:id` → Delete a review

---

## 🗄 **Database Integration**

The system will use **PostgreSQL with GORM** as the ORM.  
Database schema will include **proper foreign key relationships and constraints**.

---

## 📦 **Development & Deployment**

### 🐳 **Containerization**

- The **REST API and SQL service** will run in a **containerized** environment using **Docker**.
- A **Dockerfile** and `docker-compose.yaml` will be provided.

### 📖 **Swagger Documentation**

- API endpoints will be documented using **Swagger**.

### 🛠 **Version Control**

- A new repository will be created under **MentalArts** organization on **GitHub**.
- Repository name format:

---

## 🏆 **Extra Features (Optional)**

### 🔒 **Authentication & Authorization**

- Implement **JWT-based authentication**
- **Protected Routes** (authentication required)
- **Role-based Access Control (Admin, User)**
- User Authentication Endpoints:
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/refresh-token`

### 🚀 **Additional Enhancements**

- **Rate Limiting**: Limit requests per user/IP
- **Caching**: Use **Redis** for caching frequently accessed data
- **Logging**: Structured logging (`info`, `error`, `debug`)
- **Input Validation**: Validate request payloads
- **Error Handling**: Consistent API error responses

---

## 🧪 **Testing**

- ✅ **Unit Tests**: Business logic tests
- ✅ **Integration Tests**: API endpoint tests
- ✅ **Load Testing**: Using `k6`

---

## 📊 **Monitoring & Observability**

- **Metrics**: Implement **Prometheus** metrics
- **Health Checks**: Add API health check endpoints
- **Tracing**: Use **Jaeger** for distributed tracing
- **Dashboard**: Set up **Grafana** for monitoring

---

## 📚 **Resources**

### **Go Resources**

- [Tour of Go](https://tour.golang.org/)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [Go Installation](https://golang.org/doc/install)

### **Docker Resources**

- [Docker Installation](https://docs.docker.com/get-docker/)
- [OrbStack](https://orbstack.dev/) (MacOS Alternative)

### **Framework & Tools**

- [Gin-Gonic](https://gin-gonic.com/) - REST API Framework
- [GORM](https://gorm.io/) - ORM for Golang
- [Docker Docs](https://docs.docker.com/) - Containerization
- [Swagger](https://swagger.io/) - API Documentation
- [Gin-Swagger](https://github.com/swaggo/gin-swagger)
- [Swag](https://github.com/swaggo/swag)

---
