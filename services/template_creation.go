package services

import "fmt"

type Template struct {
	Path string
	Language string
	Packages []string
}

func CreateTemplate(template Template) {
	fmt.Println(template)
}