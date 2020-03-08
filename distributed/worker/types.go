package worker

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/uestcfaculty/parser"
	"webcrawler/distributed/config"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items:    r.Items,
		Requests: nil,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Requests: nil,
		Items:    r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseSchoolList:
		return engine.NewFuncParser(parser.ParseSchoolList, config.ParseSchoolList), nil
	case config.ParseSchool:
		return engine.NewFuncParser(parser.ParseSchool, config.ParseSchool), nil
	case config.ParseTeacherProfile:
		if args, ok := p.Args.([]interface{}); ok {
			// you can't convert a slice from json data without context
			// just assert that this variable is a slice of interface
			// then assert type of each element
			teacherName, ok1 := (args[0]).(string)
			position, ok2 := (args[1]).(string)
			if ok1 && ok2 {
				return parser.NewProfileParser(teacherName, position), nil
			} else {
				return nil, fmt.Errorf("invalid arguments: %v and %v", args[0], args[1])
			}
		} else {
			return nil, fmt.Errorf("invalid argument: %v; and its type: %s", p.Args, typeof(p.Args))
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
