# Docker Compose işlemleri için Makefile
# docker/ dizininde yer alan docker-compose.yml ve .env dosyalarını kullanmak üzere düzenlendi.

DOCKER_COMPOSE_DIR=dockercompose
DOCKER_COMPOSE=docker-compose -f $(DOCKER_COMPOSE_DIR)/docker-compose.yml --env-file $(DOCKER_COMPOSE_DIR)/.env

up:
	@$(DOCKER_COMPOSE) up --build -d
	@echo "Docker Compose is up and running."

down:
	@$(DOCKER_COMPOSE) down
	@echo "Docker Compose has been stopped."

restart:
	@$(DOCKER_COMPOSE) down
	@$(DOCKER_COMPOSE) up --build -d
	@echo "Docker Compose has been restarted."

logs:
	@$(DOCKER_COMPOSE) logs -f

ps:
	@$(DOCKER_COMPOSE) ps

pull:
	@$(DOCKER_COMPOSE) pull
	@echo "Images have been pulled."

clean:
	@$(DOCKER_COMPOSE) down --rmi all -v --remove-orphans
	@echo "Docker Compose cleaned up. Removed images, volumes, and orphans."

# Makefile hedeflerini açıklama
help:
	@echo "Usage:"
	@echo "  make up        - Start Docker Compose in detached mode (background)"
	@echo "  make down      - Stop Docker Compose and remove containers"
	@echo "  make restart   - Restart Docker Compose"
	@echo "  make logs      - View Docker Compose logs"
	@echo "  make ps        - List running Docker Compose containers"
	@echo "  make pull      - Pull the latest Docker images"
	@echo "  make clean     - Remove all containers, images, volumes, and orphans"
