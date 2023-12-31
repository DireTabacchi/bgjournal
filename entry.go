package main

import(
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
)

type Entry struct {
    TimeAndDate TimeDate
    BgLevel int
    InsulinAmount, BasalInsulinAmount int
    BasalInsulinUsed bool
}

type TimeDate struct {
    Hour, Minute int
    Day, Month, Year int
}

func writeEntry(e Entry) error {
    var fileNameParts []string
    fileNameParts = append(fileNameParts,
    strconv.FormatInt(int64(e.TimeAndDate.Year), 10))
    fileNameParts = append(fileNameParts, formatDayMonth(e.TimeAndDate.Month))
    fileNameParts = append(fileNameParts, formatDayMonth(e.TimeAndDate.Day))
    fileName := strings.Join(fileNameParts, "")
    fmt.Println(fileName)

    //jsonEntry, err := json.Marshal(e)
    jsonEntry, err := json.MarshalIndent(e, "", "  ")
    if err != nil {
        return fmt.Errorf("Error marshalling to JSON: %q", err)
    }
    fmt.Printf("%s\n", jsonEntry)
    return nil
}

func formatDayMonth(m int) string {
    if m < 10 {
        return fmt.Sprintf("0%d", m)
    }
    return fmt.Sprintf("%d", m)
}
