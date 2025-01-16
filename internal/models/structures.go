package models

type company struct {
	Name      string
	FileStart string
	FileNum   int
}

var no1Online = company{
	Name:      "Number 1 Online Solutions",
	FileStart: "/N1O_",
}

var VirtigoResources = company{
	Name:      "Virtigo Resources",
	FileStart: "/VirRes_",
}
var VirtigoAgri = company{
	Name:      "Virtigo Agri",
	FileStart: "/VirAgri_",
}

var VirtigoSandAndStone = company{
	Name:      "Virtigo Sand And Stone",
	FileStart: "/VirSand_",
}

var LutzvilleCellars = company{
	Name:      "Lutzville Cellars",
	FileStart: "/LCell_",
}

var NamakwaLogistics = company{
	Name:      "Namakwa Logistics",
	FileStart: "/NamLog_",
}

var Companies = []company{
	no1Online,
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
