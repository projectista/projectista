package project

type Project struct {
	name       string
	rootFolder string
}

func New(name, rootFolder string) *Project {
	t := &Project{
		name:       name,
		rootFolder: rootFolder,
	}
	return t
}

func (t *Project) Name() string {
	return t.name
}

func (t *Project) RootFolder() string {
	return t.rootFolder
}
