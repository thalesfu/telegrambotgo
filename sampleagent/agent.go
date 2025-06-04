package sampleagent

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/thalesfu/golangagent/mem"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetAgentRunner(ctx context.Context) (compose.Runnable[*tgbotapi.Message, string], []compose.Option, error) {
	agent, opts, err := GetAgent(ctx)
	if err != nil {
		return nil, nil, err
	}

	runner, err := agent.Compile(ctx, compose.WithMaxRunSteps(100))

	return runner, opts, err
}

func GetAgent(ctx context.Context) (*compose.Graph[*tgbotapi.Message, string], []compose.Option, error) {
	graph := compose.NewGraph[*tgbotapi.Message, string]()

	opts := []compose.Option{
		mem.GetInitMemGraphCallBackOption(),
	}

	userInputNodeKey, userInputNode, userInputNodeOpts, userInputInvokeOpts := GetUserInputNode()
	err := graph.AddLambdaNode(userInputNodeKey, userInputNode, userInputNodeOpts...)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, userInputInvokeOpts...)

	chatTemplateNodeKey, questionAnalyseChatTemplateNode, questionAnalyseChatTemplateNodeOpts := GetChatTemplateNode()
	err = graph.AddChatTemplateNode(chatTemplateNodeKey, questionAnalyseChatTemplateNode, questionAnalyseChatTemplateNodeOpts...)
	if err != nil {
		return nil, nil, err
	}

	questionAnalyseChatModelNodeKey, questionAnalyseChatModelNode, questionAnalyseChatModeNodeOpts, questionAnalyseChatModeNodeInvokeOpts := GetChatModelNode(ctx)
	err = graph.AddChatModelNode(questionAnalyseChatModelNodeKey, questionAnalyseChatModelNode, questionAnalyseChatModeNodeOpts...)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, questionAnalyseChatModeNodeInvokeOpts...)

	chatResultNodeKey, chatResultNode, chatResultNodeKeyOpts := ChatResultLambdaNode()
	err = graph.AddLambdaNode(chatResultNodeKey, chatResultNode, chatResultNodeKeyOpts...)
	if err != nil {
		return nil, nil, err
	}

	err = graph.AddEdge(compose.START, userInputNodeKey)
	if err != nil {
		return nil, nil, err
	}

	err = graph.AddEdge(userInputNodeKey, chatTemplateNodeKey)
	if err != nil {
		return nil, nil, err
	}

	err = graph.AddEdge(chatTemplateNodeKey, questionAnalyseChatModelNodeKey)
	if err != nil {
		return nil, nil, err
	}

	err = graph.AddEdge(questionAnalyseChatModelNodeKey, chatResultNodeKey)
	if err != nil {
		return nil, nil, err
	}

	err = graph.AddEdge(chatResultNodeKey, compose.END)
	if err != nil {
		return nil, nil, err
	}

	return graph, opts, nil
}
