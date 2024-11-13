# ping-service
.PHONY: run-all-in-one
# run service :-->: run all-in-one
run-all-in-one:
	go run ./app/all-in-one/main.go -conf=./app/testing-service/configs

.PHONY: testing-all-in-one
# testing service :-->: testing all-in-one
testing-all-in-one:
	curl http://127.0.0.1:9991/api/v1/ping/logger && echo "\n"
	curl http://127.0.0.1:9991/api/v1/ping/error && echo "\n"
	curl http://127.0.0.1:9991/api/v1/ping/panic && echo "\n"
	curl http://127.0.0.1:9991/api/v1/ping/say_hello && echo "\n"
	#@echo "\n"
	$(MAKE) testing-service
