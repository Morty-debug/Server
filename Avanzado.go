package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"bytes"
)

type Response struct {
	Estado        string
	Comparaciones []Matchs
}

type Matchs struct {
	Id            string
	Puntuacion    int
}

func main() {
	http.HandleFunc("/", writejson)
	http.HandleFunc("/leer", readjson)
	http.HandleFunc("/recibir", inputjson)
	http.HandleFunc("/enviar", sendjson)
	http.ListenAndServe(":8080", nil)
}

func writejson(w http.ResponseWriter, r *http.Request) {
	match0 := Matchs{"DGME-1802", 72}
	match1 := Matchs{"DGME-yo", 150}
	Matches := []Matchs{match0, match1}
	Estructura := Response {"Identificado",Matches}
	js, err := json.Marshal(Estructura)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"Estado\":\"No se logro crear JSON\",\"Comparaciones\":null}") 
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
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"Estado\":\"No se logro leer JSON de la URL "+url+"\",\"Comparaciones\":null}") 
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"Estado\":\"No se logro leer JSON de la URL "+url+"\",\"Comparaciones\":null}") 
		return
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

func inputjson(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"Estado\":\"Formato no compatible\",\"Comparaciones\":null}") 
		return
	}

	var data Response
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"Estado\":\"Estructura no compatible\",\"Comparaciones\":null}") 
		return
	}
	
	fmt.Printf("Estado: %s\n", data.Estado)
	fmt.Fprintf(w, "Estado: %s\n", data.Estado) 
	for i:=0; i<len(data.Comparaciones); i++{
		fmt.Printf("Id: %s Puntuacion: %d\n", data.Comparaciones[i].Id, data.Comparaciones[i].Puntuacion)
		fmt.Fprintf(w, "Id: %s Puntuacion: %d\n", data.Comparaciones[i].Id, data.Comparaciones[i].Puntuacion) 
	}
	return
}

func sendjson(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/recibir"
    
    var jsonStr = []byte(`{"Estado": "Identificado","Comparaciones": [{"Id": "DGME-1","Puntuacion": 100},{"Id": "DGME-2","Puntuacion": 5}]}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Fprintf(w, "Respuesta Estado: %s\n", resp.Status) 
    fmt.Fprintf(w, "Respuesta Encabezado: %s\n", resp.Header) 
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, "Respuesta Cuerpo: \n\n%s\n", string(body)) 
    return
}

