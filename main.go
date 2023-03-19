package main

import (
	"fmt"
	"os"
	"phpReindentNestedDocblocks/internal/helper"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("No file path provided")
		return
	}
	filePath := args[1]

	outputPath := filePath
	if len(args) > 2 {
		outputPath = args[2]
	}

	// read file path content

	phpFileHelper := &helper.PhpFileHelper{}

	cleanContent := phpFileHelper.GetDocblockCleanedContent(filePath)

	err := os.WriteFile(outputPath, []byte(cleanContent), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
