package main
import "fmt"

type SaldoInsuficienteError struct {
    Message string
}

func (e *SaldoInsuficienteError) Error() string {
    return e.Message
}

type ContaBancaria struct {
    Saldo float64
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.Saldo {
        return &SaldoInsuficienteError{Message: "Saldo insuficiente para saque"}
    }
    c.Saldo -= valor
    return nil
}

func main() {
    conta := ContaBancaria{Saldo: 500}

    err := conta.Sacar(600)
    if err != nil {
        fmt.Println(err)
    }
}
