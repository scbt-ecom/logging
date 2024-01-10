# Logger implementation for Совкомбанк Технологии

## Getting started
```bash
go get github.com/skbt-ecom/logging
```
## Development

### Logger initialization

```
log := logging.InitLogger(<logLevel>)
```

### Adding a hook for graylog
```
logging.AddGraylogHook(log, <graylogUrl>, <containerName>)
```