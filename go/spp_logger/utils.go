package spp_logger

// import (
// 	"errors"

// 	"github.com/fatih/structs"
// )

// func ContextToDict(context *Context) (map[string]string, error) {
// 	dict := make(map[string]string)
// 	contextFields := structs.Names(&Context{})
// 	if err := context.IsValid(); err != nil {
// 		return nil, err
// 	}
// 	dict[contextFields[0]] = context.LogLevel()
// 	dict[contextFields[1]] = context.CorrelationID()
// 	return dict, nil
// }

// func DictToContext(dict map[string]string) (*Context, error) {
// 	contextFields := structs.Names(&Context{})
// 	level := dict[contextFields[0]]
// 	correlationID := dict[contextFields[1]]

// 	if level == "" || correlationID == "" {
// 		return nil, errors.New("Invalid Dictionary")
// 	}
// 	context, err := NewContext(level, correlationID)

// 	return context, err
// }
