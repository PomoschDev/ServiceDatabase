.PHONY: all

PKG_FOLDER = "../pkg/DatabaseServicev1/"
DIR_PROTO_FOR_DOCS = "./gen-doc/proto/"
DIR_GEN_DOC = "./gen-doc"
DIR_DOC = "../doc/"
DIR_WITH_GEN_DOC = "./gen-doc/doc/*.md"
DIR_PROTO = ".\proto\DatabaseService\*.proto"
COPY_STR = ".\gen-doc\doc\*.md"
CMD = $(shell more < "./gen-doc/cmd.txt")

all: gen_go gen_python

gen_go:
		@echo Step 1: Generate Golang files...
		@protoc -I proto proto/DatabaseService/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
		@echo Files created...
		@echo Step 2: Copy files to ${PKG_FOLDER}
		@copy .\gen\go\DatabaseService\*.go ${PKG_FOLDER} >NUL
		@echo Successfully!!!

gen_python:
		@echo Step 1: Generate Python files...
		@python -m grpc_tools.protoc -I proto proto/DatabaseService/*.proto --python_out=./gen/python/ --pyi_out=./gen/python/ --grpc_python_out=./gen/python/
		@echo Successfully!!!
		
linux_gen_go:
		@echo Step 1: Generate proto files...
		@protoc -I proto proto/DatabaseService/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
		@echo Files created...
		@echo Step 2: Copy files to ${PKG_FOLDER}
		@cp .\gen\go\DatabaseService\*.go ${PKG_FOLDER}
		@echo Successfully!!!

gen_docs:
		@echo Move proto file...
		@copy ${DIR_PROTO} ${DIR_PROTO_FOR_DOCS} >NUL
		@echo Generate docs...
		@cd ${DIR_GEN_DOC} && ${CMD}
		@echo Move doc file...
		@copy ${COPY_STR} ${DIR_DOC} >NUL

