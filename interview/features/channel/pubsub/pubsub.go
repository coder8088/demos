package pubsub

import (
	"sync"
	"time"
)

type Subscriber chan interface{}        // 订阅者为一个通道
type TopicFunc func(v interface{}) bool // 主题为一个过滤器

type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列缓存容量
	timeout     time.Duration            // 发布超时时间
	subscribers map[Subscriber]TopicFunc // 订阅者信息
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[Subscriber]TopicFunc),
	}
}

// 添加一个订阅者，订阅全部主题
func (p *Publisher) Subscribe() Subscriber {
	return p.SubscribeTopic(nil)
}

// 添加一个订阅制，指定订阅的主题
func (p *Publisher) SubscribeTopic(topic TopicFunc) Subscriber {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendTopic(sub Subscriber, topic TopicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func (p Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
