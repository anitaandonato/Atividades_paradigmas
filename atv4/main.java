class Animal {
    public void som() {
        System.out.println("O animal faz um som.");
    }
}

class Cachorro extends Animal {
    @Override
    public void som() {
        System.out.println("O cachorro late.");
    }
}

class Gato extends Animal {
    @Override
    public void som() {
        System.out.println("O gato mia.");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal a1 = new Cachorro();
        Animal a2 = new Gato();

        a1.som();
        a2.som();
    }
}
