package main

type Subject interface {
	register()
	deregister()
	notifyAll()
}
