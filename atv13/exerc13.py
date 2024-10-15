class Matematica:
    @staticmethod
    def fatorial(n):
        if n == 0:
            return 1
        resultado = 1
        for i in range(1, n + 1):
            resultado *= i
        return resultado

print(f"Fatorial de 5: {Matematica.fatorial(5)}")
