# Observação

O projeto encontra-se na pasta **src**. Estando na pasta src, executar os comandos:
```
go mod tidy

#Executando a busca do cep 
go run cmd/cli/main.go [cep a ser pesquisado]
```

## Descrição do desafio

 Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta. (feito)

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou. (feito)

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido. (feito)