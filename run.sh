#!/usr/bin/env bash

# Golang
docker build -t my-golang-app .
docker run -it --rm --name my-running-app my-golang-app

# Dev
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d

# Prod
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d


# Notes
# Base docker-compose.yml: Contains the foundational service definitions that you might share between environments (e.g., the database).
# docker-compose.dev.yml: Extends services using the same service names and provides overrides specific to development. This will often include volume mounts for hot-reloading and development-focused commands.
# docker-compose.prod.yml: Offers overrides tuned for a production environment (e.g., production-ready application servers).
# -d runs it in the background

#Key Points
#
#Overriding vs. Addition: Docker Compose intelligently merges multiple Compose files. If a property exists in both the base and an override file, the override takes precedence. You can also add entirely new services in these override files.
#Selective Startup: By using docker compose -f ... command, you control which configurations get applied, giving you tailored environments.
#Flexibility: Compose files can extend multiple levels if needed (e.g., you could have staging overrides separate from dev or prod).
#Important Considerations
#
#Environment Variables: Handle secrets and environment-specific configuration using environment variables and .env files.
#Networks: You may want to adjust or separate networks for development and production configurations.