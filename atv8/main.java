import java.util.ArrayList;
import java.util.List;

class Empregado {
    String nome;
    String cargo;
    double salario;

    Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }
}

class Empresa {
    String nome;
    List<Empregado> empregados;

    Empresa(String nome) {
        this.nome = nome;
        this.empregados = new ArrayList<>();
    }

    void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }

    void listarEmpregados() {
        System.out.println("Empregados da empresa " + nome + ":");
        for (Empregado e : empregados) {
            System.out.println(e.nome + " - " + e.cargo + " - R$" + e.salario);
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Empresa empresa = new Empresa("Tech Corp");
        Empregado e1 = new Empregado("João", "Desenvolvedor", 3000.0);
        Empregado e2 = new Empregado("Maria", "Designer", 3500.0);

        empresa.adicionarEmpregado(e1);
        empresa.adicionarEmpregado(e2);

        empresa.listarEmpregados();
    }
}
