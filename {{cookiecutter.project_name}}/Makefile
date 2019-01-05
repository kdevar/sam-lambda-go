{% set camel_name = cookiecutter.project_name|replace("-", " ")|title|replace(" ", "") -%}
include config.mk

name = {{cookiecutter.project_name}}
stack_name = {{cookiecutter.project_name}}

init:
	cd api && GOOS=linux GOARCH=amd64 GO111MODULE=on go mod init

build-local:
	GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o ../artifacts/$(name) ./api

build: clean
	cd api && GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ../artifacts/$(name)

zip:
	cd artifacts && zip $(name).zip $(name)

serve: build
	sam local start-api --template ./cloudformation.yaml

lint:
	go vet ./...

.PHONY: test
test:
	aws cloudformation validate-template --template-body file://cloudformation.yaml
	go test ./... -v

.PHONY: deploy
deploy: build zip
	sam package \
		--template-file ./cloudformation.yaml \
		--output-template-file ./cloudformation_deploy.yaml \
		--s3-bucket $(ENV)-$(ACCOUNT_ID)-serverless-code
	sam deploy \
		--no-fail-on-empty-changeset \
		--capabilities CAPABILITY_NAMED_IAM \
		--template-file ./cloudformation_deploy.yaml \
		--stack-name $(ENV)-$(stack_name) \
		--parameter-overrides EnvironmentType=$(ENV)

deploy_sandbox: $(eval ENV := personal) deploy

geturl:
	@echo "find your api here: $(shell aws cloudformation list-exports --query "Exports[?Name==\`{{camel_name}}ApiUrl\`].Value" --no-paginate --output text)"

.PHONY: deploy_wait
deploy_wait: ;

.PHONY: teardown
teardown:
	aws cloudformation delete-stack --stack-name $(stack_name)

.PHONY: teardown_wait
teardown_wait:
	aws cloudformation wait stack-delete-complete --stack-name $(stack_name)

.PHONY: clean
clean:
	rm -f artifacts/$(name)
	rm -rf artifacts/$(name).zip
