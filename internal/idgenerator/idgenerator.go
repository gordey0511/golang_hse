package idgenerator

import (
	"hash/adler32"
	"hash/fnv"
)

type BookIdGenerator struct {
	genID func(title string) uint32
}

func (generator *BookIdGenerator) GeneratorId(title string) uint32 {
	return generator.genID(title)
}

func fnvID(title string) uint32 {
	generator := fnv.New32a()
	generator.Write([]byte(title))
	return generator.Sum32()
}

func NewFnvGenerator() BookIdGenerator {
	return BookIdGenerator{genID: fnvID}
}

func adlerID(title string) uint32 {
	return adler32.Checksum([]byte(title))
}

func NewAdlerGenerator() BookIdGenerator {
	return BookIdGenerator{genID: adlerID}
}
