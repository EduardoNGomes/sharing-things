migrate-diff:
	atlas migrate diff --env gorm
migrate-up:
	atlas migrate apply --env gorm

.PHONY: migrate-up-tests
migrate-up-tests:
	@if [ -z "$$DATABASE_URL" ]; then \
		echo "DATABASE_URL is required. Usage: make migrate-up-tests DATABASE_URL='<url>' REVISIONS_SCHEMA='<schema>'" >&2; \
		exit 2; \
	fi
	@if [ -z "$$REVISIONS_SCHEMA" ]; then \
		echo "REVISIONS_SCHEMA is required. Usage: make migrate-up-tests DATABASE_URL='<url>' REVISIONS_SCHEMA='<schema>'" >&2; \
		exit 2; \
	fi
	@atlas migrate apply \
		--dir "file://migrations" \
		--url "$$DATABASE_URL" \
		--revisions-schema "$$REVISIONS_SCHEMA"

db-inspect:
	atlas schema inspect --env gorm --url "env://src"
