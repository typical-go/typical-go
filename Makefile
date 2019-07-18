-include .env

SAMPLE_FOLDER = sample/hello-world

new-sample:
	@go build 
	@rm -rf $(SAMPLE_FOLDER)
	@./typical-go new $(SAMPLE_FOLDER)
	
	