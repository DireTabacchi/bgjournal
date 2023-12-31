package main

import (
    "fmt"
    "os"
)

func main() {
    myEntry := Entry{
        TimeAndDate: TimeDate {
            12, 13,
            15, 1, 2023,
        },
        BgLevel: 134,
        InsulinAmount: 20,
        BasalInsulinAmount: 38,
        BasalInsulinUsed: true,
    }

    myEntry2 := Entry{
        TimeAndDate: TimeDate {
            18, 9,
            2, 11, 2023,
        },
        BgLevel: 198,
        InsulinAmount: 23,
        BasalInsulinAmount: 0,
        BasalInsulinUsed: false,
    }

    err := writeEntry(myEntry)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error from writeEntry: %q", err)
    }

    err = writeEntry(myEntry2)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error from writeEntry: %q", err)
    }
}
