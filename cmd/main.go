package main

import (
	"fmt"
	"log"

	"github.com/AleksMedovnik/Teach-student__personal-account/controllers"
	"github.com/AleksMedovnik/Teach-student__personal-account/initializers"
	"github.com/AleksMedovnik/Teach-student__personal-account/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New(), cors.New())

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", middleware.DeserializeAdmin, controllers.SignUpUser)
		router.Post("/admin", controllers.SignUpAdmin)
		router.Post("/login", controllers.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
	})

	micro.Get("/users/profile", middleware.DeserializeUser, controllers.GetProfile)

	micro.Post("/groups", middleware.DeserializeAdmin, controllers.CreateGroup)
	micro.Get("/groups", middleware.DeserializeAdmin, controllers.GetGroups)
	micro.Get("/groups/:id", middleware.DeserializeAdmin, controllers.GetGroup)
	micro.Put("/groups/:id", middleware.DeserializeAdmin, controllers.UpdateGroup)
	micro.Delete("/groups/:id", middleware.DeserializeAdmin, controllers.DeleteGroup)

	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	log.Fatal(app.Listen(":8000"))
}
