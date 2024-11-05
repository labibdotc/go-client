cd /home/aafia/projects/go-client/priv-src/go-oapi-codegen/src ; /usr/bin/env /usr/lib/jvm/java-21-openjdk-amd64/bin/java -agentlib:jdwp=transport=dt_sock
et,server=n,suspend=y,address=localhost:34841 -cp /home/aafia/projects/go-client/priv-src/go-oapi-codegen.jar:/home/aafia/projects/go-client/priv-src/openapi-generator-cli.jar -XX:+S
howCodeDetailsInExceptionMessages -cp /home/aafia/.vscode-server/data/User/workspaceStorage/48f5cc05e11178bfd2ea309ae05334e8/redhat.java/jdt_ws/jdt.ls-java-project/bin org.openapitoo
ls.codegen.OpenAPIGenerator generate -i openapi.json.tmp -g go-oapi-codegen -o /home/aafia/projects/go-client/priv-src/onshape --type-mappings DateTime=JSONTime --additional-properti
es=packageVersion=\$\{packageVersion\} --additional-properties=useOneOfDiscriminatorLookup=true --api-name-suffix=Api --global-property apiTests=false -c /home/aafia/projects/go-clie
nt/priv-src/openapi_config.json 