package main

import (
	"log"
	"os"

	"github.com/cakazies/go-boilerplate/template"
)

func main() {
	project := os.Getenv("PROJECT")
	if project == "" {
		println("ENV null, you can set this env with `PROJECT=your_project` go run main.go")
		return
	}

	path := os.Getenv("HOME") + "/go/src/" + project
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
