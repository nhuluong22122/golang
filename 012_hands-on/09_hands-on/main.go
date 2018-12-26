package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)
//Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
//Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.
type Stock struct {
	Date time.Time
	Open,High,Low,Close,Volume,AdjClose float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	csvFile, _ := os.Open("table.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var stocks []Stock
	var first = true
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if first {
			first = false
		} else {
			date, _ := time.Parse("2006-01-02", line[0])
			open, _ := strconv.ParseFloat(line[1], 64)
			high, _ := strconv.ParseFloat(line[2], 64)
			low, _ := strconv.ParseFloat(line[3], 64)
			close, _ := strconv.ParseFloat(line[4], 64)
			volume, _ := strconv.ParseFloat(line[5], 64)
			adjclose, _ := strconv.ParseFloat(line[6], 64)

			stocks = append(stocks, Stock{
				date,
				open,
				high,
				low,
				close,
				volume,
				adjclose,
			})
		}
	}
	err := tpl.Execute(os.Stdout, stocks)
	if err != nil {
		log.Fatalln(err)
	}
}
