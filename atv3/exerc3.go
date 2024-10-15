package main
import "fmt"

type ContaBancaria struct {
    saldo   float64
    titular string
}

func (c *ContaBancaria) depositar(valor float64) {
    c.saldo += valor
}

func (c *ContaBancaria) sacar(valor float64) {
    if valor <= c.saldo {
        c.saldo -= valor
    } else {
        fmt.Println("Saldo insuficiente.")
    }
}

func (c ContaBancaria) getSaldo() float64 {
    return c.saldo
}
