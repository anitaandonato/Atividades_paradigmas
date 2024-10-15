package main
import "fmt"

type Animal interface {
    som()
}

type Cachorro struct{}

func (Cachorro) som() {
    fmt.Println("O cachorro late.")
}

type Gato struct{}

func (Gato) som() {
    fmt.Println("O gato mia.")
}

func main() {
    var a1 Animal = Cachorro{}
    var a2 Animal = Gato{}

    a1.som()
    a2.som()
}
