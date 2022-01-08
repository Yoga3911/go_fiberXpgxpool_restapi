package main

import (
	"coba/config"
	"coba/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool = config.DBConnection()

func route(app *fiber.App) {
	film := app.Group("/api/film")
	film.Get("/:id", models.GetFilm)
	film.Get("", models.GetAllFilm)
	film.Post("", models.AddFilm)
	film.Put("", models.UpdateFilm)
	film.Delete("", models.DeleteFilm)

	category := app.Group("/api/category")
	category.Get("", models.GetAllCategory)
	category.Post("", models.AddCategory)

	app.Get("/api/film_category", models.GetFilmCategory)
}

func main() {
	defer db.Close()
	app := fiber.New()
	route(app)
	app.Listen(":8080")
}
