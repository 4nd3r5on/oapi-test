# OpenAPI playground

My learning repo to learn about openapi, how to work with it, hot to integrate it into my flow.

```sh
# Bundle OpenAPI coduments 
npx @redocly/cli bundle ./pkg/api/openapi/api.yml --ext yml -o ./pkg/api/api.yml
# Run codegen
go generate ./...
```

### How did I setup it

All the OpenAPI-related documents are at `./pkg/api/openapi/`

#### Kin OpenAPI

[Kin OpenAPI](github.com/getkin/kin-openapi)

Used to integrate it into the app.

Before running codegen or importing to the project -- run [OpenAPI document validation](https://github.com/getkin/kin-openapi?tab=readme-ov-file#validating-an-openapi-document)

`go run github.com/getkin/kin-openapi/cmd/validate@latest -- <files>`

#### Codegen

[install oapi keygen](https://github.com/oapi-codegen/oapi-codegen/?tab=readme-ov-file#install)
```sh
go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

Since it's already done in this repo -- should work with just `go mod download`
