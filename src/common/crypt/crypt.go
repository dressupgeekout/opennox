// Package crypt implements Blowfish encryption algorithm with keys used in Nox.
package crypt

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/blowfish"
)

// Constants for known crypto keys.
const (
	SoundSetBin = 5
	ThingBin    = 7
	GameDataBin = 8
	ModifierBin = 13
	MonsterBin  = 23
	MapKey      = 19
)

const Block = blowfish.BlockSize

var errInvalidSize = errors.New("invalid buffer size")

// KeyForFile return crypto key for a given file. If the file is unknown, it returns false.
func KeyForFile(path string) (int, bool) {
	path = filepath.Base(path)
	path = strings.ToLower(path)
	switch path {
	case "soundset.bin":
		return SoundSetBin, true
	case "thing.bin":
		return ThingBin, true
	case "gamedata.bin":
		return GameDataBin, true
	case "modifier.bin":
		return ModifierBin, true
	case "monster.bin":
		return MonsterBin, true
	}
	switch filepath.Ext(path) {
	case ".map":
		return MapKey, true
	}
	return 0, false
}

// Encode a buffer with a given key.
func Encode(p []byte, key int) error {
	if len(p)%Block != 0 {
		return errInvalidSize
	}
	c, err := NewCipher(key)
	if err != nil {
		return err
	}
	for i := 0; i < len(p); i += Block {
		b := p[i : i+Block]
		c.Encrypt(b, b)
	}
	return nil
}

// Decode a buffer with a given key.
func Decode(p []byte, key int) error {
	if len(p)%Block != 0 {
		return errInvalidSize
	}
	c, err := NewCipher(key)
	if err != nil {
		return err
	}
	for i := 0; i < len(p); i += Block {
		b := p[i : i+Block]
		c.Decrypt(b, b)
	}
	return nil
}

// NewCipher creates a new cipher using Nox key with a given index.
func NewCipher(key int) (*blowfish.Cipher, error) {
	if key < 0 || key > maxKeyInd {
		return nil, fmt.Errorf("invalid key index: %d", key)
	}
	data := keyByInd(key)
	return blowfish.NewCipher(data)
}

const (
	tableSize     = 896
	tableKeySize  = 56
	tableElemSize = 28

	maxKeyInd = (tableSize - tableKeySize) / tableElemSize
)

func keyByInd(ind int) []byte {
	out := keys[tableElemSize*ind:]
	return out[:tableKeySize:tableKeySize]
}

var keys = [tableSize]byte{
	0x6a, 0xd2, 0x04, 0xab, 0x3c, 0x47, 0xbe, 0x4d, 0x84, 0xe5, 0x6c, 0xae, 0x1d, 0x49, 0xe2, 0xd4,
	0xc7, 0x2c, 0x68, 0x0c, 0x37, 0xd5, 0x11, 0xf5, 0x0a, 0x92, 0xe7, 0xb6, 0x72, 0xc4, 0x4b, 0x26,
	0x3d, 0x6f, 0xb0, 0xfc, 0x4e, 0xce, 0xc7, 0x26, 0xf7, 0x45, 0x99, 0x64, 0x76, 0x70, 0xf6, 0xea,
	0xc3, 0xc4, 0xfd, 0x8f, 0x0c, 0xf3, 0xad, 0x17, 0xac, 0x8d, 0x4a, 0xc2, 0xae, 0x22, 0x45, 0x4a,
	0x4a, 0x33, 0x5e, 0x06, 0x2e, 0x42, 0xd4, 0x5d, 0x96, 0xfe, 0xac, 0xc9, 0x2e, 0x7a, 0x19, 0xe2,
	0x75, 0x0e, 0x3f, 0x1f, 0x0b, 0x84, 0x7c, 0x85, 0xb5, 0xe6, 0xa0, 0xb1, 0xb1, 0x18, 0xd4, 0x14,
	0xf7, 0xf1, 0x5e, 0x59, 0xaa, 0xa7, 0x46, 0xb2, 0x41, 0x80, 0x7e, 0x31, 0xf7, 0x45, 0x99, 0xcb,
	0xb4, 0x90, 0x73, 0x3d, 0xae, 0x40, 0x31, 0x3a, 0x1e, 0xc5, 0x7c, 0x98, 0x9a, 0x66, 0x4a, 0xfe,
	0xfb, 0x4c, 0x28, 0x1f, 0x34, 0x77, 0x0e, 0x79, 0xd1, 0x7d, 0x9d, 0x0d, 0x93, 0xfb, 0x03, 0x49,
	0xb7, 0xe6, 0x95, 0xb1, 0xe9, 0x13, 0x2b, 0x69, 0xd5, 0xd7, 0x8c, 0xf7, 0xe3, 0xa1, 0x0e, 0x72,
	0x3d, 0x25, 0xc2, 0x65, 0xe8, 0xf4, 0x2b, 0x83, 0xa3, 0x96, 0x57, 0xc4, 0xf0, 0x3c, 0x9a, 0x76,
	0x5c, 0x9a, 0xa2, 0x48, 0x7b, 0x94, 0x96, 0x20, 0x28, 0x73, 0x9d, 0x57, 0x30, 0xdc, 0x88, 0x64,
	0x3e, 0xe5, 0x61, 0x56, 0xe7, 0xad, 0x48, 0xb1, 0xd6, 0xa3, 0x57, 0x68, 0x76, 0x5b, 0xd7, 0xa1,
	0x0a, 0xe5, 0xe0, 0x11, 0x1f, 0x41, 0xc7, 0x7b, 0x43, 0xae, 0x56, 0x19, 0xe9, 0xa1, 0x37, 0x58,
	0x7b, 0xa4, 0xfe, 0xb6, 0x5b, 0x23, 0xae, 0xa4, 0x6e, 0xaa, 0xd1, 0x71, 0x4a, 0x6c, 0x4c, 0xac,
	0x08, 0x43, 0x83, 0xbb, 0x4f, 0x99, 0xce, 0x4c, 0x9d, 0x09, 0x0c, 0x90, 0x6d, 0x26, 0xba, 0x80,
	0x55, 0xa2, 0xd2, 0x9f, 0x33, 0xd2, 0x86, 0x68, 0x57, 0x77, 0x70, 0x80, 0x8a, 0x57, 0xb6, 0x2c,
	0x4d, 0xae, 0xc6, 0xa7, 0x1a, 0x31, 0xce, 0x8d, 0xbb, 0x28, 0xb5, 0xd0, 0xba, 0xb7, 0x92, 0x07,
	0x30, 0xf6, 0x09, 0x82, 0xb7, 0x3b, 0xd6, 0x53, 0x39, 0xa5, 0x04, 0x51, 0x6a, 0xe9, 0x0f, 0xc7,
	0xfe, 0xfc, 0x9f, 0x2b, 0xfa, 0x50, 0xe1, 0x9f, 0x6c, 0x39, 0xd4, 0x68, 0xbc, 0x02, 0x34, 0xbb,
	0x27, 0xa7, 0x14, 0x65, 0xce, 0x94, 0x0e, 0x9e, 0x2f, 0xcc, 0x92, 0x09, 0x04, 0x29, 0xb4, 0xc1,
	0x02, 0x4f, 0xd4, 0x99, 0xe2, 0x15, 0xe0, 0xf5, 0x0f, 0xa4, 0x8b, 0x5f, 0xc3, 0xa7, 0x09, 0x3d,
	0x51, 0xbd, 0x33, 0xd1, 0xff, 0x7d, 0x13, 0x60, 0x96, 0x74, 0xe3, 0x55, 0x94, 0x68, 0xd1, 0x69,
	0x44, 0x7b, 0xeb, 0x6a, 0xac, 0xd9, 0xe4, 0xad, 0x7e, 0xe7, 0xad, 0x9c, 0x55, 0xe0, 0x73, 0x2e,
	0x11, 0xc0, 0x1a, 0xd8, 0xfe, 0x70, 0xe1, 0xe6, 0xf5, 0xf8, 0x36, 0xcc, 0xd7, 0x9d, 0xbc, 0xa9,
	0x56, 0xf8, 0xad, 0x02, 0x9b, 0x59, 0x50, 0x38, 0xc0, 0x0c, 0x70, 0xfa, 0x9b, 0x74, 0x1a, 0x38,
	0x67, 0xf7, 0x2a, 0x1d, 0x14, 0x89, 0x1a, 0x99, 0x2b, 0x21, 0xcd, 0x7f, 0x7b, 0x3e, 0x6b, 0xdd,
	0x87, 0x85, 0x08, 0x52, 0x38, 0xb7, 0xdf, 0x52, 0x4a, 0x2e, 0x90, 0x5d, 0xb2, 0x76, 0xa6, 0xa6,
	0x80, 0xf3, 0xe6, 0x13, 0x93, 0x40, 0x39, 0xc9, 0xd1, 0x37, 0x2b, 0xd8, 0xc5, 0xa8, 0x2d, 0x90,
	0xd2, 0x6b, 0x25, 0x49, 0x97, 0xc7, 0x3d, 0x55, 0x9b, 0x74, 0x7f, 0xdc, 0x76, 0xeb, 0x16, 0xe2,
	0x13, 0x7e, 0xbe, 0xc8, 0xfb, 0x5f, 0xa0, 0x46, 0xc2, 0xf0, 0xdb, 0xe2, 0x5e, 0x81, 0xaf, 0xb9,
	0x2a, 0xf3, 0xcd, 0xa5, 0xbe, 0x9b, 0xaf, 0xd7, 0xbb, 0x19, 0xf1, 0xc4, 0xfe, 0xb5, 0xc9, 0x53,
	0x1c, 0xff, 0x29, 0x8e, 0x85, 0x8f, 0x5a, 0x2c, 0xf7, 0x3e, 0x34, 0xf9, 0x8d, 0x9f, 0x24, 0x8f,
	0x13, 0x21, 0x14, 0xf1, 0x31, 0xd1, 0xd3, 0x4e, 0xe9, 0x2d, 0x83, 0x31, 0xc9, 0x6d, 0xe7, 0x09,
	0xb0, 0xc8, 0xfa, 0x46, 0x41, 0xdd, 0xd4, 0x8f, 0xc1, 0x73, 0xf7, 0x70, 0xbd, 0x7e, 0xca, 0xe6,
	0x70, 0xa7, 0xb6, 0xe2, 0x67, 0xf6, 0x5e, 0x08, 0xc2, 0x7b, 0x5b, 0x41, 0xab, 0x39, 0x37, 0xa9,
	0xbf, 0x0f, 0x63, 0x99, 0x5c, 0x09, 0x49, 0x7a, 0x27, 0xfe, 0x0d, 0x23, 0x0b, 0xd4, 0xf6, 0xdf,
	0x2c, 0x63, 0x21, 0x8f, 0x0b, 0xef, 0xa1, 0x86, 0x1b, 0xb6, 0xa3, 0x9a, 0x82, 0x2f, 0x03, 0x3b,
	0x67, 0x8b, 0xdb, 0xf8, 0x23, 0xbe, 0xe8, 0x05, 0xc5, 0x61, 0xa7, 0x98, 0xfc, 0xf9, 0x83, 0x3b,
	0xcc, 0x52, 0x76, 0x47, 0xf8, 0xbd, 0xe7, 0xb0, 0x67, 0x60, 0x81, 0xd9, 0x93, 0xf9, 0x0f, 0x45,
	0x57, 0xe6, 0xfb, 0x55, 0x81, 0x18, 0xea, 0x58, 0xd7, 0x7d, 0x8f, 0x09, 0xdf, 0x78, 0x61, 0xdf,
	0xad, 0xbe, 0x3e, 0x9e, 0xff, 0x41, 0x9f, 0x58, 0xdf, 0x75, 0x40, 0x88, 0xff, 0xba, 0x3e, 0x85,
	0xf3, 0xbd, 0x1d, 0x3c, 0x15, 0xf6, 0x16, 0xa8, 0xe8, 0x6c, 0x73, 0xd3, 0x97, 0xbe, 0x59, 0x51,
	0xc3, 0x13, 0x9a, 0xde, 0x29, 0xd4, 0x87, 0xdd, 0xc5, 0x65, 0x12, 0x04, 0x3d, 0x29, 0x9c, 0x0a,
	0x3b, 0x1e, 0x79, 0x24, 0x0a, 0x14, 0xbc, 0xc5, 0x73, 0xe7, 0x81, 0xc0, 0x30, 0x98, 0xbf, 0x84,
	0x23, 0xcc, 0xaf, 0x76, 0x23, 0x30, 0xe8, 0xdc, 0xae, 0x17, 0xbf, 0x22, 0xbc, 0x97, 0x1e, 0xb2,
	0xea, 0x20, 0xdf, 0x15, 0xf1, 0xdd, 0x4a, 0xe0, 0x27, 0x3e, 0x75, 0x93, 0xf2, 0xe9, 0xfb, 0x8e,
	0xdf, 0x16, 0x52, 0x00, 0xca, 0x9f, 0xc6, 0x52, 0x7a, 0x07, 0xf7, 0xd1, 0xf9, 0x77, 0xd3, 0x87,
	0xf7, 0x69, 0x87, 0x7c, 0x64, 0x12, 0x66, 0xc0, 0xe2, 0xd1, 0xc5, 0x6b, 0x9f, 0x9e, 0x1a, 0x10,
	0xe0, 0x79, 0x65, 0xdb, 0xc5, 0x6b, 0x10, 0x00, 0xf9, 0xbc, 0xbc, 0xc0, 0x4c, 0xc7, 0xc7, 0x2f,
	0x14, 0xee, 0x7a, 0xef, 0x8e, 0x41, 0x20, 0xd5, 0x82, 0x32, 0x17, 0x96, 0xc1, 0x2b, 0x69, 0x16,
	0xf2, 0x66, 0xe8, 0x32, 0x72, 0x84, 0xbd, 0x91, 0xeb, 0x0e, 0xe7, 0x84, 0xf7, 0x37, 0x5a, 0x23,
	0x02, 0xa0, 0xe1, 0x53, 0x2d, 0xe8, 0x63, 0xd5, 0xc6, 0x98, 0x44, 0x92, 0xfb, 0x66, 0x6f, 0xbb,
	0x29, 0x45, 0x28, 0xa9, 0x7c, 0xc9, 0x41, 0x16, 0x94, 0x74, 0x8b, 0x16, 0x2d, 0x91, 0xa2, 0x35,
	0x35, 0x53, 0x90, 0x61, 0xbb, 0x28, 0xb5, 0xd0, 0xba, 0xb7, 0x92, 0x07, 0x30, 0xf6, 0x09, 0x82,
	0xb7, 0x3b, 0xd6, 0x53, 0x39, 0xa5, 0x04, 0x51, 0x6a, 0xe9, 0x0f, 0xc7, 0xfe, 0xfc, 0x9f, 0x2b,
}
