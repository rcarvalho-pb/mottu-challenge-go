DB_LOCATION="../data-storage/db.db"

USER_SERVICE_BINARY=userApp
USER_SERVICE_ADDRESS=12345

MOTORCYCLE_SERVICE_BINARY=motorcycleApp
MOTORCYCLE_SERVICE_ADDRESS=12348

TOKEN_SERVICE_BINARY=tokenApp
TOKEN_SERVICE_ADDRESS=12346

AUTH_SERVICE_BINARY=authApp
AUTH_SERVICE_ADDRESS=12347

LOCATION_SERVICE_BINARY=locationApp
LOCATION_SERVICE_ADDRESS=12348

migration:
	@goose create -dir ./config_databases/files/migrations ${name} sql

config-db:
	@echo "Adjusting db"
	@cd ./config_databases/ && go run main.go
	@echo "DB adjusted"

motorcycle-service:
	@echo "Starting motorcycle service"
	@cd ./motorcycle_service/ && go build -o app/${MOTORCYCLE_SERVICE_BINARY} ./cmd/api
	@cd ./motorcycle_service/ && export DB_LOCATION=${DB_LOCATION} USER_SERVICE_ADDRESS=${MOTORCYCLE_SERVICE_ADDRESS} && app/${USER_SERVICE_BINARY} &
	@echo "Motorcycle service started on port ${MOTORCYCLE_SERVICE_ADDRESS}"

user-service:
	@echo "Starting user service"
	@cd ./user_service/ && go build -o app/${USER_SERVICE_BINARY} ./cmd/api
	@cd ./user_service/ && export DB_LOCATION=${DB_LOCATION} USER_SERVICE_ADDRESS=${USER_SERVICE_ADDRESS} && app/${USER_SERVICE_BINARY} &
	@echo "User service started on port ${USER_SERVICE_ADDRESS}"

location-service:
	@echo "Starting location service"
	@cd ./location_service/ && go build -o app/${LOCATION_SERVICE_BINARY} ./cmd/api
	@cd ./location_service/ && export DB_LOCATION=${DB_LOCATION} LOCATION_SERVICE_ADDRESS=${LOCATION_SERVICE_ADDRESS} && app/${LOCATION_SERVICE_BINARY} &
	@echo "User location started on port ${LOCATION_SERVICE_ADDRESS}"

token-service:
	@echo "Starting token service"
	@cd ./token_service/ && go build -o app/${TOKEN_SERVICE_BINARY} ./cmd/api
	@cd ./token_service/ && export TOKEN_SERVICE_ADDRESS=${TOKEN_SERVICE_ADDRESS} && app/${TOKEN_SERVICE_BINARY} &
	@echo "Token service started on port ${TOKEN_SERVICE_ADDRESS}"

auth-service:
	@echo "Starting auth service"
	@cd ./authentication_service/ && go build -o app/${AUTH_SERVICE_BINARY} ./cmd/api
	@cd ./authentication_service/ && export USER_SERVICE_ADDRESS=${USER_SERVICE_ADDRESS} TOKEN_SERVICE_ADDRESS=${TOKEN_SERVICE_ADDRESS} AUTH_SERVICE_ADDRESS=${AUTH_SERVICE_ADDRESS} && app/${AUTH_SERVICE_BINARY} &
	@echo "Auth service started on port ${AUTH_SERVICE_ADDRESS}"

teste:
	@cd ./test && go run main.go
run: user-service token-service auth-service
stop:
	@echo "Stoping services"
	pkill -SIGTERM -f "${USER_SERVICE_BINARY}"
	pkill -SIGTERM -f "${TOKEN_SERVICE_BINARY}"
	pkill -SIGTERM -f "${AUTH_SERVICE_BINARY}"
	@echo "Services finisheds"
