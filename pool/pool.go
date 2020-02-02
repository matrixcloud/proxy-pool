package pool

import (
	"github.com/matrixcloud/proxy-pool/db"
)

// Pool used for manipulate proxies pool periodicly
type Pool struct {
	conn             *db.Client
	maxThreshold     int
	minThreshold     int
	checkInterval    int
	validateInterval int
}

// Options ...
type Options struct {
	MaxThreshold     int
	MinThreshold     int
	CheckInterval    int
	ValidateInterval int
}

// NewPool returns a proxy pool instance
func NewPool(conn *db.Client, opts Options) *Pool {
	return &Pool{
		conn:             conn,
		maxThreshold:     opts.MaxThreshold,
		minThreshold:     opts.MinThreshold,
		checkInterval:    opts.CheckInterval,
		validateInterval: opts.ValidateInterval,
	}
}

// Start proxy pool
func (p *Pool) Start() {
	go p.check()
	go p.validate()
}

func (p *Pool) size() int64 {
	return p.conn.Length()
}

func (p *Pool) isOverThreshold() bool {
	return p.size() > int64(p.maxThreshold)
}
