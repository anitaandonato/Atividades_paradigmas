import java.util.ArrayList;
import java.util.List;

class Professor {
    String nome;

    Professor(String nome) {
        this.nome = nome;
    }
}

class Escola {
    String nome;
    List<Professor> professores;

    Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    void adicionarProfessor(Professor professor) {
        professores.add(professor);
    }

    void listarProfessores() {
        System.out.println("Professores da escola " + nome + ":");
        for (Professor p : professores) {
            System.out.println(p.nome);
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Escola escola = new Escola("UNIPÊ");
        Professor p1 = new Professor("Professor João");
        Professor p2 = new Professor("Professor Maria");

        escola.adicionarProfessor(p1);
        escola.adicionarProfessor(p2);

        escola.listarProfessores();
    }
}
