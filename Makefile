default: production

production:
	./scripts/generate-client.sh production openapi.json

staging:
	./scripts/generate-client.sh staging openapi.json

pr-%:
	./scripts/generate-client.sh staging dev/$@.json
