class Professor:
    def __init__(self, nome):
        self.nome = nome

class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []

    def adicionar_professor(self, professor):
        self.professores.append(professor)

    def listar_professores(self):
        print(f"Professores da escola {self.nome}:")
        for professor in self.professores:
            print(professor.nome)

escola = Escola("UNIPÊ")
p1 = Professor("Professor João")
p2 = Professor("Professor Maria")

escola.adicionar_professor(p1)
escola.adicionar_professor(p2)

escola.listar_professores()
