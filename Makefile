-include .env

PARENT_PATH=sample
SAMPLE_FOLDER = github.com/typical-go/hello-world

new-sample:
	@go build 
	@rm -rf $(PARENT_PATH)
	@./typical-go new $(SAMPLE_FOLDER) -parentPath=$(PARENT_PATH)
	
	