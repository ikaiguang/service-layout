.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/testing-service/cmd/store-configuration/... -conf=./app/testing-service/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/testing-service/cmd/database-migration/... -conf=./app/testing-service/configs

.PHONY: run-testing-service
# run service :-->: run testing-service
run-testing-service:
	go run ./app/testing-service/cmd/testing-service/... -conf=./app/testing-service/configs

.PHONY: testing-testing-service
# testing service :-->: testing testing-service
testing-testing-service:
	curl http://127.0.0.1:9991/api/v1/testdata/get?message=hello

.PHONY: run-service
# run service :-->: run testing-service
run-service:
	#@$(MAKE) run-testing-service
	go run ./app/testing-service/cmd/testing-service/... -conf=./app/testing-service/configs

.PHONY: testing-service
# testing service :-->: testing testing-service
testing-service:
	curl http://127.0.0.1:9991/api/v1/testdata/get?message=hello

