# go-embed

[english readme](README.md)  
基于 Go 1.16 使用 Embed 嵌入 web 静态资源， 目录约定:

```
go-embed
└── web -- 前端项目，这里使用 vitejs + vues
     ├── dist -- 生产环境打包目录
     ├── ...
├── main.go -- 后端入口，这里使用 gin
├── ...
```

## 参考
- https://jishuin.proginn.com/p/763bfbd3aa2e
- https://blog.lawrencejones.dev/golang-embed
- https://www.icode9.com/content-4-875160.html
- https://harsimranmaan.medium.com/embedding-static-files-in-a-go-binary-using-go-embed-bac505f3cb9a
