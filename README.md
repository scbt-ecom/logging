# Logger implementation for Совкомбанк Технологии

## Getting started
```bash
go get ...
```
## Development

### Logger initialization
```
log := logging.InitLogger()
```

### Adding a hook for graylog
```
logging.AddGraylogHook(log, <graylogUrl>, <containerName>)
```