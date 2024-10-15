class Carro {
    String marca;
    String modelo;
    int ano;

    Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }

    void detalhes() {
        System.out.println(marca + " " + modelo + " " + ano);
    }
}

public class Main {
    public static void main(String[] args) {
        Carro c1 = new Carro("Ford", "Fusion", 2020);
        Carro c2 = new Carro("Honda", "Civic", 2021);
        Carro c3 = new Carro("Toyota", "Corolla", 2019);

        c1.detalhes();
        c2.detalhes();
        c3.detalhes();
    }
}
