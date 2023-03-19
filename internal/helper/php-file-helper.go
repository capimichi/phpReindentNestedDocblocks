package helper

import (
	"fmt"
	"os"
	"strings"
)

type PhpFileHelper struct {
}

func (phpFileHelper *PhpFileHelper) GetDocblockCleanedContent(filePath string) string {

	content := phpFileHelper.ReadFile(filePath)

	// split content by new line
	lines := strings.Split(content, "\n")

	dockblockStarted := false
	dockblockIndentation := 0
	for i, line := range lines {

		if dockblockStarted {
			lines[i] = phpFileHelper.ReplaceLineContent(line, dockblockIndentation)
		}

		if phpFileHelper.IsDockblockStarting(line) {
			dockblockStarted = true
			dockblockIndentation = 0
		}

		if phpFileHelper.IsDockblockIncrementLevel(line) {
			dockblockIndentation++
		}

		if phpFileHelper.IsDockblockDecrementLevel(line) {
			dockblockIndentation--
		}

		if phpFileHelper.IsDockblockEnding(line) {
			dockblockStarted = false
		}
	}

	// join lines
	newContent := strings.Join(lines, "\n")

	return newContent
}

func (phpFileHelper *PhpFileHelper) ReplaceLineContent(line string, indentCount int) string {
	// remove everthing after *

	if phpFileHelper.IsDockblockEnding(line) {
		return line
	}

	if phpFileHelper.IsDockblockStarting(line) {
		return line
	}

	clean := phpFileHelper.GetCleanLine(line)

	finalString := strings.Split(line, "*")[0]

	finalString = finalString + "*"

	for i := 0; i < indentCount; i++ {
		finalString = finalString + "    "
	}

	if len(clean) > 0 && clean != "/" && indentCount == 0 {
		finalString = finalString + " "
	}

	finalString = finalString + clean

	return finalString
}

func (phpFileHelper *PhpFileHelper) GetCleanLine(line string) string {
	// return everything after *

	clean := strings.Split(line, "*")[1]

	// remove spaces
	clean = strings.TrimSpace(clean)

	return clean
}

func (phpFileHelper *PhpFileHelper) IsDockblockStarting(line string) bool {
	return strings.Contains(line, "/**")
}

func (phpFileHelper *PhpFileHelper) IsDockblockIncrementLevel(line string) bool {
	return strings.Contains(line, "(")
}

func (phpFileHelper *PhpFileHelper) IsDockblockDecrementLevel(line string) bool {
	return strings.Contains(line, ")")
}

func (phpFileHelper *PhpFileHelper) IsDockblockEnding(line string) bool {
	return strings.Contains(line, "*/")
}

func (phpFileHelper *PhpFileHelper) ReadFile(filePath string) string {

	f, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	return string(f)
}
