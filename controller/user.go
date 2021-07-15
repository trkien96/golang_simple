package controller

import (
	"fmt"
	"go_rest_api/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	ListUser := model.ListStudent()

	if r.Method == "POST" {
		// logic part of log in
		fmt.Println("name:", strings.Title(r.FormValue("name")))
		fmt.Println("class:", strings.Title(r.FormValue("class")))
		fmt.Println("sex:", strings.Title(r.FormValue("sex")))
		item := &model.Student{
			Id:    4,
			Name:  strings.Title(r.FormValue("name")),
			Class: strings.Title(r.FormValue("class")),
			Sex:   strings.Title(r.FormValue("sex")),
		}
		ListUser = append(ListUser, item)
	}

	//Parse Json
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(ListUser)

	//User template external
	tmpl, _ := template.ParseFiles("./view/index.html")
	tmpl.Execute(w, ListUser)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Fprintf(w, "id is missing in parameters")
	}

	ListUser := model.ListStudent()
	index := Find(ListUser, id)
	if index == -1 {
		tmpl, _ := template.ParseFiles("./view/404.html")
		tmpl.Execute(w, nil)
	} else {
		// Use template internal
		tmpl := `
		{{define "user"}}
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Infor UserId: {{ .Id }}</title>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css">
		</head>
		<body>
		</head>
		<body>
			<div class="container mt-5">
				<h3><a href="/user/list">HOME</a> >> Infor UserId: {{ .Id }}</h3>
				<div class="row">
				<table class="table table-striped table-bordered">
				<tr>
					<th>Id</th>
					<th>Name</th>
					<th>Class</th>
					<th>Sex</th>
				</tr>
				<tr>
					<td><a href="./{{ .Id }}">{{ .Id }}</a></td>
					<td><a href="./{{ .Id }}">{{ .Name }}</a></td>
					<td><a href="./{{ .Id }}">{{ .Class }}</a></td>
					<td><a href="./{{ .Id }}">{{ .Sex }}</a></td>
				</tr>
				</table>
			</div>
			</div>
		</body>
		</html>
		{{end}}
		`

		t, _ := template.New("show").Parse(tmpl)
		t.ExecuteTemplate(w, "user", ListUser[index])
	}
}

func Find(listUser []*model.Student, id string) int {
	Id, _ := strconv.Atoi(id)
	for index, user := range listUser {
		if user.Id == Id {
			return index
		}
	}

	return -1
}

func CreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./view/create.html")
	tmpl.Execute(w, nil)
}

func ShowPage404(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./view/404.html")
	tmpl.Execute(w, nil)
}
