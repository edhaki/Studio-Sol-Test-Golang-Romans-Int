package main

// importei algumas dependencias necessaria para aderir o projeto!
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"prova/romanos"
	"prova/router"

	"github.com/golang/gddo/httputil/header"
)

//foi criado duas struct para finalidades das fucoes, request para receber e response para devolver.
type Request struct {
	Text string
}

type Response struct {
	Number string
	Value  int
}

//func main foi criado router para gerar as rotas configuradas e subir a porta para api.
func main() {
	fmt.Println("api is port 8080")
	router := router.Gerar()
	router.HandleFunc("/search", getRomans)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRomans(w http.ResponseWriter, r *http.Request) {
	// Vai checkar se o content-type - header esta presente e check se e uma app json se nao ele retorna com erro.
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	// foi declarado uma request no struct
	var request Request

	// tentar o decode e receber no corpo do struct.
	// se ocorrer algum erro com client vai dar erro 400.
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// vai rodar o RomanNumerals e receber o Texto algorismo romano e devolver o valor de acordo com a logica aplicada.
	var number, value = romanos.RomanNumerals(request.Text)
	response := Response{number, value}

	// Enconde response para retornar como json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

// lembrando que pela logica vai ter que pegar o maior numero romano e devolver em json.
//curl -d '{"text":"AXXBLXBXX"}' -H "Content-Type: application/json" -X POST http://localhost:8080/search
//{"Number":"LX","Value":60}
// aqui por exemplo ele pegou o LX = 60 ou seja ele e maior que XX = 20. e mostrou o valor, e pegou a string maior do algorismo romano.

//PROGRAMA FOI DESENVOLVIDO POR EDERSON ALVES DA SILVA NO UBUNTU 22.04 LTS.
//FERRAMENTAS ULTILIZADAS VS CODE E GOLANG 1.18.1.
