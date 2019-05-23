.PHONY: bin

generate:
	@mkdir -p bin
	cd gen && go-bindata data/ templates/
	go build -o bin/gen ./gen
	go generate ./...

bin:
	go build -o bin/basic ./examples/basic
	go build -o bin/pairing ./examples/pairing
	go build -o bin/subscribe ./examples/subscribe
	go build -o bin/gen ./gen

graphviz:
	cat fsm.dot | dot -Tpng -o fsm.png
