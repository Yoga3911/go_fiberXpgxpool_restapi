package models

import (
	"coba/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool = config.DBConnection()
const varFilmCategory = `SELECT film.id, film.name, film.title, category.category FROM film JOIN category ON (film.category_id=category.id) ORDER BY film.id`

func GetFilmCategory(c *fiber.Ctx) error {
	films, err := db.Query(c.Context(), varFilmCategory)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "failed",
		})
	}
	var film []*FilmCategory
	for films.Next() {
		var f FilmCategory
		films.Scan(&f.ID, &f.Name, &f.Title, &f.Category)
		film = append(film, &f)
	}
	return c.JSON(fiber.Map{
		"status": "succes",
		"data":   film,
	})
}

const varFilm = `SELECT * FROM film WHERE id = $1`

func GetFilm(c *fiber.Ctx) error {
	id := c.Params("id")
	result := db.QueryRow(c.Context(), varFilm, id)
	var film Film
	err := result.Scan(&film.ID, &film.Name, &film.Title, &film.CategoryID)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "failed",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   film,
	})
}

const varAllFilm = `SELECT * FROM film ORDER BY id`

func GetAllFilm(c *fiber.Ctx) error {
	films, err := db.Query(c.Context(), varAllFilm)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	var film []*Film
	for films.Next() {
		var f Film
		films.Scan(&f.ID, &f.Name, &f.Title, &f.CategoryID)
		film = append(film, &f)
	}
	return c.JSON(fiber.Map{
		"status": "succes",
		"data":   film,
	})
}

const varAddFilm = `INSERT INTO film (name, title, category_id) VALUES ($1, $2, $3)`

func AddFilm(c *fiber.Ctx) error {
	film := new(Film)
	err := c.BodyParser(film)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "failed",
		})
	}
	_, err2 := db.Exec(c.Context(), varAddFilm, film.Name, film.Title, film.CategoryID)
	if err2 != nil {
		return c.JSON(fiber.Map{
			"message": err2.Error(),
			"status":  "failed",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
	})
}

const varUpdateFilm = `UPDATE film SET name = $2, title = $3, category_id = $4 WHERE id = $1`

func UpdateFilm(c *fiber.Ctx) error {
	film := new(Film)
	err := c.BodyParser(film)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "failed",
		})
	}
	_, err2 := db.Exec(c.Context(), varUpdateFilm, film.ID, film.Name, film.Title, film.CategoryID)
	if err2 != nil {
		return c.JSON(fiber.Map{
			"message": err2.Error(),
			"status":  "failed",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

const varDeleteFilm = `DELETE FROM film WHERE id = $1`

func DeleteFilm(c *fiber.Ctx) error {
	film := new(Film)
	err := c.BodyParser(film)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "failed",
		})
	}
	_, err2 := db.Exec(c.Context(), varDeleteFilm, film.ID)
	if err2 != nil {
		return c.JSON(fiber.Map{
			"message": err2.Error(),
			"status":  "failed",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

const varAllCategory = `SELECT * FROM category ORDER BY id`

func GetAllCategory(c *fiber.Ctx) error {
	categorys, _ := db.Query(c.Context(), varAllCategory)
	var category []*Category
	for categorys.Next() {
		var ct Category
		categorys.Scan(&ct.ID, &ct.Category)
		category = append(category, &ct)
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    category,
	})
}

const varAddCategory = `INSERT INTO category (category) VALUES ($1)`

func AddCategory(c *fiber.Ctx) error {
	category := new(Category)
	err := c.BodyParser(category)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"status":  "failed",
		})
	}
	_, err2 := db.Exec(c.Context(), varAddCategory, category.Category)
	if err2 != nil {
		return c.JSON(fiber.Map{
			"message": err2.Error(),
			"status":  "failed",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}