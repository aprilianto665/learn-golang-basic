package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork: true,
	})

	app.Use("/api", func (ctx *fiber.Ctx) error {
		fmt.Println("I'm middleware before processing request")
		err := ctx.Next()
		fmt.Println("I'm middleware after processing request")
		return err
	})

	if fiber.IsChild(){
		fmt.Println("I'm child proccess")
	} else {
		fmt.Println("I'm parent proccess")
	}

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
