class Configuracao:
    _instancia = None

    def __new__(cls):
        if cls._instancia is None:
            cls._instancia = super(Configuracao, cls).__new__(cls)
        return cls._instancia

    def exibir_configuracao(self):
        print("Exibindo a configuração...")

config1 = Configuracao()
config2 = Configuracao()

config1.exibir_configuracao()
print(config1 is config2)  # True, pois ambas são a mesma instância
