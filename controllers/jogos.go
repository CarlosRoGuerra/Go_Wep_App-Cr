package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/CarlosRoGuerra/Go_Wep_App-Cr/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsjogos := models.BuscaTodosOsjogos()
<<<<<<< HEAD
	temp.ExecuteTemplate(w, "Index", todosOsjogos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
=======
 	temp.ExecuteTemplate(w, "Index", todosOsjogos)
 }

func New(w http.ResponseWriter, r *http.Request) {
 	temp.ExecuteTemplate(w, "New", nil)
 }
>>>>>>> a781e19743a00cb9668900c168520248772941d7

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		genero := r.FormValue("genero")
		preco := r.FormValue("preco")
		plataforma := r.FormValue("plataforma")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preco", err.Error())
		}

		models.CriarNovoJogo(nome, genero, precoConvertido, plataforma)
	}
	http.Redirect(w, r, "/", 301)
}
