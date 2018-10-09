package liuyao

import (
	"../date"
)

var bagua = []string{"乾","兑","离","震","巽","坎","艮","坤"}
var chongGua = [][]string{	{"乾为天","天泽履","天火同人","天雷无妄","天风姤","天水讼","天山遁","天地否"},
							{"泽天夬","兑为泽","泽火革","泽雷随","泽风大过","泽水困","泽山咸","泽地萃"},
							{"火天大有","火泽睽","离为火","火雷噬嗑","火风井","火水未既","火山旅","火地晋"},
							{"雷天大壮","雷泽归妹","离","震","巽","坎","艮","坤"},
							{"乾","兑","离","震","巽","坎","艮","坤"},
							{"乾","兑","离","震","巽","坎","艮","坤"},
							{"乾","兑","离","震","巽","坎","艮","坤"},
							{"乾","兑","离","震","巽","坎","艮","坤"}}
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
	cGua := ChongGua{}
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
	gua := Gua{Name:bagua[GetGuaByNum(guaNum)]}
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

