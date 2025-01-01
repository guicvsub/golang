package models

// Cliente é uma struct que implementa PessoaInterface
type Cliente struct {
	Pessoa
}

func NovoCliente(nome string, idade int, endereco string) PessoaInterface {
	return &Cliente{
		Pessoa: Pessoa{
			nome:     nome,
			idade:    idade,
			endereco: endereco,
		},
	}
}
