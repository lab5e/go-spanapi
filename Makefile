all:
	go vet ./...
	go build
	mkdir -p bin
	cd examples/simple && go build -o ../../bin
	cd examples/data && go build -o ../../bin
	cd examples/interval && go build -o ../../bin
	cd examples/websocket && go build -o ../../bin
	cd examples/paging && go build -o ../../bin

clean:
	rm -fR bin
