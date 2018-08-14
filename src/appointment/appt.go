package appointment

import (
        "fmt"
        "strconv"
        "time"
)

type Appointment struct {
        Index     int
        startTime time.Time
        endTime   time.Time
}

func (appointment *Appointment) SetAppt(index int, t1, t2 time.Time) {
        // Setter
        appointment.Index = index
        appointment.startTime = t1
        appointment.endTime = t2
}

func (appointment *Appointment) PrintTimes() {
        // For debugging
        print := fmt.Println
        format := "2006-01-02 15:04:05"
        print("Appointment " + strconv.Itoa(appointment.Index) + " from " + appointment.startTime.Format(format) + " to " + appointment.endTime.Format(format))
}

func IsConflict(a1, a2 *Appointment) bool {
        // Two appointments are a conflict if the earliest end time is after the latest start time
        value := false
        maxStart := a2.startTime
        minEnd := a1.endTime

        if a2.startTime.Before(a1.startTime) {
                maxStart = a1.startTime
        }
        if a1.endTime.After(a2.endTime) {
                minEnd = a2.endTime
        }

        if minEnd.After(maxStart) {
                value = true
        }

        return value
}
