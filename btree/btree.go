package btree

import "encoding/binary"

const (
	BTREE_PAGE_SIZE    = 4096
	BTREE_MAX_KEY_SIZE = 1000
	BTREE_MAX_VAL_SIZE = 3000
)
const (
	BNODE_INTERNAL = 1
	BNODE_LEAF     = 2
)

type BNode struct {
	data []byte
}

// type and nkeys
func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data[0:2])
}
func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[2:4])
}
func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

//pointers
func (node BNode) getPtr(idx uint16) uint64 {
	if idx >= node.nkeys() {
		panic("index out of range")
	}
	pos := 4 + 8*idx
	return binary.LittleEndian.Uint64(node.data[pos:])
}
func (node BNode) setPtr(idx uint16, val uint64) {
	if idx >= node.nkeys() {
		panic("index out of range")
	}
	pos := 4 + 8*idx
	binary.LittleEndian.PutUint64(node.data[pos:], val)
}

//offsets
func (node BNode) getOffset(idx uint16) uint16 {
	pos := 4 + 8*node.nkeys() + 2*(idx)
	return binary.LittleEndian.Uint16(node.data[pos:])
}
func (node BNode) setOffset(idx uint16, offset uint16) {
	pos := 4 + 8*node.nkeys() + 2*(idx)
	binary.LittleEndian.PutUint16(node.data[pos:], offset)
}

//kv
func (node BNode) kvPos(idx uint16) uint16 {
	return uint16(4) + node.nkeys()*8 + node.nkeys()*2 + node.getOffset(idx)
}
func (node BNode) getKey(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node.data[pos:])
	return node.data[pos+4:][:klen]
}
func (node BNode) getVal(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node.data[pos+0:])
	vlen := binary.LittleEndian.Uint16(node.data[pos+2:])
	return node.data[pos+4+klen:][:vlen]
}

//size of the node
func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nkeys())
}
