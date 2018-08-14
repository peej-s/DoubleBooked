package main

import (
        "appointment"
        "bufio"
        "filereader"
        "os"
        "strconv"
)

type Pair struct {
        a, b interface{}
}

func getPairs() ([]Pair, string) {
        // Set config variables
        inputDict := filereader.LoadConfig()

        // Get list of appointments
        results := []Pair{}
        apptList := *filereader.CSV2apptlist(inputDict.InputFilename, inputDict.TimeFormat) //O(n)
        for i, e := range apptList {
                for j := 0; j < i; j++ {
                        a1 := apptList[j]
                        a2 := e
                        if appointment.IsConflict(&a1, &a2) {
                                newPair := Pair{
                                        a1,
                                        a2,
                                }
                                results = append(results, newPair)
                        }
                }
        } // O(n^2)
        return results, inputDict.OutputFilename
}

func exportPairs() {
        // Open output file
        results, fileName := getPairs()
        file, _ := os.Create("../file/" + fileName)
        defer file.Close()

        // Write to file
        writer := bufio.NewWriter(file)
        defer writer.Flush()
        writer.WriteString("First Appointment,Second Appointment\n")

        for _, e := range results {
                a1 := e.a.(appointment.Appointment)
                a2 := e.b.(appointment.Appointment)
                writer.WriteString("Appointment " + strconv.Itoa(a1.Index) + ",Appointment " + strconv.Itoa(a2.Index) + "\n")
        }
}

func main() {
        exportPairs()
}
