package main

import (
	"context"
	"desafio-multithreading/internal/infra/client/buscacep"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
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

	contextConfiguration := buscacep.NewContextConfigWithTimeOut(time.Second)

	viaCepClient := buscacep.NewViaCepClient(contextConfiguration)
	brasilApiClient := buscacep.NewBrasilApiClient(contextConfiguration)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go runBuscaCepProcessor(wg, viaCepClient, cep)
	go runBuscaCepProcessor(wg, brasilApiClient, cep)

	wg.Wait()

	fmt.Println("Busca finalizada")
	os.Exit(0)
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

func runBuscaCepProcessor(wg *sync.WaitGroup, processor buscacep.BuscaCepProcessor, cep string) {
	defer wg.Done()

	result, err := processor.BuscaCep(cep)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("ERR: A requisição excedeu o tempo limite de 1 segundo. Detalhe: %s", err.Error())
			return
		}

		if errors.Is(err, context.Canceled) {
			return
		}

		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
