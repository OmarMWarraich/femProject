# Workout CRUD API

A robust RESTful API for managing workout data with user authentication and authorization. Built with Go, PostgreSQL, and following clean architecture principles.

## 🚀 Features

- **Complete CRUD Operations** for workouts and workout entries
- **User Registration & Authentication** with JWT-like token system
- **Secure Password Hashing** using bcrypt
- **Multi-user Support** with workout ownership validation
- **Database Migrations** using Goose
- **Comprehensive Testing** with test database setup
- **Clean Architecture** with separated concerns (handlers, stores, middleware)
- **RESTful API Design** with proper HTTP status codes

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Database**: PostgreSQL 12.4
- **Router**: Chi v5
- **Password Hashing**: bcrypt
- **Database Driver**: pgx/v4
- **Migrations**: Goose v3
- **Testing**: Testify
- **Containerization**: Docker & Docker Compose

## 📋 Prerequisites

- Go 1.24 or higher
- Docker & Docker Compose
- PostgreSQL (if running locally)

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/OmarMWarraich/femProject.git
cd femProject
```

### 2. Start the Database

```bash
docker compose up -d
```

This will start:

- Main database on port `5432`
- Test database on port `5433`

### 3. Install Dependencies

```bash
go mod download
```

### 4. Run the Application

```bash
go run main.go
```

The API will be available at `http://localhost:8080`

## 📚 API Documentation

### Authentication

All workout endpoints require authentication via Bearer token.

#### Register User

```bash
POST /users
```

**Request Body:**

```json
{
  "username": "john_doe",
  "email": "john.doe@example.com",
  "password": "securepassword123",
  "bio": "Fitness enthusiast"
}
```

#### Login

```bash
POST /tokens/authentication
```

**Request Body:**

```json
{
  "username": "john_doe",
  "password": "securepassword123"
}
```

**Response:**

```json
{
  "auth_token": {
    "token": "5MNNG4VRC4VUCTJ7PHRH5PLIOML6OXHKJ3OHQNGQ7D3KANGWKMAQ",
    "expiry": "2024-01-02T15:04:05Z"
  }
}
```

### Workouts

#### Create Workout

```bash
POST /workouts
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "title": "Morning Strength Training",
  "description": "Full body strength training session",
  "duration_minutes": 60,
  "calories_burned": 500,
  "entries": [
    {
      "exercise_name": "Squats",
      "sets": 3,
      "reps": 12,
      "weight": 100.5,
      "notes": "Felt strong today",
      "order_index": 1
    }
  ]
}
```

#### Get Workout by ID

```bash
GET /workouts/{id}
Authorization: Bearer <token>
```

#### Update Workout

```bash
PUT /workouts/{id}
Authorization: Bearer <token>
```

#### Delete Workout

```bash
DELETE /workouts/{id}
Authorization: Bearer <token>
```

### Health Check

```bash
GET /health
```

## 🗄️ Database Schema

### Users Table

- `id` - Primary key
- `username` - Unique username
- `email` - Unique email address
- `password_hash` - Bcrypt hashed password
- `bio` - User biography
- `created_at`, `updated_at` - Timestamps

### Workouts Table

- `id` - Primary key
- `user_id` - Foreign key to users
- `title` - Workout title
- `description` - Workout description
- `duration_minutes` - Duration in minutes
- `calories_burned` - Estimated calories burned
- `created_at`, `updated_at` - Timestamps

### Workout Entries Table

- `id` - Primary key
- `workout_id` - Foreign key to workouts
- `exercise_name` - Name of exercise
- `sets` - Number of sets
- `reps` - Number of reps (optional)
- `duration_seconds` - Duration in seconds (optional)
- `weight` - Weight used (optional)
- `notes` - Additional notes
- `order_index` - Order within workout

### Tokens Table

- `hash` - Primary key (SHA256 hash)
- `user_id` - Foreign key to users
- `expiry` - Token expiration timestamp
- `scope` - Token scope (authentication)

## 🧪 Testing

Run the test suite:

```bash
# Start test database
docker compose up test_db -d

# Run tests
go test ./...

# Run specific test
go test ./internal/store -v
```

## 🏗️ Project Structure

```
femProject/
├── internal/
│   ├── api/           # HTTP handlers
│   ├── app/           # Application setup
│   ├── middleware/    # Authentication middleware
│   ├── routes/        # Route definitions
│   ├── store/         # Database layer
│   ├── tokens/        # Token generation/validation
│   └── utils/         # Utility functions
├── migrations/        # Database migrations
├── docker-compose.yml # Database containers
├── main.go           # Application entry point
└── README.md
```

## 🔧 Configuration

The application uses environment-based configuration. Default values:

- **Port**: 8080
- **Database Host**: localhost
- **Database Port**: 5432 (main), 5433 (test)
- **Database Name**: postgres
- **Database User**: postgres
- **Database Password**: postgres

## 🔒 Security Features

- **Password Hashing**: All passwords are hashed using bcrypt with cost factor 12
- **Token-based Authentication**: Secure token generation using crypto/rand
- **Authorization**: Users can only access their own workouts
- **SQL Injection Protection**: Parameterized queries throughout
- **Input Validation**: Comprehensive request validation

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Chi Router](https://github.com/go-chi/chi)
- Database migrations powered by [Goose](https://github.com/pressly/goose)
- Testing utilities from [Testify](https://github.com/stretchr/testify)

---

**Author**: Omar M Warraich  
**Linkedin**: [OmarMWarraich](https://www.linkedin.com/in/o-va/)
**Twitter**: [OmarMWarraich](https://twitter.com/omarwarraich1)
