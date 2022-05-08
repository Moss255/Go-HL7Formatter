package main

import (
	"fmt"
	"os"
	strings "strings"
)

var debug bool = false

func main() {

	text, err := LoadFromFile("test.hl7")

	if err != nil {
		fmt.Println(err)
		return
	}

	text = FormatMessage(text, nil)

	fmt.Println("Writing to file")
	fmt.Println(text)

	hasWritten, writeErr := WriteToFile(text, "Real.hl7")

	if writeErr != nil || hasWritten != true {
		fmt.Println("An error has occured while write to the file")
		fmt.Println(writeErr)
		return
	}
}

func LoadFromFile(fileLocation string) (string, error) {

	// Simple load file function

	data, err := os.ReadFile(fileLocation)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func FormatMessage(text string, segments []string) string {

	// You can provide your own segments. I just added frequent used ones.

	if len(segments) <= 0 {
		segments = []string{"PID", "PD1", "PV1", "NTE", "PV2", "ZVI", "AL1", "ZAL", "ORC", "RXO", "RXR", "RXE", "OBX"}
	}

	if debug {
		fmt.Println("Checking for segments", segments)
	}

	for _, segment := range segments {
		text = strings.ReplaceAll(text, segment, "\n"+segment)
	}

	return text
}

func WriteToFile(formattedText string, fileLocation string) (bool, error) {
	data := []byte(formattedText)
	err := os.WriteFile(fileLocation, data, 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}
