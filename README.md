# ajk-emoji

## ajk-func

### Prerequisites: install develop version of aws-sam-cli

install:

```bash
pip3 install git+https://github.com/awslabs/aws-sam-cli.git@develop
```

check:

```bash
sam --version
# => SAM CLI, version 0.21.0
```

### start local server for development

```bash
cd ./ajk-func
make dev
```

### deploy

```bash
make deploy
```

## ajk-front

```bash
cd ./ajk-front
yarn install
echo 'VUE_APP_API_ENDPOINT=http://localhost:3000/ajk' > .env.development.local                                  
yarn serve
```
