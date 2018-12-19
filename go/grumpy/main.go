package main

import (
	hello2 "__python__/hello"
	"grumpy"
)

func main() {
	grumpy.RunMain(hello2.Code)
	f := grumpy.NewRootFrame()
	mod, err := grumpy.ImportModule(f, "hello")
	if err != nil {
		panic(err)
	}
	hello, e := grumpy.GetAttr(f, mod[0], grumpy.InternStr("hello"), nil)
	if e != nil {
		panic(e)
	}
	hello.Call(f, nil, nil)
	//hello.Code(rootFrame, nil, nil)

}
