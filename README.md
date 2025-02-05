# Logger implementation for Совкомбанк Технологии

## Getting started
```bash
go get github.com/skbt-ecom/logging@v1.1.4
```
## Development

### Logger initialization

```
log := logging.InitLogger()
```

### Set log level
#### Optionally, default log level - Trace
```
log.SetLevel(<logLevel>)
```

### Adding a hook for graylog
```
log.AddGraylogHook(<graylogUrl>, <containerName>)
```
