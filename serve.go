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
   "encoding/json"
   "strconv"

)
// https://github.com/go-sql-driver/mysql
type Drink struct {
	Img  string `json:"img"`
	Name string `json:"name"`
	Time string `json:"time"`
	Ing  string `json:"ing"`
	Dir  string `json:"dir"`
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
  fmt.Fprintln(w, a[i].Img)
  fmt.Fprintln(w, a[i].Name)
  fmt.Fprintln(w, a[i].Time)
  fmt.Fprintln(w, a[i].Ing)
  fmt.Fprintln(w, a[i].Dir)
 }
}

func GetMatches(w http.ResponseWriter, r * http.Request) {
    vars := mux.Vars(r)
    offby := 0
    m := make(map[string]Drink)
    tfm := make(map[string]bool)
    var ingArray[] string
		var drinkIngList[] string
    ingList := vars["inglist"]
    ingList = strings.ToLower(ingList)
    ingArray = strings.Split(ingList, "+")
    db,err := sql.Open("mysql",
            "lindjac_lindjac:@tcp(:3306)/lindjac_drinkAPI") //Password and IP Missing
    if err != nil {
        log.Fatal(1)
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

    var finalArray []Drink
    canMake := 0
    for k, v := range m {
      if k == "" {
          log.Fatal(1)
      }
	    drinkIngList = strings.Split(strings.ToLower(v.Ing), ",")
			for i:=0;i < len(drinkIngList)-1;i++ {
				for j:=0;j < len(ingArray);j++ {
						if (strings.Contains(drinkIngList[i], ingArray[j])) {
								canMake++
								break
							}
						}
					}
					if canMake == (len(drinkIngList)-1)-offby{
							finalArray = append(finalArray, v)
					}
					canMake = 0
			}
      col := `:`
      comm := `,`
      bot := `}`
      if len(finalArray) > 0 {
      // Print JSON
      jsonsuccess(w, 1)
      for i:=0;i < len(finalArray);i++ {
          intcount, _ := json.Marshal(strconv.Itoa(i))
          fmt.Fprint(w, string(intcount))
          fmt.Fprintln(w, string(col))
          b, _ := json.MarshalIndent(finalArray[i], "", "  ")
          s := string(b)
          if i != len(finalArray)-1 {
            fmt.Fprintln(w, s + string(comm))
          }else {
            fmt.Fprintln(w, s)
          }
      }
      fmt.Fprintln(w, string(bot))
      } else {
        jsonsuccess(w, 0)
      }
    }

    func jsonsuccess(w http.ResponseWriter, i int) {
      top := `{`
      col := `:`
      bot := `}`
      comm := `,`
      fmt.Fprintln(w, string(top))
      bigsucc, _ := json.Marshal("Success")
      fmt.Fprint(w, string(bigsucc))
      fmt.Fprintln(w, string(col))
      fmt.Fprintln(w, string(top))
      lilsucc, _ := json.Marshal("success")
      fmt.Fprint(w, string(lilsucc))
      fmt.Fprint(w, string(col))
      yon, _ := json.Marshal(strconv.Itoa(i))
      fmt.Fprintln(w, string(yon))
      fmt.Fprintln(w, string(bot))
      if i == 0 {
        fmt.Fprint(w, string(bot))
      }
      fmt.Fprintln(w, string(comm))
    }
