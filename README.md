# TechBoard-BE

## about

エンジニア向けの掲示板、記事投稿アプリ「TechBoard」の API を提供する。

## swagger

https://growthodyssey.github.io/TechBoard-BE/dist/index.html

## setup

- 以下が導入されていること

| feature | ver     |
| ------- | ------- |
| Go      | 1.16.6  |
| Node.js | 14 以上 |

- リポジトリのクローン

```sh
$ git clone git@github.com:GrowthOdyssey/TechBoard-BE.git
$ cd TechBoard-BE
```

- パッケージインストール

```sh
$ npm install
```

- 起動

```sh
$ go run main.go
```

## mock server

- モックサーバ起動

```sh
$ npm run mock-server
```

- ヘルスチェック

```sh
$ curl http://127.0.0.1:4010/health_check
```

## DB

- DB 作成

```sh
$ ./shell/create_db.sh
```

- テーブル作成

```sh
$ ./shell/create_table.sh
```

- テストデータ作成

```sh
$ ./shell/create_testdata.sh
```
