### Go Practices

creating an example:
```sh
mkdir example/folder_name
cd example/folder_name
go mod init example/module_name
touch example.go

```

Running examples
```sh
go run example.go
```

If your project couldn't find a dependency, use:

```sh
go mod tidy
```