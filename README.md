# AWS SAM Go lambda example

## Installing the AWS SAM CLI on macOS

[Installing the AWS SAM CLI on macOS](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html)

## build

```sh
make build
```

## Start Api

```sh
sam local start-api
```

### test diffrent api

### hello-world-api-get

```sh
curl  http://localhost:3000
```

### hello-world-api-post

```sh
curl -X POST -H "Content-Type:application/json" --data '{ "message": "world" }' http://localhost:3000
```