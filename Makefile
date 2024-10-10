SWAGGER_UI_VERSION := v5.4.2
SWAGGER_UI_DIR := ./static/swagger-ui

.PHONY: generate-swagger
generate-swagger: generate-swagger-api install-swagger-ui
	@echo "Generating Swagger documentation..."
	@cp ./api/openapi/api.yaml $(SWAGGER_UI_DIR)/swagger.yaml
	@awk '{gsub("https://petstore.swagger.io/v2/swagger.json", "./swagger.yaml")}1' $(SWAGGER_UI_DIR)/swagger-initializer.js > $(SWAGGER_UI_DIR)/swagger-initializer.js.tmp && mv $(SWAGGER_UI_DIR)/swagger-initializer.js.tmp $(SWAGGER_UI_DIR)/swagger-initializer.js
	@echo "Swagger documentation generated successfully"

.PHONY: generate-swagger-api
generate-swagger-api: .ensure-dir
	@echo "Generating API code..."
	@which oapi-codegen > /dev/null || go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@oapi-codegen -config api/openapi/cfg.yaml api/openapi/api.yaml > gen/api/openapi.gen.go
	@echo "API code generated successfully."

.PHONY: install-swagger-ui
install-swagger-ui:
	@echo "Installing Swagger UI..."
	@mkdir -p $(SWAGGER_UI_DIR)
	@curl -sSL https://github.com/swagger-api/swagger-ui/archive/$(SWAGGER_UI_VERSION).tar.gz | tar -xz -C /tmp
	@cp -R /tmp/swagger-ui-$(SWAGGER_UI_VERSION:v%=%)/dist/* $(SWAGGER_UI_DIR)
	@rm -rf /tmp/swagger-ui-$(SWAGGER_UI_VERSION:v%=%)
	@echo "Swagger UI installed successfully"

.PHONY: .ensure-dir
.ensure-dir:
	@mkdir -p gen/api/

.PHONY: test
test:
	@go test ./...

.PHONY: lint
lint:
	@golangci-lint run
