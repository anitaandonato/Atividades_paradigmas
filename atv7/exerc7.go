package main
import "fmt"

type Professor struct {
    Nome string
}

type Escola struct {
    Nome       string
    Professores []Professor
}

func (e *Escola) adicionarProfessor(p Professor) {
    e.Professores = append(e.Professores, p)
}

func (e Escola) listarProfessores() {
    fmt.Println("Professores da escola", e.Nome + ":")
    for _, p := range e.Professores {
        fmt.Println(p.Nome)
    }
}

func main() {
    escola := Escola{Nome: "UNIPÊ"}
    p1 := Professor{Nome: "Professor João"}
    p2 := Professor{Nome: "Professor Maria"}

    escola.adicionarProfessor(p1)
    escola.adicionarProfessor(p2)

    escola.listarProfessores()
}
