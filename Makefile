# Host Makefile.

include Makefile.include

env-up: env-compose-up env-devcontainer     ## Start devcontainer.

env-compose-up:
	docker-compose pull
	docker-compose up --detach --renew-anon-volumes --remove-orphans

env-devcontainer:
	docker exec -it --workdir=/root/go/src/github.com/percona/pmm pmm-managed-server .devcontainer/setup.py

env-down:                                   ## Stop devcontainer.
	docker-compose down --remove-orphans

env-remove:
	docker-compose down --volumes --remove-orphans


TARGET ?= _bash

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
	make format
	bin/go-sumtype ./...
	go install -v ./...

gen-alertmanager:     # Generate Alertmanager client.
	bin/swagger generate client --model-package=ammodels --client-package=amclient --spec=api/alertmanager/openapi.yaml --target=api/alertmanager

	bin/gofumpt  -l -w ./api/alertmanager
	go install -v ./api/alertmanager/...

gen-grafana:     # Generate Grafana Unified Alerting API client.
	bin/swagger generate client --model-package=gmodels --client-package=gclient --spec=api/grafana/swagger.yaml --target=api/grafana

	bin/gofumpt  -l -w ./api/grafana
	go install -v ./api/grafana/...

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
	bin/gofumpt -l -w .
	bin/goimports -local github.com/percona/pmm -l -w .
	bin/gci write --Section Standard --Section Default --Section "Prefix(github.com/percona/pmm)" .

check:                ## Run required checkers and linters.
	bin/golangci-lint run -c=.golangci.yml
	bin/go-consistent -pedantic ./...

serve:                ## Serve API documentation with nginx.
	# http://127.0.0.1:8080/swagger-ui.html
	nginx -p . -c api/nginx/nginx.conf

descriptors:          ## Update API compatibility descriptors.
	#./prototool break descriptor-set . -o api/api.descriptor
	bin/buf build -o descriptor.bin --as-file-descriptor-set api
env:                                        ## Run `make TARGET` in devcontainer (`make env TARGET=help`); TARGET defaults to bash.
	docker exec -it --workdir=/root/go/src/github.com/percona/pmm pmm-managed-server make $(TARGET)
