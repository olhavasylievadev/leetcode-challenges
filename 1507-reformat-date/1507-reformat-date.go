import "strings"

func reformatDate(date string) string {
    myStr := strings.Split(date, " ")
    
    months := map[string]string{
        "Jan": "01",
        "Feb": "02",
        "Mar": "03",
        "Apr": "04",
        "May": "05",
        "Jun": "06",
        "Jul": "07",
        "Aug": "08",
        "Sep": "09",
        "Oct": "10",
        "Nov": "11",
        "Dec": "12",
    }
    
    day, month, year := myStr[0][:len(myStr[0])-2], myStr[1], myStr[2]
    
    if len(day) == 1 {
        day = "0"+day
    }
    return year + "-" + months[month] + "-" + day
}