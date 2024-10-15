class Calculadora:
    def somar_dois(self, a, b):
        return a + b

    def somar_tres(self, a, b, c):
        return a + b + c

calc = Calculadora()
print(calc.somar_dois(2, 3))
print(calc.somar_tres(1, 2, 3))
