export OPENAPI_GENERATOR_TAG=v7.9.0

default: production

production:
	./scripts/generate-client.sh production openapi.json

staging:
	./scripts/generate-client.sh staging openapi.json

pr-%:
	./scripts/generate-client.sh staging dev/$@.json

# Pull templates from the generator for customization
.PHONY: templates
templates:
	docker run --rm -it -v "$(shell pwd):/local" "openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_TAG)" \
		author template -g go -o /local/source-templates