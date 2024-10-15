package main
import "fmt"

type Motor struct {
    Tipo string
}

type Carro struct {
    Marca string
    Motor Motor
}

func (c Carro) detalhes() {
    fmt.Println("Carro:", c.Marca, "com motor", c.Motor.Tipo)
}

func main() {
    motor := Motor{"V8"}
    carro := Carro{"Ford", motor}
    carro.detalhes()
}
