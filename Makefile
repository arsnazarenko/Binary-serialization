.PHONY: bench

# Таргет bench
bench:
	cd ./internal/bench && go test -bench=. --benchmem