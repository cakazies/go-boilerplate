package template

import "os"

func Configs(path string) {
	pathConfig := (path + "/configs")

	os.MkdirAll(pathConfig, os.ModePerm)
	os.MkdirAll((pathConfig + "/migrations"), os.ModePerm)
	f, err := os.Create((pathConfig + "/configs.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package configs\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f, err = os.Create((pathConfig + "/DB.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package configs\n\nfunc init() {\n\n}\n\nfunc Connect() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}

func Utils(path string) {
	pathUtils := (path + "/utils")
	os.MkdirAll(pathUtils, os.ModePerm)

	f, err := os.Create((pathUtils + "/init.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package utils\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()

}

func Controllers(path string) {
	pathControllers := path + "/controllers"
	os.MkdirAll(pathControllers, os.ModePerm)
	os.MkdirAll((pathControllers + "/v1"), os.ModePerm)
	os.MkdirAll((pathControllers + "/v2"), os.ModePerm)
	f, err := os.Create((pathControllers + "/init.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package controllers\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}

func Test(path string) {
	pathControllers := path + "/test"
	os.MkdirAll(pathControllers, os.ModePerm)
	f, err := os.Create((pathControllers + "/init_test.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package test\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}

func Models(path string) {
	pathControllers := path + "/models"
	os.MkdirAll(pathControllers, os.ModePerm)
	f, err := os.Create((pathControllers + "/init.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package test\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}

func Middleware(path string) {
	pathControllers := path + "/middleware"
	os.MkdirAll(pathControllers, os.ModePerm)
	f, err := os.Create((pathControllers + "/init.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package test\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}

func CMD(path string) {
	pathControllers := path + "/cmd"
	os.MkdirAll(pathControllers, os.ModePerm)
	f, err := os.Create((pathControllers + "/init.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package test\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}
	f, err = os.Create((pathControllers + "/routes.go"))
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package test\n\nfunc init() {\n\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}

func Main(path string) {
	f, err := os.Create(path + "/main.go")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("package main\n\nfunc main() {\n    println(`Created Boilerplate succesfully`)\n}\n")
	if err != nil {
		panic(err)
	}

	f.Sync()
	f.Close()
}
