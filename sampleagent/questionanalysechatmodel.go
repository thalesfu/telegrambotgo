package sampleagent

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/thalesfu/golangagent/mem"
)

func GetChatModelNode(ctx context.Context) (string, model.ChatModel, []compose.GraphAddNodeOpt, []compose.Option) {
	chatModel := CreateDeepSeekChatModel(ctx)

	return AgentNodeChatModel,
		chatModel,
		getChatModeNodeOptions(),
		getChatModeNodeInvokeOptions()
}

func getChatModeNodeOptions() []compose.GraphAddNodeOpt {
	return []compose.GraphAddNodeOpt{
		compose.WithNodeName(AgentNodeChatModel),
	}
}

func getChatModeNodeInvokeOptions() []compose.Option {
	return []compose.Option{
		mem.GetModelMemCallBackOptionWithNodeKey(AgentNodeChatModel),
	}
}

func CreateDeepSeekChatModel(ctx context.Context) model.ChatModel {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:   "deepseek-reasoner",                   // 使用的模型版本
		APIKey:  "sk-dd70da7bbc6a410fb97a4f6847b67208", // OpenAI API 密钥
		BaseURL: "https://api.deepseek.com/v1",
	})
	if err != nil {
		log.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}

func CreateChatModel(ctx context.Context) model.ChatModel {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:   "gpt-4o-mini",                                         // 使用的模型版本
		APIKey:  "sk-pu6Iqrc0zEh0lp5WzltT0vaVlVD6sZahzBqkvHdRDE1enFMC", // OpenAI API 密钥
		BaseURL: "http://ai-proxy.fws.qa.nt.ctripcorp.com/v1",
	})
	if err != nil {
		log.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}
