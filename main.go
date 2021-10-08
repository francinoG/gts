package main

import (
	"gts/config"
	"gts/router"
)

func main() {
	config.DB()
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
