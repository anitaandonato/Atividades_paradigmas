package main
import "fmt"

type Funcionario interface {
    CalcularSalario() float64
}

type FuncionarioHorista struct {
    Nome            string
    HorasTrabalhadas float64
    ValorHora       float64
}

func (f FuncionarioHorista) CalcularSalario() float64 {
    return f.HorasTrabalhadas * f.ValorHora
}

type FuncionarioAssalariado struct {
    Nome       string
    SalarioFixo float64
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
    return f.SalarioFixo
}

func main() {
    f1 := FuncionarioHorista{"João", 160, 20}
    f2 := FuncionarioAssalariado{"Maria", 3000}

    fmt.Println(f1.Nome, "recebe: R$", f1.CalcularSalario())
    fmt.Println(f2.Nome, "recebe: R$", f2.CalcularSalario())
}
