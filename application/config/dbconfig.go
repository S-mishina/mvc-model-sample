package config

import (
    "context"
    "log"
    "fmt"
    "os"
    "mvc-model-sample/application/model/orm/ent"
    _ "github.com/go-sql-driver/mysql"
)

// TODO: 認証情報は環境変数で持たせる。
// DB init
func ConnectDB() (*ent.Client,error) {
    client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", os.Getenv("MYSQL_ROOT_ID"),os.Getenv("MYSQL_ROOT_PASSWORD"),os.Getenv("MYSQL_HOST"),os.Getenv("MYSQL_PORT"),os.Getenv("MYSQL_DATABASE")))
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
        log.Printf(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", os.Getenv("MYSQL_ROOT_ID"),os.Getenv("MYSQL_ROOT_PASSWORD"),os.Getenv("MYSQL_HOST"),os.Getenv("MYSQL_PORT"),os.Getenv("MYSQL_DATABASE")))
    }
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    return client, err
}