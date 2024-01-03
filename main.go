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
        default:
            fmt.Printf("Unrecognized option: %s\n\n", userChoice)
        }
    }
    //myEntry := Entry{
    //    TimeAndDate: TimeDate {
    //        12, 13,
    //        15, 1, 2023,
    //    },
    //    BgLevel: 134,
    //    InsulinAmount: 20,
    //    BasalInsulinAmount: 38,
    //    BasalInsulinUsed: true,
    //}

    //myEntry2 := Entry{
    //    TimeAndDate: TimeDate {
    //        18, 9,
    //        2, 11, 2023,
    //    },
    //    BgLevel: 198,
    //    InsulinAmount: 23,
    //    BasalInsulinAmount: 0,
    //    BasalInsulinUsed: false,
    //}

    //err := writeEntryFile(myEntry)
    //if err != nil {
    //    fmt.Fprintf(os.Stderr, "Error from writeEntryFile: %q\n", err)
    //}

    //err = writeEntryFile(myEntry2)
    //if err != nil {
    //    fmt.Fprintf(os.Stderr, "Error from writeEntryFile: %q\n", err)
    //}

    //entryOne, err := readEntryFile(2023, 1, 15)
    //if err != nil {
    //    fmt.Fprintf(os.Stderr, "Error from readEntryFile: %q\n", err)
    //}
    //fmt.Println(entryOne)
    //entryTwo, err := readEntryFile(2023, 11, 2)
    //if err != nil {
    //    fmt.Fprintf(os.Stderr, "Error from readEntryFile: %q\n", err)
    //}
    //fmt.Println(entryTwo)
}
