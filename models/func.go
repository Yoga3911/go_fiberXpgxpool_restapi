package models

import (
	"coba/config"
	"coba/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool = config.DBConnection()

const varUser = `SELECT * FROM users WHERE id = $1`

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	result := db.QueryRow(c.Context(), varUser, id)
	var user User
	err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GenderID, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		return c.JSON(helper.BuildResponse(err.Error(), false))
	}
	return c.JSON(helper.BuildResponse(user, true))
}

const varAllUser = `SELECT * FROM users ORDER BY id`

func GetAllUser(c *fiber.Ctx) error {
	users, err := db.Query(c.Context(), varAllUser)
	if err != nil {
		return c.JSON(helper.BuildResponse(err.Error(), false))
	}
	var user []*User
	for users.Next() {
		var u User
		users.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.GenderID, &u.CreateAt, &u.UpdateAt)
		user = append(user, &u)
	}
	return c.JSON(helper.BuildResponse(user, true))
}

const varAddUser = `INSERT INTO users (name,email,password,gender_id,create_at,update_at) 
					VALUES ($1, $2, $3, $4, now(), now())`

func AddUser(c *fiber.Ctx) error {
	user := new(User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(err.Error(), false))
	}
	errors := ValidateFilm(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(errors, false))
	}
	_, err2 := db.Exec(c.Context(), varAddUser, user.Name, user.Email, user.Password, user.GenderID)
	if err2 != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false))
	}
	return c.JSON(helper.BuildResponse(nil, true))
}

const varUpdateUser = `UPDATE users SET name = $2, email = $3, password = $4, gender_id = $5, update_at = now() WHERE id = $1`

func UpdateUser(c *fiber.Ctx) error {
	user := new(User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(err.Error(), false))
	}
	_, err2 := db.Exec(c.Context(), varUpdateUser, user.ID, user.Name, user.Email, user.Password, user.GenderID)
	if err2 != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false))
	}
	return c.JSON(helper.BuildResponse(nil, false))
}

const varDeleteUser = `DELETE FROM users WHERE id = $1`

func DeleteUser(c *fiber.Ctx) error {
	user := new(User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BuildResponse(err.Error(), false))
	}
	_, err2 := db.Exec(c.Context(), varDeleteUser, user.ID)
	if err2 != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false))
	}
	return c.JSON(helper.BuildResponse(nil, false))
}
