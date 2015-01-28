package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"os"
)

func main() {
	fmt.Println("iniciando tail")
	ExampleNewWatcher("C:\\Temp\\wq-64Bits\\log\\WQ_AUDIT_XXPC_P9001_14012015.csv")
	t, err := tail.TailFile("C:\\Temp\\wq-64Bits\\log\\WQ_AUDIT_XXPC_P9001_14012015.csv", tail.Config{Follow: true, ReOpen: true})

	if err != nil {
		fmt.Printf("Erro ... %s", err)
		os.Exit(1)
	}
	for line := range t.Lines {
		fmt.Printf("linha -> %s\n", line.Text)
	}

}
