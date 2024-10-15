class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, outro):
        return self.preco + outro.preco

p1 = Produto("Produto A", 100.0)
p2 = Produto("Produto B", 200.0)
print(f"Soma dos preços: R${p1 + p2}")
