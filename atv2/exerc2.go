package main
import "fmt"

type Carro struct {
    Marca      string
    Modelo     string
    Ano        int
    Velocidade int
}

func (c *Carro) acelerar() {
    c.Velocidade += 10
}

func (c *Carro) frear() {
    c.Velocidade -= 10
}

func (c Carro) exibirVelocidade() {
    fmt.Println("Velocidade atual:", c.Velocidade)
}

func main() {
    c := Carro{"Ford", "Fusion", 2020, 0}
    c.acelerar()
    c.exibirVelocidade()
    c.frear()
    c.exibirVelocidade()
}
