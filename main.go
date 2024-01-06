package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    //"unicode/utf8"
)

func main() {
    fmt.Println("Blood Glucose Journal")
    fmt.Println("---------------------")
    fmt.Println("Type the option in the brackets and hit Enter.")
    in := bufio.NewReader(os.Stdin)
    for {
        printCommands()
        printPrompt()
        userChoice, err := in.ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading string: %q\n", err)
        }
        userChoice = strings.TrimSuffix(userChoice, "\n")

        switch (userChoice) {
        case "q":
            os.Exit(0)
        case "1":
            err := createEntry()
            if err != nil {
                fmt.Fprintf(os.Stderr, "%q\n", err)
            }
        case "2":
            queryEntry()
        case "3":
            year, month, day, err := queryDay()
            if err != nil {
                if err.Error() == "no entries" { continue }
                fmt.Fprintf(os.Stderr, "An error occurred finding the day: %q",
                    err)
            }
            currentDir, _ := os.Getwd()
            changeEntriesDir()
            dirEntries, _ := os.ReadDir(formatDayPath(year, month, day))
            os.Chdir(currentDir)
            var entries []Entry
            for _, entryFile :=range dirEntries {
                hour, minute := parseFileName(entryFile.Name())
                entry, _ := readEntryFile(year, month, day, hour, minute)
                entries = append(entries, entry)
            }
            printDay(entries)
        case "4":
            year, month, day, err := queryMonthQuarter()
            if err != nil {
                if err.Error() == "quit" {
                    break
                } else {
                    fmt.Fprintf(os.Stderr, "An error occurred: %q\n", err)
                }
            }
            entries, _ := readEntryFilesNDays(year, month, day, 7)
            printWeek(entries)
        default:
            fmt.Printf("Unrecognized option: %s\n\n", userChoice)
        }
    }
}
