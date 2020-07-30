package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/alexmullins/zip"
)

func main() {

	var (
		ZipFilePath  string
		TargetedFile string
		FilePassword string
	)

	ZipFilePath = "/path/to/archive/file.zip"         // you can write the filename.zip if you want your archived file in the same directory as your service file
	TargetedFile = "/path/to/targeted/file.something" // you can write the filename.something if your file is in the same directory as your service file
	FilePassword = "yourPasswordForTheFile"

	zipFile, _ := os.Create(ZipFilePath) // or you can write (os.Create("./" + ZipFilePath)) if you want your archived file in the same directory as your service file
	defer zipFile.Close()
	targetFile, _ := os.Open(TargetedFile) // or you can write (os.Open("./" + TargetedFile)) if your file is in the same directory as your service file
	defer targetFile.Close()
	reader := bufio.NewReader(targetFile)

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	//compress + kasih password
	w, errZip := zipWriter.Encrypt(TargetedFile, FilePassword)
	if errZip != nil {
		log.Panic(errZip.Error())
	}

	_, err := io.Copy(w, reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Archive completed at " + time.Now().String())
}
