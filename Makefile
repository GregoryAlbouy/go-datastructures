tests:
	@t=singlylinkedlist make runtest
	@t=doublylinkedlist make runtest
	@t=queue make runtest
	@t=stack make runtest
	@t=binarysearchtree make runtest
	@t=binaryheap make runtest
	@t=graph make runtest

runtest:
	@cd $(t) && go test