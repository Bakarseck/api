package main

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	utils.LoadEnv(".env")
	port := os.Getenv("PORT")
	log.Print(os.Getenv("AUTHOR"))

	file, err := os.Open("second.json")
	utils.CheckError(err)

	defer file.Close()

	err = json.NewDecoder(file).Decode(&words)
	utils.CheckError(err)

	http.HandleFunc("/users", getUserHandler)
	p, _ := strconv.Atoi(port)
	p = utils.FindAvailablePort(p)
	log.Printf("Serveur en cours d'exécution sur http://localhost:%v\n", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", p), nil))
}
