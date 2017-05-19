docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" \
	-w /app golang:1.8.1 sh \
	-c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'

docker build -t madhatter/go-hello-http .
