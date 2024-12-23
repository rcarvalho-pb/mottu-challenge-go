USER_SERVICE_BINARY=userApp
USER_SERVICE_ADDRESS=12345

TOKEN_SERVICE_BINARY=tokenService
TOKEN_SERVICE_ADDRESS=12346

migration:
	@goose create -dir ./config_databases/files/migrations ${name} sql

user-service:
	@echo "Starting user service"
	@cd ./user_service/ && go build -o app/${USER_SERVICE_BINARY} ./cmd/api
	@cd ./user_service/ && export USER_SERVICE_ADDRESS=${USER_SERVICE_ADDRESS} && app/${USER_SERVICE_BINARY} &
	@echo "User service started on port ${USER_SERVICE_ADDRESS}"

token-service:
	@echo "Starting token service"
	@cd ./token_service/ && go build -o app/${TOKEN_SERVICE_BINARY} ./cmd/api
	@cd ./token_service/ && export USER_SERVICE_ADDRESS=${TOKEN_SERVICE_ADDRESS} && app/${TOKEN_SERVICE_BINARY} &
	@echo "Token service started on port ${TOKEN_SERVICE_ADDRESS}"


teste:
	@cd ./test && go run main.go
run: user-service token-service
stop:
	@echo "Stoping services"
	pkill -SIGTERM -f "${USER_SERVICE_BINARY}"
	pkill -SIGTERM -f "${TOKEN_SERVICE_BINARY}"
	@echo "Services finisheds"
