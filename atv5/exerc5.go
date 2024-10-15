package main

func emitirSom(animais []Animal) {
    for _, animal := range animais {
        animal.som()
    }
}

func main() {
    animais := []Animal{Cachorro{}, Gato{}}
    emitirSom(animais)
}
