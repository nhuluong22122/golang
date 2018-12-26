package main

import (
	"log"
	"os"
	"text/template"
)

//1. Create a data structure to pass to a template which
//contains information about restaurant's menu
// including Breakfast, Lunch, and Dinner items

type restaurant struct {
	Name string
	Items []item
}

type item struct {
	Name, Description, MealType string
	Price float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []restaurant{
		restaurant {
			Name: "McDonald",
			Items: []item{
				item{
					"Oatmeal",
					"yum yum",
					"Breakfast",
					4.95,
				},
				item{
					"Cheerios",
					"American eating food traditional now",
					"Lunch",
					3.95,
				},
				item{
					"Juice Orange",
					"Delicious drinking in throat squeezed fresh",
					"Breakfast",
					2.95,
				},
			},
		},
		restaurant {
			Name: "Burger King",
			Items: []item{
				item{
					"Breakfast Burritos",
					"yum yum",
					"Breakfast",
					4.95,
				},
				item{
					"Sandwich",
					"sandwich",
					"Lunch",
					3.95,
				},
				item{
					"Carbonara",
					"Pasta",
					"Dinner",
					2.95,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
