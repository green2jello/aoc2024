.PHONY: install test run-day clean

# Install Go using brew
install:
	brew update
	brew install go
	go mod init aoc2024
	go mod tidy

# Run tests for all days
test:
	go test ./days/...

# Run a specific day (usage: make run-day DAY=01)
run-day:
	@if [ -z "$(DAY)" ]; then \
		echo "Please specify a day (e.g., make run-day DAY=01)"; \
		exit 1; \
	fi
	go run ./days/day$(DAY)/main.go

# Clean build artifacts
clean:
	go clean
	rm -f */*/input.txt