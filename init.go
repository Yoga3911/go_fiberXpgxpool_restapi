package main

import (
	"coba/config"
	"coba/models"
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB = config.DBConnection()
	q *models.Queries = models.New(db)
	ctx = context.Background()
)

func getAllFilm(c *fiber.Ctx) error {
	film, _ := q.GetAllFilm(ctx)
	return c.JSON(fiber.Map{
		"message": "succes",
		"data": film,
	})
}

func addFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	q.CreateFilm(ctx, &models.CreateFilmParams{Name: film.Name,Title: film.Title, CategoryID: film.CategoryID})

	return c.JSON(fiber.Map{
		"status": "success",
		"data": film,
	})
}

func updateFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	q.UpdateFilm(ctx, &models.UpdateFilmParams{ID: film.ID, Name: film.Name, Title: film.Title, CategoryID: film.CategoryID})
	return c.JSON(fiber.Map{
		"message": "success",
		"data": film,
	})
}

func deleteFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	q.DeleteFilm(ctx, int64(film.ID))
	return c.JSON(fiber.Map{
		"message": "success",
		"data": film.ID,
	})
}

func getAllCategory(c *fiber.Ctx) error {
	category, _ := q.GetAllCategory(ctx)
	return c.JSON(fiber.Map{
		"message": "success",
		"data": category,
	})
}

func addCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	c.BodyParser(category)
	q.CreateCategory(ctx, category.Category)
	return c.JSON(fiber.Map{
		"message": "success",
		"data": category,
	})
}

func main() {	
	defer db.Close()

	app := fiber.New()
	app.Get("/film", getAllFilm)
	app.Post("/film", addFilm)
	app.Put("/film", updateFilm)
	app.Delete("/film", deleteFilm)
	app.Get("/category", getAllCategory)
	app.Post("/category", addCategory)
	app.Listen(":8080")
}