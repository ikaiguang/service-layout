override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_TESTING_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_TESTING_INTERNAL_PROTO := "app/testing-service/internal/conf/config.conf.proto"
SAAS_TESTING_PROTO_FILES := ""
ifneq ($(SAAS_TESTING_INTERNAL_PROTO), "")
	SAAS_TESTING_PROTO_FILES=$(SAAS_TESTING_API_PROTO) $(SAAS_TESTING_INTERNAL_PROTO)
else
	SAAS_TESTING_PROTO_FILES=$(SAAS_TESTING_API_PROTO)
endif
.PHONY: protoc-testing-protobuf
# protoc :-->: generate testing service protobuf
protoc-testing-protobuf:
	@echo "# generate testdata service protobuf"
	$(call protoc_protobuf,$(SAAS_TESTING_PROTO_FILES))
