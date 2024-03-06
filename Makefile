# Makefile
DC :=
GO       ?= go

MAIN_CMD = "main.go"

# Define the target file
TARGET_FILE := "local.txt"

run:
	$(DC) $(GO) run $(MAIN_CMD)

init:
	echo "[]" > $(TARGET_FILE)

