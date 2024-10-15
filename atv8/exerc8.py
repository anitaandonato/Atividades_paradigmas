class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

class Empresa:
    def __init__(self, nome):
        self.nome = nome
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def listar_empregados(self):
        print(f"Empregados da empresa {self.nome}:")
        for empregado in self.empregados:
            print(f"{empregado.nome} - {empregado.cargo} - R${empregado.salario}")

empresa = Empresa("Tech Corp")
e1 = Empregado("João", "Desenvolvedor", 3000.0)
e2 = Empregado("Maria", "Designer", 3500.0)

empresa.adicionar_empregado(e1)
empresa.adicionar_empregado(e2)

empresa.listar_empregados()
