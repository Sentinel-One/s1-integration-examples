DOCSERVER_PORT ?= 9000
THIS_DIR := $(realpath $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
SDK_DIR := $(realpath ${ROOT_DIR}../_distfiles_/SDK)
# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

all:

init:
	@echo ">>> Copying Nexus SDK distribution files..."
	@if [ ! -d "${SDK_DIR}" ]; then \
		echo "    ERROR: Please make sure the Nexus SDK distribution files are unpacked to the ${SDK_DIR} folder"; \
		echo ""; \
		exit 1; \
	fi
	@echo ""
	cp ${SDK_DIR}/include/libdfi.h ${THIS_DIR}/pkg/scanner/nexus/include/
	cp ${SDK_DIR}/lib/linux/arm64/libdfi.so ${THIS_DIR}/pkg/scanner/nexus/lib/linux/arm64/
	cp ${SDK_DIR}/lib/linux/x64/libdfi.so ${THIS_DIR}/pkg/scanner/nexus/lib/linux/amd64/
	@echo ""
	@echo ">>> Downloading required Go packages"
	@echo ""
	go mod tidy
	@echo ""
	@echo ">>> Checking for Docker 'multi-platform' builder"
	@echo ""
	@if ! docker buildx ls | grep 'multi-platform-builder'; then \
		docker buildx create \
			--name "multi-platform-builder" \
			--driver=docker-container \
			--bootstrap --use; \
	fi

s1scanner: init
	@mkdir -p "${THIS_DIR}/dist"
	go build -o "${THIS_DIR}/dist/s1scanner" "${THIS_DIR}/cmd/s1scanner"
	@if [ "$(uname -m)" = "x86_64" ]; then \
  		cp "${THIS_DIR}/pkg/scanner/nexus/lib/linux/amd64/libdfi.so" "${THIS_DIR}/dist/libnexus.unstripped.so"; \
	else \
  		cp "${THIS_DIR}/pkg/scanner/nexus/lib/linux/arm64/libdfi.so" "${THIS_DIR}/dist/libnexus.unstripped.so"; \
	fi

doc-server:
	pkgsite -http :${DOCSERVER_PORT}

clean:
	rm -rf "${THIS_DIR}/dist"

container: init
	docker buildx build \
		--load \
		--tag go/s1scanner:latest \
		-f "${THIS_DIR}/build/Dockerfile" \
		${THIS_DIR}/..

nosdk-container: init
	docker buildx build \
		--push \
		--builder multi-platform-builder \
		--tag ghcr.io/s1integrations/nexus-sdk/go/s1scanner:latest \
		--platform linux/arm64,linux/amd64 \
		-f ${THIS_DIR}/build/Dockerfile.nosdk \
		${THIS_DIR}/..
