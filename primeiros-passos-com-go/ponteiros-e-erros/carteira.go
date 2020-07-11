package main

import (
	"errors"
	"fmt"
)

// ErroSaldoInsuficiente mensagem de saldo insuficiente
var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

// Bitcoin tipo inteiro
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Carteira estrutura da carteira
type Carteira struct {
	saldo Bitcoin
}

// Depositar recebe um valor inteiro de registra na carteira
func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

// Retirar recebe um valor inteiro que deseja ser retirado do saldo
func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= quantidade
	return nil
}

// Saldo retorna um valor inteiro que representa o saldo da carteira
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
