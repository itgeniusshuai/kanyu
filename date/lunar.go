package date

import (
	"github.com/itgeniusshuai/go_common/common"
	"regexp"
	"time"
	"fmt"
)

type Lunar struct {
	solar time.Time;
	lunarYear     int;
	lunarMonth    int;
	lunarDay      int;
	isLeap       bool;
	isLeapYear   bool;
	solarYear     int;
	solarMonth    int;
	solarDay      int;
	cyclicalYear  int;
	cyclicalMonth int;
	cyclicalDay   int;
	maxDayInMonth int;

	isFinded         bool;
	isSFestival      bool;
	isLFestival      bool;
	sFestivalName   string;
	lFestivalName   string;
	description     string;
	isHoliday       bool;

}

var lunarInfo = []int{
        0x4bd8, 0x4ae0, 0xa570, 0x54d5, 0xd260, 0xd950, 0x5554, 0x56af,  
        0x9ad0, 0x55d2, 0x4ae0, 0xa5b6, 0xa4d0, 0xd250, 0xd295, 0xb54f,  
        0xd6a0, 0xada2, 0x95b0, 0x4977, 0x497f, 0xa4b0, 0xb4b5, 0x6a50,  
        0x6d40, 0xab54, 0x2b6f, 0x9570, 0x52f2, 0x4970, 0x6566, 0xd4a0,  
        0xea50, 0x6a95, 0x5adf, 0x2b60, 0x86e3, 0x92ef, 0xc8d7, 0xc95f,  
        0xd4a0, 0xd8a6, 0xb55f, 0x56a0, 0xa5b4, 0x25df, 0x92d0, 0xd2b2,  
        0xa950, 0xb557, 0x6ca0, 0xb550, 0x5355, 0x4daf, 0xa5b0, 0x4573,  
        0x52bf, 0xa9a8, 0xe950, 0x6aa0, 0xaea6, 0xab50, 0x4b60, 0xaae4,  
        0xa570, 0x5260, 0xf263, 0xd950, 0x5b57, 0x56a0, 0x96d0, 0x4dd5,  
        0x4ad0, 0xa4d0, 0xd4d4, 0xd250, 0xd558, 0xb540, 0xb6a0, 0x95a6,  
        0x95bf, 0x49b0, 0xa974, 0xa4b0, 0xb27a, 0x6a50, 0x6d40, 0xaf46,  
        0xab60, 0x9570, 0x4af5, 0x4970, 0x64b0, 0x74a3, 0xea50, 0x6b58,  
        0x5ac0, 0xab60, 0x96d5, 0x92e0, 0xc960, 0xd954, 0xd4a0, 0xda50,  
        0x7552, 0x56a0, 0xabb7, 0x25d0, 0x92d0, 0xcab5, 0xa950, 0xb4a0,  
        0xbaa4, 0xad50, 0x55d9, 0x4ba0, 0xa5b0, 0x5176, 0x52bf, 0xa930,  
        0x7954, 0x6aa0, 0xad50, 0x5b52, 0x4b60, 0xa6e6, 0xa4e0, 0xd260,  
        0xea65, 0xd530, 0x5aa0, 0x76a3, 0x96d0, 0x4afb, 0x4ad0, 0xa4d0,  
        0xd0b6, 0xd25f, 0xd520, 0xdd45, 0xb5a0, 0x56d0, 0x55b2, 0x49b0,  
        0xa577, 0xa4b0, 0xaa50, 0xb255, 0x6d2f, 0xada0, 0x4b63, 0x937f,  
        0x49f8, 0x4970, 0x64b0, 0x68a6, 0xea5f, 0x6b20, 0xa6c4, 0xaaef,  
        0x92e0, 0xd2e3, 0xc960, 0xd557, 0xd4a0, 0xda50, 0x5d55, 0x56a0,  
        0xa6d0, 0x55d4, 0x52d0, 0xa9b8, 0xa950, 0xb4a0, 0xb6a6, 0xad50,  
        0x55a0, 0xaba4, 0xa5b0, 0x52b0, 0xb273, 0x6930, 0x7337, 0x6aa0,  
        0xad50, 0x4b55, 0x4b6f, 0xa570, 0x54e4, 0xd260, 0xe968, 0xd520,  
        0xdaa0, 0x6aa6, 0x56df, 0x4ae0, 0xa9d4, 0xa4d0, 0xd150, 0xf252, 0xd520,
    };  
    var solarTermInfo = []int{  
        0, 21208, 42467, 63836, 85337, 107014, 128867, 150921,  
        173149, 195551, 218072, 240693, 263343, 285989, 308563, 331033,  
        353350, 375494, 397447, 419210, 440795, 462224, 483532, 504758,
    };  
    var Tianan = []string{  
        "甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"};
    var Deqi = []string{
        "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"};
    var Animals = []string{
        "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"};
    var solarTerm = []string{
        "小寒", "大寒", "立春", "雨水", "惊蛰", "春分",  
        "清明", "谷雨", "立夏", "小满", "芒种", "夏至",  
        "小暑", "大暑", "立秋", "处暑", "白露", "秋分",  
        "寒露", "霜降", "立冬", "小雪", "大雪", "冬至"};
    var lunarString1 = []string{
        "零", "一", "二", "三", "四", "五", "六", "七", "八", "九"};
    var lunarString2 = []string{
        "初", "十", "廿", "卅", "正", "腊", "冬", "闰"};
    /** 
     * 国历节日 *表示放假日 
     */  
    var sFtv = []string{
        "0101*元旦", "0214 情人节", "0308 妇女节", "0312 植树节",  
        "0315 消费者权益日", "0401 愚人节", "0501*劳动节", "0504 青年节",  
        "0509 郝维节", "0512 护士节", "0601 儿童节", "0701 建党节 香港回归纪念",  
        "0801 建军节", "0808 父亲节", "0816 燕衔泥节", "0909 毛泽东逝世纪念",  
        "0910 教师节", "0928 孔子诞辰", "1001*国庆节", "1006 老人节",  
        "1024 联合国日", "1111 光棍节", "1112 孙中山诞辰纪念", "1220 澳门回归纪念",  
        "1225 圣诞节", "1226 毛泽东诞辰纪念"};
    /** 
     * 农历节日 *表示放假日 
     */  
    var lFtv = []string{
        "0101*春节、弥勒佛诞", "0106 定光佛诞", "0115 元宵节",  
        "0208 释迦牟尼佛出家", "0215 释迦牟尼佛涅槃", "0209 海空上师诞",  
        "0219 观世音菩萨诞", "0221 普贤菩萨诞", "0316 准提菩萨诞",  
        "0404 文殊菩萨诞", "0408 释迦牟尼佛诞", "0415 佛吉祥日——释迦牟尼佛诞生、成道、涅槃三期同一庆(即南传佛教国家的卫塞节)",  
        "0505 端午节", "0513 伽蓝菩萨诞", "0603 护法韦驮尊天菩萨诞",  
        "0619 观世音菩萨成道——此日放生、念佛，功德殊胜",  
        "0707 七夕情人节", "0713 大势至菩萨诞", "0715 中元节",  
        "0724 龙树菩萨诞", "0730 地藏菩萨诞", "0815 中秋节",  
        "0822 燃灯佛诞", "0909 重阳节", "0919 观世音菩萨出家纪念日",  
        "0930 药师琉璃光如来诞", "1005 达摩祖师诞", "1107 阿弥陀佛诞",  
        "1208 释迦如来成道日，腊八节", "1224 小年",  
        "1229 华严菩萨诞", "0100*除夕"};
    /** 
     * 某月的第几个星期几 
     */  
    var wFtv = []string{
        "0520 母亲节", "0716 合作节", "0730 被奴役国家周"};
  
    func toInt(str string) int{
    	i,err := common.StrToInt(str)
    	if err != nil{
    		return -1
		}
       return i
    }
    var sFreg =regexp.MustCompile("^(\\d{2})(\\d{2})([\\s\\*])(.+)$")
    var wFreg = regexp.MustCompile("^(\\d{2})(\\d)(\\d)([\\s\\*])(.+)$")
  
    func (this *Lunar)findFestival() {
        //var sM = this.solarMonth;
        //var sD = this.solarDay;
        //var lM = this.solarMonth;
        //var lD = this.lunarDay;
        //var sy = this.solarYear;
        //for i := 0; i < len(sFtv); i++ {
        //	m,_ := sFreg.MatchString(sFreg,sFtv[i])
        //    m = sFreg.matcher(Lunar.sFtv[i]);
        //    if (m.find()) {
        //        if (sM == toInt(m.group(1)) && sD == toInt(m.group(2))) {
        //            this.isSFestival = true;
        //            this.sFestivalName = m.group(4);
        //            if ("*".equals(m.group(3))) {
        //                this.isHoliday = true;
        //            }
        //            break;
        //        }
        //    }
        //}
        //for i := 0; i < len(lFtv); i++ {
        //    m = Lunar.sFreg.matcher(Lunar.lFtv[i]);
        //    if (m.find()) {
        //        if (lM == Lunar.toInt(m.group(1)) && lD == Lunar.toInt(m.group(2))) {
        //            this.isLFestival = true;
        //            this.lFestivalName = m.group(4);
        //            if ("*".equals(m.group(3))) {
        //                this.isHoliday = true;
        //            }
        //            break;
        //        }
        //    }
        //}
		//
        //// 月周节日
        //int w, d;
        //for (int i = 0; i < Lunar.wFtv.length; i++) {
        //    m = Lunar.wFreg.matcher(Lunar.wFtv[i]);
        //    if (m.find()) {
        //        if (this.getSolarMonth() == Lunar.toInt(m.group(1))) {
        //            w = Lunar.toInt(m.group(2));
        //            d = Lunar.toInt(m.group(3));
        //            if (this.solar.get(Calendar.WEEK_OF_MONTH) == w
        //                    && this.solar.get(Calendar.DAY_OF_WEEK) == d) {
        //                this.isSFestival = true;
        //                this.sFestivalName += "|" + m.group(5);
        //                if ("*".equals(m.group(4))) {
        //                    this.isHoliday = true;
        //                }
        //            }
        //        }
        //    }
        //}
        //if (sy > 1874 && sy < 1909) {
        //    this.description = "光绪" + (((sy - 1874) == 1) ? "元" : "" + (sy - 1874));
        //}
        //if (sy > 1908 && sy < 1912) {
        //    this.description = "宣统" + (((sy - 1908) == 1) ? "元" : String.valueOf(sy - 1908));
        //}
        //if (sy > 1911 && sy < 1950) {
        //    this.description = "民国" + (((sy - 1911) == 1) ? "元" : String.valueOf(sy - 1911));
        //}
        //if (sy > 1949) {
        //    this.description = "共和国" + (((sy - 1949) == 1) ? "元" : String.valueOf(sy - 1949));
        //}
        //this.description += "年";
        //this.sFestivalName = this.sFestivalName.replaceFirst("^\\|", "");
        this.isFinded = true;
    }

  
    /** 
     * 返回农历年闰月月份 
     * 
     * @param lunarYear 指定农历年份(数字) 
     * @return 该农历年闰月的月份(数字,没闰返回0) 
     */  
    func getLunarLeapMonth(lunarYear int) int{
        // 数据表中,每个农历年用16bit来表示,     
        // 前12bit分别表示12个月份的大小月,最后4bit表示闰月     
        // 若4bit全为1或全为0,表示没闰, 否则4bit的值为闰月月份     
        var leapMonth int = lunarInfo[lunarYear - 1900] & 0xf;
        if leapMonth == 0xf{
        	return 0
		}
        return leapMonth;
    }  
  
    /** 
     * 返回农历年闰月的天数 
     * 
     * @param lunarYear 指定农历年份(数字) 
     * @return 该农历年闰月的天数(数字) 
     */  
    func getLunarLeapDays(lunarYear int) int{
        // 下一年最后4bit为1111,返回30(大月)     
        // 下一年最后4bit不为1111,返回29(小月)     
        // 若该年没有闰月,返回0
        if getLunarLeapMonth(lunarYear) > 0{
        	if (lunarInfo[lunarYear - 1899] & 0xf) == 0xf{
        		return 30
			}
			return 29
		}
        return 0
    }  
  
    /** 
     * 返回农历年的总天数 
     * 
     * @param lunarYear 指定农历年份(数字) 
     * @return 该农历年的总天数(数字) 
     */  
    func getLunarYearDays(lunarYear int) int{
        // 按小月计算,农历年最少有12 * 29 = 348天     
        var daysInLunarYear = 348;
        // 数据表中,每个农历年用16bit来表示,     
        // 前12bit分别表示12个月份的大小月,最后4bit表示闰月     
        // 每个大月累加一天     
        for  i := 0x8000; i > 0x8; i >>= 1 {
        	if (lunarInfo[lunarYear - 1900] & i) != 0{
				daysInLunarYear += 1
			}else{
				daysInLunarYear += 0
			}
		}
        // 加上闰月天数
        daysInLunarYear += getLunarLeapDays(lunarYear);
  
        return daysInLunarYear;  
    }  
  
    /** 
     * 返回农历年正常月份的总天数 
     * 
     * @param lunarYear 指定农历年份(数字) 
     * @param lunarMonth 指定农历月份(数字) 
     * @return 该农历年闰月的月份(数字,没闰返回0) 
     */  
    func getLunarMonthDays( lunarYear, lunarMonth int) int{
        // 数据表中,每个农历年用16bit来表示,     
        // 前12bit分别表示12个月份的大小月,最后4bit表示闰月
        lunarMonthU := uint(lunarMonth)
        if (lunarInfo[lunarYear - 1900] & (0x10000 >> lunarMonthU)) != 0{
        	return 30
		}else{
			return 29
		}
    }
  
    /** 
     * 取 Date 对象中用全球标准时间 (UTC) 表示的日期 
     * 
     * @param date 指定日期 
     * @return UTC 全球标准时间 (UTC) 表示的日期 
     */  
    func getUTCDay(time time.Time) int{
    	return time.YearDay()
    }  

  
    /** 
     * 返回全球标准时间 (UTC) (或 GMT) 的 1970 年 1 月 1 日到所指定日期之间所间隔的毫秒数。 
     * 
     * @param y 指定年份 
     * @param m 指定月份 
     * @param d 指定日期 
     * @param h 指定小时 
     * @param min 指定分钟 
     * @param sec 指定秒数 
     * @return 全球标准时间 (UTC) (或 GMT) 的 1970 年 1 月 1 日到所指定日期之间所间隔的毫秒数 
     */  
    func UTC( y,  m,  d,  h,  min,  sec int)int64 {
		month := time.Month(m)
    	t := time.Date(y,month,d,h,min,sec,0,time.UTC)
    	return t.Unix()
    }
  
    /** 
     * 返回公历年节气的日期 
     * 
     * @param solarYear 指定公历年份(数字) 
     * @param index 指定节气序号(数字,0从小寒算起) 
     * @return 日期(数字,所在月份的第几天) 
     */  
    func getSolarTermDay( solarYear, index int) int{
  
        return getUTCDay(getSolarTermCalendar(solarYear, index));
    }  
  
    /** 
     * 返回公历年节气的日期 
     * 
     * @param solarYear 指定公历年份(数字) 
     * @param index 指定节气序号(数字,0从小寒算起) 
     * @return 日期(数字,所在月份的第几天) 
     */  
    func getSolarTermCalendar( solarYear,  index int) time.Time{
        var l = (int64)( 31556925974.7 * float32((solarYear - 1900)+ solarTermInfo[index] * 60000));
        l = l/1000 + UTC(1900, 0, 6, 2, 5, 0);
        return time.Unix(l,0);
    }
  
    func(this *Lunar) init(TimeInMillis int64) {
    	this.maxDayInMonth = 29
        this.solar = time.Unix(TimeInMillis/1000,0);
        var baseDate = time.Date(1900, 1, 31,0,0,0,0,time.UTC);
        var offset = (TimeInMillis - baseDate.Unix()*1000) / 86400000;
        // 按农历年递减每年的农历天数，确定农历年份     
        this.lunarYear = 1900;  
        var daysInLunarYear = getLunarYearDays(this.lunarYear);
        for (this.lunarYear < 2100 && offset >= int64(daysInLunarYear)) {
            offset -= int64(daysInLunarYear );
			this.lunarYear += 1
            daysInLunarYear = getLunarYearDays(this.lunarYear);
        }  
        // 农历年数字     
  
        // 按农历月递减每月的农历天数，确定农历月份     
        var lunarMonth = 1;
        // 所在农历年闰哪个月,若没有返回0     
        var leapMonth = getLunarLeapMonth(this.lunarYear);
        // 是否闰年     
        this.isLeapYear = leapMonth > 0;  
        // 闰月是否递减     
        var leapDec = false;
        var isLeap = false;
        var daysInLunarMonth = 0;
        for (lunarMonth < 13 && offset > 0) {
            if (isLeap && leapDec) { // 如果是闰年,并且是闰月     
                // 所在农历年闰月的天数     
                daysInLunarMonth = getLunarLeapDays(this.lunarYear);
                leapDec = false;  
            } else {  
                // 所在农历年指定月的天数     
                daysInLunarMonth = getLunarMonthDays(this.lunarYear, lunarMonth);
            }  
            if (offset < int64(daysInLunarMonth)) {
                break;  
            }  
            offset -= int64(daysInLunarMonth);
  
            if (leapMonth == lunarMonth && isLeap == false) {  
                // 下个月是闰月     
                leapDec = true;  
                isLeap = true;  
            } else {  
                // 月份递增     
                lunarMonth++;  
            }  
        }  
        this.maxDayInMonth = daysInLunarMonth;  
        // 农历月数字     
        this.lunarMonth = lunarMonth;  
        // 是否闰月     
        this.isLeap = (lunarMonth == leapMonth && isLeap);  
        // 农历日数字     
        this.lunarDay = (int)( offset + 1);
        // 取得干支历     
        this.GetCyclicalData();
    }  
  
    /** 
     * 取干支历 不是历年，历月干支，而是中国的从立春节气开始的节月，是中国的太阳十二宫，阳历的。 
     * 
     * @param cncaData 日历对象(Tcnca) 
     */  
    func(this *Lunar) GetCyclicalData() {
        this.solarYear = this.solar.Year();
        this.solarMonth = (int)(this.solar.Month());
        this.solarDay = this.solar.Day();
        // 干支历     
        var cyclicalYear = 0;
        var cyclicalMonth = 0;
        var cyclicalDay = 0;
  
        // 干支年 1900年立春後为庚子年(60进制36)     
        var term2 = getSolarTermDay(this.solarYear, 2); // 立春日期
        // 依节气调整二月分的年柱, 以立春为界     
        if (this.solarMonth < 1 || (this.solarMonth == 1 && this.solarDay < term2)) {
            cyclicalYear = (this.solarYear - 1900 + 36 - 1) % 60;
        } else {  
            cyclicalYear = (this.solarYear - 1900 + 36) % 60;
        }  
  
        // 干支月 1900年1月小寒以前为 丙子月(60进制12)     
        var firstNode = getSolarTermDay(this.solarYear, this.solarMonth * 2); // 传回当月「节」为几日开始
        // 依节气月柱, 以「节」为界     
        if (this.solarDay < firstNode) {
            cyclicalMonth = ((this.solarYear - 1900) * 12 + this.solarMonth + 12 ) % 60;
        } else {  
            cyclicalMonth = ((this.solarYear - 1900) * 12 + this.solarMonth + 13 ) % 60;
        }  
  
        // 当月一日与 1900/1/1 相差天数     
        // 1900/1/1与 1970/1/1 相差25567日, 1900/1/1 日柱为甲戌日(60进制10)     
        cyclicalDay = (int) (UTC(this.solarYear, this.solarMonth, this.solarDay, 0, 0, 0) / 86400 + 25567 + 10) % 60;
        this.cyclicalYear = cyclicalYear;  
        this.cyclicalMonth = cyclicalMonth;  
        this.cyclicalDay = cyclicalDay;  
    }  
  
    /** 
     * 取农历年生肖 
     * 
     * @return 农历年生肖(例:龙) 
     */  
    func(this *Lunar) GetAnimalString() string{
        return Animals[(this.lunarYear - 4) % 12];
    }  
  
    /** 
     * 返回公历日期的节气字符串 
     * 
     * @return 二十四节气字符串,若不是节气日,返回空串(例:冬至) 
     */
	func(this *Lunar) GetTermString() string{
        // 二十四节气     
        var termString = "";
        if (getSolarTermDay(this.solarYear, this.solarMonth * 2) == this.solarDay) {
            termString = solarTerm[this.solarMonth * 2];
        } else if (getSolarTermDay(this.solarYear, this.solarMonth * 2 + 1) == this.solarDay) {
            termString = solarTerm[this.solarMonth * 2 + 1];
        }  
        return termString;  
    }  
  
    /** 
     * 取得干支历字符串 
     * 
     * @return 干支历字符串(例:甲子年甲子月甲子日) 
     */
	func(this *Lunar)  GetCyclicalDateString() string{
        return this.GetCyclicaYear() + "年" + this.GetCyclicaMonth() + "月"+ this.GetCyclicaDay() + "日";
    }  
  
    /** 
     * 年份天干 
     * 
     * @return 年份天干 
     */  
    func(this *Lunar) GetTiananY() int{
        return getTianan(this.cyclicalYear);
    }  
  
    /** 
     * 月份天干 
     * 
     * @return 月份天干 
     */  
    func(this *Lunar) GetTiananM() int{
        return getTianan(this.cyclicalMonth);
    }  
  
    /** 
     * 日期天干 
     * 
     * @return 日期天干 
     */  
    func(this *Lunar) GetTiananD() int {
        return getTianan(this.cyclicalDay);
    }  
  
    /** 
     * 年份地支 
     * 
     * @return 年分地支 
     */  
    func(this *Lunar) GetDeqiY() int{
        return getDeqi(this.cyclicalYear) + 1;
    }  
  
    /** 
     * 月份地支 
     * 
     * @return 月份地支 
     */  
    func(this *Lunar) GetDeqiM() int{
        return getDeqi(this.cyclicalMonth) + 1;
    }  
  
    /** 
     * 日期地支 
     * 
     * @return 日期地支 
     */  
    func(this *Lunar) GetDeqiD() int{
        return getDeqi(this.cyclicalDay) + 1;
    }  
  
    /** 
     * 取得干支年字符串 
     * 
     * @return 干支年字符串 
     */  
    func(this *Lunar) GetCyclicaYear() string{
        return getCyclicalString(this.cyclicalYear);
    }

    /**
     * 获取干支年位
     */
    func (this *Lunar)GetCyclicaY()int{
    	return this.cyclicalYear + 1;
	}

	/**
	 * 获取干支月位
	 */
	func (this *Lunar)GetCyclicaM()int{
		return this.cyclicalMonth + 1;
	}

	/**
	 * 获取干支日位
	 */
	func (this *Lunar)GetCyclicaD()int{
		return this.cyclicalDay + 1;
	}
    /** 
     * 取得干支月字符串 
     * 
     * @return 干支月字符串 
     */  
    func(this *Lunar) GetCyclicaMonth() string{
        return getCyclicalString(this.cyclicalMonth);
    }  
  
    /** 
     * 取得干支日字符串 
     * 
     * @return 干支日字符串 
     */  
    func(this *Lunar) GetCyclicaDay() string{
        return getCyclicalString(this.cyclicalDay);
    }  
  
    /** 
     * 返回农历日期字符串 
     * 
     * @return 农历日期字符串 
     */
    func(this *Lunar) GetLunarDayString() string{
        return getLunarDayString(this.lunarDay);
    }  
  
    /** 
     * 返回农历日期字符串 
     * 
     * @return 农历日期字符串 
     */
	func(this *Lunar) GetLunarMonthString() string{
		if this.isLeap {
			return "闰"+ getLunarMonthString(this.lunarMonth);
		}
        return getLunarMonthString(this.lunarMonth)
    }

    /**
     * 获取农历月
     */
	func(this *Lunar) GetLunarMonth() int{
		return this.lunarMonth
	}

	/**
	 * 获取农历年
	 */
	func(this *Lunar) GetLunarYear() int{
		return this.lunarYear
	}

	/**
	 * 获取农历日
	 */
	func(this *Lunar) GetLunarDay() int{
		return this.lunarDay
	}

/**
 * 返回农历日期字符串
 *
 * @return 农历日期字符串
 */
    func(this *Lunar) GetLunarYearString() string {
        return getLunarYearString(this.lunarYear);
    }  
  
    /** 
     * 返回农历表示字符串 
     * 
     * @return 农历字符串(例:甲子年正月初三) 
     */  
    func(this *Lunar) GetLunarDateString() string{
        return this.GetLunarYearString() + "年"+ this.GetLunarMonthString() + "月"+ this.GetLunarDayString() + "日";
    }  


    /** 
     * 当前农历月是否是大月 
     * 
     * @return 当前农历月是大月 
     */  
    func (this *Lunar)isBigMonth() bool{
        return this.GetMaxDayInMonth() > 29;
    }  
  
    /** 
     * 当前农历月有多少天 
     * 
     * @return 当前农历月有多少天 
     */  
    func(this *Lunar) GetMaxDayInMonth() int{
        return this.maxDayInMonth;  
    }  
  



  
    /** 
     * 公历月份 
     * 
     * @return 公历月份 (不是从0算起) 
     */  
    func(this *Lunar) GetSolarMonth() int{
        return this.solarMonth + 1;
    }  
  

  
    /** 
     * 星期几 
     * 
     * @return 星期几(星期日为:1, 星期六为:7) 
     */  
    func(this *Lunar) GetDayOfWeek() int {
        return int(this.solar.Weekday());
    }  
  
    /** 
     * 黑色星期五 
     * 
     * @return 是否黑色星期五 
     */  
    func(this *Lunar) isBlackFriday() bool{
        return (this.solarDay == 13 && this.solar.Weekday() == 6);
    }  
  
    /** 
     * 是否是今日 
     * 
     * @return 是否是今日 
     */  
    func(this *Lunar) isToday() bool {
        var clr = time.Now();
        return clr.Year() == this.solarYear&& int(clr.Month()) == this.solarMonth && clr.Day() == this.solarDay;
    }  
  

  
    /** 
     * 是否是农历节日 
     * 
     * @return 是否是农历节日 
     */  
    func(this *Lunar) isLFestivalF() bool{
        if (!this.isFinded) {  
            this.findFestival();  
        }  
        return this.isLFestival;  
    }  
  
    /** 
     * 是否是公历节日 
     * 
     * @return 是否是公历节日 
     */  
    func(this *Lunar) isSFestivalF() bool{
        if (!this.isFinded) {  
            this.findFestival();  
        }  
        return this.isSFestival;  
    }  
  
    /** 
     * 是否是节日 
     * 
     * @return 是否是节日 
     */  
    func(this *Lunar) isFestival() bool{
        return this.isSFestivalF() || this.isLFestivalF();
    }  
  
    /** 
     * 是否是放假日 
     * 
     * @return 是否是放假日 
     */  
    func(this *Lunar) isHolidayF() bool {
        if (!this.isFinded) {  
            this.findFestival();  
        }  
        return this.isHoliday;  
    }  
  
    /** 
     * 其它日期说明 
     * 
     * @return 日期说明(如:民国2年) 
     */  
    func (this *Lunar)getDescription() string{
        if (!this.isFinded) {  
            this.findFestival();  
        }  
        return this.description;  
    }  
  
    /** 
     * 干支字符串 
     * 
     * @param cyclicalNumber 指定干支位置(数字,0为甲子) 
     * @return 干支字符串 
     */  
    func getCyclicalString(cyclicalNumber int) string{
        return Tianan[getTianan(cyclicalNumber)] + Deqi[getDeqi(cyclicalNumber)];
    }  
  
    /** 
     * 获得地支 
     * 
     * @param cyclicalNumber 
     * @return 地支 (数字) 
     */  
    func getDeqi( cyclicalNumber int) int{
        return cyclicalNumber % 12;  
    }  
  
    /** 
     * 获得天干 
     * 
     * @param cyclicalNumber 
     * @return 天干 (数字) 
     */  
    func getTianan(cyclicalNumber int) int{
        return cyclicalNumber % 10;  
    }  
  
    /** 
     * 返回指定数字的农历年份表示字符串 
     * 
     * @param lunarYear 农历年份(数字,0为甲子) 
     * @return 农历年份字符串 
     */  
    func getLunarYearString(lunarYear int) string{
        return getCyclicalString(lunarYear - 1900 + 36);
    }  
  
    /** 
     * 返回指定数字的农历月份表示字符串 
     * 
     * @param lunarMonth 农历月份(数字) 
     * @return 农历月份字符串 (例:正) 
     */  
    func getLunarMonthString(lunarMonth int) string{
        var lunarMonthString = "";
        if (lunarMonth == 1) {  
            lunarMonthString = lunarString2[4];
        } else {  
            if (lunarMonth > 9) {  
                lunarMonthString += lunarString2[1];
            }  
            if (lunarMonth % 10 > 0) {  
                lunarMonthString += lunarString1[lunarMonth % 10];
            }  
        }  
        return lunarMonthString;  
    }  
  
    /** 
     * 返回指定数字的农历日表示字符串 
     * 
     * @param lunarDay 农历日(数字) 
     * @return 农历日字符串 (例: 廿一) 
     */  
    func getLunarDayString( lunarDay int) string{
        if (lunarDay < 1 || lunarDay > 30) {  
            return "";  
        }  
        var i1 = lunarDay / 10;
        var i2 = lunarDay % 10;
        var c1 = lunarString2[i1];
        var c2 = lunarString1[i2];
        if (lunarDay < 11) {  
            c1 = lunarString2[0];
        }  
        if (i2 == 0) {  
            c2 = lunarString2[1];
        }  
        return c1 + c2;  
    }


    func GetLunar() Lunar{
    	l := Lunar{}
    	l.init(time.Now().Unix()*1000)
    	fmt.Println(l.solarYear)
    	fmt.Println(l.solarMonth)
    	fmt.Println(l.solarDay)
    	fmt.Println(l.GetCyclicaYear())
    	fmt.Println(l.GetCyclicaMonth())
    	fmt.Println(l.GetCyclicaDay())
    	fmt.Println(l.GetCyclicalDateString())
    	fmt.Println(l.GetLunarMonthString())
    	fmt.Println(l.GetLunarYearString())
    	fmt.Println(l.GetLunarDateString())
    	fmt.Println(l.GetDeqiY())
    	fmt.Println(l.GetDeqiM())
    	fmt.Println(l.GetDeqiD())
    	fmt.Println(l.GetCyclicaD())
    	fmt.Println(l.GetCyclicaM())
    	fmt.Println(l.GetLunarDay())
    	fmt.Println(l.GetLunarMonth())
    	fmt.Println(l.GetLunarYear())
    	return l
	}

