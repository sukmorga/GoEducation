package main

import (
	"fmt"
	"net/http"
	"text/template"
	
	"github.com/gorilla/mux"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id uint16
	Title string
	Anons string
	Full_text string
}

var posts = []Article{}
var postShow = []Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/golangedu")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		// Выборка 
	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	posts = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Data is not filled")
	} else {
		db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/golangedu")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		// Установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func showPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}


	vars := mux.Vars(r)

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/golangedu")
		if err != nil {
			panic(err)
		}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	postShow = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err)
		}

		postShow = append(postShow, post)
	}

	t.ExecuteTemplate(w, "show", postShow[0])
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create/", create).Methods("GET")
	rtr.HandleFunc("/saveArticle/", saveArticle).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}/", showPost).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}