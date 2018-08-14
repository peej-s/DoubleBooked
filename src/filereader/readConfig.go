package filereader

import (
        "encoding/json"
        "io/ioutil"
        "os"
)

type Dict struct {
        InputFilename  string `json:"inputFilename"`
        OutputFilename string `json:"outputFilename"`
        TimeFormat     string `json:"timeFormat"`
}

func LoadConfig() *Dict {
        var dict Dict
        jsonFilename := "../config/config.json"

        jsonFile, err := os.Open(jsonFilename)
        check(err)
        defer jsonFile.Close()

        byteValue, _ := ioutil.ReadAll(jsonFile)
        json.Unmarshal(byteValue, &dict)

        return &dict
}
