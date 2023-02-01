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