package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
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

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

func BuscaCEP(cepParam string) (*ViaCEP, error) {
	response, error := http.Get("https://viacep.com.br/ws" + cepParam + "/json/")
	if error != nil {
		return nil, error
	}

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var cep ViaCEP
	error = json.Unmarshal(body, &cep)
	if error != nil {
		return nil, error
	}

	return &cep, nil
}
