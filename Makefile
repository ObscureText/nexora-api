run_migration:
	liquibase \
		--url="jdbc:postgresql://localhost:5432/db_nexora_local" \
		--username=postgres \
		--password=postgres \
		--changeLogFile=migrations/changelog.yaml \
		update