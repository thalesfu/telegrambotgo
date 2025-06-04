package sampleagent

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func ChatResultLambda(ctx context.Context, result *schema.Message) (string, error) {
	return result.Content, nil
}

func ChatResultLambdaNode() (string, *compose.Lambda, []compose.GraphAddNodeOpt) {
	return AgentNodeChatModelResult, compose.InvokableLambda(ChatResultLambda), getChatResultLambdaNodeOptions()
}

func getChatResultLambdaNodeOptions() []compose.GraphAddNodeOpt {
	return []compose.GraphAddNodeOpt{
		compose.WithNodeName(AgentNodeChatModelResult),
	}
}
