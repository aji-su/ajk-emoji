.PHONY: deps clean build package deploy dev event up down
PROJECT_NAME := "ajk-func"

deps:
	go get -u ./...

clean: 
	rm -rf ./build/$(PROJECT_NAME)*
	
build:
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-create   ./src/handler/create/main.go
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-show     ./src/handler/show/main.go
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-download ./src/handler/download/main.go
	GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT_NAME)-split    ./src/handler/split/main.go

package:
	sam package --profile theboss \
		--template-file template.yaml \
		--output-template-file output-template.yaml \
		--s3-bucket sam-template-store.theboss.tech

deploy: clean build package
	sam deploy --profile theboss \
		--template-file output-template.yaml \
		--stack-name sam-template-store \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides Env=prod

logs:
	sam logs --stack-name sam-template-store --name AjkFunctionCreate --tail

dev: clean build up
	env AWS_ACCESS_KEY_ID=testkey \
		AWS_SECRET_ACCESS_KEY=testsecret \
		sam local start-api \
		--docker-network $$(docker network ls -q -f name=$(PROJECT_NAME))

event:
	sam local invoke AjkFunctionSplit --event local_event.json --template template.yaml \
		--docker-network $$(docker network ls -q -f name=$(PROJECT_NAME))

up:
	docker-compose up -d

down:
	docker-compose down
