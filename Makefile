start:
	@echo "Starting app..."
	@docker compose up --build -d

stop:
	@echo "Stopping app..."
	@docker compose down

watch: start
	@echo "Watching for file changes..."
	@docker-compose watch

logs: 
	@docker-compose logs -f

create-network:
	@echo "Creating network..."
	@docker network create xtox-network

delete-network:
	@echo "Deleting network..."
	@docker network rm xtox-network

delete-data:
	@echo "Deleting data..."
	@docker volume rm xtox_db-data
	@docker volume rm xtox_cache-data
	@docker volume rm xtox_stream-data