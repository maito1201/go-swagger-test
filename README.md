# go-swagger-test

## about
go-swaggerわからんので動かして試してみるよ

# how to use

##build app

go とgo-swaggerをインストールしたコンテナがビルドされます。
ベースイメージが`node:12.7.0-alpine`なのに深い理由は無い。

```
dockder-compose build
```

## use go-swagger
go-swaggerを使う方法は、コンテナの中に入って`swagger`コマンドを使います。

```
docker-compose run --rm app sh
swagger version
```

## go-swagger command

build source

```
swagger generate server -a routemanagement -A routemanagement --exclude-main --strict-additional-properties -t server/gen f ./swagger/swagger.yaml
```

詳しくは[技術ブログ](https://future-architect.github.io/articles/20200630/)を見てね
