class Carro {
    String marca;
    String modelo;
    int ano;
    int velocidade;

    Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0;
    }

    void acelerar() {
        velocidade += 10;
    }

    void frear() {
        velocidade -= 10;
    }

    void exibirVelocidade() {
        System.out.println("Velocidade atual: " + velocidade);
    }
}
