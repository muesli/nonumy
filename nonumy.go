package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"

	"github.com/andlabs/ui"
)

var FileToHash string
var contents []byte
var err error

type hashFunc func() hash.Hash

func addButton(name string, fn hashFunc, label *ui.Entry) *ui.Button {
	btn := ui.NewButton(name)

	// Going through the options
	btn.OnClicked(func(*ui.Button) {
		hasher := fn()
		io.WriteString(hasher, string(contents))

		label.SetText(fmt.Sprintf("%x", hasher.Sum(nil)))
	})

	return btn
}

func main() {

	if len(os.Args) > 1 {
		FileToHash = os.Args[1]
	}

	err := ui.Main(func() {

		OpenFileButton := ui.NewButton("Open File")
		filepath := ui.NewLabel("")
		hashsum := ui.NewEntry()
		compareSum := ui.NewButton("Compare?")
		compare := ui.NewEntry()

		box := ui.NewVerticalBox()
		box.Append(OpenFileButton, false)
		box.Append(filepath, false)

		box.Append(addButton("MD5", md5.New, hashsum), false)
		box.Append(addButton("SHA-1", sha1.New, hashsum), false)
		box.Append(addButton("SHA-256", sha256.New, hashsum), false)
		box.Append(addButton("SHA-512", sha512.New, hashsum), false)

		box.Append(hashsum, false)
		box.Append(compareSum, false)
		box.Append(compare, false)

		window := ui.NewWindow("nonumy", 500, 250, false)
		window.SetChild(box)

		if FileToHash != "" {
			filepath.SetText(FileToHash)
			contents, err = ioutil.ReadFile(FileToHash)
			if err != nil {
				hashsum.SetText(fmt.Sprintf("Could not read out file: %s", FileToHash))
			}
		}
		OpenFileButton.OnClicked(func(*ui.Button) {
			FileToHash = ui.OpenFile(window)
			filepath.SetText(FileToHash)
			contents, err = ioutil.ReadFile(FileToHash)
			if err != nil {
				hashsum.SetText(fmt.Sprintf("Could not read out file: %s", FileToHash))
			}
		})

		// Optionally you can compare two hashsums
		compareSum.OnClicked(func(*ui.Button) {
			if compare.Text() != hashsum.Text() {
				compareSum.SetText("False!")
			} else {
				compareSum.SetText("True")
			}
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
