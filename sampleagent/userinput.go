package sampleagent

import (
	"context"
	"strconv"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/compose"
	"github.com/thalesfu/golangagent/mem"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UserInputLambda(ctx context.Context, input *tgbotapi.Message) (map[string]any, error) {
	message := make(map[string]any)

	message["user_input"] = input.Text

	session := mem.GetSessionFromContext(ctx)
	message["message_histories"] = session.GetMessages()

	return message, nil
}

func GetUserInputNode() (string, *compose.Lambda, []compose.GraphAddNodeOpt, []compose.Option) {
	return AgentNodeUserInput, compose.InvokableLambda(UserInputLambda), getUserInputNodeOptions(), getUserInputInvokeOption()
}

func getUserInputNodeOptions() []compose.GraphAddNodeOpt {
	return []compose.GraphAddNodeOpt{
		compose.WithNodeName(AgentNodeUserInput),
	}
}

func getUserInputInvokeOption() []compose.Option {
	return []compose.Option{
		getUserInputCallBackOption(),
	}
}

func getUserInputCallBackOption() compose.Option {
	handler := callbacks.NewHandlerBuilder().
		OnStartFn(initSession).
		Build()
	return compose.WithCallbacks(handler).DesignateNode(AgentNodeUserInput)
}

func initSession(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
	session := mem.GetSessionFromContext(ctx)
	if message, ok := input.(*tgbotapi.Message); ok {
		if message.From != nil {
			session.Init(strconv.FormatInt(message.From.ID, 10))
		}
	}
	return ctx
}
