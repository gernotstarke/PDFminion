
.PHONY: build
build:
	go build -o pdfminion ./cmd/pdfminion

.PHONY: install
install:
	go install ./cmd/pdfminion

.PHONY: clean
clean:
	rm -f pdfminion
	rm -f pdfminion.exe