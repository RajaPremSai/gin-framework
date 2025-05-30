# Learning Gin Framework in Go

This repository documents my journey learning the Gin Web Framework in Go. It implements a basic REST API with authentication, demonstrating various Gin features and best practices.

## 🚀 What I Learned

- Setting up a Gin server with proper project structure
- Implementing REST APIs with Gin
- Working with middleware
- Database integration with GORM
- Authentication with password hashing
- Request validation and error handling

## 🔥 Key Features

- **Authentication System**

  - User registration with password hashing
  - Login with JWT tokens
  - Password security using bcrypt

- **Notes CRUD Operations**

  - Create, Read, Update, Delete notes
  - Status filtering
  - Protected routes using middleware

- **Database Integration**
  - PostgreSQL with GORM
  - Auto migrations
  - Relationship handling

## 🛠 Technical Stack

- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM library
- **PostgreSQL** - Database
- **Docker** - Containerization
- **bcrypt** - Password hashing

## 🎯 Key Concepts Explored

1. **Middleware Implementation**

   - Authentication checking
   - Request logging
   - Error handling

2. **Route Management**

   - Group routes
   - Parameter handling
   - Query string parsing

3. **Error Handling**

   - Proper HTTP status codes
   - Consistent error responses
   - Validation errors

4. **Clean Architecture**
   - Controller-Service pattern
   - Dependency injection
   - Separation of concerns

## 🚀 Getting Started

1. Clone the repository
2. Start PostgreSQL:

```bash
docker-compose up -d
```
3. Run the application
```bash
go run main.go
```

## 📝 API Endpoints

1. **Authentication**
 * POST /auth/register - Register new user
 * POST /auth/login - Login user

2. **Notes**
 * GET /notes - Get all notes
 * POST /notes - Create new note
 * PUT /notes - Update note
 * DELETE /notes/:id - Delete note
 * GET /notes/:id - Get note by ID

## 🎓 Lessons Learned
1. **Gin's Flexibility**

* Easy to set up and configure
* Extensive middleware support
* Great performance

2. **GORM Integration**
* Simplified database operations
* Auto migrations
* Query building

3. **Project Structure**
* Importance of clean architecture
* Separation of concerns
* Maintainable codebase

4. **Security Best Practices**
* Password hashing
* Protected routes
* Input validation


## 📄 License
**MIT License - feel free to use this project for learning purposes!**
