package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		//this will run when error is there
		fmt.Fprint(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request sucesfull\n")

	//we will take values into the variable
	name := r.FormValue("name") //this is the name of the form which will be stored into the variable name
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address =  %s\n", address)

}

//* is a pointer which is poining request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello")
}

func main() {
	// ":" this is a shortform goland operator which is used to define and use a variable.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)          //this is root route
	http.HandleFunc("/form", formHandler) //it will be used to handle the form route
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is running on port http://localhost:8080\n")
	//error can be there or it couldbe nil
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
