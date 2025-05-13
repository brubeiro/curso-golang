package main

import (
	"calculator/math"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	// Usado só para conseguirmos importar o pacote
	// Não precisa rodar o echo
	e := echo.New()

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	// Função privada (com letra minúscula)
	// não pode ser usada fora do pacote math
	sub := math.sub(1, 2)
	fmt.Println(sub)
}
