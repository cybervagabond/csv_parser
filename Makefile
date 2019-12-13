# Makefile for releasing simpay
#
# The release version is controlled from pkg/version

TAG=latest
NAME:=csv_parser
GOPRIVATE:=github.com
DOCKER_REPOSITORY:=docker.io/cybervagabond
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

vendor:
	export GOPRIVATE
	go mod init gitlab.rdl-telecom.com/companion/back/simpay
	go mod vendor

update-vendor:
	export GOPRIVATE
	rm -r vendor go.mod go.sum
	make vendor

commit:
	git add .
	git commit -m "autocommit"
	git push -u origin master

run:
	GO111MODULE=on go run -ldflags "-s -w -X gitlab.rdl-telecom.com/companion/back/simpay/pkg/version.REVISION=$(GIT_COMMIT)" cmd/simpay/* --level=debug

test:
	GO111MODULE=on go test -v -race ./...

build:
	GO111MODULE=on GIT_COMMIT=$$(git rev-list -1 HEAD) && GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor  -ldflags "-s -w -X gitlab.rdl-telecom.com/companion/back/simpay/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/simpay ./cmd/simpay/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

stop-container:
	@docker stop $(NAME)
	@docker ps

clear-container:
	@docker rm -f $(NAME)

run-compose:
	docker-compose up -d

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latet
s
screen:
	@echo Running $(NAME) binary
	screen -S $(NAME) ./bin/$(NAME)

version-set:
	@next="$(TAG)" && \
	current="$(VERSION)" && \
	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/simpay/values.yaml && \
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/simpay/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/simpay/Chart.yaml && \
	sed -i '' "s/simpay:$$current/simpay:$$next/g" kustomize/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

swagger:
	GO111MODULE=on go get github.com/swaggo/swag/cmd/swag
	cd pkg/api && $$(go env GOPATH)/bin/swag init -g server.go

rmi-win:
	@echo In Windows PowerShell
	#	docker rmi $(docker images --format "{{.Repository}}:{{.Tag}}" | findstr "gitlab.rdl-telecom.com:4567/companion/back/simpay")

rmi:
	docker rmi $(docker images | grep $(DOCKER_IMAGE_NAME))

list:
	docker images $(DOCKER_IMAGE_NAME)

logs:
	docker logs $(NAME)

in:
	docker exec -it $(NAME) sh
