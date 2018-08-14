package filereader

import (
        "appointment"
        "bufio"
        "encoding/csv"
        "io"
        "log"
        "os"
        "time"
)

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func CSV2apptlist(fileName, format string) *[]appointment.Appointment {
        // Load csv
        parser := time.Parse
        csvFile, err := os.Open("../file/" + fileName)
        check(err)

        reader := csv.NewReader(bufio.NewReader(csvFile))
        defer csvFile.Close()

        var count int = 0
        values := []appointment.Appointment{}

        for {
                line, error := reader.Read()

                if error == io.EOF {
                        break
                } else if error != nil {
                        log.Fatal(error)
                }

                if count == 0 {
                        // Ignore header

                } else {
                        // Parse times
                        t1, _ := parser(format, line[0])
                        t2, _ := parser(format, line[1])

                        // Set new appointment
                        temp := &appointment.Appointment{}
                        temp.SetAppt(count, t1, t2)

                        // Append to final list
                        values = append(values, *temp)
                }
                count += 1
        }

        return &values
}
