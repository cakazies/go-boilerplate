package main

import (
	"log"
	"os"

	"github.com/cakazies/go-boilerplate/template"
)

func main() {
	path := os.Getenv("HOME") + "/go/src/try"
	go template.Configs(path)
	go template.Utils(path)
	go template.Controllers(path)
	template.Test(path)
	template.Models(path)
	template.Middleware(path)
	template.CMD(path)
	template.Main(path)
	log.Println("Create Boilerplate Success")
}
