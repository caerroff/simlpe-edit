package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Welcome to Simple Edit!")
	a := app.New()
	w := a.NewWindow("Simple Edit")
	dirPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error when trying to access path of executable")
		return
	}
	os.Chdir(dirPath)
	workDirWidget := widget.NewLabel(dirPath)

	contents, err := GetContentOfDirectory(dirPath)
	if err != nil {
		return
	}

	var index int = 0
	var openedFileContent *os.DirEntry = &contents[index]
	editor := CreateTextEditor(openedFileContent)

	w.SetContent(container.NewVBox(
		workDirWidget,
		container.NewGridWithColumns(
			2,
			CreateFileExplorer(contents, &index),
			editor,
		),
	))

	w.ShowAndRun()
}

func CreateTextEditor(file *os.DirEntry) fyne.CanvasObject {
	var editor fyne.CanvasObject
	actualFile := *file
	if actualFile.IsDir() {
		return editor
	}
	text, err := os.ReadFile(actualFile.Name())
	if err != nil {
		fmt.Println("Error while opening the file")
	}

	editor = widget.NewRichTextWithText(string(text))
	return editor
}

func CreateFileExplorer(contents []os.DirEntry, index *int) *fyne.Container {
	fileExplorer := container.NewGridWithRows(
		len(contents),
	)
	for i, entry := range contents {
		fileExplorer.Add(widget.NewButton(entry.Name(), func() {
			*index = i
			fmt.Printf("Index %d\n", *index)
		}))
	}

	fileExplorer.Resize(fyne.NewSize(300, 600))

	return fileExplorer
}

func GetContentOfDirectory(path string) ([]os.DirEntry, error) {
	// if !path.IsDir() {
	// 	return contents, errors.New("Path is not a directory.")
	// }

	contents, error := os.ReadDir(path)
	if error != nil {
		fmt.Printf("Error when trying to open dir %s", path)
		return make([]os.DirEntry, 0), error
	}

	return contents, nil
}
