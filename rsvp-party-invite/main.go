package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type RSVP struct {
	Name, Email, Phone string
	WillAttend         bool
}

type FormData struct {
	*RSVP
	Errors []string
}

var responses = make([]*RSVP, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")

		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("hit the endpoint")
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, req *http.Request) {
	templates["list"].Execute(writer, responses)
}

func formHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		templates["form"].Execute(writer, FormData{
			RSVP: &RSVP{}, Errors: []string{},
		})
	} else if req.Method == http.MethodPost {
		req.ParseForm()
		responseData := RSVP{
			Name:       req.Form["name"][0],
			Email:      req.Form["email"][0],
			Phone:      req.Form["phone"][0],
			WillAttend: req.Form["willattend"][0] == "true",
		}

		errors := []string{}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your Email")
		}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your Name")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your Phone")
		}

		if len(errors) > 0 {
			templates["form"].Execute(writer, FormData{
				RSVP: &responseData, Errors: errors,
			})
		} else {
			responses = append(responses, &responseData)

			if responseData.WillAttend {
				templates["thanks"].Execute(writer, responseData.Name)
			} else {
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}

	}
}

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
