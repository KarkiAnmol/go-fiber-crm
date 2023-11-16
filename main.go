package main

import "github.com/gofiber/fiber"

func SetUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
func initDB() {

}
func main() {
	app := fiber.New()
	SetUpRoutes(app)
	app.Listen(3000)

}
