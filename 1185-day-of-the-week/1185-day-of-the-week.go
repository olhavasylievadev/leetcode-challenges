import(
    "time"
    "fmt"
    "strconv"
)

func dayOfTheWeek(day int, month int, year int) string {
    layoutISO := "2006-01-02"
    
    var m string
    var d string 
    
    if month < 10 {
        m = fmt.Sprintf("0%v", month)
    } else {
        m = strconv.Itoa(month)
    }
    
    if day < 10 {
        d = fmt.Sprintf("0%v", day)
    } else {
        d = strconv.Itoa(day)
    }
    
    date := fmt.Sprintf("%v-%v-%v", year, m, d)
    
    t, err := time.Parse(layoutISO, date)
    if err!= nil {
        return ""
    }
    
    weekday := t.Weekday()
    
    return weekday.String()
}