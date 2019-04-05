BUILD_DIR=build
PROJECT_NAME=who-to-blame
DEFAULT_TARGET=$(BUILD_DIR)/$(PROJECT_NAME)

.PHONY: clean build install

clean:
	rm -rf $(BUILD_DIR)

build:
	mkdir -p $(BUILD_DIR)
	# default build
	go build -o $(DEFAULT_TARGET)
	# build for specific operating systems
	BUILD_DIR=$(BUILD_DIR) PROJECT_NAME=$(PROJECT_NAME) ./build.sh "linux/amd64" "darwin/amd64"

install: build
	cp $(DEFAULT_TARGET) /usr/local/bin
