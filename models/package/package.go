package _package

type Package struct {
	Id    string
	Path  string
	Epoch string
}

func (Package) TableName() string {
	return "rhnpackage"
}
