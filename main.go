package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)



func main(){
	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/",getAll)
	routes.HandleFunc("/cars",create)
	port := ":3000"

	fmt.Println("Servidor Iniciado na porta ", port)
	log.Fatal(http.ListenAndServe(port,routes))

}

type car struct {
	Model string
	Color string
	Manufactor string
}

var carros = []car{
	car{"Fusca","Preto","Volkswagen"},
	car{"Opala","Cinza","Chevrolet"},
	car{"Gol","Branco","Volkswagen"},
}


func getAll(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(carros)
}

func create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var c car

	body, err :=  ioutil.ReadAll(r.Body);

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		
		if err := json.NewEncoder(w).Encode(err); err != nil{
			panic(err)
		}
	}

	json.Unmarshal(body,&c)

	carros =  append(carros,c)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}