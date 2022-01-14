package pool

import (
	"github.com/matrixcloud/proxy-pool/db"
)

// Pool used for manipulate proxies pool periodicly
type Pool struct {
	db               *db.Client
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
		db:               conn,
		maxThreshold:     opts.MaxThreshold,
		minThreshold:     opts.MinThreshold,
		checkInterval:    opts.CheckInterval,
		validateInterval: opts.ValidateInterval,
	}
}

// Start proxy pool
func (p *Pool) Start() {
	p.check()
	p.validate()
}

func (p *Pool) size() int {
	return p.db.Length()
}

func (p *Pool) isOverThreshold() bool {
	return p.size() > p.maxThreshold
}
