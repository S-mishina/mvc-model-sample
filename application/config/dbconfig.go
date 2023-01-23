package config

import (
    "context"
    "log"
    "learn-gin-mvc/application/model/orm/ent"
    _ "github.com/go-sql-driver/mysql"
)

// TODO: 認証情報は環境変数で持たせる。
// DB init
func ConnectDB() (*ent.Client,error) {
    client, err := ent.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/test?parseTime=True")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    return client, err
}