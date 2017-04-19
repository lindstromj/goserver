package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
  "strings"
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
	router.HandleFunc("/GET/{inglist}", GetMatches)

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
			ing += scanner.Text()[2:]
			ing += ","
		}
    if scanner.Text()[0] == '4' {
			dir = scanner.Text()[2:]
      s := Drink{img,name,time,ing,dir}
			a = append(a, s)
			name = ""
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

func GetMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  var ingArray []string
	ingList := vars["inglist"]
  ingArray = strings.Split(ingList,"+")
  for i:=0;i<len(ingArray);i++ {
		for j:=0;j<len(a);j++ {
			if strings.Contains(a[j].ing, ingArray[i]) {
				fmt.Fprintln(w, a[j].name)
			}
		}
	}
}
