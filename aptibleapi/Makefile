# This should be run from the root of the repo with make aptibleapi-client
client:
	docker run --rm -it -v "$(shell pwd):/local" openapitools/openapi-generator-cli:v7.6.0 \
	  generate -g go -i https://documentation-production.s3.amazonaws.com/openapi/v1/openapi.json \
		-o /local --template-dir /local/templates --name-mappings _type=MetaType \
		--git-user-id aptible --git-repo-id aptible-api-go/aptibleapi \
		--additional-properties packageName=aptibleapi,disallowAdditionalPropertiesIfNotPresent=false
