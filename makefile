.PHONY: bin

CMDS := examples/factory_reset examples/pairing examples/send_command examples/subscribe gen

tools:
	go get github.com/kevinburke/go-bindata
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

generate:
	@mkdir -p bin/examples
	cd gen && go-bindata data/ templates/
	go build -o bin/gen ./gen
	go generate ./...
	$(MAKE) fmt

bin: $(CMDS)

.PHONY: $(CMDS)
$(CMDS): generate
	go build -o bin/$@ ./$@

graphviz:
	cat fsm.dot | dot -Tpng -o fsm.png

fmt:
	goimports -w .
	gofmt -s -w .

lint:
	golangci-lint run
