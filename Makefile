migrate_up:
	@echo "Migrating database..."
	sh ./scripts/migrate_up.sh

migrate_down:
	@echo "Rolling back database..."
	sh ./scripts/migrate_down.sh
