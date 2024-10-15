class Motor:
    def __init__(self, tipo):
        self.tipo = tipo

class Carro:
    def __init__(self, marca, motor):
        self.marca = marca
        self.motor = motor

    def detalhes(self):
        print(f"Carro: {self.marca} com motor {self.motor.tipo}")

motor = Motor("V8")
carro = Carro("Ford", motor)
carro.detalhes()
