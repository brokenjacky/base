package times

import (
    "encoding/json"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "log"
    "strconv"
    "time"
)

var (
    dayMap = make(map[string]int)
)

func getTodayStatus(today string) int {
    url := fmt.Sprintf("http://tool.bitefu.net/jiari?d=%s&back=json", today)
    _, body, _ := gorequest.New().Get(url).End()
    log.Println("get today status result", body)
    var result map[string]interface{}
    json.Unmarshal([]byte(body), &result)
    if stat, ok := result[today]; ok {
        switch stat.(type) {
        case string:
            s, _ := strconv.Atoi(stat.(string))
            return s
        case int:
            return stat.(int)
        case float64:
            return int(stat.(float64))
        }
        return 0
    } else {
        return 0
    }

}

func IsWorkDay(t time.Time) bool {

    startStr := fmt.Sprintf("%d-%02d-%02d 09:25:00+08:00", t.Year(), t.Month(), t.Day())
    endStr := fmt.Sprintf("%d-%02d-%02d 11:30:00+08:00", t.Year(), t.Month(), t.Day())

    startStr1 := fmt.Sprintf("%d-%02d-%02d 13:00:00+08:00", t.Year(), t.Month(), t.Day())
    endStr1 := fmt.Sprintf("%d-%02d-%02d 15:00:00+08:00", t.Year(), t.Month(), t.Day())

    start, _ := time.Parse("2006-01-02 15:04:05Z07:00", startStr)
    end, _ := time.Parse("2006-01-02 15:04:05Z07:00", endStr)

    start1, _ := time.Parse("2006-01-02 15:04:05Z07:00", startStr1)
    end1, _ := time.Parse("2006-01-02 15:04:05Z07:00", endStr1)
    if !((t.Unix() > start.Unix() && t.Unix() < end.Unix()) || (t.Unix() > start1.Unix() && t.Unix() < end1.Unix())) {
        return false
    }

    today := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
    var status int
    if stat, ok := dayMap[today]; ok {
        status = stat
    } else {
        status = getTodayStatus(today)
        dayMap[today] = status
    }

    if status != 0 {
        return false
    }
    weekday := t.Weekday()
    if weekday == 0 || weekday == 6 {
        return false
    }

    return true
}
