package main

import "fmt"

// builder interface
type iBuilder interface {
	setMaxTotal(maxTotal int) iBuilder
	setMaxIdle(maxIdle int) iBuilder
	setMinIdle(minIdle int) iBuilder
	getResourcePoolConfig() ResourcePoolConfig
}

// getBuilder
func getBuilder(builderType string) iBuilder {
	if builderType == "default" {
		return &defaultBuilder{}
	}
	if builderType == "urgent" {
		return &urgentBuilder{}
	}
	return nil
}

// exact builder
type defaultBuilder struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

func newDefaultBuilder() *defaultBuilder {
	return &defaultBuilder{}
}

func (b *defaultBuilder) setMaxTotal(maxTotal int) iBuilder {
	b.maxTotal = maxTotal
	return b
}

func (b *defaultBuilder) setMaxIdle(maxIdle int) iBuilder {
	b.maxIdle = maxIdle
	return b
}

func (b *defaultBuilder) setMinIdle(minIdle int) iBuilder {
	b.minIdle = minIdle
	return b
}

func (b *defaultBuilder) getResourcePoolConfig() ResourcePoolConfig {
	return ResourcePoolConfig{
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}
}

// urgentBuilder exact builder
type urgentBuilder struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

func newUrgentBuilder() *urgentBuilder {
	return &urgentBuilder{}
}

func (b *urgentBuilder) setMaxTotal(maxTotal int) iBuilder {
	b.maxTotal = maxTotal * 2
	return b
}

func (b *urgentBuilder) setMaxIdle(maxIdle int) iBuilder {
	b.maxIdle = maxIdle * 2
	return b
}

func (b *urgentBuilder) setMinIdle(minIdle int) iBuilder {
	b.minIdle = minIdle * 2
	return b
}

func (b *urgentBuilder) getResourcePoolConfig() ResourcePoolConfig {
	return ResourcePoolConfig{
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}
}

// ResourcePoolConfig exact product
type ResourcePoolConfig struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

// director
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildConfig() ResourcePoolConfig {
	return d.builder.setMaxTotal(16).setMaxIdle(10).setMinIdle(1).getResourcePoolConfig()
}

// main
func main() {
	defaultBuilder := getBuilder("default")
	director := newDirector(defaultBuilder)
	defaultResourcePoolConfig := director.buildConfig()
	fmt.Printf("Default config maxTotal: %d\n", defaultResourcePoolConfig.maxTotal)
	fmt.Printf("Default config maxIdle: %d\n", defaultResourcePoolConfig.maxIdle)
	fmt.Printf("Default config minIdle: %d\n", defaultResourcePoolConfig.minIdle)

	urgentBuilder := getBuilder("urgent")
	director.setBuilder(urgentBuilder)
	urgentResourcePoolConfig := director.buildConfig()
	fmt.Printf("\nUrgent config maxTotal: %d\n", urgentResourcePoolConfig.maxTotal)
	fmt.Printf("Urgent config maxIdle: %d\n", urgentResourcePoolConfig.maxIdle)
	fmt.Printf("Urgent config minIdle: %d\n", urgentResourcePoolConfig.minIdle)

}
