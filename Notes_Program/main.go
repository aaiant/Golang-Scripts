package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/Notes_Program/note"
	"example.com/Notes_Program/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)

	if err != nil {
		return
	}
}

func getUserInput(prompt string) (text string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return
}

func getNoteData() (title string, content string) {
	title = getUserInput("Note title: ")
	content = getUserInput("Note content: ")
	return
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
