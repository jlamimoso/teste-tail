package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"os"
)

func main() {
	fmt.Println("iniciando tail")
	t, err := tail.TailFile("teste.txt", tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Printf("Erro ... %s", err)
		os.Exit(1)
	}
	for line := range t.Lines {
		fmt.Printf("linha -> %s\n", line.Text)
	}
}
