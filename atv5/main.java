class Main {
    public static void emitirSom(Animal[] animais) {
        for (Animal animal : animais) {
            animal.som();
        }
    }

    public static void main(String[] args) {
        Animal[] animais = { new Cachorro(), new Gato() };
        emitirSom(animais);
    }
}
