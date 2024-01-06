package main

func bgAverage(entries []Entry) float64 {
    var entriesCount, bgSum int
    var bgAverage float64

    for _, entry := range entries {
        entriesCount++
        bgSum += entry.BgLevel
    }

    bgAverage = float64(bgSum) / float64(entriesCount)

    return bgAverage
}
