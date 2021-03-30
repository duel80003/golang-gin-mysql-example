# Build Executable File

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .
```

# Build Docker Image

```
docker build -t hhong/go-gin-aip-scratch .
```

# Run Image

```
docker run -p 8080:8080 -d hhong/go-gin-aip-scratch 
```

# API Doc
[Swagger Doc](https://app.swaggerhub.com/apis-docs/duel80003/gin-mysql-example/1.0.0)

