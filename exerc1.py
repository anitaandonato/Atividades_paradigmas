class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano

    def detalhes(self):
        print(f"{self.marca} {self.modelo} {self.ano}")

c1 = Carro("Ford", "Fusion", 2020)
c2 = Carro("Honda", "Civic", 2021)
c3 = Carro("Toyota", "Corolla", 2019)

c1.detalhes()
c2.detalhes()
c3.detalhes()
