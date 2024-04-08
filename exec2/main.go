package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type funcionario struct {
	arquivo   int
	nome      string
	sobrenome string
	rg        int
	numeroTel int
	endereco  string
}

func main() {

	funcionarios := leArquivo("./customers.txt")

	// defer func() {
	// 	err := recover()

	// 	if err != nil {
	// 		fmt.Println("tratando erro de panic:", err)
	// 	}
	// }()

	novoId := geraIdArquivo()

	if novoId == 0 {
		panic("nao foi possivel gerar o IDArquivo")
	}

	novoFunc := funcionario{arquivo: novoId, nome: "Teste1", sobrenome: "Steste", rg: 40834, numeroTel: 1199999, endereco: "rua jajaja"}

	insereFunc(novoFunc, funcionarios)

	fmt.Println("Programa finalizado")

}

func leArquivo(arqDir string) []funcionario {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("tratando erro de panic:", err)
		}
	}()

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

		for i, campo := range values {
			switch {
			case i == 0:
				colaborador.arquivo, err = strconv.Atoi(campo)
				if err != nil {
					panic(err)
				}
			case i == 1:
				colaborador.nome = campo
			case i == 2:
				colaborador.sobrenome = campo
				if err != nil {
					panic(err)
				}
			case i == 3:
				colaborador.rg, err = strconv.Atoi(campo)
				if err != nil {
					panic(err)
				}
			case i == 4:
				colaborador.numeroTel, err = strconv.Atoi(campo)
				if err != nil {
					panic(err)
				}
			case i == 5:
				colaborador.endereco = campo

			default:
				panic("loop for leu mais campos do que deveria")

			}

		}

		colaboradores = append(colaboradores, colaborador)
	}
	return colaboradores

}

func geraIdArquivo() (id int) {

	return rand.Intn(9)
	// teste panic tarefa 1
	//return 0

}

func insereFunc(novoFunc funcionario, funcionarios []funcionario) []funcionario {
	for _, funcionario := range funcionarios {
		if funcionario.arquivo == novoFunc.arquivo {
			panic("Chave Arquivo já cadastrado no arquivo de funcionario")
			//return funcionarios, errors.New(fmt.Sprintf("Chave Arquivo já cadastrado no arquivo de funcionario %v", novoFunc.arquivo))
		}
	}
	funcionarios = append(funcionarios, novoFunc)
	return funcionarios
	//, nil
}
