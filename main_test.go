package main

import "testing"

func TestSayHello(t * testing.T) {
	wants := "Bonjour"
	value := sayHello();

	if (value != wants) {
		t.Errorf("sayHello()=%s, want %s",value,wants)
	}
}
