package spp_logger

import (
	"github.com/fatih/structs"
)

func ContextToDict(context *Context) map[string]string {
	dict := make(map[string]string)
	contextFields := structs.Names(&Context{})

	dict[contextFields[0]] = context.LogLevel()
	dict[contextFields[1]] = context.CorrelationID()
	return dict
}

func DictToContext(dict map[string]string) (*Context, error) {
	contextFields := structs.Names(&Context{})
	context, err := NewContext(dict[contextFields[0]], dict[contextFields[1]])
	return context, err
}
