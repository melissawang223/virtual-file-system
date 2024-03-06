DC :=
GO       ?= go

MAIN_CMD = "./main.go"

run: ## Run application
	$(DC) $(GO) run $(MAIN_CMD)

