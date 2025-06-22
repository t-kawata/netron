build:
	mkdir -p dist
	./sh/build
build-linux-amd64:
	mkdir -p dist
	./sh/build -o linux -a amd64
