build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/jomtxd-linux-amd64 ./cmd/jomtxd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/jomtxd-linux-arm64 ./cmd/jomtxd/main.go

do-checksum-linux:
	cd build && sha256sum \
		jomtxd-linux-amd64 jomtxd-linux-arm64 \
		> jomtx-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/jomtxd-darwin-amd64 ./cmd/jomtxd/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/jomtxd-darwin-arm64 ./cmd/jomtxd/main.go

build-all: build-linux build-darwin

do-checksum-darwin:
	cd build && sha256sum \
		jomtxd-darwin-amd64 jomtxd-darwin-arm64 \
		> jomtx-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-with-checksum: build-linux-with-checksum build-darwin-with-checksum
