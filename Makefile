default: help

help:                 ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

init:                 ## Install tools.
	rm -rf bin
	go build -modfile=tools/go.mod -o bin/buf github.com/bufbuild/buf/cmd/buf
	go build -modfile=tools/go.mod -o bin/protoc-gen-go google.golang.org/protobuf/cmd/protoc-gen-go
	go build -modfile=tools/go.mod -o bin/protoc-gen-go-grpc google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go build -modfile=tools/go.mod -o bin/protoc-gen-govalidators github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	go build -modfile=tools/go.mod -o bin/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go build -modfile=tools/go.mod -o bin/protoc-gen-swagger github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go build -modfile=tools/go.mod -o bin/swagger github.com/go-swagger/go-swagger/cmd/swagger
	go build -modfile=tools/go.mod -o bin/swagger-order github.com/Percona-Lab/swagger-order
	go build -modfile=tools/go.mod -o bin/go-sumtype github.com/BurntSushi/go-sumtype

	# Download third-party proto files
	$(eval GO_PROTO_VALIDATOR=$(shell go list -f '{{ .Version }}' -m github.com/mwitkow/go-proto-validators))
	curl --create-dirs -L https://raw.githubusercontent.com/mwitkow/go-proto-validators/$(GO_PROTO_VALIDATOR)/validator.proto -o ../third_party/github.com/mwitkow/go-proto-validators/validator.proto


gen: clean         ## Generate files.
#	# generated by descriptors target
	#./bin/prototool break check api -f api/api.descriptor
	bin/buf breaking --against descriptor.bin api

	bin/buf generate -v api

	for API in api/agentlocalpb api/serverpb api/inventorypb api/managementpb api/managementpb/dbaas api/managementpb/ia api/managementpb/backup api/managementpb/azure api/qanpb api/platformpb ; do \
		set -x ; \
		bin/swagger mixin $$API/json/header.json $$API/*.swagger.json --output=$$API/json/$$(basename $$API).json --keep-spec-order; \
		bin/swagger flatten --with-flatten=expand --with-flatten=remove-unused $$API/json/$$(basename $$API).json --output=$$API/json/$$(basename $$API).json ; \
		bin/swagger validate $$API/json/$$(basename $$API).json ; \
		bin/swagger generate client --with-flatten=expand --with-flatten=remove-unused --spec=$$API/json/$$(basename $$API).json --target=$$API/json \
			--additional-initialism=aws \
			--additional-initialism=db \
			--additional-initialism=ok \
			--additional-initialism=pmm \
			--additional-initialism=psmdb \
			--additional-initialism=pxc \
			--additional-initialism=pt \
			--additional-initialism=qan \
			--additional-initialism=rds \
			--additional-initialism=sql \
			--additional-initialism=ha ; \
	done

	# generate public API spec, omit agentlocalpb (always private),
	# and managementpb/dbaas, managementpb/ia, managementpb/backup , managementpb/azure and qanpb (not v1 yet)
	bin/swagger mixin --output=api/swagger/swagger.json \
		api/swagger/header.json \
		api/serverpb/json/serverpb.json \
		api/inventorypb/json/inventorypb.json \
		api/managementpb/json/managementpb.json
	bin/swagger validate api/swagger/swagger.json

	bin/swagger-order --output=api/swagger/swagger.json api/swagger/swagger.json

	# generate API spec with all PMM Server APIs (omit agentlocalpb)
	bin/swagger mixin --output=api/swagger/swagger-dev.json \
		api/swagger/header-dev.json \
		api/serverpb/json/serverpb.json \
		api/inventorypb/json/inventorypb.json \
		api/managementpb/json/managementpb.json \
		api/managementpb/dbaas/json/dbaas.json \
		api/managementpb/ia/json/ia.json \
		api/managementpb/backup/json/backup.json \
		api/managementpb/azure/json/azure.json \
		api/qanpb/json/qanpb.json \
		api/platformpb/json/platformpb.json
	bin/swagger validate api/swagger/swagger-dev.json

	bin/swagger-order --output=api/swagger/swagger-dev.json api/swagger/swagger-dev.json

	# generate API spec with only dev PMM Server APIs specifically for readme.io (omit agentlocalpb)
	bin/swagger mixin --output=api/swagger/swagger-dev-only.json \
		api/swagger/header-dev.json \
		api/managementpb/dbaas/json/dbaas.json \
		api/managementpb/ia/json/ia.json \
		api/managementpb/backup/json/backup.json \
		api/managementpb/azure/json/azure.json \
		api/qanpb/json/qanpb.json \
		api/platformpb/json/platformpb.json
	bin/swagger validate api/swagger/swagger-dev-only.json

	bin/swagger-order --output=api/swagger/swagger-dev-only.json api/swagger/swagger-dev-only.json

	make clean_swagger
	go fmt ./...
	bin/go-sumtype ./...
	go install -v ./...

gen-alertmanager:     # Generate Alertmanager client.
	bin/swagger generate client --model-package=ammodels --client-package=amclient --spec=api/alertmanager/openapi.yaml --target=api/alertmanager

	go fmt ./api/alertmanager/...
	go install -v ./api/alertmanager/...

clean_swagger:
	find api -name '*.swagger.json' -print -delete

clean: clean_swagger  ## Remove generated files.
	find api -name '*.pb.go' -print -delete
	find api -name '*.pb.gw.go' -print -delete

	for API in api/agentlocalpb api/serverpb api/inventorypb api/managementpb api/managementpb/dbaas api/managementpb/ia api/managementpb/backup api/qanpb api/platformpb ; do \
		rm -fr $$API/json/client $$API/json/models $$API/json/$$(basename $$API).json ; \
	done
	rm -f api/swagger/swagger.json api/swagger/swagger-dev.json api/swagger/swagger-dev-only.json

test:                 ## Run tests
	go test ./...

format:               ## Format source code
	go fmt ./...

serve:                ## Serve API documentation with nginx.
	# http://127.0.0.1:8080/swagger-ui.html
	nginx -p . -c api/nginx/nginx.conf

descriptors:          ## Update API compatibility descriptors.
	#./prototool break descriptor-set . -o api/api.descriptor
	bin/buf build -o descriptor.bin --as-file-descriptor-set api
