package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	sseTransport, err := transport.NewSSE("http://127.0.0.1:8080/sse")
	if err != nil {
		fmt.Printf("Failed to create SSE transport: %v", err)
	}
	if err := sseTransport.Start(ctx); err != nil {
		fmt.Printf("Failed to start SSE transport: %v", err)
	}
	sseClient := client.NewClient(sseTransport)
	defer sseClient.Close()
	if err != nil {
		fmt.Println("Connect Server Failed!")
		return	
	}
	
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name: "client demo",
		Version: "1.0.0",
	}
	initResult, err := sseClient.Initialize(ctx, initRequest)
	if err != nil {
		fmt.Println("init client failed, err=", err.Error())
		return
	}
	fmt.Printf("初始化成功，服务器信息: %s %s\n", initResult.ServerInfo.Name, initResult.ServerInfo.Version)
	toolsRequest := mcp.ListToolsRequest{}
	tools, err := sseClient.ListTools(ctx, toolsRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println()
    fmt.Println("资源列表:")
    resourcesRequest := mcp.ListResourcesRequest{}
    resources, err := sseClient.ListResources(ctx, resourcesRequest)
    if err != nil {
        panic(err)
    }
    for _, resource := range resources.Resources {
        fmt.Printf("- uri: %s, name: %s, description: %s, MIME类型: %s\n", resource.URI, resource.Name, resource.Description, resource.MIMEType)
    }

	fmt.Println("工具列表:")
	for _, tool := range tools.Tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		fmt.Println("参数:", tool.InputSchema.Properties)
	}

	toolRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	toolRequest.Params.Name = "calculate"
	toolRequest.Params.Arguments = map[string]any{
		"operation": "add",
		"x":         1,
		"y":         1,
	}

	result, err := sseClient.CallTool(ctx, toolRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("调用工具结果:", result.Content[0].(mcp.TextContent).Text)
}