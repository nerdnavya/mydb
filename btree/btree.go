package btree

<<<<<<< HEAD
import "encoding/binary"
=======
import (
	"bytes"
	"encoding/binary"
)
>>>>>>> 9753a43 (insert function is added)

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

<<<<<<< HEAD
//pointers
=======
// pointers
>>>>>>> 9753a43 (insert function is added)
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

<<<<<<< HEAD
//offsets
=======
// offsets
>>>>>>> 9753a43 (insert function is added)
func (node BNode) getOffset(idx uint16) uint16 {
	pos := 4 + 8*node.nkeys() + 2*(idx)
	return binary.LittleEndian.Uint16(node.data[pos:])
}
func (node BNode) setOffset(idx uint16, offset uint16) {
	pos := 4 + 8*node.nkeys() + 2*(idx)
	binary.LittleEndian.PutUint16(node.data[pos:], offset)
}

<<<<<<< HEAD
//kv
=======
// kv
>>>>>>> 9753a43 (insert function is added)
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
<<<<<<< HEAD

//size of the node
func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nkeys())
}
=======
func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nkeys())
}
func nodeLookupLE(node BNode, key []byte) uint16 {
	nkeys := node.nkeys()
	found := uint16(0)
	for i := uint16(1); i < nkeys; i++ {
		cmp := bytes.Compare(node.getKey(i), key)
		if cmp <= 0 {
			found = i
		}
		if cmp >= 0 {
			break
		}
	}

	return found
}
func nodeAppendKV(new BNode, idx uint16, ptr uint64, key []byte, val []byte) {
	new.setPtr(idx, ptr)
	pos := new.kvPos(idx)
	binary.LittleEndian.PutUint16(new.data[pos+0:], uint16(len(key)))
	binary.LittleEndian.PutUint16(new.data[pos+2:], uint16(len(val)))
	copy(new.data[pos+4:], key)
	copy(new.data[pos+4+uint16(len(key)):], val)
	new.setOffset(idx+1, new.getOffset(idx)+uint16(4+len(key)+len(val)))

}
func nodeAppendRange(new BNode, old BNode, dstNew uint16, srcOld uint16, n uint16) {
	for i := uint16(0); i < n; i++ {
		nodeAppendKV(new, dstNew+i, old.getPtr(srcOld+i), old.getKey(srcOld+i), old.getVal(srcOld+i))
	}
}

func leafInsert(new BNode, old BNode, idx uint16, key []byte, val []byte) {
	new.setHeader(BNODE_LEAF, old.nkeys()+1)
	nodeAppendRange(new, old, 0, 0, idx)
	nodeAppendKV(new, idx, 0, key, val)
	nodeAppendRange(new, old, idx+1, idx, old.nkeys()-idx)
}
>>>>>>> 9753a43 (insert function is added)
