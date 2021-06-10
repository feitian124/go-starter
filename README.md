# go-embed

[中文说明](README-cn.md)  
Use Go 1.16 `embed` feature to include web resource to binary file.

folder structure:
```
go-embed
└── web -- web app，here use vitejs + vue3 for example
     ├── dist -- product build output
     ├── ...
├── main.go -- backend enry，here use gin for exmple
├── ...
```

## references

- [Go 1.16 使用 Embed 嵌入静态资源 | 视频文字稿](https://jishuin.proginn.com/p/763bfbd3aa2e)
- [Embed a Javascript website inside a binary with Go 1.16](https://blog.lawrencejones.dev/golang-embed)
- [Golang 1.16新特性-embed包及其使用](https://www.cnblogs.com/niuben/p/14461973.html)
- [Embedding static files in a go binary using go embed](https://harsimranmaan.medium.com/embedding-static-files-in-a-go-binary-using-go-embed-bac505f3cb9a)
