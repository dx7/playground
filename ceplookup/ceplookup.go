package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (address Address) String() string {
	return fmt.Sprintf("CEP:         %s\n", address.Cep) +
		fmt.Sprintf("Logradouro:  %s\n", address.Logradouro) +
		fmt.Sprintf("Complemento: %s\n", address.Complemento) +
		fmt.Sprintf("Bairro:      %s\n", address.Bairro) +
		fmt.Sprintf("Localidade:  %s\n", address.Localidade) +
		fmt.Sprintf("UF:          %s\n", address.Uf)
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar CEP: %v\n", err)
			return
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
			return
		}
		req.Body.Close()

		var address Address
		err = json.Unmarshal(res, &address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar JSON: %v\n", err)
			return
		}

		fmt.Printf("%+v\n", address)
	}
}
