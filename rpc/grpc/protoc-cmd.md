protoc参数基本含义
- --go_out用于指定protoc的工作方式和go代码生成的位置
  - 参数（用,分开）:
    1. plugins: 生成go代码所用插件
    2. paths: go代码生成的位置:
        - import:${path} 
          - 按照生成的go代码的包的全路径来创建目录层级
          - 例如:在demo.proto中定义了option go_package="project/demo"，那么就会在生成代码指令的路径下创建"${path}/project/demo/demo.pb.go"
        - source_relative:${path}
          - 按照**proto源文件的目录层级**去创建go代码的目录层级
          - 例如：demo.proto定在在"/pb/demo"目录下，当前目录在/pb，那么就会在生成代码指令的路径下创建"${path}/demo/demo.pb.go"
- --go-grpc_out与--go_out类似
- --proto_path/-I:指定proto文件的目录
- 如果你想编译所有proto文件（假设生成Go语言），正常的命令应该是这样的：
```
protoc --proto_path=.  --go_out=.  proto/*.proto proto/user/*proto proto/greeter/*proto
```
  但是有的朋友可能会想偷懒，想直接这样：
```
protoc --proto_path=.  --go_out=.  proto/*.proto
```
  答案是不行的，因为protoc-gen-go不支持这种形式，最终只会编译common.proto
