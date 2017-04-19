package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Drink struct {
	img  string
	name string
	time string
	ing  string
	dir  string
}

var a []Drink
var img string
var name string
var time string
var ing string
var dir string

func main() {
	ReadFile()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func ReadFile() {
	file, err := os.Open("unforgettables.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    if scanner.Text()[0] == '0' {
			img = scanner.Text()[2:]
		}
    if scanner.Text()[0] == '1' {
			name = scanner.Text()[2:]
		}
    if scanner.Text()[0] == '2' {
			time = scanner.Text()[2:]
		}
    if scanner.Text()[0] == '3' {
			ing = scanner.Text()[2:]
		}
    if scanner.Text()[0] == '4' {
			dir = scanner.Text()[2:]
      s := Drink{img,name,time,ing,dir}
			a = append(a, s)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  for i:=0;i<len(a);i++ {
  fmt.Fprintln(w, a[i].img)
  fmt.Fprintln(w, a[i].name)
  fmt.Fprintln(w, a[i].time)
  fmt.Fprintln(w, a[i].ing)
  fmt.Fprintln(w, a[i].dir)
 }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
