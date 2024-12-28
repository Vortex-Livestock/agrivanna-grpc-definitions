# ------------------------------------------------------------------------------
#  Variables
# ------------------------------------------------------------------------------
PROTO_SRC_DIR   := proto
PROTO_OUT_DIR   := internal/pb

# ------------------------------------------------------------------------------
# Install dependencies
# ------------------------------------------------------------------------------
.PHONY: install
install:
	@echo "Installing dependencies..."
	go mod tidy

	@echo "Dependencies installed."

# ------------------------------------------------------------------------------
#  Generate protobuf stubs
# ------------------------------------------------------------------------------
.PHONY: proto
proto:
	@echo "Generating protobuf stubs from '$(PROTO_SRC_DIR)' into '$(PROTO_OUT_DIR)'..."
	protoc \
		-I $(PROTO_SRC_DIR) \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_SRC_DIR)/*.proto

	@echo "Protobuf generation completed."

# ------------------------------------------------------------------------------
# Clean up
# ------------------------------------------------------------------------------
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(PROTO_OUT_DIR)/*.pb.go

	@echo "Clean up completed."