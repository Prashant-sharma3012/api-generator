package writers

// responsible for creating folder structure befotre writers start writing to them

// for now hardcoding thinsg

type Folders struct {
	Name     string   `json:"name"`
	Contents []string `json:"contents"`
}

type FolderStructure struct {
	RootFolders []Folders `json:"rootFolders"`
	RootFiles   []string  `json:"rootFiles"`
}
