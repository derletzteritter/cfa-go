package services

import "fmt"

type Template struct {
	Path     string
	Language string
}

func CreateTemplate(template Template) {
	fmt.Println(template)
	fmt.Println("Hello bro")
}
