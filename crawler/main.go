package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"crawler.me/screenshot"
)

func fatalOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func exePath() string {
	exe, err := os.Executable()
	fatalOnError(err, " could not locate executable")
	return exe
}

func main() {
	dir, err := ioutil.TempDir("/tmp", "chromedp-example-")
	fatalOnError(err, "creating temporary directory")

	fmt.Println("图片存放路径: ", dir)
	file := filepath.Join(dir, "example.png")
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	fatalOnError(err, "open temporary file")
	sh := screenshot.NewScreenshot()
	err = sh.PickImage("https://www.google.com", "#main", 800, 600, f)
	fatalOnError(err, "failed to pick image")

}
