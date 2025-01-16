package main

type company struct {
	name      string
	fileStart string
	fileNum   int
}

var no1Online = company{
	name:      "Number 1 Online Solutions",
	fileStart: "/N1O_",
}

var VirtigoResources = company{
	name:      "Virtigo Resources",
	fileStart: "/VirRes_",
}
var VirtigoAgri = company{
	name:      "Virtigo Agri",
	fileStart: "/VirAgri_",
}

var VirtigoSandAndStone = company{
	name:      "Virtigo Sand And Stone",
	fileStart: "/VirSand_",
}

var LutzvilleCellars = company{
	name:      "Lutzville Cellars",
	fileStart: "/LCell_",
}

var NamakwaLogistics = company{
	name:      "Namakwa Logistics",
	fileStart: "/NamLog_",
}

var companies = []company{
	no1Online,
	VirtigoResources,
	VirtigoAgri,
	VirtigoSandAndStone,
	LutzvilleCellars,
	NamakwaLogistics,
}
