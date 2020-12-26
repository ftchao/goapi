package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello, Iris"})
	})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen(":8080")
}
