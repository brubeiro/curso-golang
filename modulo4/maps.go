package main

import "fmt"

func main() {
	idade["bruno"] = 32
	idade["rebeca"] = 28

	fmt.Println(idade)
	fmt.Println(idade["bruno"])
	fmt.Println(idade["rebeca"])

	anoNasc := map[string]int{}
		"bruno":  1992,
		"rebeca": 1996,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["bruno"])
	fmt.Println(anoNasc["rebeca"])

}
