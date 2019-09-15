# ajk-emoji

## ajk-func

### Prerequisites: install master version of aws-sam-cli

clone: https://github.com/awslabs/aws-sam-cli.git

install:

```bash
python3 setup.py install
```

check:

```bash
sam --version
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
