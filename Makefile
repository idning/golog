all:
	go test -race

bench:
	go run benchmark/log_benchmark.go

clean:
	rm *.log
