%: src/%.go
	go build $<

all: grep cat
