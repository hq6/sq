all: sq .completions

install:
	go install .

sq: sq.go
	go build .

completions/.built: completions/generate-completions.go
	go run completions/generate-completions.go
	touch $@
