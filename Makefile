# shell - это выполнение команды в оболочке bash/sh в UNIX системах
# APP=$(shell basename ${shell git remote get-url origin})
APP= $(shell basename ${PWD})
REGISTRY=ghcr.io
REPO_OWNER=great-start
VERSION=${shell git describe --tags --abbrev=0}-${shell git rev-parse --short HEAD}
TARGETOS=linux #linux darwin windows
TARGETARCH=arm64 #amd64

# устанвока зависимостей
install:
	go get

# форматирование
format:
	gofmt -s -w ./

# проверка стиля
lint:
	golint

# тесты
test:
	go test -v

# build (но перед этим выполняет install и format)
build: install format
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X=github.com/{REPO_OWNER}/kbot/cmd.appVersion=${VERSION}"

image:
	docker buildx build . -t ${REGISTRY}/${REPO_OWNER}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${REPO_OWNER}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

# удаления сгенерированных артефактов сборки или временных файлов, чтобы очистить рабочее окружение перед повторной сборкой
clean:
	rm -rf kbot