package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

type Emp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Person struct {
	Id   []int
	Name []string
	City []string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", show)
	http.HandleFunc("/index", index)
	http.HandleFunc("/process", process)
	fs := http.FileServer(http.Dir("static/"))
	// fmt.Println(fs)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("mysql", "root:root@/goblog")
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected\n")
	defer db.Close()

	name1 := r.FormValue("name")
	city2 := r.FormValue("city")

	fmt.Println(name1)

	insert, err := db.Query("insert into employee(name,city) values(?,?)", (name1), (city2))

	if err != nil {
		panic(err)
	}

	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func show(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/goblog")
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected\n")
	defer db.Close()

	var emp Emp
	var person Person
	res, err := db.Query("select id,name,city from employee")
	for res.Next() {

		err = res.Scan(&emp.Id, &emp.Name, &emp.City)
		if err != nil {
			panic(err)
		}

		person.Name = append(person.Name, emp.Name)
		person.City = append(person.City, emp.City)
		person.Id = append(person.Id, emp.Id)

		fmt.Printf("%v\t%v\t%v\n", emp.Id, emp.Name, emp.City)
	}

	data := Person{
		Id:   person.Id,
		Name: person.Name,
		City: person.City,
	}
	tpl.ExecuteTemplate(w, "Processor.html", data)

	if err != nil {
		panic(err)
	}

}
