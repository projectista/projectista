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

// Write will write the provided bytes into the provided file
func (t *Writer) Write(file string, fileBytes []byte) string {
	completePathToFile := t.outDirectory + string(os.PathSeparator) + file

	err := os.WriteFile(
		completePathToFile, fileBytes, 0644,
	)
	if err != nil {
		panic(err)
	}

	return completePathToFile
}
