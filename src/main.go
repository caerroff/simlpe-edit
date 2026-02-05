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

	contents, err := openDirectory(dirPath)
	if err != nil {
		return
	}
	w.SetContent(container.NewVBox(
		workDirWidget,
		container.NewGridWithColumns(
			2,
			createFileExplorer(contents),
			createTextEditor(contents[0]),
		),
	))

	w.ShowAndRun()
}

func createTextEditor(file os.DirEntry) fyne.CanvasObject {
	var editor fyne.CanvasObject
	if file.IsDir() {
		return editor
	}
	text, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error while opening the file")
	}

	editor = widget.NewRichTextWithText(string(text))
	return editor
}

func createFileExplorer(contents []os.DirEntry) *fyne.Container {
	fileExplorer := container.NewGridWithRows(
		len(contents),
	)
	for _, entry := range contents {
		fileExplorer.Add(widget.NewButton(entry.Name(), func() {}))
	}

	fileExplorer.Resize(fyne.NewSize(300, 600))

	return fileExplorer
}

func openDirectory(path string) ([]os.DirEntry, error) {
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
