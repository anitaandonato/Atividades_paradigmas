abstract class Funcionario {
    protected String nome;
    public Funcionario(String nome) {
        this.nome = nome;
    }
    public abstract double calcularSalario();
}

class FuncionarioHorista extends Funcionario {
    private double horasTrabalhadas;
    private double valorHora;

    public FuncionarioHorista(String nome, double horasTrabalhadas, double valorHora) {
        super(nome);
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorHora = valorHora;
    }

    @Override
    public double calcularSalario() {
        return horasTrabalhadas * valorHora;
    }
}

class FuncionarioAssalariado extends Funcionario {
    private double salarioFixo;

    public FuncionarioAssalariado(String nome, double salarioFixo) {
        super(nome);
        this.salarioFixo = salarioFixo;
    }

    @Override
    public double calcularSalario() {
        return salarioFixo;
    }
}

public class Main {
    public static void main(String[] args) {
        Funcionario f1 = new FuncionarioHorista("João", 160, 20);
        Funcionario f2 = new FuncionarioAssalariado("Maria", 3000);

        System.out.println(f1.nome + " recebe: R$" + f1.calcularSalario());
        System.out.println(f2.nome + " recebe: R$" + f2.calcularSalario());
    }
}
