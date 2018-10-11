package liuyao
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
	初支地支4位
 */
var DanUpGuaDesc = []int{}
var DanDownGuaDesc = []int{}

