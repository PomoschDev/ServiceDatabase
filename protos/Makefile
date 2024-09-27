PKG_FOLDER = ..\pkg\DatabaseServicev1\

generate:
		@echo Step 1: Generate proto files...
		@protoc -I proto proto/DatabaseService/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
		@echo Files created...
		@echo Step 2: Copy files to $(PKG_FOLDER)
		@copy .\gen\go\DatabaseService\*.go $(PKG_FOLDER)
		@echo Successfully!!!