package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"github.com/go-fsnotify/fsnotify"
	"log"
	"os"
)

func ExampleNewWatcher(filename string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

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
