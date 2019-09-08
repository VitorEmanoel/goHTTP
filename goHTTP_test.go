package goHTTP

import "testing"

func TestRouter1(t *testing.T){
	if !isRouter("router/test", "router/test"){
		t.Fail()
	}
	t.Log("Router test 1 success 'router/test'")
}

func TestRouter2(t *testing.T){
	if !isRouter("example/123", "example/:name"){
		t.Fail()
	}
	t.Log("Router test 2 success 'example/:name'")
}
