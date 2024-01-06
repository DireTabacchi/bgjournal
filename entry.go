package main

import(
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Entry struct {
    TimeDate
    BgLevel int
    InsulinAmount, BasalInsulinAmount int
    BasalInsulinUsed bool
}

type TimeDate struct {
    Hour, Minute int
    Year, Month, Day int
}

// writeEntry takes an Entry and writes it to a file. The filename is the Hour,
// and Minute of the Entry, and is found in the Year/Month/Day directory.
func writeEntryFile(e Entry) error {
    currentDir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("writeEntry: error getting directory: %q", err)
    }

    if err = changeEntriesDir(); err != nil {
        return err
    }

    yearDirName := strconv.FormatInt(int64(e.Year), 10)
    if _, err = os.ReadDir(yearDirName); err != nil {
        os.Mkdir(yearDirName, 0766)
    }
    os.Chdir(yearDirName)

    monthDirName := formatTime(e.Month)
    if _, err = os.ReadDir(monthDirName); err != nil {
        os.Mkdir(monthDirName, 0766)
    }
    os.Chdir(monthDirName)

    dayDirName := formatTime(e.Day)
    if _, err = os.ReadDir(dayDirName); err != nil {
        os.Mkdir(dayDirName, 0766)
    }
    os.Chdir(dayDirName)

    fileName := formatTime(e.Hour) +
        formatTime(e.Minute)

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

// readEntryFile will find the Entry file associated with the given year,
// month, day, hour, and minute, read it, and return an Entry with that file's
// information.
func readEntryFile(year, month, day, hour, minute int) (Entry, error) {
    currentDir, err := os.Getwd()
    if err != nil {
        return Entry{}, fmt.Errorf("readEntryFile: error getting directory: %q",
            err)
    }

    if err = changeEntriesDir(); err != nil {
        return Entry{}, err
    }
    yearDirName := strconv.FormatInt(int64(year), 10)
    monthDirName := formatTime(month)
    dayDirName := formatTime(day)
    fileName := formatTime(hour) + formatTime(minute)
    
    if err := os.Chdir(yearDirName); err != nil {
        return Entry{}, fmt.Errorf("Error finding year directory: %q", err)
    }

    if err := os.Chdir(monthDirName); err != nil {
        return Entry{}, fmt.Errorf("Error finding month directory: %q", err)
    }

    if err := os.Chdir(dayDirName); err != nil {
        return Entry{}, fmt.Errorf("Error finding day directory: %q", err)
    }
    
    contents, err := os.ReadFile(fileName)
    if err != nil {
        return Entry{}, fmt.Errorf("Error finding file %s: %q", fileName, err)
    }

    var entry Entry
    err = json.Unmarshal(contents, &entry)
    if err != nil {
        return Entry{}, fmt.Errorf("Error unmarhalling json: %q", err)
    }

    os.Chdir(currentDir)
    return entry, nil
}

// Given an integer, return a string representing the numeric representation of
// that day/month/time-component. Numbers < 10 get a prepended "0".
func formatTime(m int) string {
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

func printEntry(e Entry) {
    fmt.Printf("\n========================\n")
    fmt.Printf("%d-%s-%s %s:%s\n",
        e.Year,
        formatTime(e.Month),
        formatTime(e.Day),
        formatTime(e.Hour),
        formatTime(e.Minute),
    )
    fmt.Println("----------------")
    fmt.Printf("Blood Glucose Level: %d mg/dL\n", e.BgLevel)
    fmt.Printf("Insulin taken: %d\n", e.InsulinAmount)
    if e.BasalInsulinUsed {
        fmt.Printf("Basal Insulin taken: %d\n", e.BasalInsulinAmount)
    }
    fmt.Printf("========================\n\n")
}

func printDay(entries []Entry) {
    year := entries[0].Year
    month := formatTime(entries[0].Month)
    day := formatTime(entries[0].Day)
    average := bgAverage(entries)

    fmt.Printf("\n========================================================" + 
        "=========\n")
    fmt.Printf("%d-%s-%s\n", year, month, day)
    fmt.Println("----------")
    fmt.Printf("Time - Blood Glucose Level >> " +
        "Insulin Taken (Basal Insulin Taken)\n")
    for _, entry := range entries {
        fmt.Printf("%s:%s - %d mg/dL >> %d",
        formatTime(entry.Hour), formatTime(entry.Minute), entry.BgLevel,
            entry.InsulinAmount)
        if entry.BasalInsulinUsed {
            fmt.Printf(" (%d)\n", entry.BasalInsulinAmount)
        } else {
            fmt.Println()
        }
    }

    fmt.Printf("\nDay Average Blood Glucose Level: %.2f mg/dL\n", average)
    fmt.Printf("============================================================" + 
        "=====\n\n")
}

func printWeek(entries []Entry) {
    average := bgAverage(entries)
    fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
        "~~~~~~~~~\n")

    for day := entries[0].Day; day < entries[0].Day + 7 && day <= 31; day++ {
        var dayEntries []Entry
        //dayEntries = make([]Entry, 0, 3)
        for _, entry := range entries {
            if entry.Day == day {
                dayEntries = append(dayEntries, entry)
            }
        }
        if len(dayEntries) > 0 {
            printDay(dayEntries)
        }
    }

    fmt.Printf("Week Average Blood Glucose Level: %.2f mg/dL\n", average)
    fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
        "~~~~~~~~~\n\n")
}

// readEntryFilesN reads n days of entry files and returns a slice of Entry
// read from those files. nDays includes the day given.
func readEntryFilesNDays(year, month, day, nDays int) ([]Entry, error) {
    currentDir, err := os.Getwd()
    if err != nil {
        return []Entry{}, fmt.Errorf("readEntryFile: error getting directory: %q",
            err)
    }

    if err = changeEntriesDir(); err != nil {
        return []Entry{}, err
    }
    yearDirName := strconv.FormatInt(int64(year), 10)
    monthDirName := formatTime(month)
    
    if err := os.Chdir(yearDirName); err != nil {
        return []Entry{}, fmt.Errorf("Error finding year directory: %q", err)
    }

    if err := os.Chdir(monthDirName); err != nil {
        return []Entry{}, fmt.Errorf("Error finding month directory: %q", err)
    }

    var entries []Entry

    for i := day; i < day + nDays && i <= 31; i++ {
        dayDirName := formatTime(i)
        dayEntries, err := os.ReadDir(dayDirName)
        if err != nil {
            if strings.Contains(err.Error(), "no such file or directory") {
                continue
            } else {
                return []Entry{}, err
            }
        }
        os.Chdir(dayDirName)
        for _, entry := range dayEntries {
            var newEntry Entry
            contents, _ := os.ReadFile(entry.Name())
            json.Unmarshal(contents, &newEntry)
            entries = append(entries, newEntry)
        }
        os.Chdir("..")
    }
    
    os.Chdir(currentDir)
    return entries, nil

}

func parseFileName(name string) (hour int, minute int) {
    tmp, _ := strconv.ParseInt(name, 10, 0)
    hour = int(tmp) / 100
    minute = int(tmp) - (hour * 100)
    return hour, minute
}

func formatDayPath(year, month, day int) string {
    currentDir, _ := os.Getwd()
    if err := changeEntriesDir(); err != nil {
        return ""
    }
    dayPath := strconv.FormatInt(int64(year), 10) + "/" +
        formatTime(month) + "/" +
        formatTime(day)
    if err := os.Chdir(dayPath); err != nil {
        return ""
    }

    os.Chdir(currentDir)
    return dayPath
}
