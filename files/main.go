package main

import (
	"fmt"

	"qmmr.xyz/go/files/fileutils"
)

func main() {
	var filename string
	fmt.Println("Enter the filename to read...")
	fmt.Scanf("%s", &filename)
	content, err := fileutils.ReadTextFile(filename)

	if err != nil {
		fmt.Printf("Oops, there was an error:\n%v", err)
	}

	fmt.Println(content)

	newContent := fmt.Sprintf("Original text: %v\nDuplicate text: %v\nThe end.", content, content)

	fileutils.WriteToFile("./output.txt", newContent)
}
