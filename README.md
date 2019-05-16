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

## Test diffrent functions/events

### hello-world-api-get

```sh
curl  http://localhost:3000
```

### hello-world-api-get-query

```sh
curl -X GET -H 'content-type: application/json'   http://localhost:3000/query?name=foo
```

### hello-world-api-get-path-param

```sh
curl -X GET -H 'content-type: application/json'   http://localhost:3000/path/foo
```

### hello-world-api-post

```sh
curl -X POST -H "Content-Type:application/json" --data '{ "message": "world" }' http://localhost:3000
```

### hello-world-api-put

```sh
curl -X PUT -H "Content-Type:application/json" --data '{ "message": "world" }' http://localhost:3000
```

### hello-world-api-delete

```sh
curl -X DELETE http://127.0.0.1:3000
```