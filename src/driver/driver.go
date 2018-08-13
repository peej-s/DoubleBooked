package main

import (
  "fmt"
  "appointment"
  "filereader"
)


type Pair struct {
    a, b interface{}
}

func getPairs(fileName string) []Pair{
    results := []Pair{}
    apptList := filereader.CSV2apptlist(fileName) //O(n)
    for i, e := range apptList{
        for j := 0; j < i; j++ {
            a1 := e
            a2 := apptList[j]
            if(appointment.IsConflict(&a1, &a2)){
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
    results := getPairs("appointments.csv")
    fmt.Println(len(results))
}
	
