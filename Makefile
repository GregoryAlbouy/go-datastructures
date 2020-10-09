test:
	@go test ./... --cover

runtest:
	@cd $(t) && go test
