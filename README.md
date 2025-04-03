# dify-retriever-mcp
dify 知识库检索工具

## 安装
### 若本地有go环境
```bash
go install github.com/wangle201210/dify-retriever-mcp@latest
```
### 若没有go环境
[下载二进制文件](https://github.com/wangle201210/dify-retriever-mcp/releases/latest)
把下载的可执行文件放到自己的 PATH 目录即可


## 用法
```json
{
  "mcpServers": {
      "dify-retriever-mcp": {
        "command": "dify-retriever-mcp",
        "args": [],
        "env": {
          "DIFY_DATASET_API_KEY": "dataset-iDBz1qVkmI2dU6KnirLCxO9K",
          "DIFY_ENDPOINT": "http://127.0.0.1/v1",
          "DIFY_DATASET_ID": "b667a65b-f40f-4dd0-8b34-8406dc8be96e",
          "DIFY_DATASET_NAME": "这里写知识库名字"
        },
        "descriptions": "***知识库检索"
      }
  } 
}
```
