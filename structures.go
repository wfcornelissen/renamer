package main

type company struct {
	name      string
	fileStart string
	fileNum   int
}

var no1Online = company{
	name:      "Number 1 Online Solutions",
	fileStart: "/N1O_",
	fileNum:   0,
}

var VirtigoResources = company{
	name:      "Virtigo Resources",
	fileStart: "/VirRes_",
	fileNum:   0,
}
var VirtigoAgri = company{
	name:      "Virtigo Agri",
	fileStart: "/VirAgri_",
	fileNum:   0,
}

var VirtigoSandAndStone = company{
	name:      "Virtigo Sand And Stone",
	fileStart: "/VirSand_",
	fileNum:   0,
}

var LutzvilleCellars = company{
	name:      "Lutzville Cellars",
	fileStart: "/LCell_",
	fileNum:   0,
}

var NamakwaLogistics = company{
	name:      "Namakwa Logistics",
	fileStart: "/NamLog_",
	fileNum:   0,
}

var companies = []company{
	no1Online,
	VirtigoResources,
	VirtigoAgri,
	VirtigoSandAndStone,
	LutzvilleCellars,
	NamakwaLogistics,
}
