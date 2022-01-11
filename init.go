package main

import (
	"coba/config"
	"coba/controllers"
	"coba/models"
	"coba/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db *pgxpool.Pool = config.DBConnection()
	s services.AuthService = services.NewAuthService(db)
	c controllers.AuthController = controllers.NewAuthController(s)
)

func route(app *fiber.App) {
	user := app.Group("/api")
	user.Get("/user/:id", models.GetUser)
	user.Get("/users", models.GetAllUser)
	user.Post("/user", models.AddUser)
	user.Put("/user", models.UpdateUser)
	user.Delete("/user", models.DeleteUser)
	
	auth := user.Group("/auth")
	auth.Post("/register", c.Register)
	auth.Post("/login", c.Login)
}

func main() {
	defer db.Close()
	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
		Prefork: true,
	})
	route(app)
	app.Listen(":8080")
}
