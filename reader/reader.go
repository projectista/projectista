package reader

// Reader is an helper to read
// files from the filesystem, and using
// them as templates, using the go templating system
type Reader struct {
	sourceDirectory string
}

// New creates a new instance of the Reader helper
func New(sourceDirectory string) *Reader {
	t := &Reader{sourceDirectory: sourceDirectory}

	return t
}

// SourceDirectory returns the directory from which the files will be read.
func (t *Reader) SourceDirectory() string {
	return t.sourceDirectory
}
