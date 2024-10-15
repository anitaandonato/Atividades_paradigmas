package main
import "fmt"

type Imprimivel interface {
    Imprimir()
}

type Relatorio struct{}

func (Relatorio) Imprimir() {
    fmt.Println("Imprimindo relatório...")
}

type Contrato struct{}

func (Contrato) Imprimir() {
    fmt.Println("Imprimindo contrato...")
}

func main() {
    var r Imprimivel = Relatorio{}
    var c Imprimivel = Contrato{}

    r.Imprimir()
    c.Imprimir()
}
