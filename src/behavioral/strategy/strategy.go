package main

import "fmt"

// iStrategy
type iStrategy interface {
	evict(*cache)
}

// fifo concrete strategy
type fifo struct{}

func (f *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

// lru concrete strategy
type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

// lfu concrete strategy
type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

// cache
type cache struct {
	storage     map[string]string
	strategy    iStrategy
	capacity    int
	maxCapacity int
}

func newCache(strategy iStrategy, capacity int, maxCapacity int) *cache {
	return &cache{
		storage:     make(map[string]string),
		strategy:    strategy,
		capacity:    capacity,
		maxCapacity: maxCapacity,
	}
}

func (c *cache) setStrategy(strategy iStrategy) {
	c.strategy = strategy
}

func (c *cache) get(key string) string {
	// if val, ok := c.storage[key]; ok {
	// 	return val, nil
	// }
	// return nil, erros.New("")
	return c.storage[key]
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) evict() {
	c.strategy.evict(c)
	c.capacity--
}

// default_strategy
var defaultStrategy iStrategy = &lru{}

// strategyFactory
type strategyFactory struct {
	strategies map[string]iStrategy
}

func newStrategyFactory() *strategyFactory {
	return &strategyFactory{
		make(map[string]iStrategy),
	}
}

func (s *strategyFactory) getStrategy(t string) iStrategy {
	if val, ok := s.strategies[t]; ok {
		return val
	}
	return defaultStrategy
}

var strFact *strategyFactory

// init
func init() {
	strFact = newStrategyFactory()
	strFact.strategies["lru"] = defaultStrategy
	strFact.strategies["fifo"] = &fifo{}
	strFact.strategies["lfu"] = &lfu{}
}

// main
func main() {
	strategy := strFact.getStrategy("unknown")
	cache := newCache(strategy, 0, 2)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	strategy = strFact.getStrategy("lfu")
	cache.setStrategy(strategy)

	cache.add("d", "4")

	strategy = strFact.getStrategy("fifo")
	cache.setStrategy(strategy)

	cache.add("e", "5")
}
