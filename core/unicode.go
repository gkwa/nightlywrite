package core

import (
	"fmt"
	"os"
	"unicode/utf8"
)

type UnicodeDetector interface {
	IsUnicode(filePath string) (bool, string, error)
}

type unicodeDetector struct{}

func NewUnicodeDetector() UnicodeDetector {
	return &unicodeDetector{}
}

func (ud *unicodeDetector) IsUnicode(filePath string) (bool, string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return false, "", fmt.Errorf("error reading file: %v", err)
	}

	if !utf8.Valid(content) {
		return false, "", fmt.Errorf("the file contains invalid UTF-8")
	}

	if utf8.RuneCount(content) != 1 {
		return false, "", fmt.Errorf("the text file does not contain a single character")
	}

	r, _ := utf8.DecodeRune(content)
	codePoint := fmt.Sprintf("U+%04X", r)

	return true, codePoint, nil
}
