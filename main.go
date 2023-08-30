package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bakarseck/api/internals/utils"
)

type Words struct {
	Id      int    `json:"id"`
	French  string `json:"français"`
	English string `json:"anglais"`
	Chinois string `json:"chinois"`
}

var words []Words

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

func main() {

	http.NewServeMux()

	utils.LoadEnv(".env")
	port := os.Getenv("PORT")
	log.Print(os.Getenv("AUTHOR"))

	file, err := os.Open("second.json")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer file.Close()

	json.NewDecoder(file).Decode(&words)

	http.HandleFunc("/users", getUserHandler)

	log.Printf("Serveur en cours d'exécution sur http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
