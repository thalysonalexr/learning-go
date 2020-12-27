package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type Pokemon struct {
	Name     string   `json:"name"`
	Power    int      `json:"power"`
	Avatar   string   `json:"avatar"`
	Specials []string `json:"specials"`
}

func listPokemonsJson() []Pokemon {
	jsonFile, _ := os.Open("pokemons.json")
	var pokemons []Pokemon
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &pokemons)
	defer jsonFile.Close()
	return pokemons
}

func HandlePokemons(w http.ResponseWriter, r *http.Request) {
	pokemons := listPokemonsJson()
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	response, _ := json.Marshal(pokemons)
	w.Write(response)
}

func HandleTemplatePokemons(w http.ResponseWriter, r *http.Request) {
	pokemons := listPokemonsJson()
	t := template.Must(template.ParseFiles("pokemons.html"))
	t.Execute(w, pokemons)
}

func createServer(port int) {
	http.HandleFunc("/pokemons/json", HandlePokemons)
	http.HandleFunc("/pokemons/html", HandleTemplatePokemons)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func main() {
	createServer(3333)
}
