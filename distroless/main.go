/* golang
   very simple listen HTTP port 8080
   reply with message "ok"

   posibilitatea de a trimite argumente in URL
   de genul http://localhost:8080/?msg=Salutare

   ruta /hostname afiseaza hostname

   access log to stdout
*/


package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func logRequest(r *http.Request) {
	clientIP := r.RemoteAddr
	requestedURL := r.URL.String()
	log.Printf("Request from %s for %s\n", clientIP, requestedURL)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Logăm detaliile despre request
	logRequest(r)

	// Extragem valoarea parametrului "msg" din URL
	message := r.URL.Query().Get("msg")
	if message == "" {
		// Dacă parametrul "msg" nu este specificat, folosim un mesaj implicit
		message = "ok"
	}
	fmt.Fprintln(w, message)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	// Logăm detaliile despre request
	logRequest(r)

	// Executăm comanda "hostname" pentru a obține numele gazdei
	out, err := exec.Command("hostname").Output()
	if err != nil {
		http.Error(w, "Eroare la obținerea hostname-ului", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(out))
}

func main() {
	http.HandleFunc("/", messageHandler)
	http.HandleFunc("/hostname", hostnameHandler)

	fmt.Println("Serverul rulează pe http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Eroare la pornirea serverului:", err)
	}
}

