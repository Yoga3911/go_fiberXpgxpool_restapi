package main

import (
	"coba/config"
	"coba/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var	db *pgxpool.Pool = config.DBConnection()

const varFilm = `SELECT * FROM film WHERE id = $1`
func getFilm(c *fiber.Ctx) error {
	id := c.Params("id")
	result:= db.QueryRow(c.Context(), varFilm, id)
	var film models.Film
	result.Scan(&film.ID, &film.Name, &film.Title, &film.CategoryID)
	return c.JSON(fiber.Map{
		"message": "success",
		"data": film,
	})
}

const varAllFilm = `SELECT * FROM film ORDER BY id`
func getAllFilm(c *fiber.Ctx) error {
	films, _ := db.Query(c.Context(), varAllFilm)
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

const varAddFilm = `INSERT INTO film (name, title, category_id) VALUES ($1, $2, $3)`
func addFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	db.Exec(c.Context(), varAddFilm, film.Name, film.Title, film.CategoryID)
	return c.JSON(fiber.Map{
		"status": "success",
	})
}

const varUpdateFilm = `UPDATE film SET name = $2, title = $3, category_id = $4 WHERE id = $1`
func updateFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	db.Exec(c.Context(), varUpdateFilm, film.ID, film.Name, film.Title, film.CategoryID)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

const varDeleteFilm = `DELETE FROM film WHERE id = $1`
func deleteFilm(c *fiber.Ctx) error {
	film := new(models.Film)
	c.BodyParser(film)
	db.Exec(c.Context(), varDeleteFilm, film.ID)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

const varAllCategory = `SELECT * FROM category ORDER BY id`
func getAllCategory(c *fiber.Ctx) error {
	categorys, _ := db.Query(c.Context(), varAllCategory)
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

const varAddCategory = `INSERT INTO category (category) VALUES ($1)`
func addCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	c.BodyParser(category)
	db.Exec(c.Context(), varAddCategory, category.Category)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func route(app *fiber.App) {
	film := app.Group("/api/film")
	film.Get("/:id", getFilm)
	film.Get("", getAllFilm)
	film.Post("", addFilm)
	film.Put("", updateFilm)
	film.Delete("", deleteFilm)

	category := app.Group("/api/category")
	category.Get("", getAllCategory)
	category.Post("", addCategory)
}

func main() {	
	defer db.Close()
	app := fiber.New()
	route(app)
	app.Listen(":8080")
}