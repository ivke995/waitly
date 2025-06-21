.PHONY: watch-api
watch-api:
	@export $$(cat .env | xargs) && \
	/home/ivan/go/bin/air -c .air.toml