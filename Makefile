.PHONY: build deploy push-container build-and-push

potato-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/potato

potato-osx:
	go build -o bin/potato

clean:
	rm -R bin

build:
	docker-compose build

push-container:
	docker push itoto/patata-caldo

compile-build-and-push: potato-linux build push-container

deploy:
	kubectl create -f service.yaml -f deployment.yaml

	