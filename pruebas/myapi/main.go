package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Estructura para un recurso
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Datos simulados
var articles []Article

// Obtener todos los artículos
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Obtener un solo artículo
func getArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Crear un nuevo artículo
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		log.Printf("Error decoding article: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Article received: %+v", article)
	articles = append(articles, article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

// Eliminar un artículo
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(articles)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	// Datos simulados iniciales
	articles = append(articles, Article{ID: "1", Title: "Artículo Uno", Content: "Contenido del artículo uno"})
	articles = append(articles, Article{ID: "2", Title: "Artículo Dos", Content: "Contenido del artículo dos"})

	// Crear enrutador
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":8000", router))
}
