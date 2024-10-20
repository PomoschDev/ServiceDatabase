PKG_FOLDER = ..\pkg\DatabaseServicev1\

generate_go:
		@echo Step 1: Generate proto files...
		@protoc -I proto proto/DatabaseService/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
		@echo Files created...
		@echo Step 2: Copy files to $(PKG_FOLDER)
		@copy .\gen\go\DatabaseService\*.go $(PKG_FOLDER)
		@echo Successfully!!!

generate_python:
		@echo Step 1: Generate proto files...
		@python -m grpc_tools.protoc -I proto proto/DatabaseService/*.proto --python_out=./gen/python/ --pyi_out=./gen/python/ --grpc_python_out=./gen/python/
		@echo Successfully!!!
		
linux_generate_go:
		@echo Step 1: Generate proto files...
		@protoc -I proto proto/DatabaseService/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
		@echo Files created...
		@echo Step 2: Copy files to $(PKG_FOLDER)
		@cp .\gen\go\DatabaseService\*.go $(PKG_FOLDER)
		@echo Successfully!!!