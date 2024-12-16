AuthBin=authApp
UserApp=userApp

migration:
	@goose create -dir ./config_databases/files/migrations ${name} sql

run-config-db:
	@cd config_databases/ && go build -o ./app/configDBApp .
	@cd config_databases && ./app/configDBApp

run-authentication-service:
	@cd authentication_service && go build -o ./app/${AuthBin} ./cmd/api && ./app/${AuthBin}

run-user-service:
	@cd user_service && go build -o ./app/${UserApp} ./cmd/api && ./app/${UserApp}
