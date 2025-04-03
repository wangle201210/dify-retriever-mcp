package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/retriever/dify"
	"github.com/mark3labs/mcp-go/mcp"
)

func getDifyTool() mcp.Tool {
	datasetName := os.Getenv("DIFY_DATASET_NAME")
	return mcp.NewTool("dify_retriever",
		mcp.WithDescription(fmt.Sprintf("检索 %s 知识库", datasetName)), // 在环境变量中的知识库名字，增加tool命中率
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("检索内容"),
		),
	)
}

func difyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, ok := request.Params.Arguments["query"].(string)
	if !ok {
		return nil, errors.New("query must be a string")
	}
	ret, err := dify.NewRetriever(ctx, &dify.RetrieverConfig{
		APIKey:    os.Getenv("DIFY_DATASET_API_KEY"),
		Endpoint:  os.Getenv("DIFY_ENDPOINT"),
		DatasetID: os.Getenv("DIFY_DATASET_ID"),
	})
	if err != nil {
		return nil, err
	}
	// 开始检索
	docs, err := ret.Retrieve(ctx, query)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(docs)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(marshal)), nil
}
