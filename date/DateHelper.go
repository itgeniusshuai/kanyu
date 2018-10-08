package date

import "time"

type YinLiDate struct {
	Year int
	Mouth int
	Day int
	Hour int

	YearZhi int
	MonthZhi int
	DayZhi int
	HourZhi int

	YearGan int
	MonthGan int
	DayGan int
	HourGan int

}

func toYinli(t time.Time) *YinLiDate{

	return nil
}

func getYinliYear(year int)int{
	// year - 1977 = 4Q + R   ---R < 4
	year = year - 1977
	var q int = 0
	var r int = year
	for ;r<4;q++{
		r = year - year / 4
	}
	return r
}
