package pathtree_test

import (
	"fmt"
	"mockingbird/pathtree"
	"testing"
)

type Dynamic struct {
	WildCard []string
	Values   []string
}

type Handler func(dynamic *Dynamic)

func TestPathTreeRouter(t *testing.T) {
	node := pathtree.New()

	node.Add("/api/v1/ping", func(dynamic *Dynamic) {
		fmt.Println("Inside Ping")
	})
	node.Add("/api/v1/pong", func(dynamic *Dynamic) {
		fmt.Println("Inside Pong")
	})

	node.Add("/api/v1/:dynamic", func(dynamic *Dynamic) {
		fmt.Println("Inside Ping Pong")
	})

	{
		leaf, value := node.Find("/api/v1/ping")
		dynamic := &Dynamic{
			WildCard: leaf.Wildcards,
			Values:   value,
		}

		handler, ok := leaf.Value.(func(*Dynamic))
		if !ok {
			fmt.Println("Handler not found")
		} else {
			if nil != handler {
				handler(dynamic)
			}
		}
	}

	{
		leaf, value := node.Find("/api/v1/pong")
		dynamic := &Dynamic{
			WildCard: leaf.Wildcards,
			Values:   value,
		}
		handler, ok := leaf.Value.(func(*Dynamic))
		if !ok {
			fmt.Println("Handler not found")
		} else {
			handler(dynamic)
		}
	}

	{
		leaf, value := node.Find("/api/v1/pingpong")
		dynamic := &Dynamic{
			WildCard: leaf.Wildcards,
			Values:   value,
		}
		handler, ok := leaf.Value.(func(*Dynamic))
		if !ok {
			fmt.Println("Handler not found")
		} else {
			handler(dynamic)
		}
	}
}

func TestPathTree(t *testing.T) {

	node := pathtree.New()
	node.Add("/hello/world", "static_path")
	node.Add("/hello/world", "duplicate_path")
	node.Add("/:x/:y", "dynamic_path")

	{
		leaf, values := node.Find("/hello/world")
		fmt.Println(leaf, values)
	}

	{
		leaf, values := node.Find("/goodbye/world")
		fmt.Println(leaf.Value)
		fmt.Println(leaf, values)
	}

	{
		leaf, values := node.Find("/hello/goodbye/world")
		//fmt.Println(leaf.Value)
		fmt.Println(leaf, values)
	}
}
