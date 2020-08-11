# go-swagger-test

## about
go-swaggerわからんので動かして試してみるよ

# how to use

## use go-swagger
go-swaggerを任意の方法でインストールしてね。
インストール方法は[こちら](https://goswagger.io/install.html)

```
swagger version
```

## go-swagger command

makeファイルにユーティリティを記載

build source

```
cd server
make generate server

※初回はmakeコマンド使わずにexclude-mainオプション使わ無い方が良い
swagger generate server -a factory -A factory --exclude-main -f ../swagger.yaml -t gen
```

run server
```
make run-dev
```

詳しくは[技術ブログ](https://future-architect.github.io/articles/20200630/)を見てね
