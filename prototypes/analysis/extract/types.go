package extract

type Language string

const (
	Go         Language = "golang"
	Rust       Language = "rust"
	Java       Language = "java"
	JavaScript Language = "javascript"
	TypeScript Language = "typescript"
)

type ExtractResult struct {
	TypeDecls  map[string]TypeDecl
	Members    map[string]Member
	Interfaces map[string]TypeDecl
	Functions  map[string]Function
	NamedTypes map[string]Named
	Files      map[string]File
}

type Extractor interface {
	Extract(pkgstr string, dir string) (ExtractResult, error)
}

type Node struct {
	Name     string
	Code     string
	FilePath string
	Pos      int
	End      int
}

type Kind byte

const (
	Interface Kind = iota
	Struct
	Alias
)

type Named struct {
	Name       string
	QName      string
	TypeQName  string
	Underlying string
	Code       string
	Doc        Doc
	Pos        int
	End        int
	Filepath   string
}

type TypeDecl struct {
	Name            string
	QName           string
	TypeQName       string
	Underlying      string
	ImplementsQName string
	Code            string
	Doc             Doc
	Kind            Kind
	Pos             int
	End             int
	Filepath        string
	// Package    string
}

type ExtractNodesResult struct{}

type File struct {
	Filename string
	Package  string
	Imports  []string
	Language Language
}

type Namespace struct {
	Node
}

type Member struct {
	Name        string
	QName       string
	TypeQName   string
	ParentQName string
	Code        string
	Doc         Doc
	Pos         int
	End         int
	Filepath    string
}

type Function struct {
	Name         string
	QName        string
	ParentQName  string
	Calls        []string
	Doc          Doc
	Code         string
	Pos          int
	End          int
	Filepath     string
	ReturnQNames []string
	ParamQNames  []string
}

type DocType byte

const (
	SingleLine DocType = iota
	MultiLine
	Block
	Inline
)

type Import struct {
	Name string
	Path string
	Doc  Doc
}

type Doc struct {
	Comment string
	OfQName string
	Type    DocType
}
