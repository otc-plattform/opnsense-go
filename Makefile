# HAProxy uses a handwritten generic-object/action client. Its schema exists
# solely to register HAProxy() in the aggregate client, not to generate a
# controller implementation.
handwritten_modules := haproxy
modules := $(filter-out $(handwritten_modules),$(notdir $(basename $(wildcard schema/*.yml))))
generator := $(shell find internal/generate)

PKG ?=
TEST ?=

.PHONY: testacc
testacc:
ifdef PKG
	go test -v -p 1 -timeout 120m $(if $(TEST),-run $(TEST)) ./pkg/$(PKG)/...
else
	go test -v -p 1 -timeout 120m $(if $(TEST),-run $(TEST)) ./pkg/...
endif

pkg/opnsense/client.go: $(generator) $(wildcard schema/*.yml)
	@echo "Generating opnsense client"
	go generate ./pkg/opnsense

pkg/%/controller.go: schema/%.yml $(generator) pkg/%/generate.go
	@echo "Generating $* controller"
	go generate ./$(@D)

$(modules): %: pkg/%/controller.go

all: $(modules) pkg/opnsense/client.go
