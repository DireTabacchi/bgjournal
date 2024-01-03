package main

import (
    "fmt"
    "os"
    "strconv"
)

func dailyAverage(year, month, day int) (float64, error) {
    currentDir, _ := os.Getwd()
    if err := changeEntriesDir(); err != nil {
        return 0, err
    }
    dayPath := strconv.FormatInt(int64(year), 10) + "/" +
        formatTime(month) + "/" +
        formatTime(day)
    if err := os.Chdir(dayPath); err != nil {
        return 0, err
    }

    dayEntries, err := os.ReadDir(".")
    if err != nil {
        fmt.Fprintf(os.Stderr, "An error occurred reading %s directory: %q\n",
        dayPath, err)
    }

    var entriesCount, bgSum int
    var bgAverage float64

    for _, entryFile := range dayEntries {
        entriesCount++
        hour, minute := parseFileName(entryFile.Name())
        entry, err := readEntryFile(year, month, day, hour, minute)
        if err != nil {
            return 0.0, err
        }
        bgSum += entry.BgLevel
    }

    bgAverage = float64(bgSum) / float64(entriesCount)

    os.Chdir(currentDir)
    return bgAverage, nil
}
