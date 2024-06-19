# Aptible API Go Client

Go API client generated from Aptible's [OpenAPI specs](https://www.openapis.org/) using the [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator). See the [`examples/`](./examples) directory for examples on how to get started using the client and the generated [`aptibleapi/README.md`](./aptibleapi/README.md) for details on the available endpoints.

# Using the Client

Import the module:

```shell
go get github.com/aptible/aptible-api-go
```

Import the `aptibleapi` package in `.go` files:

```go
import "github.com/aptible/aptible-api-go/aptibleapi"
```

# Contributing

Use the various `make` targets to re-generate the client from the latest OpenAPI specs. See the `Makefile` for options.

Most of the client's contents are generated. If a generated file needs modified (see [`aptibleapi/.openapi-generator/FILES`](./aptibleapi/.openapi-generator/FILES) for a list of generated files) it should be done by modifying the generator's templates. See [`aptibleapi/templates`](./aptibleapi/templates) for details and examples.

# Release

To release the client, create a tagged release in GitHub following semantic versioning convention.
