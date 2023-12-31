package main

import(
    "encoding/json"
    "fmt"
    "os"
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

// writeEntry takes an Entry and writes it to a file. The filename is the Year,
// Month, and Day of the TimeAndDate in the Entry.
func writeEntryFile(e Entry) error {
    currentDir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("writeEntry: error getting directory: %q", err)
    }

    if err = changeEntriesDir(); err != nil {
        return err
    }

    var fileNameParts []string
    fileNameParts = append(fileNameParts,
    strconv.FormatInt(int64(e.TimeAndDate.Year), 10))
    fileNameParts = append(fileNameParts, formatDayMonth(e.TimeAndDate.Month))
    fileNameParts = append(fileNameParts, formatDayMonth(e.TimeAndDate.Day))
    fileName := strings.Join(fileNameParts, "")

    jsonEntry, err := json.Marshal(e)
    if err != nil {
        return fmt.Errorf("Error marshalling to JSON: %q", err)
    }

    err = os.WriteFile(fileName, jsonEntry, 0666)
    if err != nil {
        return fmt.Errorf("Error writing entry file: %q", err)
    }

    os.Chdir(currentDir)
    return nil
}

func formatDayMonth(m int) string {
    if m < 10 {
        return fmt.Sprintf("0%d", m)
    }
    return fmt.Sprintf("%d", m)
}

// changeEntriesDir changes cwd to the $HOME/bgjournal/entries directory.
// If bgjournal/entries does not exist, changeEntriesDir attempts to create
// it. Returns error on any errors os reports.
func changeEntriesDir() error {
    userHomeDir, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf("writeEntry: error getting user home dir: %q", err)
    }
    
    err = os.Chdir(userHomeDir)
    if err != nil {
        return fmt.Errorf("Error changing to user home dir: %q", err)
    }

    _, err = os.ReadDir("bgjournal")
    if err != nil &&
        err.Error() == "open bgjournal: no such file or directory"  {
        if err = os.MkdirAll("bgjournal/entries", 0766); err != nil {
            return fmt.Errorf("Error creating bgjournal/entries: %q", err)
        }
    } else if err != nil {
        return fmt.Errorf("Error finding bgjournal directory: %q", err)
    }

    err = os.Chdir("bgjournal/entries")
    if err != nil {
        return fmt.Errorf("Error changing to bgjournal/entries: %q", err)
    }

    return nil
}
