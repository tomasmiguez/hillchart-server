### Run on development

```sh
go run .
```

### Build for production

```sh
go build -tags netgo -ldflags '-s -w' -o app
```
