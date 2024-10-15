class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0

    def acelerar(self):
        self.velocidade += 10

    def frear(self):
        self.velocidade -= 10

    def exibir_velocidade(self):
        print(f"Velocidade atual: {self.velocidade}")

c = Carro("Ford", "Fusion", 2020)
c.acelerar()
c.exibir_velocidade()
c.frear()
c.exibir_velocidade()
