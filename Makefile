test: tree.c
	@$(CC) $^ -std=c99 -o $@
	@./test

.PHONY: test
