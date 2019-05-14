.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world-api-get/dist
	rm -rf ./hello-world-api-post/dist
	rm -rf ./hello-world-api-delete/dist
	
build:
	mkdir -p hello-world-api-get/dist; cd hello-world-api-get; GOOS=linux GOARCH=amd64 go build -o dist/hello-world-api-get ./; cd dist; zip hello-world-api-get.zip hello-world-api-get
	mkdir -p hello-world-api-post/dist; cd hello-world-api-post; GOOS=linux GOARCH=amd64 go build -o dist/hello-world-api-post ./; cd dist; zip hello-world-api-post.zip hello-world-api-post
	mkdir -p hello-world-api-delete/dist; cd hello-world-api-delete; GOOS=linux GOARCH=amd64 go build -o dist/hello-world-api-delete ./; cd dist; zip hello-world-api-delete.zip hello-world-api-delete