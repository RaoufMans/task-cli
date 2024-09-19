package types


type Tasks struct {
    Tasks []Task `json:"tasks"`
    
}

type Task struct {
    Id int 
    Description string 
    Status string     
}
