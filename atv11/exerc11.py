from abc import ABC, abstractmethod

class Funcionario(ABC):
    def __init__(self, nome):
        self.nome = nome

    @abstractmethod
    def calcular_salario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, nome, horas_trabalhadas, valor_hora):
        super().__init__(nome)
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_hora = valor_hora

    def calcular_salario(self):
        return self.horas_trabalhadas * self.valor_hora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, nome, salario_fixo):
        super().__init__(nome)
        self.salario_fixo = salario_fixo

    def calcular_salario(self):
        return self.salario_fixo

f1 = FuncionarioHorista("João", 160, 20)
f2 = FuncionarioAssalariado("Maria", 3000)

print(f"{f1.nome} recebe: R${f1.calcular_salario()}")
print(f"{f2.nome} recebe: R${f2.calcular_salario()}")
