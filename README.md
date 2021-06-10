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

- https://jishuin.proginn.com/p/763bfbd3aa2e
- https://blog.lawrencejones.dev/golang-embed
- https://www.icode9.com/content-4-875160.html
- https://harsimranmaan.medium.com/embedding-static-files-in-a-go-binary-using-go-embed-bac505f3cb9a
