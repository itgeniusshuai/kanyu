package liuyao

import (
	"../date"
)

var baguaNames = []string{"乾","兑","离","震","巽","坎","艮","坤"}
var chongGuaNames = [][]string{	{"乾为天","天泽履","天火同人","天雷无妄","天风姤","天水讼","天山遁","天地否"},
							{"泽天夬","兑为泽","泽火革","泽雷随","泽风大过","泽水困","泽山咸","泽地萃"},
							{"火天大有","火泽睽","离为火","火雷噬嗑","火风鼎","火水未既","火山旅","火地晋"},
							{"雷天大壮","雷泽归妹","雷火丰","震","雷风恒","雷水解","雷山小过","雷地豫"},
							{"风天小畜","风泽中孚","风火家人","风雷益","巽为风","风水涣","风山渐","风地观"},
							{"水天需","水泽节","水火既济","水雷屯","水风井","坎为水","水山蹇","水地比"},
							{"山天大畜","山泽损","山火贲","山雷颐","山风蛊","山水蒙","艮为山","山地剥"},
							{"地天泰","地泽临","地火明夷","地雷覆","地风升","地水师","地山谦","坤为地"}}
// 金0水1木2火3土4
var wuxingSheng = []string{"金","水","木","火","土"}
var wuxingKe = []string{"木","火","土","金","水"}
var guaWuxing = []int{0,0,3,2,2,1,4,4}
/**
	卦信息，卦五行3位，世3位，应3位，顺逆1位，初支地支4位

 */
var chongGuaDesc = [][]int{{},{},{},{},{},{},{},{}}
// 取挂
// 取动
// 取卦支
// 取六亲
// 取世应
func QiGuaByTime() *FinalGua{

	l := date.GetLunar()
	// 获取年月日
	y := l.GetCyclicaY()
	m := l.GetLunarMonth()
	d := l.GetLunarDay()
	h := l.GetLunlarHour()
	// 上挂

	var upNum = y + m + d
	var downNum = y + m + d + h

	return GetFinalGua(upNum,downNum)
}

func GetChongGua(upNum, downNum int) *ChongGua{
	cGua := ChongGua{Name:chongGuaNames[upNum-1][downNum-1]}
	upGuaNum := GetGuaByNum(upNum)
	downGuaNum := GetGuaByNum(downNum)
	// 获取上挂
	upGua := GetDanGuaByGuaNum(upGuaNum)
	downGua := GetDanGuaByGuaNum(downGuaNum)
	cGua.UpGua = *upGua
	cGua.DownGua = *downGua
	// 获取下挂
	return &cGua
}

func GetFinalGua(upNum,downNum int) *FinalGua{
	f := FinalGua{}
	chongGua := GetChongGua(upNum,downNum)
	f.ZhuGua = *chongGua
	return &f
}

func GetDanGuaByGuaNum(guaNum int) *Gua{
	gua := Gua{Name:baguaNames[GetGuaByNum(guaNum)]}
	return &gua
}

func GetGuaByNum(num int) int{
	if num % 8 == 0{
		return 8;
	}
	return num
}


func GetDongYaoNum(num int) int{
	if num % 6 == 0{
		return 6;
	}
	return num
}

