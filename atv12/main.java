class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public double somarPreco(Produto outro) {
        return this.preco + outro.preco;
    }

    public static void main(String[] args) {
        Produto p1 = new Produto("Produto A", 100.0);
        Produto p2 = new Produto("Produto B", 200.0);
        System.out.println("Soma dos preços: R$" + p1.somarPreco(p2));
    }
}
