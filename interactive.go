package main

import (
	"bufio"
	"fmt"
    "os"
    "strconv"
    "strings"
)

func printCommands() {
    fmt.Println("[1] Create new entry")
    fmt.Println("[q] Quit")
}

func printPrompt() {
    fmt.Print("--,^^,*> ")
}

func createEntry() error {
    var newEntry Entry
    var fieldSet bool

    fmt.Println("\nEnter the following to create a new entry.")

    // Get the Year
    for fieldSet = false; !fieldSet; {
        fmt.Print("Year: ")
        year, set, err := promptNumberField()
        if err != nil {
            return err
        }
        fieldSet = set
        if fieldSet {
            newEntry.TimeAndDate.Year = year
        }
    }

    //Get the Month
    for fieldSet = false; !fieldSet; {
        fmt.Print("Month: ")
        month, set, err := promptNumberField()
        if err != nil {
            return err
        }
        fieldSet = set
        if month < 1 || month > 12 {
            fieldSet = false
            fmt.Println("Please enter a valid month.")
        }
        if fieldSet {
            newEntry.TimeAndDate.Month = month
        }
    }

    // Get the Day
    for fieldSet = false; !fieldSet; {
        fmt.Print("Day: ")
        day, set, err := promptNumberField()
        if err != nil {
            return err
        }
        fieldSet = set
        if day < 1 || day > 31 {
            fieldSet = false
            fmt.Println("Please enter a valid day.")
        }
        if fieldSet {
            newEntry.TimeAndDate.Day = day
        }
    }

    // Get the Hour
    for fieldSet = false; !fieldSet; {
        fmt.Print("Hour (24-hour; 0 is midnight): ")
        hour, set, err := promptNumberField()
        if err != nil {
            return err
        }
        fieldSet = set
        if hour < 0 || hour > 23 {
            fieldSet = false
            fmt.Println("Please enter a valid hour.")
        }
        if fieldSet {
            newEntry.TimeAndDate.Hour = hour
        }
    }

    // Get the Minute
    for fieldSet = false; !fieldSet; {
        fmt.Print("Minute: ")
        minute, set, err := promptNumberField()
        if err != nil {
            return err
        }
        fieldSet = set
        if minute < 0 || minute > 59 {
            fieldSet = false
            fmt.Println("Please enter a valid minute.")
        }
        if fieldSet {
            newEntry.TimeAndDate.Minute = minute
        }
    }

    for fieldSet = false; !fieldSet; {
        fmt.Print("Blood Glucose level: ")
        bg, set, err := promptNumberField()
        if err != nil {
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
        }
        fmt.Print("Units of basal insulin: ")
        bInsulin, set, err := promptNumberField()
        if err != nil {
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
        newEntry.TimeAndDate.Year,
        formatTwoDigitDateTime(newEntry.TimeAndDate.Month),
        formatTwoDigitDateTime(newEntry.TimeAndDate.Day),
        formatTwoDigitDateTime(newEntry.TimeAndDate.Hour),
        formatTwoDigitDateTime(newEntry.TimeAndDate.Minute))

    return nil
}

func promptNumberField() (int,  bool, error,) {
    in := bufio.NewReader(os.Stdin)
    userData, err := in.ReadString('\n')
    if err != nil {
        return 0, false, fmt.Errorf("An error occurred: %q", err)
    }
    userData = strings.TrimSuffix(userData, "\n")
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
