package main

import (
    "fmt"
    "time"
)

type Appointment struct {
    startTime time.Time
    endTime time.Time
}

func (appointment *Appointment) printTimes(){
    print := fmt.Println
    format := "2006-01-02 15:04:05"
    print("Appointment from " + appointment.startTime.Format(format) + " to " + appointment.endTime.Format(format))
}

func isConflict (a1, a2 *Appointment) bool{
    value := false
    maxStart := a1.startTime
    minEnd := a1.endTime
    
    if a1.startTime.Before(a2.startTime){
        maxStart = a2.startTime
    }
    if a1.endTime.After(a2.endTime){
        minEnd = a2.endTime
    }
    
    if minEnd.Before(maxStart){
        value = true
    }
    
    return value
}
	
func main() {
    print := fmt.Println
    parser := time.Parse
    
    format := "2006-01-02 15:04:05"
    startTime := "2017-03-27 14:00:00"
    endTime := "2017-03-27 15:30:00"
    
    t0, err := parser(format, startTime)
    if err != nil {
        fmt.Println("ERROR:", err.Error())
    }
    t1, err := parser(format, endTime)
        if err != nil {
            fmt.Println("ERROR:", err.Error())
    }
    
    var myAppointment Appointment = Appointment{
        t0,
        t1,
    }
    
    startTime = "2017-03-27 15:00:00"
    endTime = "2017-03-27 16:30:00"
    
    t2, err := parser(format, startTime)
    if err != nil {
        fmt.Println("ERROR:", err.Error())
    }
    t3, err := parser(format, endTime)
        if err != nil {
            fmt.Println("ERROR:", err.Error())
    }
    
    var myConflict Appointment = Appointment{
        t2,
        t3,
    }
    startTime = "2017-03-28 15:00:00"
    endTime = "2017-03-28 16:30:00"
        
    t4, err := parser(format, startTime)
    if err != nil {
        fmt.Println("ERROR:", err.Error())
    }
    t5, err := parser(format, endTime)
        if err != nil {
            fmt.Println("ERROR:", err.Error())
    }
       
    var myFreedom Appointment = Appointment{
        t4,
        t5,
    }
    
    myAppointment.printTimes()
    myConflict.printTimes()
    
    print(isConflict(&myAppointment, &myConflict))
    print(isConflict(&myConflict, &myAppointment))
    print(isConflict(&myAppointment, &myAppointment))
    print(isConflict(&myAppointment, &myFreedom))
}