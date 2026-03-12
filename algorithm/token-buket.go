package algorithm

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   float64 // 桶的最大容量
	rate       float64 // 令牌生成的速率（个/秒）
	tokens     float64 // 当前桶里的令牌数
	lastRefill int64   // 上次补水的时间戳（纳秒）
	mu         sync.Mutex
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	//计算从此刻到上次补水时间，产生了多少token
	now := time.Now().UnixNano()
	passedNano := now - b.lastRefill
	deltaTokens := float64(passedNano) * b.rate / 1e9

	//补水
	newTokens := b.tokens + deltaTokens
	if deltaTokens > b.capacity {
		newTokens = b.capacity
	}
	b.capacity = newTokens
	b.lastRefill = now

	//准出
	if b.tokens >= 1 {
		b.tokens -= 1
		return true
	}
	return false
}
