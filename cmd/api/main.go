package main

import (
	"log"

	"github.com/Vadim-12/tszh-backend/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
