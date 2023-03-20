package container_test

import (
	"errors"
	"testing"

	"github.com/farzai/container-go"
)

func TestContainerServiceSuite(t *testing.T) {
	t.Run("Bind and resolve", func(t *testing.T) {
		service := container.New()

		service.Bind("foo", func(c *container.ContainerService) (interface{}, error) {
			return "bar", nil
		})
		foo, err := service.Resolve("foo")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if foo != "bar" {
			t.Fatalf("Expected foo to be \"bar\", got %v", foo)
		}
	})

	t.Run("Should be singleton", func(t *testing.T) {
		service := container.New()

		service.Singleton("quux", func(c *container.ContainerService) (interface{}, error) {
			return "corge", nil
		})

		if !service.IsSingleton("quux") {
			t.Fatalf("Expected quux to be a singleton")
		}
	})

	t.Run("Should be singleton when is binding before", func(t *testing.T) {
		service := container.New()

		service.Bind("quux", func(c *container.ContainerService) (interface{}, error) {
			return "corge", nil
		})

		if service.IsSingleton("quux") {
			t.Fatalf("Expected quux to be a binding")
		}

		if !service.IsBinding("quux") {
			t.Fatalf("Expected quux to be a binding")
		}

		service.Singleton("quux", func(c *container.ContainerService) (interface{}, error) {
			return "corge", nil
		})

		if !service.IsSingleton("quux") {
			t.Fatalf("Expected quux to be a singleton")
		}

		if service.IsBinding("quux") {
			t.Fatalf("Expected quux to be a binding")
		}
	})

	t.Run("Singleton and resolve", func(t *testing.T) {
		service := container.New()

		service.Singleton("baz", func(c *container.ContainerService) (interface{}, error) {
			return "qux", nil
		})
		baz1, err := service.Resolve("baz")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		baz2, err := service.Resolve("baz")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if baz1 != baz2 {
			t.Fatalf("Expected baz1 and baz2 to be the same instance, got %v and %v", baz1, baz2)
		}
	})

	t.Run("Resolve non-existent binding", func(t *testing.T) {
		service := container.New()

		_, err := service.Resolve("nonexistent")
		if !errors.Is(err, container.ErrNoBinding) {
			t.Fatalf("Expected ErrNoBinding, got %v", err)
		}
	})

	t.Run("Binding after singleton", func(t *testing.T) {
		service := container.New()

		service.Singleton("quux", func(c *container.ContainerService) (interface{}, error) {
			return "corge", nil
		})
		service.Bind("quux", func(c *container.ContainerService) (interface{}, error) {
			return "grault", nil
		})
		quux, err := service.Resolve("quux")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if quux != "grault" {
			t.Fatalf("Expected quux to be \"grault\", got %v", quux)
		}
	})

	t.Run("Resolving binding after singleton", func(t *testing.T) {
		service := container.New()

		bar := "baz"
		service.Singleton("foo", func(c *container.ContainerService) (interface{}, error) {
			return bar, nil
		})
		bar = "qux"
		foo, err := service.Resolve("foo")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if foo != "qux" {
			t.Fatalf("Expected foo to be \"qux\", got %v", foo)
		}
	})

	t.Run("Unbind", func(t *testing.T) {
		service := container.New()

		service.Bind("foo", func(c *container.ContainerService) (interface{}, error) {
			return "bar", nil
		})

		if !service.IsBinding("foo") {
			t.Fatalf("Expected foo to be a binding")
		}

		service.Unbind("foo")

		if service.IsBinding("foo") {
			t.Fatalf("Expected foo to be a binding")
		}

		_, err := service.Resolve("foo")
		if !errors.Is(err, container.ErrNoBinding) {
			t.Fatalf("Expected ErrNoBinding, got %v", err)
		}
	})
}
