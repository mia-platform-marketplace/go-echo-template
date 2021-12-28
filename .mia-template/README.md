# mia_template_service_name_placeholder

This is a simple Go application template with a pre-configured [logger]("https://github.com/mia-platform/glogger") and various other configurations.
It also contains basic dependencies for testing and http request.
By default the module name is "service", if you want to change it, please remember to change it in the imports too.

## Development Local

To develop the service locally you need:
    - Go 1.17+

To start the application locally

```go
go run mia_template_service_name_placeholder
```

By default the service will run on port 8080, to change the port please set `HTTP_PORT` env variable

## Testing

To test the application use:

```go
go test -v
```