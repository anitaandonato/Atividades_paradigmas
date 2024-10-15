package main
import "fmt"

type Produto struct {
    Nome  string
    Preco float64
}

func (p Produto) SomarPreco(outro Produto) float64 {
    return p.Preco + outro.Preco
}

func main() {
    p1 := Produto{"Produto A", 100.0}
    p2 := Produto{"Produto B", 200.0}

    fmt.Println("Soma dos preços: R$", p1.SomarPreco(p2))
}
