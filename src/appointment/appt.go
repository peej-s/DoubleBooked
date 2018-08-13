package appointment

import (
    "fmt"
    "time"
)

type Appointment struct {
    startTime time.Time
    endTime time.Time
}

func (appointment *Appointment) SetTimes(t1, t2 time.Time){
    appointment.startTime = t1
    appointment.endTime = t2
}

func (appointment *Appointment) PrintTimes(){
    print := fmt.Println
    format := "2006-01-02 15:04:05"
    print("Appointment from " + appointment.startTime.Format(format) + " to " + appointment.endTime.Format(format))
}

func IsConflict (a1, a2 *Appointment) bool{
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