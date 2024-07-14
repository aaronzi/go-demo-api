# Go Demo API

## Rebuild the Swagger Documentation

To build the Swagger documentation, run the following command:

```bash
swag init -g cmd/movie-api/main.go --parseDependency --parseInternal -o docs
```

This will generate the `docs` directory with the Swagger documentation.

> **Note:** You should run this command every time you make changes to the API.
