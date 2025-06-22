package blockchain

import "sync"

type MemoryStore struct {
	sync.Mutex
	blocks    []Block
	maxBlocks int
}

func NewMemoryStore(maxBlocks int) *MemoryStore {
	return &MemoryStore{
		blocks:    make([]Block, 0),
		maxBlocks: maxBlocks,
	}
}

func (m *MemoryStore) Add(b Block) {
	m.Lock()
	defer m.Unlock()

	m.blocks = append(m.blocks, b)
	if len(m.blocks) > m.maxBlocks {
		m.blocks = m.blocks[len(m.blocks)-m.maxBlocks:]
	}
}

func (m *MemoryStore) Len() int {
	m.Lock()
	defer m.Unlock()

	if len(m.blocks) == 0 {
		return 0
	}
	return m.blocks[len(m.blocks)-1].Index
}

func (m *MemoryStore) Last() Block {
	m.Lock()
	defer m.Unlock()

	if len(m.blocks) == 0 {
		return Block{}
	}
	return m.blocks[len(m.blocks)-1]
}

// type MemoryStore struct {
// 	sync.Mutex
// 	block *Block
// }

// func (m *MemoryStore) Add(b Block) {
// 	m.Lock()
// 	m.block = &b
// 	m.Unlock()
// }

// func (m *MemoryStore) Len() int {
// 	m.Lock()
// 	defer m.Unlock()
// 	if m.block == nil {
// 		return 0
// 	}
// 	return m.block.Index
// }

// func (m *MemoryStore) Last() Block {
// 	m.Lock()
// 	defer m.Unlock()
// 	return *m.block
// }
