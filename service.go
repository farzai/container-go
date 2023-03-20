package container

import (
	"fmt"
)

type ContainerService struct {
	bindings   map[string]ConcreteResolver
	singletons map[string]ConcreteResolver

	singletonInstances map[string]interface{}
}

func (c *ContainerService) Bind(name string, resolver ConcreteResolver) {
	c.bindings[name] = resolver

	// Check if the binding is a singleton
	if _, ok := c.singletons[name]; ok {
		delete(c.singletons, name)
	}
}

func (c *ContainerService) Singleton(name string, resolver ConcreteResolver) {
	c.singletons[name] = resolver

	// Check if the singleton is a binding
	if _, ok := c.bindings[name]; ok {
		delete(c.bindings, name)
	}
}

func (c *ContainerService) Has(name string) bool {
	return c.IsBinding(name) || c.IsSingleton(name)
}

func (c *ContainerService) IsBinding(name string) bool {
	_, ok := c.bindings[name]
	return ok
}

func (c *ContainerService) IsSingleton(name string) bool {
	_, ok := c.singletons[name]
	return ok
}

func (c *ContainerService) Unbind(name string) {
	if _, ok := c.bindings[name]; ok {
		delete(c.bindings, name)
	}

	if _, ok := c.singletons[name]; ok {
		delete(c.singletons, name)
	}
}

func (c *ContainerService) Resolve(name string) (interface{}, error) {
	// Check if the name is a singleton
	if resolver, ok := c.singletons[name]; ok {
		// Check if the singleton has already been resolved
		if instance, ok := c.singletonInstances[name]; ok {
			return instance, nil
		}

		// Resolve the singleton
		instance, err := resolver(c)
		if err != nil {
			return nil, err
		}

		// Save the singleton instance
		c.singletonInstances[name] = instance

		return instance, nil
	}

	// Check if the name is a binding
	if resolver, ok := c.bindings[name]; ok {
		return resolver(c)
	}

	return nil, fmt.Errorf("%w: %s", ErrNoBinding, name)
}

var container *ContainerService

func init() {
	container = New()
}
