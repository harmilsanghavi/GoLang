package main

import (
	"fmt"

	"log"
	"net/http"
	"os"
)

type page struct {
	Tittle string
	Body   []byte
}

func (p *page) save() error {
	filename := p.Tittle + ".txt"
	return os.WriteFile("files/"+filename, p.Body, 0600)
}

func loadpage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("files/" + filename)
	if err != nil {
		return nil, err

	}
	return &page{Tittle: title, Body: body}, nil

}

func main() {
	// p1 := &page{Tittle: "Conclusion",Body: []byte("This Is a conclusion")}
	// p1.save()

	// p2,_ := loadpage("Conclusion")

	// fmt.Println(string(p2.Body))
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/view/", viewhandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func viewhandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p2, _ := loadpage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p2.Tittle, p2.Body)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadpage(title)
	if err != nil {
		p = &page{Tittle: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Tittle, p.Tittle, p.Body)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &page{Tittle: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
