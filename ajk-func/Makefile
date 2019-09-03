.PHONY: deps clean build package deploy dev
PROJECT_NAME := "ajk-func"

deps:
	go get -u ./...

clean: 
	rm -rf ./build/$(PROJECT_NAME)*
	
build:
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-create ./src/handler/create/main.go
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-show ./src/handler/show/main.go
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-download ./src/handler/download/main.go

package:
	sam package --profile theboss \
		--template-file template.yaml \
		--output-template-file output-template.yaml \
		--s3-bucket sam-template-store.theboss.tech

deploy:
	sam deploy --profile theboss \
		--template-file output-template.yaml \
		--stack-name sam-template-store \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides Env=prod

dev:
	env AWS_ACCESS_KEY_ID=testkey \
		AWS_SECRET_ACCESS_KEY=testsecret \
		sam local start-api \
		--docker-network $$(docker network ls -q -f name=$(PROJECT_NAME))
