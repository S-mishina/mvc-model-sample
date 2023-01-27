# introduction

このリポジトリは、MVCモデルで作成されたSample WEB Applicationです。 本プログラムは、Go言語のWEB Application Framework Gin を利用して作成されています。

## MVCとは？

後日記入

## ディレクトリ構造

### 全体のディレクトリ構成

```terminal:全体のディレクトリ構成
├── application/
├── infra/
```

`application/`: Web Applicationに関するコードが格納されています。

`infra/`: infraに関するコードが格納されています。

### applicationのディレクトリ構成

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

applicationディレクトリ直下に置かれているファイルは基本的に、Goアプリケーションを構成する際に必要なファイルが置かれています。

ex) go.mod、go.sum、main.go、Dockerfile、docker-compose.yamlなど

### infraのディレクトリ構成

```terminal:ディレクトリ構造
├── infra/
    ├── database/
```

`database/`: データベース(mysql)コンテナを起動させるのに必要なファイルが置かれています。ex) mysql.conf、Dockerfile

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

```terminal:アプリケーションの起動方法
> go run main.go
```


### DBのマイグレーション

アプリケーションを初回実行した場合にはDBにマイグレーションを行う必要があります。

```terminal:
# application配下まで移動する
> application/tools/run_migration.sh -v <migration version> -e root:pass@localhost:3306/test -f migration/
```

## 参考文献

後日記入
