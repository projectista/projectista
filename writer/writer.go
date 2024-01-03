package writer

import "os"

// Writer is an helper used to write
// files on the filesystem, in a specific outDirectory
type Writer struct {
	outDirectory string
}

// New creates a new instace of the Writer helper
func New(outDirectory string) *Writer {

	// Exit with error if the outDirectory does not exists
	if _, err := os.Stat(outDirectory); os.IsNotExist(err) {
		println("Folder", outDirectory, "does not exists")
		os.Exit(1)
	}

	t := &Writer{outDirectory: outDirectory}

	return t
}

// OutDirectory returns the directory to which the writer will write files
func (t *Writer) OutDirectory() string {
	return t.outDirectory
}
