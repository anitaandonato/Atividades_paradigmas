class ContaBancaria:
    def __init__(self, titular, saldo_inicial):
        self.__saldo = saldo_inicial
        self.titular = titular

    def depositar(self, valor):
        self.__saldo += valor

    def sacar(self, valor):
        if valor <= self.__saldo:
            self.__saldo -= valor
        else:
            print("Saldo insuficiente.")

    def get_saldo(self):
        return self.__saldo
