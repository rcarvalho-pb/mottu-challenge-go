UserServiceAddress='12345'
TokenServiceAddress=12346
AuthApp=authApp
UserApp=userApp
TokenApp=tokenApp

migration:
	@goose create -dir ./config_databases/files/migrations ${name} sql

run-config-db:
	@cd config_databases/ && go build -o ./app/configDBApp .
	@cd config_databases && ./app/configDBApp

run-auth-service:
	@cd authentication_service && USER_SERVICE_PORT=${UserServiceAddress} TOKEN_SERVICE_PORT=${TokenServiceAddress} go build -o ./app/${AuthApp} ./cmd/api && ./app/${AuthApp}

run-user-service:
	@cd user_service && go build -o ./app/${UserApp} ./cmd/api && PORT=${UserServiceAddress} ./app/${UserApp}

run-token-service:
	@cd token_service && env PORT=${TokenServiceAddress} go build -o ./app/${TokenApp} ./cmd/api && ./app/${TokenApp}
