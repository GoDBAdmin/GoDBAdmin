.PHONY: all build-backend build-frontend build-package clean test install

VERSION ?= 1.0.0
BUILD_DIR = build
DEB_DIR = debian

all: build-backend build-frontend

build-backend:
	@echo "Building backend..."
	@mkdir -p $(BUILD_DIR)
	cd backend && go mod download
	cd backend && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ../$(BUILD_DIR)/mysql-admin-tool ./cmd/server
	@echo "Backend built successfully!"

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm ci
	cd frontend && npm run build
	@echo "Frontend built successfully!"

build-package: build-backend build-frontend
	@echo "Building Debian package..."
	@if [ ! -d "$(DEB_DIR)" ]; then \
		echo "Error: debian directory not found"; \
		exit 1; \
	fi
	dpkg-buildpackage -us -uc -b
	@echo "Package built successfully!"

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)/*
	rm -rf frontend/dist
	rm -rf frontend/node_modules
	rm -rf backend/vendor
	rm -f ../mysql-admin-tool_*.deb
	rm -f ../mysql-admin-tool_*.changes
	rm -f ../mysql-admin-tool_*.buildinfo
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	cd backend && go test ./...
	@echo "Tests complete!"

install: build-package
	@echo "Installing package..."
	sudo dpkg -i ../mysql-admin-tool_*.deb || sudo apt-get install -f -y
	@echo "Installation complete!"

dev-backend:
	@echo "Starting backend in development mode..."
	cd backend && go run ./cmd/server/main.go

dev-frontend:
	@echo "Starting frontend in development mode..."
	cd frontend && npm run dev

