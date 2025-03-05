package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// GetCepArgument extrai o cep dos argumentos da linha de comando, remove tudo que nao for numero,
// converte para inteiro e formata com 8 digitos, preenchendo com zeros a esquerda.
// Retorna o cep formatado e um erro se houver.
func GetCepArgument(args []string) (string, error) {

	if len(args) < 2 {
		return "", errors.New("informe o cep como argumento")
	}

	// remover tudo que nao for numero
	regex, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "", err
	}
	replaced := regex.ReplaceAllString(args[1], "")

	// converter o replaced para inteiro
	cepInt, err := strconv.Atoi(replaced)
	if err != nil {
		return "", err
	}

	// formatar o cep
	sprintf := fmt.Sprintf("%08d", cepInt)

	return sprintf, err
}
