from typing import Protocol

class Imprimivel(Protocol):
    def imprimir(self) -> None:
        ...

class Relatorio:
    def imprimir(self) -> None:
        print("Imprimindo relatório...")

class Contrato:
    def imprimir(self) -> None:
        print("Imprimindo contrato...")

def imprimir_documento(doc: Imprimivel) -> None:
    doc.imprimir()

relatorio = Relatorio()
contrato = Contrato()

imprimir_documento(relatorio)
imprimir_documento(contrato)
