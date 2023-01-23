package model

import (
    "time"
    "context"
    "log"
    "mvc-model-sample/application/config"
    "mvc-model-sample/application/model/orm/ent/todolist"
    "encoding/json"
)

type Todolist struct {
    ID          int         `json:"id"`
    Title       string      `json:"title"`
    Body        string      `json:"body"`
    Priority    int         `json:"priority"`
    Delete_flag int         `json:"delete_flag"`
    CreatedAt   time.Time   `json:"created_at"`
}

func CreateTodo(title string, body string, priority int) {
    client, err := config.ConnectDB()
    if err != nil {
        // エラー処理
    }
    ctx := context.Background()
    _, err = client.Todolist.
    Create().
    SetTitle(title).
    SetBody(body).
    SetPriority(priority).
    Save(ctx)
    if err != nil {
        log.Fatalf("failed creating CreateTodo: %v", err)
    }
    defer client.Close()
}

func GetTodo()[]map[string]interface{} {
    client, err := config.ConnectDB()
    if err != nil {
        // エラー処理
    }
    ctx := context.Background()
    u, err := client.Todolist.Query().Where(todolist.DeleteFlag(0)).All(ctx)
    if err != nil {
        log.Fatalf("failed select GetTodo: %v", err)
    }
    defer client.Close()
    todolistdata, _ := json.Marshal(u)
    var todolist []map[string]interface{}
    json.Unmarshal(todolistdata, &todolist)
    return todolist
}

func GetDetails(id int)[]map[string]interface{}{
    client, err := config.ConnectDB()
    if err != nil {
        // エラー処理
    }
    ctx := context.Background()
    u, err := client.Todolist.Query().Where(todolist.ID(id)).All(ctx)
    if err != nil {
        log.Fatalf("failed select GetDetails: %v", err)
    }
    defer client.Close()
    detailsdata, _ := json.Marshal(u)
    var tododetails []map[string]interface{}
    json.Unmarshal(detailsdata, &tododetails)
    log.Println(tododetails)
    return tododetails
}

func UpdateDetail(id int, title string, body string, priority int){
    client, err := config.ConnectDB()
    if err != nil {
        // エラー処理
    }
    ctx := context.Background()
    err = client.Todolist.
        Update().
        SetTitle(title).
        SetBody(body).
        SetPriority(priority).
        Where(todolist.ID(id)).
        Exec(ctx)
        if err != nil {
            log.Fatalf("not update select UpdateDetail: %v", err)
        }
    defer client.Close()
}

func DeleteDetail(id int){
    client, err := config.ConnectDB()
    if err != nil {
        // エラー処理
    }
    ctx := context.Background()
    err = client.Todolist.
        Update().
        SetDeleteFlag(1).
        Where(todolist.ID(id)).
        Exec(ctx)
        if err != nil {
            log.Fatalf("not update select DeleteDetail: %v", err)
        }
    defer client.Close()
}