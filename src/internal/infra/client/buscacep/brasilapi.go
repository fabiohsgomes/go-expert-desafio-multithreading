package buscacep

import (
	"desafio-multithreading/internal/infra/client"
	"encoding/json"
	"fmt"
)

type BrasilApiClient struct {
	uri           string
	contextConfig ContextConfig
}

type BrasilApiClientResult struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func NewBrasilApiClient(contextConfig ContextConfig) *BrasilApiClient {
	return &BrasilApiClient{
		uri:           "https://brasilapi.com.br/api/cep/v1/",
		contextConfig: contextConfig,
	}
}

func (c *BrasilApiClient) BuscaCep(cep string) (cepResult CepResult, err error) {
	ctx, cancel := c.contextConfig.Get()
	defer cancel()

	uri := fmt.Sprintf("%s%s", c.uri, cep)

	response, err := client.Get(ctx, uri)
	if err != nil {
		return cepResult, err
	}

	var brasilApiClientResult BrasilApiClientResult
	err = json.Unmarshal(response.GetBody(), &brasilApiClientResult)
	if err != nil {
		return cepResult, err
	}

	cepResult.API = "BrasilApi"
	cepResult.Cep = brasilApiClientResult.Cep
	cepResult.Logradouro = brasilApiClientResult.Street
	cepResult.Bairro = brasilApiClientResult.Neighborhood
	cepResult.Localidade = brasilApiClientResult.City
	cepResult.Uf = brasilApiClientResult.State

	return cepResult, err
}
