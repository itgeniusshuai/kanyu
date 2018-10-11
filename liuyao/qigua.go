package liuyao

import (
	"../date"
	"github.com/itgeniusshuai/go_common/common"
	"fmt"
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
	gua := Gua{Name:BaguaNames[GetGuaByNum(guaNum)]}
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


