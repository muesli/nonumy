package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andlabs/ui"
)

var FileToHash string
var contents []byte
var err error

func main() {

	if len(os.Args) > 1 {
		FileToHash = os.Args[1]
	}

	err := ui.Main(func() {

		OpenFileButton := ui.NewButton("Open File")
		filepath := ui.NewLabel("")
		md5Button := ui.NewButton("MD5")
		sha1Button := ui.NewButton("SHA1")
		sha256Button := ui.NewButton("SHA256")
		sha512Button := ui.NewButton("SHA512")
		hashsum := ui.NewEntry()
		compareSum := ui.NewButton("Compare?")
		compare := ui.NewEntry()

		box := ui.NewVerticalBox()
		box.Append(OpenFileButton, false)
		box.Append(filepath, false)
		box.Append(md5Button, false)
		box.Append(sha1Button, false)
		box.Append(sha256Button, false)
		box.Append(sha512Button, false)
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

		// Going through the options
		md5Button.OnClicked(func(*ui.Button) {
			hashsum.SetText(fmt.Sprintf("%x", md5.Sum(contents)))
		})
		sha1Button.OnClicked(func(*ui.Button) {
			hashsum.SetText(fmt.Sprintf("%x", sha1.Sum(contents)))
		})
		sha256Button.OnClicked(func(*ui.Button) {
			hashsum.SetText(fmt.Sprintf("%x", sha256.Sum256(contents)))
		})
		sha512Button.OnClicked(func(*ui.Button) {
			hashsum.SetText(fmt.Sprintf("%x", sha512.Sum512(contents)))
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
