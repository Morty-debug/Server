package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
)

type Response struct {
  Estado        string
  Comparaciones []Matchs
}

type Matchs struct {
  Id             string
  Puntuacion     int
}

func main() {
	http.HandleFunc("/", writejson)
	http.HandleFunc("/leer", readjson)
	http.ListenAndServe(":8080", nil)
}

func writejson(w http.ResponseWriter, r *http.Request) {
    match0 := Matchs{"DGME-1802", 72}
	match1 := Matchs{"DGME-yo", 150}
	Matches := []Matchs{match0, match1}
	Estructura := Response {"Identificado",Matches}
	js, err := json.Marshal(Estructura)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
    return
}

func readjson(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
	url := "http://localhost:8080/"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data Response
	json.Unmarshal(body, &data)
	fmt.Printf("Estado: %s\n", data.Estado)
	fmt.Fprintf(w, "Estado: %s\n", data.Estado) 
	for i:=0; i<len(data.Comparaciones); i++{
		fmt.Printf("Id: %s Puntuacion: %d\n", data.Comparaciones[i].Id, data.Comparaciones[i].Puntuacion)
		fmt.Fprintf(w, "Id: %s Puntuacion: %d\n", data.Comparaciones[i].Id, data.Comparaciones[i].Puntuacion) 
	}
    return
}
