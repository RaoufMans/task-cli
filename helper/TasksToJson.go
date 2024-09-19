package helper

import (
    "encoding/json" 
    "projects/task-cli/types"
    "os"
    "log"
) 


func TasksToJson(tasks types.Tasks) { 
    jsonData, err := json.Marshal(tasks)
    if err != nil {
        log.Fatal(err)
    }

    jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR, 0644)
    if err != nil {
       log.Fatal(err) 
    }
    
    defer jsonFile.Close()
    jsonFile.Seek(0,0)
    jsonFile.Truncate(0)
    jsonFile.Write(jsonData)

}


