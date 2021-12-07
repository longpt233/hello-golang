package driver
// khong quan tam ten file la gi, no chi quan tam package 

import "golang-web-module/internal/model"

var Todos []model.Todo

func init() {
	Todos = []model.Todo{
		{ID: 1, Name: "Monday", Content: "Learn Maths"},
		{ID: 2, Name: "Tuesday", Content: "Learn Literature"},
		{ID: 3, Name: "Wednesday", Content: "Learn Physics"},
		{ID: 4, Name: "Thursday", Content: "Learn Chemistry"},
		{ID: 5, Name: "Friday", Content: "Learn English"},
	}
}
