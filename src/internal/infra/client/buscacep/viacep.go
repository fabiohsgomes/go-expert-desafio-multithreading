package buscacep

import (
	"desafio-multithreading/internal/infra/client"
	"encoding/json"
	"fmt"
	"strings"
)

type ViaCepClient struct {
	uri           string
	contextConfig ContextConfig
}

type ViaCepClientResult struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        string `json:"erro"`
}

func NewViaCepClient(contextConfig ContextConfig) *ViaCepClient {
	return &ViaCepClient{
		uri:           "http://viacep.com.br/ws/",
		contextConfig: contextConfig,
	}
}

func (c *ViaCepClient) BuscaCep(cep string) (cepResult CepResult, err error) {
	ctx, cancel := c.contextConfig.Get()
	defer cancel()

	uri := fmt.Sprintf("%s%s/json/", c.uri, cep)

	response, err := client.Get(ctx, uri)
	if err != nil {
		return cepResult, err
	}

	var viaCepClientResult ViaCepClientResult
	err = json.Unmarshal(response.GetBody(), &viaCepClientResult)
	if err != nil {
		return cepResult, err
	}

	if len(viaCepClientResult.Erro) > 0 && viaCepClientResult.Erro == "true" {
		return cepResult, fmt.Errorf("ERR: Cep inválido")
	}

	cepResult.API = "ViaCep"
	cepResult.Cep = strings.ReplaceAll(viaCepClientResult.Cep, "-", "")
	cepResult.Logradouro = viaCepClientResult.Logradouro
	cepResult.Bairro = viaCepClientResult.Bairro
	cepResult.Localidade = viaCepClientResult.Localidade
	cepResult.Uf = viaCepClientResult.Uf

	return cepResult, err
}
