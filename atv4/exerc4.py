class Animal:
    def som(self):
        print("O animal faz um som.")

class Cachorro(Animal):
    def som(self):
        print("O cachorro late.")

class Gato(Animal):
    def som(self):
        print("O gato mia.")

a1 = Cachorro()
a2 = Gato()

a1.som()
a2.som()
