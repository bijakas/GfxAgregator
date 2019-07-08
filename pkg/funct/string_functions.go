package funct

import (
    "fmt"
    "strconv"
    "encoding/json"
)
func JsonEscape(i string) string {
    b,err :=json.Marshal(i)
    if err != nil {
        panic(err)
    }
    s:= string(b)
    return s[1:len(s)-1]
} 

func FormatMinutesDuration (inMinutes string) string {
    intMinutes, _ := strconv.Atoi(inMinutes)
    hour := intMinutes / 60
    minute := intMinutes % 60
    str := fmt.Sprintf("%02d:%02d:00", hour, minute)
    return str
}