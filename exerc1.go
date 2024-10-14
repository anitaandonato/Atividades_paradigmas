package main
import "fmt"

type Carro struct {
    Marca  string
    Modelo string
    Ano    int
}

func (c Carro) detalhes() {
    fmt.Println(c.Marca, c.Modelo, c.Ano)
}

func main() {
    c1 := Carro{"Ford", "Fusion", 2020}
    c2 := Carro{"Honda", "Civic", 2021}
    c3 := Carro{"Toyota", "Corolla", 2019}

    c1.detalhes()
    c2.detalhes()
    c3.detalhes()
}
