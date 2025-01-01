package main

import (
	"barbearia/models"
	"fmt"
)

func main() {
	// Criar Cliente através do construtor
	cliente := models.NovoCliente("Carlos", 30, "Rua das Flores, 123")
	fmt.Println("Nome:", cliente.Nome())
	fmt.Println("Idade:", cliente.Idade())
	fmt.Println("Endereço:", cliente.Endereco())

	// Não é possível criar uma Pessoa diretamente (não exposta como pública)
	// pessoa := models.Pessoa{} // Erro
}
