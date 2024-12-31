package day09

import (
	"slices"
	"strconv"
)

type DiskMap struct {
	indexed []*int
}

func NewDiskMap(input string) *DiskMap {
	var indexed []*int
	id := 0
	for i, block := range input {
		digit, _ := strconv.Atoi(string(block))
		if i%2 == 0 { // file
			for range digit {
				id := id
				indexed = append(indexed, &id)
			}
			id++
		} else { // empty
			for range digit {
				indexed = append(indexed, nil)
			}
		}
	}
	return &DiskMap{indexed}
}

func (dm *DiskMap) String() (res string) {
	for _, d := range dm.indexed {
		if d == nil {
			res += "."
		} else {
			res += strconv.Itoa(*d)
		}
	}
	return
}

func (dm *DiskMap) Compact() *DiskMap {
	for lp, rp := 0, len(dm.indexed)-1; lp < rp; {
		if dm.indexed[lp] != nil {
			lp++
			continue
		}
		if dm.indexed[rp] == nil {
			rp--
			continue
		}
		dm.indexed[lp] = dm.indexed[rp]
		dm.indexed[rp] = nil
	}
	return dm
}

type Slot struct {
	Index int
	Size  int
	Chunk []*int
}

func (dm *DiskMap) Compact2() *DiskMap {
	// scan
	var frees, files []*Slot
	for i := 0; i < len(dm.indexed); {
		j := i + 1
		if dm.indexed[i] == nil { //free
			for ; j < len(dm.indexed) && dm.indexed[j] == nil; j++ {
			}
			frees = append(frees, &Slot{i, j - i, dm.indexed[i:j]})
		} else { //file
			for ; j < len(dm.indexed) && dm.indexed[j] != nil && *dm.indexed[i] == *dm.indexed[j]; j++ {
			}
			files = append(files, &Slot{i, j - i, dm.indexed[i:j]})
		}
		i += j - i
	}
	slices.Reverse(files)
	for _, file := range files {
		for _, free := range frees {
			if file.Index < free.Index {
				break
			}
			if file.Size > free.Size {
				continue
			}
			for i := range file.Chunk {
				free.Chunk[0] = file.Chunk[i]
				file.Chunk[i] = nil
				free.Chunk = free.Chunk[1:]
				free.Size--
			}
			break
		}
	}
	return dm
}

func (dm *DiskMap) Checksum() (res int64) {
	for i, r := range dm.indexed {
		if r != nil {
			res += int64(i) * int64(*r)
		}
	}
	return
}
