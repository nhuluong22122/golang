package main

import (
	"log"
	"os"
	"text/template"
)


//1. Create a data structure to pass to a template which
//* contains information about California hotels including Name, Address, City, Zip, Region
//* region can be: Southern, Central, Northern
//* can hold an unlimited number of hotels


type hotel struct {
	Name string
	Address string
	City string
	Zip int
	Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	california_hotels := []hotel{
		hotel{
			"Test",
			"Test Address",
			"Test City",
			96100,
			"Southern",
		},
		hotel{
			"Test Hotel 2",
			"Test Address 2",
			"Test City 2 ",
			96122,
			"Northern",
		},
	}

	err := tpl.Execute(os.Stdout, california_hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
