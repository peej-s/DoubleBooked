package main

import (
    "bufio"
    "encoding/csv"
    "io"
    "log"
    "os"
    "fmt"
    "time"
)

type Appointment struct {
    startTime time.Time
    endTime time.Time
}

type Pair struct {
    a, b interface{}
}

func (appointment *Appointment) printTimes(){
    print := fmt.Println
    format := "2006-01-02 15:04:05"
    print("Appointment from " + appointment.startTime.Format(format) + " to " + appointment.endTime.Format(format))
}

func isConflict (a1, a2 *Appointment) bool{
    value := false
    maxStart := a2.startTime
    minEnd := a1.endTime
    
    if a2.startTime.Before(a1.startTime){
        maxStart = a1.startTime
    }
    if a1.endTime.After(a2.endTime){
        minEnd = a2.endTime
    }
    
    if minEnd.After(maxStart){
        value = true
    }
    
    return value
}

func csv2apptlist(fileName string) []Appointment{
    // Load csv
    parser := time.Parse
    csvFile, _ := os.Open("./file/" + fileName)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    defer csvFile.Close()

    format := "2006-01-02 15:04:05"
    var count int = 0
    values := []Appointment{}
    
    
    for {
    	line, error := reader.Read()
    	
    	if error == io.EOF {
    	    break
    	} else if error != nil {
    	    log.Fatal(error)
    	}
    	
    	if(count == 0){
    	    
    	} else{

            t1, _ := parser(format, line[0])
            t2, _ := parser(format, line[1])
            temp := Appointment{
                t1,
                t2,
	        }
	    values = append(values, temp)
    	}
    	count += 1
    }

    return values
}

func csv2apptlist(fileName string) []Pair{
    results := []Pair{}
    apptList := csv2apptlist("appointments.csv") //O(n)
    for i, e := range apptList{
        for j := 0; j < i; j++ {
            a1 := e
            a2 := apptList[j]
            if(isConflict(&a1, &a2)){
                newPair := Pair{
                    a1,
                    a2,
                }
                results = append(results, newPair)
            }
        } 
    } // O(n^2)
    return results
}


func main(){
    
}
	
