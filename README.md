# introduction

このリポジトリは、MVCモデルで作成されたSample WEB Applicationです。 本プログラムは、Go言語のWEB Application Framework Gin を利用して作成されています。

## MVCとは？

後日記入

## ディレクトリ構造

```terminal:ディレクトリ構造
├── application/
│   ├── config/
│   ├── controller/
│   ├── docs/
│   ├── migration/
│   ├── model/
│   │   └── orm/
│   │       └── ent/
│   ├── tools/
│   └── view/
```

application配下のディレクトリについて解説します。

`config/`: backendシステム(DBなど)の接続情報と接続に関するコードが格納されています。

`docs/`: アプリケーションに関するドキュメンテーションのデータが格納されています。ex) ER図、Go docなど

`migration/`: migrationファイルが格納されています。

`tools/`: アプリケーションに関するツール類が格納されています。

`model/`: MVCモデルにおけるM(Model)に関係するコードが格納されています。※ORMのコードもモデルに含まれています。

`view/`: MVCモデルにおけるV(View)に関係するコードが格納されています。

`controller/`: MVCモデルにおける(Controller)に関係するコード(Router、controller)が格納されています。

## 本プログラムの実行方法

### dockerを使ってアプリケーションを立ち上げる場合

```terminal:terminal
>　docker-compose build
>　docker-compose up -d
```

### dockerを使わずにアプリケーションを起動させる場合

1. DBは、docker containerで起動させる。

```terminal:terminal
> docker-compose -f infra/mysql.yaml build
> docker-compose -f infra/mysql.yaml up -d
```

2. アプリケーションを起動させる。

```
> go run main.go
```

## 参考文献

後日記入
