package main

import (
	"coba/config"
	"coba/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db *pgxpool.Pool = config.DBConnection()
	ctx = context.Background()
)

func getFilm(c *fiber.Ctx) error {
	id := c.Params("id")
	result:= db.QueryRow(ctx, "SELECT * FROM film WHERE id = $1", id)
	var film models.Film
	result.Scan(&film.ID, &film.Name, &film.Title, &film.CategoryID)

	return c.JSON(fiber.Map{
		"message": "success",
		"data": film,
	})
}

func getAllFilm(c *fiber.Ctx) error {
	films, _ := db.Query(ctx, "SELECT * FROM film ORDER BY id")
	var film []*models.Film
	for films.Next() {
		var f models.Film
		films.Scan(&f.ID, &f.Name, &f.Title, &f.CategoryID)
		film = append(film, &f)
	}
	return c.JSON(fiber.Map{
		"message": "succes",
		"data": film,
	})
}

func addFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	db.QueryRow(ctx, "INSERT INTO film (name, title, category_id) VALUES ($1, $2, $3)", &film.Name, &film.Title, &film.CategoryID)

	return c.JSON(fiber.Map{
		"status": "success",
		"data": film,
	})
}

func updateFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	db.QueryRow(ctx, "UPDATE film SET name = $2, title = $3, category_id = $4 WHERE id = $1", &film.ID, &film.Name, &film.Title, &film.CategoryID)
	return c.JSON(fiber.Map{
		"message": "success",
		"data": film,
	})
}

func deleteFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	db.QueryRow(ctx, "DELETE FROM film WHERE id = $1", &film.ID)
	return c.JSON(fiber.Map{
		"message": "success",
		"data": film.ID,
	})
}

func getAllCategory(c *fiber.Ctx) error {
	categorys, _ := db.Query(ctx, "SELECT * FROM category ORDER BY id")
	var category []*models.Category
	for categorys.Next() {
		var ct models.Category
		categorys.Scan(&ct.ID, &ct.Category)
		category = append(category, &ct)
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data": category,
	})
}

func addCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	c.BodyParser(category)
	db.QueryRow(ctx, "INSERT INTO category (category) VALUES ($1)", category.Category)
	return c.JSON(fiber.Map{
		"message": "success",
		"data": category,
	})
}

func main() {	
	defer db.Close()

	app := fiber.New()
	app.Get("/film/:id", getFilm)
	app.Get("/film", getAllFilm)
	app.Post("/film", addFilm)
	app.Put("/film", updateFilm)
	app.Delete("/film", deleteFilm)
	app.Get("/category", getAllCategory)
	app.Post("/category", addCategory)
	app.Listen(":8080")
}