.PHONY: setup build test clean

setup:
	@echo "Installing necessary tools and depencendies..."
	@./scripts/install_dependencies.sh
build:
	@echo "Building the project"
	go build -o myapp ./cmd/myapp
test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Nothing to clean up"
	# placeholder
