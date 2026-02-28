package main

var months = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func getMonthByNumber(month int) string {
	if month < 1 || month > 12 {
		return ""
	}
	return months[month-1]
}
