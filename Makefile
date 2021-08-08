build:
	go build -o ascii-banner
test:
	go test ./banner
test_bench:
	go test ./banner -bench=.
clean:
	rm -rf ascii-banner
