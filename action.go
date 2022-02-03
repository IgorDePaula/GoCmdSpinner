package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
)

type actionDone struct{}
type actionStep struct{}
type Logs struct {
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
}

func runAction(m model) <-chan tea.Msg {

	ch := make(chan tea.Msg, 0)

	go func() {
		/*time.Sleep(time.Second * 1)
		ch <- actionStep{}
		time.Sleep(time.Second * 2)
		ch <- actionStep{}
		time.Sleep(time.Second * 3)
		ch <- actionStep{}*/
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		Logs := Logs{}
		Logs.InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Logs.WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Logs.ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		Logs.InfoLogger.Println("diretorio " + m.Dir)

		cmd := exec.Command("composer", "require", "monolog/monolog")
		//cmd.Dir = ".." //diretorio anterior ao projeto
		cmd.Dir = m.Dir //diretorio anterior ao projeto
		err = cmd.Run()
		if err != nil {
			log.Panic(err)
		}
		ch <- actionDone{}

		close(ch)
	}()

	return ch
}
