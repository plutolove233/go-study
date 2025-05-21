package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("server", "1.0.0")
	// 添加工具
	{
		calculateTool := mcp.NewTool(
			"calculate",
			mcp.WithDescription("基本算数运算"),
			mcp.WithString(
				"operation",
				mcp.Required(),
				mcp.Enum("add", "subtract", "multiply", "divide"),
			),
			mcp.WithNumber(
				"x",
				mcp.Required(),
				mcp.Description("第一个数字"),
			),
			mcp.WithNumber(
				"y",
				mcp.Required(),
				mcp.Description("第二个数字"),
			),
		)
		s.AddTool(calculateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			op := request.Params.Arguments["operation"].(string)
			x := request.Params.Arguments["x"].(float64)
			y := request.Params.Arguments["y"].(float64)

			var result float64
			switch op {
			case "add":
				result = x + y
			case "subtract":
				result = x - y
			case "multiply":
				result = x * y
			case "divide":
				if y == 0 {
					return nil, errors.New("除数不允许为0")
				}
				result = x / y
			}
			return mcp.FormatNumberResult(result), nil
		})
	}
	// 添加资源
	{
		resource := mcp.NewResource(
			"docs://readme",
			"说明文档",
			mcp.WithResourceDescription("项目README文档"),
			mcp.WithMIMEType("text/readme"),
		)
		s.AddResource(resource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			content, err := os.ReadFile("./content/readme.md")
			if err != nil {
				return nil, err
			}
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI: "docs://readme",
					MIMEType: "text/markdown",
					Text: string(content),
				},
			}, nil
		})
	}
	// if server.ServeStdio(s) != nil {
	// 	fmt.Println("Server Start error")
	// }
	sseServer := server.NewSSEServer(s)
	if sseServer.Start(":8080") != nil {
		fmt.Println("Server Start error")
	}
}
