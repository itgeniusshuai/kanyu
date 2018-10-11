package liuyao


type Gua struct{
	Yaos []Yao

	Name string
	GuaNum int
}

type ChongGua struct{
	UpGua Gua
	DownGua Gua
	Name string
	WuXingName string
	Wuxing int
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

	DiZhi int
	LiuQin int

	Xiang string
	DizhiName string
	LiuQinName string
}
