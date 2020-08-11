#############
# Build
#############

.PHONY: deps clean build unittest test zip deploy-dev

build: deps
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/lambda cmd/lambda/lambda.go

deps:
	go mod download
	go mod tidy

clean:
	rm -rf ./bin/*

unittest:
	go test -v ./

test:
	go mod tidy
	go test -race -v ./

#############
# Code generate
#############

generate: generate-server generate-client

generate-server:
	swagger generate server -a factory -A factory --exclude-main -f ../swagger.yaml -t gen

generate-client:
	swagger generate client -a factory -A factory -f ../swagger.yaml -t ../gen

#############
# Run as Dev server
#############
run-dev:
	go run gen/cmd/factory-server/main.go --host 0.0.0.0 --port 3000

# #############
# # Deployment
# #############
# zip: build
# 	zip -j bin/lambda.zip bin/lambda

# # update lambda function. if you wanna create new function, do not use belows.
# deploy-dev: zip
# 	aws lambda update-function-code --profile nicigas_dev --region ap-northeast-1 --function-name dev-nicigas-mddm-api --zip-file fileb://bin/lambda.zip

# deploy-stg: zip
# 	aws lambda update-function-code --profile nicigas_dev --region ap-northeast-1 --function-name stg-nicigas-mddm-api --zip-file fileb://bin/lambda.zip

# deploy-prod: zip
# 	aws lambda update-function-code --profile nicigas_prod --region ap-northeast-1 --function-name prod-nicigas-mddm-api --zip-file fileb://bin/lambda.zip
