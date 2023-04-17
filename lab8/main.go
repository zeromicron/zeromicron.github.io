package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type PageData struct {
    Title   string
    Name    string
}

type FormData struct {
    Name     string
    Id       string
    Email    string
    PhoneNum string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if r.Method != http.MethodPost {

            data := PageData{
                Title: "lab8",
                Name:  "",
            }


            templatePath := "static/index.html"

            tmpl, err := template.ParseFiles(templatePath)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            err = tmpl.Execute(w, data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            return
        }

        name := r.FormValue("name")
        id := r.FormValue("id")
        email := r.FormValue("email")
        phoneNum := r.FormValue("phoneNum")

        formData := FormData{
            Name:     name,
            Id:     id,
            Email:    email,
            PhoneNum: phoneNum,
        }


        templatePath := "static/secondpage.html"


        tmpl, err := template.ParseFiles(templatePath)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        err = tmpl.Execute(w, formData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    fmt.Println("Starting server on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
