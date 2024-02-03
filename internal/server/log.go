package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex // used for concurrent access to the log
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint16, error) {
	c.mu.Lock() // lock so that 2 goroutines don't append at the same time
	defer c.mu.Unlock() // will unlock after the function returns
	record.Offset = uint16(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint16) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint16(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint16 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
