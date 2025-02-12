# ------------------------------------------------------------------------------
#  Variables
# ------------------------------------------------------------------------------
PROTO_BASE_DIR := proto/v1
GENERATED_DIR  := generated
MICROSERVICES  := livestock user farm

# ------------------------------------------------------------------------------
#  Install dependencies
# ------------------------------------------------------------------------------
.PHONY: install
install:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed."

# ------------------------------------------------------------------------------
#  Generate protobuf stubs for each microservice
# ------------------------------------------------------------------------------
.PHONY: proto
proto:
	@echo "Generating protobuf stubs..."
	@mkdir -p $(GENERATED_DIR)
	@for svc in $(MICROSERVICES); do \
		echo "Generating stubs for $$svc..."; \
		protoc \
			-I $(PROTO_BASE_DIR)/$$svc \
			--go_out=paths=source_relative:$(GENERATED_DIR)/$$svc \
			--go-grpc_out=paths=source_relative:$(GENERATED_DIR)/$$svc \
			$(PROTO_BASE_DIR)/$$svc/*.proto; \
	done
	@echo "Protobuf generation completed."

# ------------------------------------------------------------------------------
#  Clean up
# ------------------------------------------------------------------------------
.PHONY: clean
clean:
	@echo "Cleaning up generated stubs..."
	@for svc in $(MICROSERVICES); do \
		rm -rf $(GENERATED_DIR)/$$svc/*; \
	done
	@echo "Clean up completed."
