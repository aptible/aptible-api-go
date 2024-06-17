# Aptible API Go Clients

Clients are generated from Aptible's [OpenAPI specs](https://www.openapis.org/) using the [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator). See the `examples/` directory in each client directory for examples on how to get started using the client and the generated `README.md`.

## aptibleapi

Client for managing Aptible resources such as Apps, Databases, and Environments.

# Contributing

## Updating Clients

Use the `%-client` `make` target to re-generate clients from the latest OpenAPI spec. For example, to update the `aptibleapi` client use `make aptibleapi-client`.

## Adding Clients

To add a new OpenAPI-generated client, create a new directory with a `Makefile` containing a `client` target.

Be sure to include at least one example in the client's `examples/` showing how to authenticate and configure the client.

# Release

To release the clients, create a tagged release in GitHub following semantic versioning convention.
