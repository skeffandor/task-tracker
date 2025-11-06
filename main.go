package main

import (
	"log"

	"github.com/skeffandor/go-project-1/internal/cli"
	"github.com/skeffandor/go-project-1/internal/manager"
)

func main() {
	tm := manager.NewTaskManager()

	if err := manager.Load(tm, "tm.json"); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := manager.Save(tm, "tm.json"); err != nil {
			log.Fatal(err)
		}
	}()

	cli.InitCLI(tm)
}
