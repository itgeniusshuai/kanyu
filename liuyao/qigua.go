package liuyao

import (
	"../date"
)


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
	cGua := ChongGua{Name:ChongGuaNames[upNum-1][downNum-1]}
	upGuaNum := GetGuaByNum(upNum)
	downGuaNum := GetGuaByNum(downNum)
	// 获取上挂
	upGua := GetDanGuaByGuaNum(upGuaNum,true)
	downGua := GetDanGuaByGuaNum(downGuaNum,false)
	cGua.UpGua = *upGua
	cGua.DownGua = *downGua
	ParseChongGuaDesc(cGua)
	// 获取下挂
	return &cGua
}

func GetFinalGua(upNum,downNum int) *FinalGua{
	f := FinalGua{}
	chongGua := GetChongGua(upNum,downNum)
	f.ZhuGua = *chongGua
	return &f
}

func GetDanGuaByGuaNum(guaNum int, isUp bool) *Gua{
	return ParseDanGuaDesc(isUp,guaNum)
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


