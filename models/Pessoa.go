package models

// PessoaInterface define o contrato de uma "Pessoa"
type PessoaInterface interface {
	Nome() string
	Idade() int
	Endereco() string
}

// Pessoa é a struct base
type Pessoa struct {
	nome     string
	idade    int
	endereco string
}

func (p Pessoa) Nome() string {
	return p.nome
}

func (p Pessoa) Idade() int {
	return p.idade
}

func (p Pessoa) Endereco() string {
	return p.endereco
}
