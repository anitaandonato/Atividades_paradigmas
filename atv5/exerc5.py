def emitir_som(animais):
    for animal in animais:
        animal.som()

animais = [Cachorro(), Gato()]
emitir_som(animais)
