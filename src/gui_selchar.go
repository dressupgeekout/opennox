package nox

/*
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME3_2.h"
extern uint32_t dword_5d4594_1563080;
extern uint32_t dword_5d4594_1563084;
extern uint32_t dword_5d4594_1563088;
extern uint32_t dword_5d4594_1563092;
extern uint32_t dword_5d4594_1563096;
extern uint32_t dword_5d4594_830872;
extern uint32_t dword_5d4594_830972;
extern uint32_t dword_5d4594_831224;
extern unsigned int dword_5d4594_251744;
extern unsigned int nox_gameDisableMapDraw_5d4594_2650672;
*/
import "C"
import (
	"encoding/binary"
	"io"

	"nox/v1/common/alloc"
	"nox/v1/common/datapath"
	noxflags "nox/v1/common/flags"
	"nox/v1/common/log"
	"nox/v1/common/memmap"
)

var (
	saveLog = log.New("save")
)

//export sub_413A00
func sub_413A00(a1 C.int) {
	if noxflags.HasGame(2048) {
		if a1 != 0 {
			noxflags.SetGame(0x40000)
		} else {
			if C.dword_5d4594_251744 == 0 {
				noxflags.UnsetGame(0x40000)
				nox_ticks_reset_416D40()
			}
		}
	}
}

//export sub_448640
func sub_448640() { C.sub_44A400() }

//export sub_4DCE60
func sub_4DCE60(a1 C.int) {
	*memmap.PtrUint32(0x5D4594, 1563100) = uint32(a1)
}

//export sub_4DCE80
func sub_4DCE80(a1 *C.char) {
	sub4DCE80(GoString(a1))
}

func sub4DCE80(a1 string) {
	ptr := memmap.PtrOff(0x5D4594, 1563104)
	alloc.Memset(ptr, 0, 20)
	StrCopy((*C.char)(ptr), 20, a1)
}

//export sub_41D090
func sub_41D090(a1 *C.char) C.int {
	v, err := sub41D090(GoString(a1))
	if err != nil {
		saveLog.Println(err)
		return 0
	}
	return C.int(v)
}

func sub41D090(path string) (uint32, error) {
	if err := cryptFileOpen(path, 1, 27); err != nil {
		return 0, err
	}
	defer cryptFileClose()
	var buf [4]byte
	for {
		_, err := cryptFileReadWrite(buf[:4])
		if err == io.EOF {
			return 0, nil
		} else if err != nil {
			return 0, err
		}
		a1 := binary.LittleEndian.Uint32(buf[:])
		if a1 == 0 {
			return 0, nil
		}
		cryptFileReadMaybeAlign(buf[:4])
		v3 := binary.LittleEndian.Uint32(buf[:])
		if a1 == 10 {
			return sub_41D110()
		}
		nox_xxx_cryptSeekCur(int64(v3))
	}
}

func sub_41D110() (uint32, error) {
	if !noxflags.HasGame(2048) {
		return 0, nil
	}
	var buf [4]byte
	v2 := uint16(5)
	binary.LittleEndian.PutUint16(buf[:], v2)
	_, err := cryptFileReadWrite(buf[:2])
	v2 = binary.LittleEndian.Uint16(buf[:])
	if int16(v2) <= 5 && int16(v2) >= 5 {
		buf = [4]byte{}
		_, err = cryptFileReadWrite(buf[:4])
		v3 := binary.LittleEndian.Uint32(buf[:])
		return v3, err
	}
	return uint32(v2), err
}

//export sub_4505B0
func sub_4505B0() {
	sub_450580()
	C.nox_gameDisableMapDraw_5d4594_2650672 = 0
	v0 := nox_client_getIntroScreenDuration_44E3B0()
	C.nox_client_screenFadeTimeout_44DAB0(v0, 1, (*[0]byte)(C.sub_44E320))
	C.nox_gameDisableMapDraw_5d4594_2650672 = 1
}

//export sub_450580
func sub_450580() {
	sub_44D8F0()
	*memmap.PtrUint32(0x5D4594, 832488) = 1
	C.dword_5d4594_831224 = 0
	*memmap.PtrUint32(0x5D4594, 831292) = 0
	*memmap.PtrUint32(0x5D4594, 831296) = 0
}

//export sub_4DB170
func sub_4DB170(a1, a2, a3 C.int) {
	C.dword_5d4594_1563092 = C.uint(a3)
	C.dword_5d4594_1563088 = C.uint(gameFrame())
	C.dword_5d4594_1563084 = C.uint(a2)
	C.dword_5d4594_1563080 = C.uint(a1)
	C.dword_5d4594_1563096 = C.uint(bool2int(a2 != 0))
	if a1 == 0 {
		sub_4DCBD0(0)
	}
}

func sub_4DCBD0(a1 int) {
	*memmap.PtrUint32(0x5D4594, 1563076) = uint32(a1)
}

//export sub_44D8F0
func sub_44D8F0() {
	C.dword_5d4594_830872 = 0
	C.dword_5d4594_830972 = 0
}

//export nox_savegame_nameFromPath_4DC970
func nox_savegame_nameFromPath_4DC970(src, dst *C.char, max C.int) {
	name := datapath.SaveNameFromPath(GoString(src))
	StrCopy(dst, int(max), name)
}
