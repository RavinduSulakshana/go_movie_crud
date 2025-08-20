package main
import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`git remote add origin git@github.com:RavinduSulakshana/go_movie_crud.git 

}

var movies []Movie

func main(){
	r := mux.NewRouter()
	movies =append(movies,Movie{id:"1",Isbn:"43453423",Title:"Movie One",Director:&Director})
	r.HandleFunc("/movies".getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Method("DELETE")

	fmt.Println("Server at start on port 8000")
	log.Fatal(http.ListenAndServe(":8080",r))

}

