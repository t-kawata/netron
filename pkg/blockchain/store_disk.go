package blockchain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/peterbourgon/diskv"
)

type DiskStore struct {
	chain      *diskv.Diskv
	maxBlocks  int
	blockCount int
	sync.Mutex
}

// func NewDiskStore(d *diskv.Diskv) *DiskStore {
// 	store := &DiskStore{chain: d, maxBlocks: 0}
// 	store.initializeBlockCount()
// 	return store
// }

func NewDiskStoreWithLimit(d *diskv.Diskv, maxBlocks int) *DiskStore {
	store := &DiskStore{
		chain:     d,
		maxBlocks: maxBlocks,
	}
	store.initializeBlockCount()
	return store
}

func (m *DiskStore) initializeBlockCount() {
	count, err := m.chain.Read("index")
	if err != nil {
		m.blockCount = 0
		return
	}
	latestIndex, _ := strconv.Atoi(string(count))
	if countData, err := m.chain.Read("block_count"); err == nil {
		m.blockCount, _ = strconv.Atoi(string(countData))
	} else {
		m.blockCount = m.countExistingBlocks(latestIndex)
		m.chain.Write("block_count", fmt.Append(nil, m.blockCount))
	}
}

func (m *DiskStore) countExistingBlocks(latestIndex int) int {
	count := 0
	for i := 0; i <= latestIndex; i++ {
		if _, err := m.chain.Read(fmt.Sprint(i)); err == nil {
			count++
		}
	}
	return count
}

func (m *DiskStore) Add(b Block) {
	m.Lock()
	defer m.Unlock()

	bb, _ := json.Marshal(b)
	m.chain.Write(fmt.Sprint(b.Index), bb)
	m.chain.Write("index", fmt.Append(nil, b.Index))

	m.blockCount++
	m.chain.Write("block_count", fmt.Append(nil, m.blockCount))

	if m.maxBlocks > 0 && m.blockCount > m.maxBlocks {
		blocksToDelete := m.blockCount - m.maxBlocks
		oldestIndex := b.Index - m.blockCount + 1
		for i := range blocksToDelete {
			m.chain.Erase(fmt.Sprint(oldestIndex + i))
		}

		m.blockCount = m.maxBlocks
		m.chain.Write("block_count", fmt.Append(nil, m.blockCount))
	}
}

func (m *DiskStore) Len() int {
	m.Lock()
	defer m.Unlock()

	if m.maxBlocks > 0 {
		return m.blockCount
	} else {
		count, err := m.chain.Read("index")
		if err != nil {
			return 0
		}
		c, _ := strconv.Atoi(string(count))
		return c
	}
}

func (m *DiskStore) Last() Block {
	m.Lock()
	defer m.Unlock()

	// 最新ブロックのインデックス取得
	indexBytes, err := m.chain.Read("index")
	if err != nil {
		return Block{} // ブロックが存在しない場合は空を返す
	}

	index, err := strconv.Atoi(string(indexBytes))
	if err != nil {
		return Block{} // 不正なインデックス値の場合は空を返す
	}

	// 最新ブロックデータの読み込み
	blockBytes, err := m.chain.Read(fmt.Sprint(index))
	if err != nil {
		return Block{}
	}

	var block Block
	if err := json.Unmarshal(blockBytes, &block); err != nil {
		return Block{}
	}

	return block
}

// type DiskStore struct {
// 	chain *diskv.Diskv
// }

// func NewDiskStore(d *diskv.Diskv) *DiskStore {
// 	return &DiskStore{chain: d}
// }

// func (m *DiskStore) Add(b Block) {
// 	bb, _ := json.Marshal(b)
// 	m.chain.Write(fmt.Sprint(b.Index), bb)
// 	m.chain.Write("index", []byte(fmt.Sprint(b.Index)))

// }

// func (m *DiskStore) Len() int {
// 	count, err := m.chain.Read("index")
// 	if err != nil {
// 		return 0
// 	}
// 	c, _ := strconv.Atoi(string(count))
// 	return c

// }

// func (m *DiskStore) Last() Block {
// 	b := &Block{}

// 	count, err := m.chain.Read("index")
// 	if err != nil {
// 		return *b
// 	}

// 	dat, _ := m.chain.Read(string(count))
// 	json.Unmarshal(dat, b)

// 	return *b
// }
