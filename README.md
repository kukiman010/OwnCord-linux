Это форк проекта [OwnCord](https://github.com/J3vb/OwnCord) слегка модифецированого для компиляции сервера под debian 12

сборка происходит примерно так же, как и для винды
``` bash
cd Server
go build -o chatserver -ldflags "-s -w -X main.version=1.0.0" .
```

Так же в этой сборке сервер работает на порту 8444
