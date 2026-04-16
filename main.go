package main

import (
 "database/sql"
 "os"
 "time"

 "github.com/gin-gonic/gin"
 "github.com/golang-jwt/jwt/v5"
 _ "github.com/lib/pq"
 "golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
 UserID string `json:"user_id"`
 Email string `json:"email"`
 jwt.RegisteredClaims
}

func main() {
 var err error
 db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
 if err != nil {
  panic(err)
 }

 r := gin.Default()

 // AUTH
 r.POST("/auth/register", register)
 r.POST("/auth/login", login)

 api := r.Group("/")
 api.Use(authMiddleware())
 {
  api.GET("/projects", getProjects)
  api.POST("/projects", createProject)
  api.GET("/projects/:id", getProjectByID)
  api.PATCH("/projects/:id", updateProject)
  api.DELETE("/projects/:id", deleteProject)

  api.GET("/projects/:id/tasks", getTasks)
  api.POST("/projects/:id/tasks", createTask)

  api.PATCH("/tasks/:id", updateTask)
  api.DELETE("/tasks/:id", deleteTask)
 }

 r.Run(":8080")
}
