package container

import "errors"

type ConcreteResolver func(c *ContainerService) (interface{}, error)

var ErrNoBinding = errors.New("No binding found for name")

func Bind(name string, resolver ConcreteResolver) {
	container.Bind(name, resolver)
}

func Singleton(name string, resolver ConcreteResolver) {
	container.Singleton(name, resolver)
}

func Resolve(name string) (interface{}, error) {
	return container.Resolve(name)
}

func New() *ContainerService {
	return &ContainerService{
		bindings:           make(map[string]ConcreteResolver),
		singletons:         make(map[string]ConcreteResolver),
		singletonInstances: make(map[string]interface{}),
	}
}
