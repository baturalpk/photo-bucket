APP_EXEC_SRC_PATH   = cmd/server/server.go
APP_BUILD_DEST_PATH = build/photo_bucket

DARWIN  = $(APP_BUILD_DEST_PATH)_darwin
LINUX   = $(APP_BUILD_DEST_PATH)_linux
WINDOWS = $(APP_BUILD_DEST_PATH)_windows.exe

VERSION = $(shell git describe --tags --always --long)

deployment:
	docker build -t "photo-bucket:$(VERSION)" .
	docker tag "photo-bucket:$(VERSION)" "photo-bucket:latest"
	cd deploy; docker compose -p photo-bucket-project up -d

darwin_build:
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" ${APP_EXEC_SRC_PATH}

linux_build:
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" ${APP_EXEC_SRC_PATH}

windows_build:
	env GOOS=windows GOARCH=amd64 go.exe build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" ${APP_EXEC_SRC_PATH}

# test:
	# TODO: Prepare & Execute "./scripts/run_tests.sh"
