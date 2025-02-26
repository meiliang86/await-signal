.PHONY: bins clean

# default target
default: bins

bins/await-signal:
	go build -o bins/worker ./worker/main.go
	go build -o bins/starter ./starter/main.go
	go build -o bins/signal ./signal/main.go

bins: clean bins/await-signal

clean:
	rm -rf bins
