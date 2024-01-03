package writer

import "os"

type Writer struct {
	outDirectory string
}

func New(outDirectory string) *Writer {

	// Exit with error if the outDirectory does not exists
	if _, err := os.Stat(outDirectory); os.IsNotExist(err) {
		println("Folder", outDirectory, "does not exists")
		os.Exit(1)
	}

	t := &Writer{outDirectory: outDirectory}

	return t
}
