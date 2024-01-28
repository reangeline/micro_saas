# .PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/app/main.go

deploy_prod: build
	serverless deploy --stage prod

# clean:
# 	rm -rf ./bin ./vendor go.sum

# deploy: clean build
# 	sls deploy --verbose


# gomodgen:
# 	chmod u+x gomod.sh
# 	./gomod.sh
