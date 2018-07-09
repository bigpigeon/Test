/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"math"
)

const abortIndex int = math.MaxInt8 / 2

type Request struct {
	Auth bool
	//...
}

type apiService struct {
}

type HandlerFunc func(c *Context) error
type HandlersChain []HandlerFunc
type Context struct {
	handlers HandlersChain
	index    int8
	// put Request to context
	Request Request
	value   map[interface{}]interface{}
	//err      error
}

func NewContext(request Request, handlers ...HandlerFunc) *Context {
	return &Context{
		handlers: handlers,
		index:    -1,
		Request:  request,
	}
}

func (c *Context) Next() error {
	c.index++
	var err error
	for s := int8(len(c.handlers)); c.index < s; c.index++ {
		err = c.handlers[c.index](c)
		// if have error in current handler, stop
		if err != nil {
			c.Abort()
		}
	}
	return err
}

func (c *Context) IsAborted() bool {
	return len(c.handlers) >= abortIndex
}

func (c *Context) Abort() {
	c.index = int8(abortIndex)
}

func (s *apiService) authorizationCheck(ctx *Context) error {
	// do sth
	fmt.Println("authorization")
	if ctx.Request.Auth == false {
		return fmt.Errorf("authorization failure")
	}
	return nil
}

func (s *apiService) get(ctx *Context) error {

	//method logic
	fmt.Println("get")
	return nil
}

func (s *apiService) put(ctx *Context) error {

	//method logic
	fmt.Println("put")
	return nil
}

func main() {

	{
		request := Request{Auth: true}
		api := apiService{}
		getFunc := NewContext(request, api.authorizationCheck, api.get)
		getFunc.Next()
		putFunc := NewContext(request, api.authorizationCheck, api.put)
		putFunc.Next()
	}

	// no auth request
	{
		request := Request{Auth: false}
		api := apiService{}
		getFunc := NewContext(request, api.authorizationCheck, api.get)
		err := getFunc.Next()
		if err != nil {
			fmt.Printf("err %s\n", err)
		}
		putFunc := NewContext(request, api.authorizationCheck, api.put)
		err = putFunc.Next()
		if err != nil {
			fmt.Printf("err %s\n", err)
		}
	}
}
