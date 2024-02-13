# Makefile for running Swag command

SWAG_TARGET := docs/swagger.json  # Output file for generated Swagger spec

# Rule to generate your Swagger documentation
$(SWAG_TARGET):
	swag init -g ../../cmd/kontokompass/main.go -o ./docs -d ./internal/handler

# Convenience target to directly run the 'swag init' command
run-swag: $(SWAG_TARGET)

# Standard target for cleaning generated files
clean:
	rm -rf docs
