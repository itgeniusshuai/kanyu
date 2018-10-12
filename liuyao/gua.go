package liuyao

import "bytes"
import "../date"

type Gua struct{
	Yaos []*Yao

	Name string
	GuaNum int
}

type ChongGua struct{
	UpGua Gua
	DownGua Gua
	Name string
	WuXingName string
	Wuxing int
	DongYaoNums []int
}

type FinalGua struct{
	ZhuGua ChongGua
	BianGua ChongGua

	Lunar date.Lunar
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

func (this *FinalGua)String()string{
	var buf bytes.Buffer
	// 时间
	lunar := this.Lunar
	buf.WriteString(lunar.GetLunarYearString()+"年")
	buf.WriteByte('\t')
	buf.WriteString(lunar.GetCyclicaMonth()+"月")
	buf.WriteByte('\t')
	buf.WriteString(lunar.GetCyclicaDay()+"日")
	buf.WriteByte('\n')
	// 主卦名称 +变卦名称
	zhuGua := this.ZhuGua
	bianGua := this.BianGua
	buf.WriteString(zhuGua.Name)
	buf.WriteString("\t\t\t\t\t\t\t\t\t")
	buf.WriteString(bianGua.Name)
	buf.WriteByte('\n')
	// 主卦上挂爻 + 下卦爻
	for i,e := range zhuGua.UpGua.Yaos{
		e1 := bianGua.UpGua.Yaos[i]
		buf.WriteString(e.String())
		buf.WriteString("\t\t\t\t")
		buf.WriteString(e1.String())
		buf.WriteByte('\n')
	}
	for i,e := range zhuGua.DownGua.Yaos{
		e1 := bianGua.DownGua.Yaos[i]
		buf.WriteString(e.String())
		buf.WriteString("\t\t\t\t")
		buf.WriteString(e1.String())
		buf.WriteByte('\n')
	}

	return buf.String()
}

func (this *Gua)String()string{
	var buf bytes.Buffer
	//buf.Write([]byte(this.Name))
	for i := 0; i < 3; i++{
		buf.Write([]byte(this.Yaos[i].String()))
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	return buf.String()
}

func (this *Yao)String()string{
	var buf bytes.Buffer
	buf.WriteString(this.Xiang)
	if this.Prop == 1{
		if this.IsDong{
			buf.WriteString("  x")
		}else{
			buf.WriteString("  ﹑")
		}
	}else{
		if this.IsDong{
			buf.WriteString("  o")
		}else{
			buf.WriteString(" ﹑﹑")
		}
	}
	buf.WriteByte('\t')
	buf.WriteString("\t"+this.LiuQinName)
	buf.WriteString(this.DizhiName+WuxingSheng[DizhiWuxings[this.DiZhi]])
	if this.IsShi {
		buf.WriteString("\t世")
	}else{
		buf.WriteString("\t")
	}
	if this.IsYing {
		buf.WriteString("应\t")
	}else{
		buf.WriteString("\t")
	}
	return buf.String()
}

func (this *ChongGua)String()string{

	return this.Name +"\n"+this.UpGua.String()+this.DownGua.String()
}