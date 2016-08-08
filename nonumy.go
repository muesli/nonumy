package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/andlabs/ui"
)

var FileToHash string

func main() {

	if len(os.Args) > 1 {
		FileToHash = os.Args[1]
	}

	// Handling command line input

	err := ui.Main(func() {

		file := ui.NewButton("Open File")
		filepath := ui.NewLabel("")
		md5Button := ui.NewButton("MD5")
		sha1Button := ui.NewButton("SHA1")
		sha256Button := ui.NewButton("SHA256")
		sha512Button := ui.NewButton("SHA512")
		hashsum := ui.NewEntry()
		compareSum := ui.NewButton("Compare?")
		compare := ui.NewEntry()

		box := ui.NewVerticalBox()
		box.Append(file, false)
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
		}
		// Choosing the file to hash
		file.OnClicked(func(*ui.Button) {
			FileToHash = ui.OpenFile(window)
			filepath.SetText(FileToHash)
		})

		// Going through the options
		md5Button.OnClicked(func(*ui.Button) {
			md5sum := md5.Sum(OpenFile(FileToHash))
			hashsum.SetText(fmt.Sprintf("%x", md5sum))
		})
		sha1Button.OnClicked(func(*ui.Button) {
			sha1sum := sha1.Sum(OpenFile(FileToHash))
			hashsum.SetText(fmt.Sprintf("%x", sha1sum))
		})
		sha256Button.OnClicked(func(*ui.Button) {
			sha256sum := sha256.Sum256(OpenFile(FileToHash))
			hashsum.SetText(fmt.Sprintf("%x", sha256sum))
		})
		sha512Button.OnClicked(func(*ui.Button) {
			sha512sum := sha512.Sum512(OpenFile(FileToHash))
			hashsum.SetText(fmt.Sprintf("%x", sha512sum))
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
	check(err, "Something failed")
}

func OpenFile(filepath string) []byte {
	// Currently im loading the whole file to memory
	file, err := ioutil.ReadFile(filepath)
	check(err, fmt.Sprintf("Could not read out file: %s", filepath))
	return file
}

func check(e error, s string) {
	if e != nil {
		log.Fatalln(s, e)
	}
}
