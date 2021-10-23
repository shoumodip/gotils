%: src/%.go
	go build $<

all: grep cat tac echo seq yes factor head tail sleep
