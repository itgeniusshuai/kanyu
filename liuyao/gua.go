package liuyao


type Gua struct{
	Yao1 Yao
	Yao2 Yao
	Yao3 Yao

	Name string
}

type ChongGua struct{
	UpGua Gua
	DownGua Gua
	Name string
}

type FinalGua struct{
	ZhuGua ChongGua
	BianGua ChongGua
}

type Yao struct{
	Prop byte
	IsShi bool
	IsYing bool
	IsDong bool

	DiZhi byte
	LiuQin byte
}
