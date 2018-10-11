package liuyao

import (
	"fmt"
	"github.com/itgeniusshuai/go_common/common"
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
	卦信息，卦五行3位，应3位，世3位，顺逆1位，
	如乾卦，五行为金0，世位置1爻，应4爻，阳顺1 最后值为 000 100 001 1 = ox0043
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
	如乾卦，五行为金0，世位置1爻，应4爻，阳顺1 最后值为 000 100 001 1 = ox0043
 */
func ParseChongGuaDesc(upGuaNum,downGuanNum int){
	var desc = ChongGuaDesc[upGuaNum-1][downGuanNum-1]
	// 解析顺逆
	var isShun = desc & 0x01
	// 解析世位置
	var shiPos = desc>>1 & 0x07
	// 解析应位置
	var yingPos = desc>>4 & 0x7
	// 解析五行属性
	var wuxing = desc>>7 & 0x7

	fmt.Println("是否顺:"+common.IntToStr(isShun))
	fmt.Println("世位置:"+common.IntToStr(shiPos))
	fmt.Println("应位置:"+common.IntToStr(yingPos))
	fmt.Println("五行:"+WuxingSheng[wuxing])
}

func ParseDanGuaDesc(isUp bool,downGuaNum int){
	var desc int
	if isUp{
		desc = DanUpGuaDesc[downGuaNum-1]
	}else{
		desc = DanDownGuaDesc[downGuaNum-1]
	}


	var dizhi = desc & 0xf
	var yaoxiang = desc>>4 & 0x7
	var isShun = desc>>7 & 0x1
	fmt.Println("是否顺:"+common.IntToStr(isShun))
	fmt.Println("地支:"+Dizhis[dizhi])
	var i uint
	for i = 0; i < 3; i++{
		if (yaoxiang>>(2-i) & 0x1) == 0 {
			fmt.Println("- -")
		}else{
			fmt.Println("---")
		}
	}
}


