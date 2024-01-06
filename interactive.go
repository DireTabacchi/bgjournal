package main

import (
	"bufio"
	"fmt"
    "os"
    "strconv"
    "strings"
)

func printCommands() {
    fmt.Println("[1] Create a new entry")
    fmt.Println("[2] Look for an entry")
    fmt.Println("[3] Stats for a day")
    fmt.Println("[4] Stats for a week")
    fmt.Println("[q] Quit")
}

func printPrompt() {
    fmt.Print("--,^^,*> ")
}

func createEntry() error {
    var newEntry Entry
    var fieldSet bool

    fmt.Println("\nEnter the following to create a new entry.\n" +
        "Enter q to quit.")

    // Get the Year
    for fieldSet = false; !fieldSet; {
        fmt.Print("Year: ")
        year, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if fieldSet {
            newEntry.Year = year
        }
    }

    //Get the Month
    for fieldSet = false; !fieldSet; {
        fmt.Print("Month: ")
        month, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if month < 1 || month > 12 {
            fieldSet = false
            fmt.Println("Please enter a valid month.")
        }
        if fieldSet {
            newEntry.Month = month
        }
    }

    // Get the Day
    for fieldSet = false; !fieldSet; {
        fmt.Print("Day: ")
        day, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if day < 1 || day > 31 {
            fieldSet = false
            fmt.Println("Please enter a valid day.")
        }
        if fieldSet {
            newEntry.Day = day
        }
    }

    // Get the Hour
    for fieldSet = false; !fieldSet; {
        fmt.Print("Hour (24-hour; 0 is midnight): ")
        hour, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if hour < 0 || hour > 23 {
            fieldSet = false
            fmt.Println("Please enter a valid hour.")
        }
        if fieldSet {
            newEntry.Hour = hour
        }
    }

    // Get the Minute
    for fieldSet = false; !fieldSet; {
        fmt.Print("Minute: ")
        minute, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if minute < 0 || minute > 59 {
            fieldSet = false
            fmt.Println("Please enter a valid minute.")
        }
        if fieldSet {
            newEntry.Minute = minute
        }
    }

    for fieldSet = false; !fieldSet; {
        fmt.Print("Blood Glucose level: ")
        bg, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if bg < 0 {
            fieldSet = false
            fmt.Println("Please enter a positive number or zero.")
        }
        if fieldSet {
            newEntry.BgLevel = bg
        }
    }

    for fieldSet = false; !fieldSet; {
        fmt.Print("Units of insulin taken: ")
        insulin, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if insulin < 0 {
            fieldSet = false
            fmt.Println("Please enter a positive number or zero.")
        }
        if fieldSet {
            newEntry.InsulinAmount = insulin
        }
    }

    for fieldSet = false; !fieldSet; {
        in := bufio.NewReader(os.Stdin)
        fmt.Print("Was basal insulin taken? (y,n): ")
        userData, err := in.ReadString('\n')
        if err != nil {
            return fmt.Errorf("An error occurred: %q", err)
        }
        userData = strings.TrimSuffix(userData, "\n")
        if userData == "n" {
            newEntry.BasalInsulinUsed = false
            newEntry.BasalInsulinAmount = 0
            break
        } else if userData == "q" {
            return nil
        }
        fmt.Print("Units of basal insulin: ")
        bInsulin, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        if bInsulin < 1 {
            fieldSet = false
            fmt.Println("Please enter a positive number. If no basal insulin" +
                "was taken, please enter 'n'.")
        }
        if fieldSet {
            newEntry.BasalInsulinUsed = true
            newEntry.BasalInsulinAmount = bInsulin
        }
    }

    err := writeEntryFile(newEntry)
    if err != nil {
        return fmt.Errorf("Error writing new Entry: %q", err)
    }

    fmt.Printf("\nNew entry for %d-%s-%s at %s:%s created.\n\n",
        newEntry.Year,
        formatTime(newEntry.Month),
        formatTime(newEntry.Day),
        formatTime(newEntry.Hour),
        formatTime(newEntry.Minute))

    return nil
}

// queryEntry finds and displays the requested entry from the user.
func queryEntry() error {
    currentDir, err := os.Getwd()
    if err != nil {
        return err
    }
    err = changeEntriesDir()
    if err != nil {
        return err
    }
    dirEntries, err := os.ReadDir(".")
    if err != nil {
        return fmt.Errorf("Error reading bgjournal/entries: %q", err)
    }

    if len(dirEntries) < 1 {
        fmt.Println("No entries to find. Create an entry to find entries.")
        return nil
    }


    var year, month, day, hour, minute int

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Year:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptYear, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        promptYear -= 1
        if promptYear < 0 || promptYear >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }
        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptYear].Name(), 10, 0)
            if err != nil {
                return err
            }
            year = int(tmp)
            os.Chdir(dirEntries[promptYear].Name())
        }
    }

    dirEntries, err = os.ReadDir(".")
    if err != nil {
        return err
    }

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Month:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptMonth, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        promptMonth -= 1
        if promptMonth < 0 || promptMonth >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }
        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptMonth].Name(), 10, 0)
            if err != nil {
                return err
            }
            month = int(tmp)
            os.Chdir(dirEntries[promptMonth].Name())
        }
    }

    dirEntries, err = os.ReadDir(".")
    if err != nil {
        return err
    }

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Day:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptDay, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        promptDay -= 1
        if promptDay < 0 || promptDay >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptDay].Name(), 10, 0)
            if err != nil {
                return err
            }
            day = int(tmp)
            os.Chdir(dirEntries[promptDay].Name())
        }
    }

    dirEntries, err = os.ReadDir(".")
    if err != nil {
        return err
    }

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Entry:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptEntry, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return nil
            }
            return err
        }
        fieldSet = set
        promptEntry -= 1
        if promptEntry < 0 || promptEntry >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptEntry].Name(), 10, 0)
            if err != nil {
                return err
            }
            hour = int(tmp) / 100
            minute = int(tmp) - (hour * 100)
            os.Chdir(dirEntries[promptEntry].Name())
        }
    }
    os.Chdir(currentDir)

    entry, err := readEntryFile(year, month, day, hour, minute)
    if err != nil {
        return err
    }

    printEntry(entry)

    return nil
}

func promptNumberField() (int,  bool, error) {
    in := bufio.NewReader(os.Stdin)
    userData, err := in.ReadString('\n')
    if err != nil {
        return 0, false, fmt.Errorf("An error occurred: %q", err)
    }
    userData = strings.TrimSuffix(userData, "\n")
    if strings.Contains(userData, "q") {
        return -1, true, fmt.Errorf("quit")
    }
    tmp, err := strconv.ParseInt(userData, 10, 0)
    if err != nil {
        if strings.Contains(err.Error(), "invalid syntax") {
            fmt.Println("please enter a number." )
            return 0, false, nil
        } else {
            return 0, false,
                fmt.Errorf("An error occurred parsing int: %q", err)
        }
    }
    return int(tmp), true, nil
}

func queryDay() (int, int, int, error) {
    currentDir, err := os.Getwd()
    if err != nil {
        return 0, 0, 0, err
    }
    err = changeEntriesDir()
    if err != nil {
        return 0, 0, 0, err
    }

    dirEntries, err := os.ReadDir(".")
    if err != nil {
        return 0, 0, 0, fmt.Errorf("Error reading bgjournal/entries: %q", err)
    }

    if len(dirEntries) < 1 {
        fmt.Println("No entries to find. Create an entry to find entries.")
        return 0, 0, 0, fmt.Errorf("no entries")
    }

    var year, month, day int

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Year:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptYear, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return 0, 0, 0, nil
            }
            return 0, 0, 0, err
        }
        fieldSet = set
        promptYear -= 1
        if promptYear < 0 || promptYear >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptYear].Name(), 10, 0)
            if err != nil {
                return 0, 0, 0, err
            }
            year = int(tmp)
            os.Chdir(dirEntries[promptYear].Name())
        }
    }

    dirEntries, err = os.ReadDir(".")
    if err != nil {
        return 0, 0, 0, err
    }

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Month:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptMonth, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return 0, 0, 0, nil
            }
            return 0, 0, 0, err
        }
        fieldSet = set
        promptMonth -= 1
        if promptMonth < 0 || promptMonth >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptMonth].Name(), 10, 0)
            if err != nil {
                return 0, 0, 0, err
            }
            month = int(tmp)
            os.Chdir(dirEntries[promptMonth].Name())
        }
    }

    dirEntries, err = os.ReadDir(".")
    if err != nil {
        return 0, 0, 0, err
    }

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Day:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptDay, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return 0, 0, 0, nil
            }
            return 0, 0, 0, err
        }
        fieldSet = set
        promptDay -= 1
        if promptDay < 0 || promptDay >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptDay].Name(), 10, 0)
            if err != nil {
                return 0, 0, 0, err
            }
            day = int(tmp)
        }
    }

    os.Chdir(currentDir)
    return year, month, day, nil
}

// queryMonthQuarter helps the user find a specific week of a month. Each month
// starts on the 1st, and every 7th day after the 1st (8th, 15th, etc.) may be
// selected. queryMonthQuarter return the year, month, and day for the selected
// week.
func queryMonthQuarter() (int, int, int, error) {
    currentDir, _ := os.Getwd()
    changeEntriesDir()

    dirEntries, _ := os.ReadDir(".")

    if len(dirEntries) < 1 {
        fmt.Println("No entries to find. Create an entry to find entries.")
        return 0, 0, 0, fmt.Errorf("no entries")
    }
    
    var year, month, day int

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Year:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s \n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptYear, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return 0, 0, 0, fmt.Errorf("quit")
            }
            return 0, 0, 0, err
        }
        fieldSet = set
        promptYear -= 1
        if promptYear < 0 || promptYear >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, _ := strconv.ParseInt(dirEntries[promptYear].Name(), 10, 0)
            year = int(tmp)
            os.Chdir(dirEntries[promptYear].Name())
        }
    }

    dirEntries, _ = os.ReadDir(".")

    for fieldSet := false; !fieldSet; {
        fmt.Println("Choose the Month:")
        for i, de := range dirEntries {
            fmt.Printf("[%d] %s\n", i+1, de.Name())
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptMonth, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return 0, 0, 0, fmt.Errorf("quit")
            }
            return 0, 0, 0, err
        }
        fieldSet = set
        promptMonth -= 1
        if promptMonth < 0 || promptMonth >= len(dirEntries) {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptMonth].Name(), 10, 0)
            if err != nil {
                return 0, 0, 0, err
            }
            month = int(tmp)
            os.Chdir(dirEntries[promptMonth].Name())
        }
    }

    dirEntries, _ = os.ReadDir(".")

    for fieldSet := false; !fieldSet; {
        optNum := 1
        fmt.Println("Choose the Day:")
        for i := 0; i < len(dirEntries); i+=7 {
            fmt.Printf("[%d] %s\n", optNum, dirEntries[i].Name())
            optNum++
        }
        fmt.Println("[q] Quit")
        printPrompt()
        promptDay, set, err := promptNumberField()
        if err != nil {
            if err.Error() == "quit" {
                return 0, 0, 0, fmt.Errorf("quit")
            }
            return 0, 0, 0, err
        }
        fieldSet = set
        promptDay -= 1
        if promptDay < 0 || promptDay >= optNum {
            fieldSet = false
            fmt.Println("Please choose a valid option.")
        }

        if fieldSet {
            tmp, err := strconv.ParseInt(dirEntries[promptDay*7].Name(), 10, 0)
            if err != nil {
                return 0, 0, 0, err
            }
            day = int(tmp)
        }
    }

    os.Chdir(currentDir)
    return year, month, day, nil
}
