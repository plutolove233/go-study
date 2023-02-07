## metadata
1. 定义在RPC请求和响应过程中需要但是不属于业务的信息（例如身份验证...）。采用键值对的形式保存数据
   ```go
    type MD map[string][]string
    ```
   gRPC中的 metadata 类似于我们在 HTTP headers中的键值对，元数据可以包含认证token、请求标识和监控标签等。
2. metadata中的大小写不敏感，由字母和特殊字符-、_、. 组成。
   <br>
   <font color="red">**不能以grpc-开头**</font>
   <br>
   二进制值的键值名必须以 **-bin** 结尾
3. 元数据对 gRPC 本身是不可见的，我们通常是在应用程序代码或中间件中处理元数据，我们不需要在.proto文件中指定元数据。
4. 元数据的处理方式：
   1. 获取元数据：
      - 服务端：
        ```go
         // 从上下文中获取

         // 普通调用
         func (s *server)SomeRPC(ctx context.Context, in *pb.someReq) (*pb.someResp, error){
         md, ok := metadata.FromInComingContext(ctx)
         }
      
         // 流式调用
         func (s *server) SomeStreamRPC(stream pb.Service_SomeStreamRPCServer) error{
         md, ok := metadata.FromInComingContext(stream.Context())
         }
         ```
      - 客户端：
        ```go
        // 普通调用
        r, err := client.SomeRPC(
           ctx,
           someReq,
           grpc.Header(&header),
           grpc.Trailer(&trailer),
        )
                 
        // 流式调用
        stream, err := client.SomeStreamRPC(ctx)
        header, err := stream.Header()
        trailer, err := stream.Trailer()
        ```
   2. 发送元数据：
      - 服务端：
        ```go
        // 普通调用
        header := metadata.Pairs("header-key", "val")
        grpc.SendHeader(ctx, header)
        trailer := metadata.Pairs("trailer-key", "val")
        grpc.SetTrailer(ctx, trailer)
                 
        // 流式调用
        header := metadata.Pairs("header-key", "val")
        stream.SetHeader(header)
        trailer := metadata.Pairs("trailer-key", "val")
        stream.SetTrailer(trailer)
        ```
      - 客户端：
        ```go
        // 创建带有metadata的context
        ctx := metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")
        // 添加一些 metadata 到 context (e.g. in an interceptor)
        ctx := metadata.AppendToOutgoingContext(ctx, "k3", "v4")
        // 发起普通RPC请求
        response, err := client.SomeRPC(ctx, someRequest)
        ```