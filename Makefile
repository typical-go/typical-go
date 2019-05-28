-include .env

BIN_TARGET=bin

build:
	@mkdir -p $(BIN_TARGET)
	# @echo "  >  Building $(BIN_TARGET)/typigen"
	# @go build -o $(BIN_TARGET)/typigen ./cli/typigen
	@echo "  >  Building $(BIN_TARGET)/typigo"
	@go build -o $(BIN_TARGET)/typigo ./cli/typigo

clean:
	@echo "  >  Remove $(BIN_TARGET)"
	@rm -rf $(BIN_TARGET)
	
	