package main

import (
	"context"
	"desafio-multithreading/internal/infra/client/buscacep"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	argumentos := os.Args
	cep := argumentos[1]

	err := validaCep(cep)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	viaCepResult := make(chan buscacep.CepResult)
	brasilApiResult := make(chan buscacep.CepResult)

	contextConfiguration := buscacep.NewContextConfigWithTimeOut(time.Second)

	viaCepClient := buscacep.NewViaCepClient(contextConfiguration)
	brasilApiClient := buscacep.NewBrasilApiClient(contextConfiguration)

	go runBuscaCepProcessor(viaCepResult, viaCepClient, cep)
	go runBuscaCepProcessor(brasilApiResult, brasilApiClient, cep)

	select {
	case result := <-viaCepResult:
		printResult(&result)
	case result := <-brasilApiResult:
		printResult(&result)
	case <-time.After(time.Second):
		log.Println("Timeout: A requisição excedeu o tempo limite de 1 segundo.")
	}

	os.Exit(0)
}

func runBuscaCepProcessor(channel chan buscacep.CepResult, processor buscacep.BuscaCepProcessor, cep string) {
	result, err := processor.BuscaCep(cep)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = fmt.Errorf("ERR: A requisição excedeu o tempo limite de 1 segundo. Detalhe: %s", err.Error())
		}

		if errors.Is(err, context.Canceled) {
			close(channel)
			return
		} else {
			fmt.Println(err.Error())
			close(channel)
			return
		}
	}

	channel <- result
}

func validaCep(cep string) (err error) {
	if len(cep) < 8 || len(cep) > 8 {
		return fmt.Errorf("ERR: O cep deve conter somente 8 digitos numéricos")
	}

	if _, err := strconv.Atoi(cep); err != nil {
		return fmt.Errorf("ERR: O cep informado não é válido")
	}

	return err
}

func printResult(result *buscacep.CepResult) {
	if (buscacep.CepResult{}) != *result {
		fmt.Println(result)
		fmt.Println("Busca finalizada")
	}
}
