// For GCC.exe compiler you'll need to download TDM-GCC-64 and make sure the path to GCC.exe is found
// test GCC.exe by gcc.exe --version

package main

import (

	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

type Page struct{
	Name string
	DBStatus bool
}

func main(){

	templates:= template.Must(template.ParseFiles("template/index.html"))

	db,_:=sql.Open("sqlite3","dev.db")

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		p:= Page{Name: "Gopher",}

		if name:=r.FormValue("name"); name != ""{
			p.Name = name
		}

		p.DBStatus = db.Ping()  == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		db.Close()
	})
	fmt.Println(http.ListenAndServe(":8080",nil))
}