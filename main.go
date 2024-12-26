package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter , r *http.Request){
	
	if r.URL.Path != "/form" {
		http.Error(w , "404 not found" ,http.StatusNotFound)
		return
	}
	
	if err:= r.ParseForm() ; err != nil {
		fmt.Fprintf(w , "parseform error %v" ,err)
		return
	}

	fmt.Fprint(w,"POST REQUEST SUCCESSFUL")
	name:= r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintf(w,"user name is %s\n",name)
	fmt.Fprintf(w,"password is %s\n" , password)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET" {

		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w , "hello")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/" ,fileServer )
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Println("starting server at port 8080")

	if err:=http.ListenAndServe(":8080", nil) ; err != nil {
		log.Fatal(err)
	}
}