AuthBin=authApp
UserApp=userApp

migration:
	@goose create -dir ${migration_files} ${name} sql

run-config-db:
	@cd config_databases/ && go build -o configDBApp .
	@cd config_databases && configDBApp

run-authentication-service:
	@cd authentication_service && go build -o ./app/${AuthBin} ./cmd/api && ./app/${AuthBin}

run-user-service:
	@echo "Starting user service"
	@cd user_service && go build -o ./app/${UserApp} ./cmd/api && ./app/${UserApp}
	@echo "User service started"
