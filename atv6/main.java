class Motor {
    private String tipo;

    Motor(String tipo) {
        this.tipo = tipo;
    }

    public String getTipo() {
        return tipo;
    }
}

class Carro {
    private String marca;
    private Motor motor;

    Carro(String marca, Motor motor) {
        this.marca = marca;
        this.motor = motor;
    }

    public void detalhes() {
        System.out.println("Carro: " + marca + " com motor " + motor.getTipo());
    }
}

public class Main {
    public static void main(String[] args) {
        Motor motor = new Motor("V8");
        Carro carro = new Carro("Ford", motor);
        carro.detalhes();
    }
}
