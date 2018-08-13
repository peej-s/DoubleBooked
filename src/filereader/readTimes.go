package filereader

import (
    "bufio"
    "encoding/csv"
    "io"
    "log"
    "os"
    "time"
    "appointment"
)
    
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CSV2apptlist(fileName string) []appointment.Appointment{
    // Load csv
    parser := time.Parse
    csvFile, err := os.Open("../file/" + fileName)
    check(err)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    defer csvFile.Close()

    format := "2006-01-02 15:04:05"
    var count int = 0
    values := []appointment.Appointment{}
    
    
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
            temp := &appointment.Appointment{}
            temp.SetTimes(t1,t2)
	    values = append(values, *temp)
    	}
    	count += 1
    }

    return values
}
