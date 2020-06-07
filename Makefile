.PHONY: migrate_db prepopulate_db start_db start_prepopulated_db stop_db start_server check_db

migrate_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/migrate_db.sh

prepopulate_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/prepopulate_db.sh

start_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/start_db.sh

start_prepopulated_db: ENV_FILE := cli-env
start_prepopulated_db: | start_db migrate_db prepopulate_db

stop_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/stop_db.sh

start_server: ENV_FILE := cli-env
start_server:
	ENV_FILE="${ENV_FILE}" \
	./scripts/start_server.sh

check_db: ENV_FILE := cli-env
check_db:
	ENV_FILE="${ENV_FILE}" \
	./scripts/check_db.sh
