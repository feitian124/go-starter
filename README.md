# go-starter

This is a skeleton project for a go application, espeically web application, which captures the best practise including:

- following suggested project layout `cmd`, `internal` and easy to extend.  
   there are some discussions about if `pkg` should be used, see [issue #1](https://github.com/feitian124/go-starter/issues/1).
   if you prefer less folder nest, please ignore it; if you want your main folder more clean, you can use it.

- uses a makefile to drive the build and a dockerfile to build a docker image.

- uses `embed` to package web resources to one binary file

folder structure:
```
go-starter
└── web -- web app，here use vitejs + vue3 for example
     ├── dist -- product build output
     ├── build.go -- embed.FS
     ├── ...
└── cmd
     ├── main.go -- backend enry，here use gin for exmple
├── ...
```

## start
1. git clone https://github.com/feitian124/go-starter

2. build web
```shell
cd web/
yarn
yarn build
cd ..
```

this will generate web product build to web/dist

3. start go server, **including the web**!
```shell
± |main ✓| → go run cmd/app/main.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /                         --> main.setupRouter.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080

```

## Building

### Customizing it

To use this, simply copy these files and make the following changes:

Makefile:
   - change `BINS` to your binary name(s)
   - replace `cmd/myapp-*` with one directory for each of your `BINS`
   - change `REGISTRY` to the Docker registry you want to use
   - maybe change `SRC_DIRS` if you use some other layout
   - choose a strategy for `VERSION` values - git tags or manual

Dockerfile.in:
   - maybe change or remove the `USER` if you need

### How to build

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all supported architectures.

Run `make push` to push the container image to `REGISTRY`.  Run `make all-push`
to push the container images for all architectures.

Run `make clean` to clean up.

Run `make help` to get a list of available targets.

## references
### go embed
- [Go 1.16 使用 Embed 嵌入静态资源 | 视频文字稿](https://jishuin.proginn.com/p/763bfbd3aa2e)
- [Embed a Javascript website inside a binary with Go 1.16](https://blog.lawrencejones.dev/golang-embed)
- [Golang 1.16新特性-embed包及其使用](https://www.cnblogs.com/niuben/p/14461973.html)
- [Embedding static files in a go binary using go embed](https://harsimranmaan.medium.com/embedding-static-files-in-a-go-binary-using-go-embed-bac505f3cb9a)

### project structure
- https://github.com/golang-standards/project-layout.git
- https://github.com/katzien/go-structure-examples
- https://github.com/thockin/go-build-template
- https://github.com/1024casts/snake
