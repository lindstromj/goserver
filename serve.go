package main

import (
  "log"
  "net/http"
  "fmt"
  "github.com/gorilla/mux"
  "bufio"
  "os"
)

type Drink struct {
  name string
  ing string
}

  var a []Drink

func main() {
  ReadFile()
  s := &Drink{"Bob", "gaysex"}
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", s.Index)
  router.HandleFunc("/todos", TodoIndex)
  router.HandleFunc("/todos/{todoId}", TodoShow)

log.Fatal(http.ListenAndServe(":8080", router))
}

func ReadFile() {
file, err := os.Open("filein.txt")
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
       if scanner.Text()[0] == '0'{
         s := Drink{scanner.Text()[1:], ""}
         a = append(a,s)
       }
   }

   if err := scanner.Err(); err != nil {
       log.Fatal(err)
   }

}

func (s *Drink) Index(w http.ResponseWriter, r *http.Request) {

fmt.Fprintln(w, s.name)
fmt.Fprintln(w, s.ing)

}


func TodoIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, a[0].name)
  fmt.Fprintln(w, a[1].name)
}


func TodoShow(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   todoId := vars["todoId"]
   fmt.Fprintln(w, "Todo show:", todoId)
}
