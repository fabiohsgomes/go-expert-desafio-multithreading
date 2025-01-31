package buscacep

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type ContextConfig struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type CepResult struct {
	API        string `json:"api"`
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type BuscaCepProcessor interface {
	BuscaCep(cep string) (cepResult CepResult, err error)
}

func (r CepResult) String() string {
	return fmt.Sprintf(`
	API:%s,
	Cep:%s,
	Logradouro:%s,
	Bairro:%s,
	Localidade:%s,
	UF:%s
	`, fmt.Sprintf("%s %s", strings.Repeat(".", 15-len("API:")), r.API),
		fmt.Sprintf("%s %s", strings.Repeat(".", 15-len("Cep:")), r.Cep),
		fmt.Sprintf("%s %s", strings.Repeat(".", 15-len("Logradouro:")), r.Logradouro),
		fmt.Sprintf("%s %s", strings.Repeat(".", 15-len("Bairro:")), r.Bairro),
		fmt.Sprintf("%s %s", strings.Repeat(".", 15-len("Localidade:")), r.Localidade),
		fmt.Sprintf("%s %s", strings.Repeat(".", 15-len("Uf:")), r.Uf))
}

func NewContextConfigWithTimeOut(duraction time.Duration) ContextConfig {
	ctx, cancel := context.WithTimeout(context.Background(), duraction)
	return ContextConfig{ctx: ctx, cancel: cancel}
}

func (c *ContextConfig) Get() (ctx context.Context, cancel context.CancelFunc) {
	return c.ctx, c.cancel
}
