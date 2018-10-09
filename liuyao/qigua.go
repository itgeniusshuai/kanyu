package liuyao

import (
	"../date"
	"github.com/andlabs/ui"
)

var bagua = []string{"乾","坎","艮","震","巽","离","坤","兑"}

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

	return GetChongGua(upNum,downNum)
}

func GetChongGua(upNum, downNum int) *FinalGua{
	upGuaNum := GetGuaByNum(upNum)
	downGuaNum := GetGuaByNum(downNum)
	// 获取上挂
	upGua := GetDanGuaByGuaNum(upGuaNum)
	downGua := GetDanGuaByGuaNum(downGuaNum)
	// 获取下挂
	return GetFinalGua(*upGua,*downGua)
}

func GetFinalGua(upGua,downGua Gua) *FinalGua{
	f := FinalGua{}

	return &f
}

func GetDanGuaByGuaNum(guaNum int) *Gua{
	gua := Gua{}

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