.PHONY: watch-api start-api

# For local dev
watch-api:
	@export $$(cat .env | xargs) && \
	/home/ivan/go/bin/air -c .air.toml

# For Koyeb deployment
start-api:
	go run main.go
