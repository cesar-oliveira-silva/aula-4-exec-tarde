package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type funcionario struct {
	nome   string
	divida string
}

func main() {

	colaboradores := leArquivo("./customers.txt")

	for _, col := range colaboradores {
		fmt.Printf("Funcionario %s possui divida de %s\n", col.nome, col.divida)
	}

}

func leArquivo(arqDir string) []funcionario {

	var colaboradores []funcionario

	arquivo, err := os.Open(arqDir)
	if err != nil {
		panic(fmt.Sprintf("O arquivo %s indicado não foi encontrado ou está danificado. Erro: %v", arqDir, err))

	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	scanner.Scan() // pula cabecalho

	for scanner.Scan() {
		var colaborador = funcionario{}
		values := strings.Split(scanner.Text(), ";")
		colaborador.nome = values[0]
		colaborador.divida = values[1]

		colaboradores = append(colaboradores, colaborador)
	}
	return colaboradores

}
