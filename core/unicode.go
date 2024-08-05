package core

import (
	"fmt"
	"io"
	"unicode/utf8"
)

type UnicodeDetector interface {
	IsUnicode(reader io.Reader, writer io.Writer) error
}

type unicodeDetector struct{}

func NewUnicodeDetector() UnicodeDetector {
	return &unicodeDetector{}
}

func (ud *unicodeDetector) IsUnicode(reader io.Reader, writer io.Writer) error {
	content, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading content: %v", err)
	}

	if !utf8.Valid(content) {
		fmt.Fprintf(writer, "The content contains invalid UTF-8\n")
		return nil
	}

	if utf8.RuneCount(content) != 1 {
		fmt.Fprintf(writer, "The content does not contain a single character\n")
		return nil
	}

	r, _ := utf8.DecodeRune(content)
	codePoint := fmt.Sprintf("U+%04X", r)

	fmt.Fprintf(writer, "Is Unicode: true\nCode Point: %s\n", codePoint)
	return nil
}
