class Configuracao {
    private static Configuracao instancia;

    private Configuracao() {
        // Construtor privado
    }

    public static Configuracao getInstancia() {
        if (instancia == null) {
            instancia = new Configuracao();
        }
        return instancia;
    }

    public void exibirConfiguracao() {
        System.out.println("Exibindo a configuração...");
    }
}

public class Main {
    public static void main(String[] args) {
        Configuracao config1 = Configuracao.getInstancia();
        Configuracao config2 = Configuracao.getInstancia();

        config1.exibirConfiguracao();
        System.out.println(config1 == config2);  // true, porque ambas são a mesma instância
    }
}
