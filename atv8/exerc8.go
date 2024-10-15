package main
import "fmt"

type Empregado struct {
    Nome   string
    Cargo  string
    Salario float64
}

type Empresa struct {
    Nome      string
    Empregados []Empregado
}

func (e *Empresa) adicionarEmpregado(empregado Empregado) {
    e.Empregados = append(e.Empregados, empregado)
}

func (e Empresa) listarEmpregados() {
    fmt.Println("Empregados da empresa", e.Nome + ":")
    for _, empregado := range e.Empregados {
        fmt.Println(empregado.Nome, "-", empregado.Cargo, "-", empregado.Salario)
    }
}

func main() {
    empresa := Empresa{Nome: "Tech Corp"}
    e1 := Empregado{"João", "Desenvolvedor", 3000.0}
    e2 := Empregado{"Maria", "Designer", 3500.0}

    empresa.adicionarEmpregado(e1)
    empresa.adicionarEmpregado(e2)

    empresa.listarEmpregados()
}
