BUILD_DIR=build

.PHONY: clean build

clean:
	rm -rf $(BUILD_DIR)

build:
	mkdir -p $(BUILD_DIR)
	BUILD_DIR=$(BUILD_DIR) ./build.sh "linux/amd64" "darwin/amd64"
