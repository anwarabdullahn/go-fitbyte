.PHONY: swagger-gen
swagger-gen:
	swag init --generalInfo src/app.go --output docs

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  swagger-gen    Generate Swagger documentation"