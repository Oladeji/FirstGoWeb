package main

import ("fmt"
        "net/http")

func handleFunc(w http.ResponseWriter , r *http.Request ){
fmt.Println("Hello World- lister")
fmt.Fprint(w, "<h1>Welcome to world wide be</h1>")
}
func main() {
	http.HandleFunc("/",handleFunc)
	http.ListenAndServe(":3000",nil)
	//fmt.Println("Hello World- And everybody")
}
