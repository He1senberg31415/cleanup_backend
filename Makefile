# Name of the executable
APP_NAME=pocketbase

# Go build command with CGO disabled for static linking
build:
	CGO_ENABLED=0 go build -o $(APP_NAME)

# Start the server after building
run: build
	./$(APP_NAME) serve --http "0.0.0.0:80"

# Clean the executable
clean:
	rm -f $(APP_NAME)

# Rebuild and run
rebuild: clean run