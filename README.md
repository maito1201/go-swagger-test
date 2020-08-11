# go-swagger-test

## about
go-swaggerわからんので動かして試してみる

# how to use

## use go-swagger
go-swaggerを任意の方法でインストールしてね。
インストール方法は[こちら](https://goswagger.io/install.html)

## go-swagger command

makeファイルにユーティリティを記載

build source

```
cd server
make generate server

※初回はmakeコマンド使わずにexclude-mainオプション無しでgenerateする
swagger generate server -a factory -A factory --exclude-main -f ../swagger.yaml -t gen
```

run server
```
cd server
make run-dev
```

詳しくは[技術ブログ](https://future-architect.github.io/articles/20200630/)を見てね

## hello world

サーバー立ち上げ後にブラウザでlocalhostにアクセス

ビルド済のためswaggerインストールしなくてもmake run-devコマンド叩けば動く
```
http://localhost:3000/hello?name=go-swagger
```
