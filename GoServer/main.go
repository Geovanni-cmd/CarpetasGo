package main
import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.Responsewrite, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return 
	}
	fmt.Fprintf(w, "POST exitoso")
	name := r.formvalue("name")
	address := r.formvalue("address")
	fmt.Fprintf(w,"Name : %s\n", name)
	fmt.Fprintf(w,"Address : %s\n", address)
}

func helloHandler(w http.Responsewrite, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "el metodo no esta", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))	
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Iniciando Servidor en el puerto 8080\n")
	if err := http.ListenAdnServe(":8080", nil); err !=nil{
		log.Fatal(err)
	}

}