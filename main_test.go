package main

import "testing"

func TestSayHello(t * testing.T) {
	wants := "BonjouR"
	value := sayHello();

	if (value != "BonjouR") {
		t.Errorf("sayHello()=%s, want %s",value,wants)
	}
}