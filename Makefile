BINARY_NAME=tunaflow

all: build

build: asset
	go build -o ./${BINARY_NAME} .

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

asset:
	go run tools/go-assets-builder.go templates -o assets/asset_files.go -p assets -v AssetFiles
