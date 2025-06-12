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
	@echo ">>> Checking for Nexus SDK distribution files..."
	@if [ ! -d "${SDK_DIR}" ]; then \
		echo "    ERROR: Please make sure the Nexus SDK distribution files are unpacked to the ${SDK_DIR} folder"; \
		echo ""; \
		exit 1; \
	fi
	@echo ""
	@echo ">>> Checking for Docker 'multi-platform' builder"
	@echo ""
	@if ! docker buildx ls | grep 'multi-platform-builder'; then \
		docker buildx create \
			--name "multi-platform-builder" \
			--driver=docker-container \
			--bootstrap --use; \
	fi

container: init
	docker buildx build \
		--load \
		--tag python/s1scanner:latest \
		-f "${THIS_DIR}/build/Dockerfile" \
		${THIS_DIR}/..

nosdk-container: init
	docker buildx build \
		--push \
		--builder multi-platform-builder \
		--tag ghcr.io/s1integrations/nexus-sdk/python/s1scanner:latest \
		--platform linux/arm64,linux/amd64 \
		-f ${THIS_DIR}/build/Dockerfile.nosdk \
		${THIS_DIR}/..
