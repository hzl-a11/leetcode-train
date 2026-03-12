package algorithm

import (
	"sync"
	"time"
)

// IPLimiter 基于滑动窗口的 IP 限流器。
type IPLimiter struct {
	limit      int      // 窗口内允许的最大请求数
	windowNano int64    // 窗口大小（纳秒）
	ips        sync.Map // 并发安全 map (key: string, value: *ipWindow)
}

// ipWindow 保存单个 IP 的请求时间戳队列。
type ipWindow struct {
	mu          sync.Mutex
	ts          []int64 // 存储历史请求时间戳，是个环形数组
	lastTimeIdx int     // 环形数组当前最后一个有效元素的位置，也就是上一次请求的时间戳
}

// NewIPLimiter 创建限流器。
// limit<=0 时默认 1，window<=0 时默认 1 秒。
func NewIPLimiter(limit int, window time.Duration) *IPLimiter {
	if limit <= 0 {
		limit = 1
	}
	if window <= 0 {
		window = time.Second
	}
	return &IPLimiter{
		limit:      limit,
		windowNano: int64(window),
	}
}

// Allow 判断某个 IP 当前请求是否放行。
// 逻辑：
// 1. 清理窗口外（过期）时间戳；
// 2. 若窗口内请求数已达上限则拒绝；
// 3. 否则记录本次请求并放行。

func (l *IPLimiter) Allow(ip string) bool {

	// 1. 原子化加载或创建窗口数据
	// LoadOrStore 会检查 key 是否存在，不存在则存入新创建的 ipWindow
	v, _ := l.ips.LoadOrStore(ip, &ipWindow{
		ts: make([]int64, l.limit),
	})
	window := v.(*ipWindow)

	window.mu.Lock()
	defer window.mu.Unlock()

	//计算当前的窗口范围
	now := time.Now().UnixNano()
	cutoff := now - int64(l.windowNano)

	nextIdx := (window.lastTimeIdx + 1) % len(window.ts) //指向窗口中最久远的一条时间戳，同时也是新ip要写入的位置
	if window.ts[nextIdx] > cutoff {
		// 若窗口中最久远的一条时间戳，如果在cutoff里，说明窗口已经满了，不能加新请求，拒绝
		return false
	} else {
		window.ts[nextIdx] = now
		window.lastTimeIdx = nextIdx
		return true
	}
}
