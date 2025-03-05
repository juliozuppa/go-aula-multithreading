package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juliozuppa/go-aula-multithreading/configs"
	"github.com/juliozuppa/go-aula-multithreading/internal/dto"
	"github.com/juliozuppa/go-aula-multithreading/internal/util"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	cep, err := util.GetCepArgument(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Carregando configurações")
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	log.Printf("Realizando consulta do cep %s\n", cep)

	brasilApiChain := make(chan dto.SearchCepOutput)
	viaCepChain := make(chan dto.SearchCepOutput)

	go SearchCepViaCep(ctx, cfg, cep, viaCepChain)
	go SearchCepBrasilApi(ctx, cfg, cep, brasilApiChain)

	select {

	case viaCepResult := <-viaCepChain:
		log.Println("CEP encontrado utilizando ViaCep")
		log.Printf("Cep: %s, %s, Bairro %s, %s, %s\n",
			viaCepResult.Cep, viaCepResult.Street, viaCepResult.Neighborhood, viaCepResult.City, viaCepResult.State)

	case brasilApiResult := <-brasilApiChain:
		log.Println("CEP encontrado utilizando BrasilApi")
		log.Printf("Cep: %s, %s, Bairro %s, %s, %s\n",
			brasilApiResult.Cep, brasilApiResult.Street, brasilApiResult.Neighborhood, brasilApiResult.City, brasilApiResult.State)

	case <-time.After(time.Duration(cfg.Application.SearchTimeout) * time.Millisecond):
		log.Printf("Timeout. Nenhuma resposta recebida em %d ms\n", cfg.Application.SearchTimeout)
	}
}

func SearchCepBrasilApi(ctx context.Context, cfg *configs.Config, cep string, chain chan dto.SearchCepOutput) {
	response, err := SearchCep(ctx, cfg.BrasilApi.URL, cep, cfg.Application.SearchTimeout)
	if err != nil {
		log.Println("BrasilApi: CEP nao encontrado")
		return
	}
	if response != nil {
		if response["cep"] == "" {
			log.Println("BrasilApi: CEP nao encontrado")
			return
		}
		chain <- dto.NewSearchCepOutput(
			response["cep"],
			response["state"],
			response["city"],
			response["neighborhood"],
			response["street"],
		)
	}
}

func SearchCepViaCep(ctx context.Context, cfg *configs.Config, cep string, chain chan dto.SearchCepOutput) {
	response, err := SearchCep(ctx, cfg.ViaCep.URL, cep, cfg.Application.SearchTimeout)
	if err != nil {
		log.Println("ViaCep: CEP nao encontrado")
		return
	}
	if response != nil {
		if response["cep"] == "" {
			log.Println("ViaCep: CEP nao encontrado")
			return
		}
		chain <- dto.NewSearchCepOutput(
			response["cep"],
			response["uf"],
			response["localidade"],
			response["bairro"],
			response["logradouro"],
		)
	}
}

func SearchCep(ctx context.Context, url, cep string, timeout int) (map[string]string, error) {
	request, err := SearchCepRequest(ctx, fmt.Sprintf(url, cep), timeout)
	if err != nil {
		return nil, err
	}
	response, err := ParseSearchCepResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// SearchCepRequest função que realiza a requisição HTTP para o serviço externo
func SearchCepRequest(ctx context.Context, url string, timeout int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer CloseResponseBody(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func ParseSearchCepResponse(dataJson []byte) (map[string]string, error) {
	var response map[string]string
	err := json.Unmarshal(dataJson, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// CloseResponseBody fecha o corpo da resposta HTTP, e se houver um erro ao fechar,
// registra o erro com log.Fatal.
func CloseResponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Println(err)
		return
	}
}
