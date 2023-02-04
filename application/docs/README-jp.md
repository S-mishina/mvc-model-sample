# applicationのディレクトリ構成

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
