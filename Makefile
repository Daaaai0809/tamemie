migrate_up:
	@echo "Migrating database..."
	sh ./scripts/migrate_up.sh

migrate_down:
	@echo "Rolling back database..."
	sh ./scripts/migrate_down.sh

# Usage: make create_migration name=<migration_name>
create_migration:
	@echo "Creating migration..."
	sh ./scripts/create_migration.sh $(name)

generate_model:
	@echo "Generating DB Model & Queries..."
	sh ./scripts/generate_model.sh
