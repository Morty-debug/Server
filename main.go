package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
    "path/filepath"
    "github.com/skratchdot/open-golang/open"
)

func subidor(w http.ResponseWriter, r *http.Request) {
	//mostrar en consola server
	fmt.Printf("Invocacion de %q/upload \n", r.Host) 

	//mostrar en cliente
	fmt.Fprintf(w, "Metodo: %s URL: %s Protocolo: %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Servidor = %q\n", r.Host)
	fmt.Fprintf(w, "Cliente = %q\n", r.RemoteAddr)

	// maixmo 1MB de subida
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
	err := r.ParseMultipartForm(1024*1024)
	if err != nil {
		fmt.Fprintf(w, "El archivo es demasiado grande, excede 1MB o no contiene Multiples parte\n")
		fmt.Println("El archivo es demasiado grande, excede 1MB o no contiene Multiples parte")
		return
	}
	
	// input text con nombre Nombre
	variable := r.FormValue("Nombre")
	fmt.Fprintf(w, "Variable Nombre: %s\n",variable)
	fmt.Println("Variable Nombre:",variable)
		
	// inputfile de nombre Archivo
	file, handler, err := r.FormFile("Archivo")
	if err != nil {
		fmt.Fprintf(w, "Error al obtener el archivo el servidor\n")
		fmt.Println("Error al obtener el archivo del cliente")
		return
	}
	defer file.Close()
	fmt.Printf("Archivo: %+v\n", handler.Filename)
	fmt.Printf("Tamaño: %+v\n", handler.Size)
	fmt.Printf("MIME Encabezado: %+v\n", handler.Header)

	// generamos un nombre de archivo ramdom
	tempFile, err := ioutil.TempFile("subidos", "*"+filepath.Ext(handler.Filename))
	if err != nil {
		fmt.Fprintf(w, "No se logro ramdomizar el nombre\n")
		fmt.Println("No se logro ramdomizar el nombre")
		return
	}
	defer tempFile.Close()

	// obtenemos el archivo desde archivos temporales
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "No se logro leer el archivo\n")
		fmt.Println("No se logro leer el archivo")
		return
	}

	// escribimos el archivo en el servidor
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Subida Completada '%s'\n",tempFile.Name())

	//mostrar en consola server
	fmt.Println("Subida Completada (",tempFile.Name(),")")
}

type Html struct {
	UrlServer string 
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Invocacion de %q \n", r.Host)
	
	tmplt := template.New("index.html") //creamos una nueva plantilla
	tmplt, _ = tmplt.ParseFiles("index.html") //seteamos contenido para reemplazar variables en plantilla
	conte := Html{UrlServer: r.Host} //reemplazamos variables de la plantilla por contenido
	
	tmplt.Execute(w, conte) //mostramos el contenido
}

func main() {
	fmt.Println("localhost:8080")
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", subidor)
	open.Run("http://localhost:8080")
	//http.Handle("/", http.FileServer(http.Dir("./paginasweb")))
	http.ListenAndServe(":8080", nil)
}
