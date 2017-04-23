package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
  "strings"
	"github.com/gorilla/mux"
  "database/sql"
   _ "github.com/go-sql-driver/mysql"
)
// https://github.com/go-sql-driver/mysql
type Drink struct {
	img  string
	name string
	time string
	ing  string
	dir  string
}

var a []Drink
var m map[string]Drink
var tfm map[string]bool
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
			ing = ""
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

func GetMatches(w http.ResponseWriter, r * http.Request) {
    vars := mux.Vars(r)
    m := make(map[string]Drink)
    tfm := make(map[string]bool)
    var ingArray[] string
    ingList := vars["inglist"]
    ingList = strings.ToLower(ingList)
    ingArray = strings.Split(ingList, "+")
    db,err := sql.Open("mysql",
            "lindjac_lindjac:@/lindjac_drinkAPI") //Password and IP Missing
    if err != nil {
        fmt.Fprintln(w, "error lol")
    }
    defer db.Close()
    var qImg string
    var qName string
    var qTime string
    var qIng string
    var qDir string
    var sIng string
    for i := 0; i < len(ingArray); i++ {
        if ingArray[i] != "" {
            sIng = "'%" + ingArray[i] + "%'"
            q := "SELECT * FROM drinks WHERE ingredients LIKE " + sIng
            rows, err := db.Query(q)
            if err != nil {
                log.Fatal(err)
            }
            defer rows.Close()
            for rows.Next() {
                err = rows.Scan(&qImg, &qName, &qTime, &qIng, &qDir)
                if err != nil {
                    log.Fatal(err)
                }
                ok := tfm[qName]
                if ok == false {
                  tfm[qName] = true
                  m[qName] = Drink{qImg, qName, qTime, qIng, qDir}
                }
            }

            if err = rows.Err();
            err != nil {
                  log.Fatal(err)
            }
        }
    }
    //Start
    for k, v := range m {


    }
    //Stop
}
