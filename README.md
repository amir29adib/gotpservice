# 📲 Go OTP Service

A lightweight backend service built with Go (Gin + GORM) to handle **OTP-based authentication**, user management, and JWT-protected endpoints.

---

## 🔧 Features

- ✅ OTP-based login/registration via phone number
- 🔐 JWT authentication
- 📦 User management with pagination and search
- 🧱 PostgreSQL database (with Docker)
- 🐳 Docker & Docker Compose ready
- 📝 Swagger API docs available at `/docs/index.html`
- ⏳ Rate limiting: 3 OTPs per 10 minutes per phone
- 🧠 In-memory OTP cache with 1-minute expiration

---

## 📁 Project Structure

```
│
├── cmd/ # App entrypoint (main.go)
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── service/ # Business logic
│ ├── repository/ # DB and memory storage
│ ├── dto/ # Request/response DTOs
│ ├── model/ # DB models (User)
│ └── middleware/ # JWT auth
├── pkg/
│ └── utils/ # JWT handling
├── docs/ # Swagger generated docs
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/amir29adib/gotpservice.git
cd gotpservice
```

### 2. Setup .env

```
PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=otpuser
DB_PASSWORD=otppass
DB_NAME=otpdb
JWT_SECRET=super-secret-key
```

### 3. Run with Docker

```
docker-compose up --build
```

App will be available at http://localhost:8080
Swagger: http://localhost:8080/docs/index.html


---


## 🧪 API Examples

### Request OTP
```
POST /auth/request-otp
Content-Type: application/json

{
  "phone": "+989012345678"
}
```

### Verify OTP
```
POST /auth/verify-otp
Content-Type: application/json

{
  "phone": "+989012345678",
  "code": "123456"
}
```


### 🔐 Secured Endpoints

To access /users routes, you must provide JWT:
```
Authorization: Bearer <token>
```
