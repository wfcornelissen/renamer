package models

type Company struct {
	Name      string
	FileStart string
	FileNum   int
}

var No1Online = Company{
	Name:      "Number 1 Online Solutions",
	FileStart: "/N1O_",
}

var VirtigoResources = Company{
	Name:      "Virtigo Resources",
	FileStart: "/VirRes_",
}
var VirtigoAgri = Company{
	Name:      "Virtigo Agri",
	FileStart: "/VirAgri_",
}

var VirtigoSandAndStone = Company{
	Name:      "Virtigo Sand And Stone",
	FileStart: "/VirSand_",
}

var LutzvilleCellars = Company{
	Name:      "Lutzville Cellars",
	FileStart: "/LCell_",
}

var NamakwaLogistics = Company{
	Name:      "Namakwa Logistics",
	FileStart: "/NamLog_",
}

var Companies = []Company{
	No1Online,
	VirtigoResources,
	VirtigoAgri,
	VirtigoSandAndStone,
	LutzvilleCellars,
	NamakwaLogistics,
}

type EntryMethod map[int]string

var EntryMethods = EntryMethod{
	1: "Manual",
	2: "Auto",
}
