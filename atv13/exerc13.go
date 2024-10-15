package main
import "fmt"

func Fatorial(n int) int {
    if n == 0 {
        return 1
    }
    resultado := 1
    for i := 1; i <= n; i++ {
        resultado *= i
    }
    return resultado
}

func main() {
    fmt.Println("Fatorial de 5:", Fatorial(5))
}
