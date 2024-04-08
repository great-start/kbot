# shell - это выполнение команды в оболочке bash/sh в UNIX системах
# APP=$(shell basename ${shell git remote get-url origin})
APP= $(shell basename ${PWD})
REGISTRY=evgenbill
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
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X=github.com/great-start/kbot/cmd.appVersion=${VERSION}"

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

# удаления сгенерированных артефактов сборки или временных файлов, чтобы очистить рабочее окружение перед повторной сборкой
clean:
	rm -rf kbot