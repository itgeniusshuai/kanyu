package liuyao

import (
	"bytes"
)

const(
	YIN  = 0
	YANG = 1
)

const(
	ZHI = iota
	CHOU
	YIN2
	MON
	CHEN
	SHI
	WU
	WEI
	SHEN
	YOU
	XU
	HAI
)

const(
	JIA = iota
	YI
	BING
	DING
	WU4
	JI
	GENG
	XIN
	REN
	GUI
)

const (
	QIAN = iota
	KAN
	GEN
	ZHEN
	XUN
	LI
	KUN
	DUI
)

const(
	BIJIE = iota
	FUMU
	ZHISUN
	QICAI
	GUANGUI
)

var Dizhis []string = []string{"子","丑","寅","卯","辰","巳","午","未","申","酉","戌","亥"}
var TianGans []string = []string{"甲","乙","丙","丁","戊","己","庚","辛","壬","癸"}
var LiuQins []string = []string{"兄弟","子孙","妻财","官鬼","父母"}
var BaguaNames = []string{"乾","兑","离","震","巽","坎","艮","坤"}
var ChongGuaNames = [][]string{	{"乾为天","天泽履","天火同人","天雷无妄","天风姤","天水讼","天山遁","天地否"},
								{"泽天夬","兑为泽","泽火革","泽雷随","泽风大过","泽水困","泽山咸","泽地萃"},
								{"火天大有","火泽睽","离为火","火雷噬嗑","火风鼎","火水未既","火山旅","火地晋"},
								{"雷天大壮","雷泽归妹","雷火丰","震","雷风恒","雷水解","雷山小过","雷地豫"},
								{"风天小畜","风泽中孚","风火家人","风雷益","巽为风","风水涣","风山渐","风地观"},
								{"水天需","水泽节","水火既济","水雷屯","水风井","坎为水","水山蹇","水地比"},
								{"山天大畜","山泽损","山火贲","山雷颐","山风蛊","山水蒙","艮为山","山地剥"},
								{"地天泰","地泽临","地火明夷","地雷覆","地风升","地水师","地山谦","坤为地"}}
// 金0 水1 木2 火3 土4
var WuxingSheng = []string{"金","水","木","火","土"}
var WuxingKe = []string{"木","火","土","金","水"}
var GuaWuxing = []int{0,0,3,2,2,1,4,4}
/**
	卦信息，卦五行4位，应4位，世4位，
	如乾卦，五行为金0，世位置1爻，应4爻，阳顺1 最后值为 0000 0100 0001  = ox0041
 */
var ChongGuaDesc = [][]int{{0x0043},{},{},{},{},{},{},{}}
/**
	初支爻象三位 地支4位,顺逆1
	如乾 1	111 0000  0xf 0
	兑   0	011 0101 0x35
	离 	0	101 0011   0x53
	震 	1	001 0000   0x90
	巽 	0	110 0001  0x61
	坎 	1	010 0010  0xa2
	艮 	1	100 0100  0xc4
	坤 	0	000 0111  0x0e
 */
var DanUpGuaDesc = []int{0xf6,0x3b,0x59,0x96,0x67,0xa8,0xca,0x01}
var DanDownGuaDesc = []int{0xf0,0x35,0x53,0x90,0x61,0xa2,0xc4,0x07}

/**
	卦信息，卦五行3位，应3位，世3位，顺逆1位，
	如乾卦，五行为金0，应位置4爻，世4爻， 最后值为 0000 0100 0001 = ox0041
 */
func ParseChongGuaDesc(gua ChongGua) *ChongGua{
	upGua := gua.UpGua
	downGua := gua.DownGua

	var desc = ChongGuaDesc[upGua.GuaNum-1][downGua.GuaNum-1]
	// 解析世位置
	var shiPos = desc & 0x0f
	// 解析应位置
	var yingPos = desc>>4 & 0xf
	// 解析五行属性
	var wuxing = desc>>8 & 0xf
	if shiPos > 3{
		upGua.Yaos[shiPos-4].IsShi = true
	}else{
		downGua.Yaos[shiPos-1].IsShi = true
	}
	if yingPos > 3{
		upGua.Yaos[shiPos-4].IsYing = true
	}else{
		upGua.Yaos[shiPos].IsYing = true
	}
	gua.Wuxing = wuxing
	gua.WuXingName = WuxingSheng[wuxing]
	return &gua
}

func ParseDanGuaDesc(isUp bool,guaNum int) *Gua{
	gua := Gua{GuaNum:guaNum,Name:BaguaNames[guaNum-1]}
	var desc int
	if isUp{
		desc = DanUpGuaDesc[guaNum-1]
	}else{
		desc = DanDownGuaDesc[guaNum-1]
	}

	var dizhi = desc & 0xf
	var yaoxiang = desc>>4 & 0x7
	var isShun = desc>>7 & 0x1
	var i uint
	var yaos []Yao
	for i = 0; i < 3; i++{
		yao := Yao{DiZhi:dizhi,DizhiName:Dizhis[dizhi]}
		if (yaoxiang>>(2-i) & 0x1) == 0 {
			yao.Prop = 0
			yao.Xiang = "- -"
		}else{
			yao.Prop = 1
			yao.Xiang = "---"
		}
		if isShun == 1{
			dizhi += 2
			if dizhi >= 12 {
				dizhi = dizhi%12
			}
		}
		if isShun == 0{
			dizhi -= 2
			if dizhi < 0{
				dizhi += 12
			}
		}
		yaos = append(yaos, yao)
	}
	gua.Yaos = yaos
	return &gua
}

func (this *Gua)String()string{
	var buf bytes.Buffer
	buf.Write([]byte(this.Name))
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
	buf.WriteString("  "+this.LiuQinName)
	buf.WriteString("  "+this.DizhiName)
	if this.IsShi {
		buf.WriteString("  世")
	}
	if this.IsYing {
		buf.WriteString("  应")
	}
	return buf.String()
}

func (this *ChongGua)String()string{
	return this.UpGua.String()+this.DownGua.String()
}


