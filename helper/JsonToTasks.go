package helper

import (
    "os"
    "log"
    "io"
    "encoding/json"
    "projects/task-cli/types"
)

func JsonToTasks () types.Tasks {
    tasksFile,err:= os.Open("tasks.json")
    if err != nil {
        log.Fatal(err)
    }
    defer tasksFile.Close()
    jsonData, err:= io.ReadAll(tasksFile)
    if len(jsonData) == 0 {
        defaultData := `{"tasks":[]}`
        jsonData = []byte(defaultData)
    }
    if err != nil {
        log.Fatal(err)
    }

    var tasks types.Tasks
    err = json.Unmarshal(jsonData, &tasks)
    if err != nil {
        log.Fatal(err)
    } 
    
   return tasks 
}
