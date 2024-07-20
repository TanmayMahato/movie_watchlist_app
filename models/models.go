package models

type Mvd struct {
	Name string
	Gen  string
	Cat  string
	Exp  int
}

type Mvdata struct {
	Mvid int
	Name string
	Gen  string
	Cat  string
	Exp  int
}

type Mvwdata struct {
	Mvid int
	Name string
	Gen  string
	Cat  string
	Rate int
}

type Id struct {
	Movieid int
}
