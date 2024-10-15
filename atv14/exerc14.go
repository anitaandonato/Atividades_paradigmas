package main
import (
    "fmt"
    "sync"
)

type Configuracao struct{}

var instancia *Configuracao
var once sync.Once

func GetInstancia() *Configuracao {
    once.Do(func() {
        instancia = &Configuracao{}
    })
    return instancia
}

func (c *Configuracao) ExibirConfiguracao() {
    fmt.Println("Exibindo a configuração...")
}

func main() {
    config1 := GetInstancia()
    config2 := GetInstancia()

    config1.ExibirConfiguracao()
    fmt.Println(config1 == config2)  // true, porque ambas são a mesma instância
}
