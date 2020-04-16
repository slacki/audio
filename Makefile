default:
	@echo ">>> [DEV] Composing default"
	docker-compose up -d

up: down
	@echo ">>> [DEV] Composing default"
	docker-compose up -d

build:
	@echo ">>> [DEV] Building default"
	docker-compose build

debug:
	@echo ">>> [DEV] Composing with debugger"
	docker-compose -f docker-compose.yaml -f docker-compose.debug.yaml up api

logs_api: 
	@echo ">>> [DEV] Follow API logs"
	docker-compose logs -f api

cleanup:
	@echo ">>> [DEV] Cleaning up _uploads and _data"

down: 
	@echo ">>> [DEV] Going down"
	docker-compose down