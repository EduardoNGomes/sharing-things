migrate-diff:
	atlas migrate diff --env gorm
migrate-up:
	atlas migrate apply --env gorm
db-inspect:
	atlas schema inspect --env gorm --url "env://src"
