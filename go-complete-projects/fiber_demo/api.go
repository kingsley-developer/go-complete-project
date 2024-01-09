package serve

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func WebServer() {
	//REMEMBER POST, PUT AND DELETE ROUTE
	app := fiber.New(fiber.Config{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Fiber",
		AppName:           "Kingsley app demo",
		EnablePrintRoutes: true, //This does not work
	})
	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Use("/:id", func(c *fiber.Ctx) error {
		if c.Params("id") == "12" {
			fmt.Println("id is 12")
			fmt.Println("Secret key: 555ggggyhy")
		} else {
			fmt.Println("Not allowed")
		}
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("All middleware route")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		return c.SendString("<h1>Welcome to the homepage</h1>")
	})

	app.Get("/download", func(c *fiber.Ctx) error { //WORK ON THIS
		txt_path := "./test.txt" //NOT WORKING
		return c.Download(txt_path)
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Your Id " + c.Params("id"))
	})

	app.Get("/user/find/", func(c *fiber.Ctx) error {
		m := c.Queries()
		return c.SendString(m["name"])
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		payload := struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			fmt.Println(err)
			return err
		}
		return c.JSON(payload)
	})
	app.Post("/set_cookies", func(c *fiber.Ctx) error {

		payload := struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			fmt.Println(err)
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    payload.Password,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.SendString("Cookies set")
	})

	app.Delete("/delete_cookies", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name: "token",
			// Set expiry date to the past
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
			SameSite: "lax",
		})
		// ...
		return c.SendString("Cookies deleted")
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return c.SendString("I'm a update request!")
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return c.SendString("I'm a delete request!")
	})

	app.All("*", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Not found")
	})
	app.Listen(":7800")
}

func Static_Server() { //Not working
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Kingsley app demo",
	})
	app.Static("/", "./files")
	app.Listen(":7800")
}
