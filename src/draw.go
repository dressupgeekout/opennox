package nox

/*
#include "GAME1_2.h"
#include "GAME2_3.h"
#include "GAME3_1.h"
#include "defs.h"
#include "GAME2_1.h"
#include "client__system__client.h"
#include "client__gui__guiquit.h"
#include "client__gui__guiggovr.h"
#include "client__video__draw_common.h"
#include "client__draw__selectdw.h"
#include "client__draw__drawwin.h"
#include "client__draw__debugdraw.h"
extern nox_draw_viewport_t nox_draw_viewport;
extern nox_drawable* nox_xxx_drawablePlayer_1046600;
extern unsigned int nox_client_drawFrontWalls_80812;
extern unsigned int nox_client_translucentFrontWalls_805844;
extern unsigned int nox_client_highResFrontWalls_80820;
extern unsigned int nox_client_highResFloors_154952;
extern unsigned int nox_client_lockHighResFloors_1193152;
extern unsigned int nox_client_texturedFloors_154956;
extern unsigned int nox_gui_console_translucent;
extern unsigned int nox_client_renderGlow_805852;
extern unsigned int nox_client_fadeObjects_80836;
extern unsigned int nox_client_renderBubbles_80844;
extern unsigned int nox_client_renderGUI_80828;
extern unsigned int nox_client_showTooltips_80840;
extern unsigned int nox_xxx_xxxRenderGUI_587000_80832;
extern unsigned int nox_profiled_805856;
extern unsigned int nox_xxx_waypointCounterMB_587000_154948;
extern unsigned char nox_arr_84EB20[280*57*4];
extern unsigned int dword_5d4594_816440;
extern unsigned int dword_5d4594_816456;
extern unsigned int dword_5d4594_816460;
extern void* nox_draw_defaultFont_816492;
extern unsigned int dword_5d4594_2523804;
extern unsigned int dword_5d4594_2650676;
extern unsigned int dword_5d4594_2650680;
extern unsigned int dword_5d4594_811896;
extern unsigned int dword_5d4594_811904;
extern unsigned int nox_client_gui_flag_815132;
extern unsigned int dword_5d4594_1096432;
extern unsigned int dword_5d4594_3799452;
extern unsigned int dword_5d4594_3799468;
extern unsigned int dword_5d4594_3799484;
extern unsigned int dword_5d4594_3799476;
extern unsigned int dword_5d4594_3799508;
extern unsigned int nox_client_texturedFloors2_154960;
extern unsigned int dword_5d4594_1193156;
extern unsigned int dword_5d4594_1193188;
extern void* nox_draw_sprite_dstPtr_3799540;
extern unsigned int dword_5d4594_3799552;
extern unsigned char* nox_video_cur_pixdata_3799444;
extern int dword_5d4594_3799524;
extern unsigned int nox_client_gui_flag_1556112;
extern unsigned int nox_gameDisableMapDraw_5d4594_2650672;
extern void* dword_5d4594_810640;
extern void* nox_draw_colorTablesRev_3804668;
void nox_xxx_tileDrawMB_481C20_A(nox_draw_viewport_t* vp, int v3);
void nox_xxx_tileDrawMB_481C20_B(nox_draw_viewport_t* vp, int v78);
void nox_xxx_tileDrawMB_481C20_C_textured(nox_draw_viewport_t* vp, int v72, int v78);
void nox_xxx_tileDrawMB_481C20_C_solid(nox_draw_viewport_t* vp, int v72, int v78);
void sub_4C86B0();
void  nox_xxx_cliLight16_469140(nox_drawable* dr, nox_draw_viewport_t* vp);
void nox_xxx_clientDrawAll_436100_draw_A();
void nox_xxx_clientDrawAll_436100_draw_B();
void nox_xxx_drawAllMB_475810_draw_A(nox_draw_viewport_t* vp);
int nox_xxx_drawAllMB_475810_draw_B(nox_draw_viewport_t* vp);
void nox_xxx_drawAllMB_475810_draw_C(nox_draw_viewport_t* vp, int v36, int v7);
void nox_xxx_drawAllMB_475810_draw_D(nox_draw_viewport_t* vp);
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"image"
	"sort"
	"unicode"
	"unicode/utf16"
	"unsafe"

	"nox/v1/client/input"
	"nox/v1/common"
	"nox/v1/common/alloc"
	noxcolor "nox/v1/common/color"
	noxflags "nox/v1/common/flags"
	"nox/v1/common/memmap"
	"nox/v1/common/noximage"
	"nox/v1/common/types"
)

var nox_draw_colorTablesRev_3804668 []byte // map[Color16]byte

func getViewport() *Viewport {
	return asViewport(&C.nox_draw_viewport)
}

func asViewport(p *C.nox_draw_viewport_t) *Viewport {
	return (*Viewport)(unsafe.Pointer(p))
}

type Viewport C.nox_draw_viewport_t

func (vp *Viewport) C() *C.nox_draw_viewport_t {
	return (*C.nox_draw_viewport_t)(unsafe.Pointer(vp))
}

func (vp *Viewport) toScreenPos(pos types.Point) types.Point { // sub_4739E0
	return pos.Add(types.Point{
		X: int(vp.x1) - int(vp.field_4),
		Y: int(vp.y1) - int(vp.field_5),
	})
}

func nox_xxx_guiFontHeightMB_43F320(a1 int) int {
	return int(C.nox_xxx_guiFontHeightMB_43F320(C.int(a1)))
}

func detectBestVideoSettings() { // nox_setProfiledMode_4445C0
	const cfg = 450
	cut := 0
	*memmap.PtrUint32(0x587000, 80808) = 1
	*memmap.PtrUint32(0x587000, 80816) = 1
	C.nox_client_showTooltips_80840 = 1
	C.nox_client_renderGUI_80828 = 1
	C.nox_xxx_xxxRenderGUI_587000_80832 = 1
	if cfg == 0 {
		//v4 = 8
		cut = 75
		C.nox_client_drawFrontWalls_80812 = 0
		C.nox_client_translucentFrontWalls_805844 = 0
		C.nox_client_highResFrontWalls_80820 = 0
		C.nox_client_highResFloors_154952 = 0
		C.nox_client_lockHighResFloors_1193152 = 0
		C.nox_client_texturedFloors_154956 = 1
		C.nox_gui_console_translucent = 0
		C.nox_client_renderGlow_805852 = 0
		C.nox_client_fadeObjects_80836 = 0
		resetEngineFlag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE)
		C.nox_client_renderBubbles_80844 = 0
	} else if cfg == 200 {
		//v4 = 8
		cut = 85
		C.nox_client_drawFrontWalls_80812 = 1
		C.nox_client_translucentFrontWalls_805844 = 0
		C.nox_client_highResFrontWalls_80820 = 0
		C.nox_client_highResFloors_154952 = 0
		C.nox_client_lockHighResFloors_1193152 = 0
		C.nox_client_texturedFloors_154956 = 1
		C.nox_gui_console_translucent = 0
		C.nox_client_renderGlow_805852 = 0
		C.nox_client_fadeObjects_80836 = 0
		resetEngineFlag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE)
		C.nox_client_renderBubbles_80844 = 0
	} else if cfg == 266 {
		//v4 = 8
		cut = 100
		C.nox_client_drawFrontWalls_80812 = 1
		C.nox_client_translucentFrontWalls_805844 = 1
		C.nox_client_highResFrontWalls_80820 = 1
		C.nox_client_highResFloors_154952 = 1
		C.nox_client_lockHighResFloors_1193152 = 0
		C.nox_client_texturedFloors_154956 = 1
		C.nox_gui_console_translucent = 0
		C.nox_client_renderGlow_805852 = 1
		C.nox_client_fadeObjects_80836 = 1
		setEngineFlag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE)
		C.nox_client_renderBubbles_80844 = 1
	} else if cfg == 300 {
		//v4 = 16
		cut = 100
		C.nox_client_drawFrontWalls_80812 = 1
		C.nox_client_translucentFrontWalls_805844 = 1
		C.nox_client_highResFrontWalls_80820 = 1
		C.nox_client_highResFloors_154952 = 1
		C.nox_client_lockHighResFloors_1193152 = 0
		C.nox_client_texturedFloors_154956 = 1
		C.nox_gui_console_translucent = 0
		C.nox_client_renderGlow_805852 = 1
		C.nox_client_fadeObjects_80836 = 1
		setEngineFlag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE)
		C.nox_client_renderBubbles_80844 = 1
	} else if cfg == 450 {
		//v4 = 16
		cut = 100
		C.nox_client_drawFrontWalls_80812 = 1
		C.nox_client_translucentFrontWalls_805844 = 1
		C.nox_client_highResFrontWalls_80820 = 1
		C.nox_client_highResFloors_154952 = 1
		C.nox_client_lockHighResFloors_1193152 = 1
		C.nox_client_texturedFloors_154956 = 1
		C.nox_gui_console_translucent = 1
		C.nox_client_renderGlow_805852 = 1
		C.nox_client_fadeObjects_80836 = 1
		setEngineFlag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE)
		C.nox_client_renderBubbles_80844 = 1
	}
	C.nox_xxx_tileSetDrawFn_481420()
	if !getEngineFlag(NOX_ENGINE_FLAG_ENABLE_WINDOWED_MODE) {
		videoUpdateGameMode(types.Size{
			W: noxDefaultWidth,
			H: noxDefaultHeight,
		})
	}
	C.nox_video_setCutSize_4766A0(C.int(cut))
	if noxflags.HasGame(0x10000000) {
		nox_draw_setCutSize_476700(C.int(cut), 0)
	}
	C.nox_profiled_805856 = 1
}

func nox_draw_initColorTablesRev_434DA0() {
	const max = 0x7FFF
	const (
		rshift = 7 // -10+3
		gshift = 2 // -5+3
		bshift = 3 // -0+3

		rmask = 0x7c00
		gmask = 0x03e0
		bmask = 0x001f
	)
	nox_draw_colorTablesRev_3804668 = make([]byte, max+3)
	for i := 0; i <= max; i++ {
		cr := uint32((i & rmask) >> rshift)
		cg := uint32((i & gmask) >> gshift)
		cb := uint32((i & bmask) << bshift)
		nox_draw_colorTablesRev_3804668[i] = byte((28*(cb|7) + 150*(cg|7) + 76*(cr|7)) >> 8)
	}
	ptr, _ := alloc.Bytes(uintptr(len(nox_draw_colorTablesRev_3804668)))
	copy(ptr, nox_draw_colorTablesRev_3804668)
	C.nox_draw_colorTablesRev_3804668 = unsafe.Pointer(&ptr[0])
}

func nox_draw_freeColorTables_433C20() {
	if C.dword_5d4594_810640 != nil {
		alloc.Free(C.dword_5d4594_810640)
		C.dword_5d4594_810640 = nil
	}
	if C.nox_draw_colorTablesRev_3804668 != nil {
		alloc.Free(C.nox_draw_colorTablesRev_3804668)
		C.nox_draw_colorTablesRev_3804668 = nil
	}
	if p := C.nox_draw_colors_r_3804672; p != nil {
		alloc.Free(unsafe.Pointer(p))
		C.nox_draw_colors_r_3804672 = nil
	}
	if p := C.nox_draw_colors_g_3804656; p != nil {
		alloc.Free(unsafe.Pointer(p))
		C.nox_draw_colors_g_3804656 = nil
	}
	if p := C.nox_draw_colors_b_3804664; p != nil {
		alloc.Free(unsafe.Pointer(p))
		C.nox_draw_colors_b_3804664 = nil
	}
	*memmap.PtrUint32(0x973F18, 5232) = 0
}

func sub_4338D0() int {
	noxcolor.SetMode(noxcolor.ModeRGBA5551)
	C.dword_975240 = (*[0]byte)(C.sub_435280)
	C.dword_975380 = (*[0]byte)(C.sub_434E80)
	copy(byte_5D4594_3804364[:], byte_581450_9176[:])
	copy(unsafe.Slice((*uint32)(unsafe.Pointer(&C.byte_5D4594_3804364[0])), 40), byte_581450_9176[:])
	ptr := C.ptr_5D4594_3799572
	ptr.field_13 = 0
	ptr.field_14 = 0
	ptr.field_17 = 0
	ptr.field_24 = 0
	ptr.field_25 = 0
	ptr.field_26 = 0
	ptr.field_44 = 0
	ptr.field_45 = 0
	ptr.field_46 = 0
	ptr.field_54 = 0
	ptr.field_55 = 0
	ptr.field_56 = 0
	ptr.field_58 = 0
	ptr.field_59 = 0
	ptr.field_60 = 0
	ptr.field_61 = 0
	ptr.field_259 = 255
	ptr.field_260 = 0xFF00FF
	ptr.field_261 = 0xFF00FF
	ptr.field_262 = 0
	C.sub_434990(25, 25, 25)
	nox_draw_initColorTables_434CC0()
	nox_draw_initColorTablesRev_434DA0()
	if C.dword_5d4594_823772 == 0 {
		C.sub_4353C0()
	}
	*memmap.PtrUint32(0x5D4594, 809596) = 0
	C.dword_5d4594_808568 = 0
	C.dword_5d4594_810628 = 0
	if C.dword_5d4594_823772 != 0 {
		v2, freeV2 := alloc.Malloc(256 * 3)
		defer freeV2()
		C.sub_435150((*C.uchar)(v2), (*C.char)(memmap.PtrOff(0x973F18, 3880)))
		C.sub_4347F0((*C.char)(v2), 256)
	} else {
		C.sub_4347F0((*C.char)(memmap.PtrOff(0x581450, 8404)), 256)
	}
	if C.sub_434FB0() == 0 {
		return 0
	}
	C.sub_4340A0(0, 0, 0, 255)
	C.sub_4340A0(1, 0, 255, 0)
	C.sub_4340A0(2, 0, 255, 255)
	C.sub_4340A0(3, 255, 0, 0)
	C.sub_4340A0(4, 255, 0, 255)
	C.sub_4340A0(5, 255, 255, 0)
	C.sub_4340A0(6, 255, 255, 255)
	C.sub_4340A0(7, 0, 0, 0)
	C.sub_4340A0(8, 0, 0, 128)
	C.sub_4340A0(9, 0, 128, 0)
	C.sub_4340A0(10, 0, 128, 128)
	C.sub_4340A0(11, 128, 0, 0)
	C.sub_4340A0(12, 128, 0, 128)
	C.sub_4340A0(13, 128, 128, 0)
	C.sub_4340A0(14, 128, 128, 128)
	C.sub_4340A0(15, 0, 0, 0)
	// TODO gamma control
	C.dword_5d4594_808564 = 0
	C.sub_4348C0()
	C.sub_4352E0()
	*memmap.PtrUint32(0x973F18, 5232) = 0
	return 1
}

var noxrend = NewNoxRender()

type drawOpFunc func(dst []uint16, src []byte, op byte, val int) (outDst []uint16, outSrc []byte)

func cgoSetRenderData(p *C.nox_render_data_t) {
	C.ptr_5D4594_3799572 = p
}

type NoxRender struct {
	p   *C.nox_render_data_t
	pix *noximage.Image16

	colors struct {
		mode noxcolor.Mode
		R    [256]uint16
		G    [256]uint16
		B    [256]uint16
	}

	draw27 drawOpFunc
	draw4  drawOpFunc
	draw5  drawOpFunc
	draw6  drawOpFunc

	particles struct {
		byOpts   map[particleOpt]*Particle
		byHandle map[unsafe.Pointer]*Particle
	}
}

func (r *NoxRender) PixBuffer() *noximage.Image16 {
	return r.pix
}

func (r *NoxRender) SetPixBuffer(pix *noximage.Image16) {
	r.pix = pix
}

func newNoxRenderData() (*C.nox_render_data_t, func()) {
	p, free := alloc.Malloc(unsafe.Sizeof(C.nox_render_data_t{}))
	return (*C.nox_render_data_t)(p), free
}

func NewNoxRender() *NoxRender {
	r := &NoxRender{}
	r.SetColorMode(noxcolor.ModeRGBA5551)
	return r
}

func (r *NoxRender) shouldDrawGlow() bool {
	return C.nox_client_renderGlow_805852 != 0
}

func (r *NoxRender) Frame() uint32 {
	return gameFrame()
}

func (r *NoxRender) SetData(p *C.nox_render_data_t) {
	r.p = p
}

func (r *NoxRender) Rect() types.Rect {
	return types.Rect{
		Left:   int(r.p.clip.left),
		Top:    int(r.p.clip.top),
		Right:  int(r.p.clip.right),
		Bottom: int(r.p.clip.bottom),
	}
}

func (r *NoxRender) IsAlphaEnabled() bool {
	return r.p.field_13 != 0
}

func (r *NoxRender) SetAlphaEnabled(enabled bool) { // nox_client_drawEnableAlpha_434560
	if enabled != r.IsAlphaEnabled() {
		r.p.field_13 = C.uint(bool2int(enabled))
	}
}

func (r *NoxRender) SetAlpha(v byte) { // nox_client_drawSetAlpha_434580
	if byte(r.p.field_259) != v {
		r.p.field_259 = C.uint(v)
		v2 := uint64(v) | (uint64(v) << 16)
		v2 &= 0xffffffff
		v2 <<= 16
		v2 = (v2 & 0xffffffff00000000) | uint64(v) | (v2 & 0xffffffff)
		v2 <<= 16
		r.p.field_260 = C.uint(uint32(v) | uint32(v2))
		r.p.field_261 = C.uint(v2 >> 32)
	}
}

func (r *NoxRender) setField17(v int) { // 	nox_xxx_draw_434600
	r.p.field_17 = C.uint(v)
}

func (r *NoxRender) SelectColor(a1 uint32) { // nox_xxx_drawSelectColor_434350
	r.p.field_58 = C.uint(a1)
}

func (r *NoxRender) SetTextColor(a1 uint32) { // nox_xxx_drawSetTextColor_434390
	r.p.field_59 = C.uint(a1)
}

func (r *NoxRender) SetColor(a1 uint32) { // nox_xxx_drawSetColor_4343E0
	r.p.field_60 = C.uint(a1)
}

func (r *NoxRender) SetColor2(a1 uint32) { // nox_client_drawSetColor_434460
	r.p.field_61 = C.uint(a1)
}

func (r *NoxRender) field66(ind int) []uint32 {
	const (
		size = 12
		max  = 16
	)
	arr := unsafe.Slice((*uint32)(unsafe.Pointer(&r.p.field_66)), size*max)
	return arr[ind*size : (ind+1)*size]
}

func (r *NoxRender) sub434080(a1 int) { // sub_434080
	r.p.field_262 = C.uint(a1)
}

func (r *NoxRender) GetLightColor() NoxLight { // sub_434A10
	return NoxLight{
		R: int(r.p.field_62),
		G: int(r.p.field_63),
		B: int(r.p.field_64),
	}
}

func (r *NoxRender) sub49EFA0(pos types.Point) { // sub_49EFA0
	C.sub_49EFA0(C.int(pos.X), C.int(pos.Y))
}

func (r *NoxRender) DrawPoint(pos types.Point, rad int) { // nox_xxx_drawPointMB_499B70
	switch rad {
	case 0, 1:
		r.sub49EFA0(pos)
	case 2:
		r.DrawRectFilledOpaque(pos.X, pos.Y, 2, 2)
	case 3:
		r.sub49EFA0(pos.Add(types.Point{Y: -1}))
		r.DrawRectFilledOpaque(pos.X-1, pos.Y, 3, 1)
		r.sub49EFA0(pos.Add(types.Point{Y: +1}))
	case 4:
		r.DrawRectFilledOpaque(pos.X, pos.Y-1, 2, 1)
		r.DrawRectFilledOpaque(pos.X-1, pos.Y, 4, 2)
		r.DrawRectFilledOpaque(pos.X, pos.Y+2, 2, 1)
	case 5:
		r.DrawRectFilledOpaque(pos.X-1, pos.Y-2, 3, 1)
		r.DrawRectFilledOpaque(pos.X-2, pos.Y-1, 5, 3)
		r.DrawRectFilledOpaque(pos.X-1, pos.Y+2, 3, 1)
	case 6:
		r.DrawRectFilledOpaque(pos.X, pos.Y-2, 2, 1)
		r.DrawRectFilledOpaque(pos.X-1, pos.Y-1, 4, 1)
		r.DrawRectFilledOpaque(pos.X-2, pos.Y, 6, 2)
		r.DrawRectFilledOpaque(pos.X-1, pos.Y+2, 4, 1)
		r.DrawRectFilledOpaque(pos.X, pos.Y+3, 2, 1)
	default:
		C.sub_4B0BC0(C.int(pos.X), C.int(pos.Y), C.int(rad/2))
	}
}

func (r *NoxRender) nox_client_drawAddPoint_49F500(pos types.Point) {
	C.nox_client_drawAddPoint_49F500(C.int(pos.X), C.int(pos.Y))
}

func (r *NoxRender) DrawLineFromPoints(arr ...types.Point) { // nox_client_drawLineFromPoints_49E4B0
	for _, p := range arr {
		r.nox_client_drawAddPoint_49F500(p)
	}
	C.nox_client_drawLineFromPoints_49E4B0()
}

func (r *NoxRender) DrawRectFilledOpaque(x, y, w, h int) { // nox_client_drawRectFilledOpaque_49CE30
	C.nox_client_drawRectFilledOpaque_49CE30(C.int(x), C.int(y), C.int(w), C.int(h))
}

func (r *NoxRender) DrawRectFilledAlpha(x, y, w, h int) { // nox_client_drawRectFilledAlpha_49CF10
	C.nox_client_drawRectFilledAlpha_49CF10(C.int(x), C.int(y), C.int(w), C.int(h))
}

func (r *NoxRender) DrawString(font unsafe.Pointer, str string, pos types.Point) { // nox_xxx_drawString_43F6E0
	if getEngineFlag(NOX_ENGINE_FLAG_DISABLE_TEXT_RENDERING) {
		return
	}
	if font == nil {
		font = C.nox_draw_defaultFont_816492
		if font == nil {
			return
		}
	}
	C.dword_5d4594_816460 = C.uint(pos.X)
	r.drawStringLine(font, str, pos)
}

func (r *NoxRender) drawStringLine(font unsafe.Pointer, str string, pos types.Point) { // nox_xxx_guiDrawString_4407F0
	for _, c := range str {
		if c == '\n' || c == '\r' {
			continue
		}
		pos.X = r.drawChar(font, c, pos)
	}
}

func rune2wchar(r rune) uint16 {
	// TODO: should we drop the second one?
	r1, _ := utf16.EncodeRune(r)
	if r1 == unicode.ReplacementChar {
		return uint16(r)
	}
	return uint16(r1)
}

func nox_xxx_FontGetChar_43FE30(font unsafe.Pointer, r rune) unsafe.Pointer {
	p := C.nox_xxx_FontGetChar_43FE30(font, C.ushort(rune2wchar(r)))
	if p != nil {
		return p
	}
	p = C.nox_xxx_FontGetChar_43FE30(font, C.ushort('?'))
	if p != nil {
		return p
	}
	return nil
}

func (r *NoxRender) drawChar(font unsafe.Pointer, c rune, pos types.Point) int {
	return int(C.nox_xxx_StringDraw_43FE90(font, C.short(rune2wchar(c)), C.int(pos.X), C.int(pos.Y)))
}

func (r *NoxRender) DrawStringWrapped(font unsafe.Pointer, s string, x, y, maxW, maxH int) { // nox_xxx_drawString_43FAF0
	if getEngineFlag(NOX_ENGINE_FLAG_DISABLE_TEXT_RENDERING) {
		return
	}
	if font == nil {
		font = C.nox_draw_defaultFont_816492
		if font == nil {
			return
		}
	}
	if s == "" {
		return
	}
	dy := int(*(*uint32)(unsafe.Add(font, 28)))
	C.dword_5d4594_816460 = C.uint(x)
	str := []rune(s)

	var (
		wordX int    // relative to X
		wordY int    // relative to Y
		word  []rune // current word
		wordN int    // words in the current line
		addX  int    // relative to wordX
	)

	drawWord := func() {
		cx := x + wordX
		cy := y + wordY
		for _, c := range word {
			cx = r.drawChar(font, c, types.Point{X: cx, Y: cy})
		}
		word = word[:0]
		wordX = cx - x
		addX = 0
		wordN++
	}

	for i, c := range str {
		switch c {
		case '\t':
			drawWord()
			tab := int(C.dword_5d4594_816456)
			wordX += tab
			wordX -= wordX % tab
			continue
		case '\r', '\n':
			if i != 0 && str[i-1] == '\r' {
				continue
			}
			drawWord()
			wordX = 0
			wordN = 0
			wordY += dy
			if maxH > 0 && wordY >= maxH {
				return
			}
			continue
		}
		if maxW > 0 {
			cv := nox_xxx_FontGetChar_43FE30(font, c)
			if cv == nil {
				continue
			}
			dx := int(*(*byte)(cv))
			addX += dx
			if wordN != 0 && wordX+addX > maxW {
				// word is too long, but not the first in a line
				wordX = 0
				wordN = 0
				wordY += dy
				if maxH > 0 && wordY >= maxH {
					return
				}
			}
		}
		word = append(word, c)
		if c == ' ' {
			drawWord()
		}
	}
	drawWord()
}

func (r *NoxRender) GetStringSize(a1 int, a2 string, a5 int) types.Size { // nox_xxx_drawGetStringSize_43F840
	sp, freeS := CWString(a2)
	defer freeS()
	pp, freeP := alloc.Malloc(unsafe.Sizeof(C.nox_point{}))
	defer freeP()
	p := (*C.nox_point)(pp)
	C.nox_xxx_drawGetStringSize_43F840(C.int(a1), sp, &p.x, &p.y, C.int(a5))
	return types.Size{W: int(p.x), H: int(p.y)}
}

func (r *NoxRender) DrawCircle(a1, a2, a3 int) {
	if r.IsAlphaEnabled() {
		C.nox_video_drawCircle16Alpha_4B2480(C.int(a1), C.int(a2), C.int(a3))
	} else {
		C.nox_video_drawCircle16Opaque_4B1380(C.int(a1), C.int(a2), C.int(a3))
	}
}

func (r *NoxRender) DrawCircleColored(a1, a2, a3 int, a4 uint32) {
	r.SetColor2(a4)
	r.DrawCircle(a1, a2, a3)
}

//export nox_xxx_drawString_43FAF0
func nox_xxx_drawString_43FAF0(font unsafe.Pointer, sp *C.wchar_t, x, y, a5, a6 C.int) C.int {
	s := GoWString(sp)
	noxrend.DrawStringWrapped(font, s, int(x), int(y), int(a5), int(a6))
	return 1
}

//export nox_video_drawAnimatedImageOrCursorAt_4BE6D0
func nox_video_drawAnimatedImageOrCursorAt_4BE6D0(a1, a2, a3 C.int) {
	noxrend.nox_video_drawAnimatedImageOrCursorAt(asImageRefP(unsafe.Pointer(uintptr(a1))), types.Point{X: int(a2), Y: int(a3)})
}

func (r *NoxRender) nox_video_drawAnimatedImageOrCursorAt(ref *noxImageRef, pos types.Point) {
	if v3 := asImageP(unsafe.Pointer(uintptr(C.sub_4BE640(C.int(uintptr(unsafe.Pointer(ref.C()))), 0)))); v3 != nil {
		if C.dword_5d4594_3798728 != 0 {
			r.noxDrawCursor(v3, pos)
		} else {
			r.DrawImageAt(v3, pos)
		}
	}
}

func nox_xxx_client_435F80_draw(inp *input.Handler) bool {
	mpos := inp.GetMousePos()
	if C.nox_xxx_serverIsClosing_446180() != 0 {
		C.sub_446190()
	}
	if sub_437060(inp) == 0 && !noxflags.HasGame(8) {
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		return false
	}
	C.sub_437100()
	if C.sub_478030() == 0 && !nox_xxx_guiCursor_477600() {
		C.sub_470DE0()
	}
	ctrlEvent.nox_xxx_input_42D220(inp)
	if C.nox_xxx_get_430890() == 1 {
		C.sub_430880(0)
	}
	C.nox_xxx_clientEnumHover_476FA0()
	vp := noxGetViewport()
	ipx := int(vp.field_4) + mpos.X - int(vp.x1)
	ipy := int(vp.field_5) + mpos.Y - int(vp.y1)
	if !clientSendInput(31, uint16(ipx), uint16(ipy)) {
		return true
	}
	if !noxflags.HasGame(1) {
		C.sub_40B970()
		C.sub_40B790()
	}
	if noxflags.HasGame(1) {
		C.nox_xxx_spriteDeleteSomeList_49C4B0()
		if C.nox_netlist_receiveCli_494E90(31) == 0 {
			return true
		}
	} else {
		C.sub_43CCA0()
	}
	C.sub_49BD70(vp)
	C.sub_49BBC0()
	C.nox_xxx_polygonDrawColor_421B80()
	if C.nox_client_isConnected_43C700() != 0 {
		C.nox_xxx_cliToggleObsWindow_4357A0()
	}
	nox_xxx_clientDrawAll_436100_draw()
	C.sub_49BB40()
	C.sub_49BA70()
	maybeScreenshot()
	if memmap.Uint32(0x5D4594, 826068) != 0 {
		C.nox_xxx_motd_4467F0()
	}
	if C.sub_42EBA0() == 1 {
		C.sub_42EB90(0)
		C.nox_xxx_consoleEsc_49B7A0()
	}
	C.sub_49B6E0()
	return C.int(memmap.Uint32(0x587000, 85720)) != 0
}

func nox_xxx_clientDrawAll_436100_draw() {
	v0 := platformTicks()
	isTick := false
	if int(v0-memmap.Uint64(0x5D4594, 814532)) >= int(memmap.Int32(0x587000, 85748)) {
		isTick = true
		nox_ticks_reset_416D40()
	}
	if !(memmap.Uint32(0x587000, 85724) == 0 || isTick || !noxflags.HasGame(1) || nox_ticks_check_416D70() ||
		C.nox_client_gui_flag_815132 != 0 || nox_xxx_checkGameFlagPause_413A50() || getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		setEngineFlag(NOX_ENGINE_FLAG_PAUSE)
		return
	}
	resetEngineFlag(NOX_ENGINE_FLAG_PAUSE)
	*memmap.PtrUint64(0x5D4594, 814532) = v0
	*memmap.PtrUint32(0x5D4594, 811916) = gameFrame()
	vp := getViewport()
	v6 := int(vp.x1)
	v7 := int(vp.y1)
	if memmap.Uint32(0x587000, 85744) != 0 {
		vp.height = vp.width * C.int(nox_win_height) / C.int(nox_win_width)
		v6 = (nox_win_width - int(vp.width)) / 2
		v7 = (nox_win_height - int(vp.height)) / 2
		vp.x1 = C.int(v6)
		vp.y1 = C.int(v7)
		vp.x2 = C.int(v6) + vp.width - 1
		vp.y2 = C.int(v7) + vp.height - 1
	}
	C.sub_430B50(C.int(v6), C.int(v7), vp.x2, vp.y2)
	if id := clientPlayerNetCode(); id != 0 {
		*memmap.PtrPtr(0x852978, 8) = unsafe.Pointer(C.nox_xxx_netSpriteByCodeDynamic_45A6F0(C.int(id)))
	}
	if getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING) {
		C.nox_xxx_clientDrawAll_436100_draw_A()
	} else if memmap.Uint32(0x852978, 8) != 0 && C.nox_client_isConnected_43C700() != 0 {
		nox_xxx_drawAllMB_475810_draw(vp)
		C.nox_xxx_drawMinimapAndLines_4738E0()
	} else {
		noxrend.SelectColor(memmap.Uint32(0x85B3FC, 952))
		nox_client_clearScreen_440900()
	}
	if C.dword_5d4594_811896 != 0 {
		C.sub_4365C0()
	}
	if C.dword_5d4594_811904 != 0 {
		C.sub_436F50()
	}
	if noxflags.HasGame(8) {
		v13 := platformTicks() - memmap.Uint64(0x5D4594, 811908)
		if v13 > 10000 && !noxflags.HasGame(1) {
			noxflags.UnsetGame(8)
		}
	}
	if noxflags.HasGame(8) {
		C.nox_xxx_clientDrawAll_436100_draw_B()
	}
	if memmap.Uint32(0x587000, 85744) != 0 {
		C.sub_430B50(0, 0, C.int(nox_win_width-1), C.int(nox_win_height-1))
		nox_client_clearScreen_440900()
		*memmap.PtrUint32(0x587000, 85744) = 0
	}
}

func nox_xxx_drawAllMB_475810_draw(vp *Viewport) {
	C.nox_xxx_drawAllMB_475810_draw_A(vp.C())
	if int32(vp.field_12) < 0 {
		vp.field_12 = C.int(-1 - int32(vp.field_12))
	} else if int32(vp.field_12) > 0 {
		vp.field_12 = C.int(1 - int32(vp.field_12))
	}
	*memmap.PtrInt32(0x5D4594, 1096428) = int32(C.int(vp.field_4) - vp.x1)
	C.dword_5d4594_1096432 = C.uint(C.int(vp.field_5) - vp.y1)
	xmin := int(vp.field_4) / common.GridStep
	ymin := int(vp.field_5) / common.GridStep
	nox_wallsYyy = nox_wallsYyy[:0]
	C.nox_xxx_drawBlack_496150(vp.C())
	disableDraw := false
	if !noxflags.HasGame(noxflags.GameModeCoop) && (C.nox_xxx_testCD_413830() == 0) ||
		noxflags.HasGame(noxflags.GameModeCoop) && (C.nox_xxx_testCDAndSolo_413840() == 0) ||
		asDrawable((*C.nox_drawable)(*memmap.PtrPtr(0x852978, 8))).CheckFlag31(2) ||
		C.nox_gameDisableMapDraw_5d4594_2650672 != 0 {
		disableDraw = true
	}
	if C.nox_client_gui_flag_1556112 != 0 || disableDraw {
		noxrend.SelectColor(memmap.Uint32(0x85B3FC, 952))
		nox_client_clearScreen_440900()
		C.sub_437290()
		C.dword_5d4594_3799524 = 1
		return
	}
	if memmap.Uint32(0x5D4594, 1096520) != 0 {
		noxrend.SelectColor(memmap.Uint32(0x5D4594, 2523948))
		nox_client_clearScreen_440900()
		noxrend.SelectColor(memmap.Uint32(0x85B3FC, 952))
		*memmap.PtrUint32(0x5D4594, 1096520) = 0
		C.sub_437290()
		C.dword_5d4594_3799524 = 1
		return
	}
	sub_468F80(vp)
	v10 := C.nox_xxx_drawAllMB_475810_draw_B(vp.C()) != 0
	C.sub_4765F0(vp.C())
	sub_4754F0(vp)
	if v10 {
		nox_xxx_tileDrawMB_481C20(vp)
		C.sub_4C5500(vp.C())
	} else {
		nox_client_clearScreen_440900()
	}
	sub_475F10(vp)
	nox_client_queueWallsDraw(vp, xmin, ymin)
	nox_client_drawBackWalls(vp)
	sub_475FE0(vp)
	C.nox_video_drawCursorSelectCircle_4773C0(vp.C())
	nox_xxx_drawAllMB_475810_draw_E(vp)
	C.sub_4AFD40()
	C.sub_4C5060(vp.C())
	nox_client_maybeDrawFrontWalls(vp)
	C.nox_client_procFade_44D9F0(0)
	if getEngineFlag(NOX_ENGINE_FLAG_ENABLE_SHOW_AI) {
		C.sub_476270(vp.C())
	}
	C.sub_45AB40()
	C.sub_437290()
	*memmap.PtrUint32(0x973F18, 68) = 1
	C.sub_476680()
}

type NoxLight struct {
	R, G, B int
}

const (
	// TODO: the 4x factor is for high-res; figure out what those values are
	lightGridW = 57 * 4
	lightGridH = 45 * 4
	lightGrid  = common.GridStep
)

var (
	nox_arr2_853BC0 [lightGridW][lightGridH]NoxLight
	lightsOutBuf    []uint32
)

func nox_xxx_get_57AF20() int {
	return int(C.dword_5d4594_2523804)
}

func sub_468F80(vp *Viewport) {
	// TODO: values here are similar to lightGridW and lightGridH
	C.dword_5d4594_2650676 = C.uint(46*((int(vp.field_4)+11)/46-1) - 11)
	C.dword_5d4594_2650680 = C.uint(46*((int(vp.field_5)+11)/46) - 57)
	if getEngineFlag(NOX_ENGINE_DISABLE_SOFT_LIGHTS) {
		for i := 0; i < lightGridW; i++ {
			for j := 0; j < lightGridH; j++ {
				nox_arr2_853BC0[i][j].R = 255 << 16
				nox_arr2_853BC0[i][j].G = 255 << 16
				nox_arr2_853BC0[i][j].B = 255 << 16
			}
		}
	} else {
		cl := noxrend.GetLightColor()
		if nox_xxx_get_57AF20() != 0 {
			cl = NoxLight{R: 50, G: 50, B: 50}
		}
		for i := 0; i < lightGridW; i++ {
			for j := 0; j < lightGridH; j++ {
				nox_arr2_853BC0[i][j].R = cl.R << 16
				nox_arr2_853BC0[i][j].G = cl.G << 16
				nox_arr2_853BC0[i][j].B = cl.B << 16
			}
		}
		const add = 100
		rect := image.Rect(
			int(vp.field_4)-add,
			int(vp.field_5)-add,
			int(vp.width)+int(vp.field_4)+add,
			int(vp.height)+int(vp.field_5)+add,
		)
		nox_xxx_forEachSprite(rect, func(dr *Drawable) {
			C.nox_xxx_cliLight16_469140(dr.C(), vp.C())
		})
	}
}

//export sub_469920
func sub_469920(p *C.nox_point) *C.char {
	if getEngineFlag(NOX_ENGINE_DISABLE_SOFT_LIGHTS) {
		return (*C.char)(unsafe.Pointer(&lightsOutBuf[0]))
	}

	x := int(int32(p.x) - int32(uint32(C.dword_5d4594_2650676)))
	y := int(int32(p.y) - int32(uint32(C.dword_5d4594_2650680)))

	xd := x / lightGrid
	yd := y / lightGrid
	if xd < 0 || yd < 0 || xd+1 >= lightGridW || yd+1 >= lightGridH {
		return (*C.char)(unsafe.Pointer(&lightsOutBuf[0]))
	}

	xr := x % lightGrid
	yr := y % lightGrid

	c00 := nox_arr2_853BC0[xd+0][yd+0]
	c10 := nox_arr2_853BC0[xd+1][yd+0]
	c01 := nox_arr2_853BC0[xd+0][yd+1]
	c11 := nox_arr2_853BC0[xd+1][yd+1]

	var cr1 NoxLight
	cr1.R = c00.R + xr*(c10.R-c00.R)/lightGrid
	cr1.G = c00.G + xr*(c10.G-c00.G)/lightGrid
	cr1.B = c00.B + xr*(c10.B-c00.B)/lightGrid

	var cr2 NoxLight
	cr2.R = c01.R + xr*(c11.R-c01.R)/lightGrid
	cr2.G = c01.G + xr*(c11.G-c01.G)/lightGrid
	cr2.B = c01.B + xr*(c11.B-c01.B)/lightGrid

	var res NoxLight
	res.R = cr1.R + yr*(cr2.R-cr1.R)/lightGrid
	res.G = cr1.G + yr*(cr2.G-cr1.G)/lightGrid
	res.B = cr1.B + yr*(cr2.B-cr1.B)/lightGrid

	dst := lightsOutBuf[3:]
	dst[0] = uint32(res.R >> 16)
	dst[1] = uint32(res.G >> 16)
	dst[2] = uint32(res.B >> 16)
	return (*C.char)(unsafe.Pointer(&dst[0]))
}

//export sub_4814F0
func sub_4814F0(a1 *C.nox_point) C.int {
	c1 := nox_arr2_853BC0[a1.x][a1.y+0]
	v2 := c1.R >> 8
	v3 := c1.G >> 8
	v4 := c1.B >> 8

	c2 := nox_arr2_853BC0[a1.x][a1.y+1]
	v2b := c2.R >> 8
	v3b := c2.G >> 8
	v4b := c2.B >> 8

	v8 := int(memmap.Int32(0x8529A0, uintptr(1020+4*((v2b-v2)>>8))))
	v9 := int(memmap.Int32(0x8529A0, uintptr(1020+4*((v3b-v3)>>8))))
	v10 := int(memmap.Int32(0x8529A0, uintptr(1020+4*((v4b-v4)>>8))))

	for i := 0; i < lightGrid; i++ {
		v6 := unsafe.Slice((*uint32)(unsafe.Pointer(&C.nox_arr_84EB20[280*int(a1.x)+12*i+4])), 3)
		v6[0] = uint32(int32(v2))
		v6[1] = uint32(int32(v3))
		v6[2] = uint32(int32(v4))
		v2 += v8
		v3 += v9
		v4 += v10
	}

	v6p := (*int32)(unsafe.Pointer(&C.nox_arr_84EB20[280*a1.x]))
	*v6p = int32(a1.y)
	return 7 * a1.x
}

func sub_4C1C60(a1, a2 int) int {
	return int((int64(a1) * int64(a2)) >> 16)
}

//export sub_4695E0
func sub_4695E0(a1, a2 C.int, a3p *C.int, a4, a5 C.int) {
	v5 := int(a4)
	if a5 != 0 {
		v5 = -int(a4)
	}
	a3 := unsafe.Slice((*int32)(unsafe.Pointer(a3p)), 3)
	v6 := sub_4C1C60(v5, int(a3[0])) << 8
	v7 := sub_4C1C60(v5, int(a3[1])) << 8
	v8 := sub_4C1C60(v5, int(a3[2])) << 8

	ptr := &nox_arr2_853BC0[a1][a2]

	var res NoxLight
	res.R = v6 + ptr.R
	res.G = v7 + ptr.G
	res.B = v8 + ptr.B

	if res.R <= 0xFF0000 {
		if res.R < 0 {
			res.R = 0
		}
	} else {
		res.R = 0xFF0000
	}
	if res.G <= 0xFF0000 {
		if res.G < 0 {
			res.G = 0
		}
	} else {
		res.G = 0xFF0000
	}
	if res.B <= 0xFF0000 {
		if res.B < 0 {
			res.B = 0
		}
	} else {
		res.B = 0xFF0000
	}
	*ptr = res
}

func nox_xxx_tileDrawMB_481C20(vp *Viewport) {
	C.nox_xxx_waypointCounterMB_587000_154948++
	dx := int(vp.field_4) - int(vp.x1)
	dy := int(vp.field_5) - int(vp.y1)
	if C.nox_client_texturedFloors_154956 == 0 && C.dword_5d4594_1193156 == 1 {
		C.nox_client_texturedFloors2_154960 = 0
		C.nox_client_texturedFloors_154956 = 1
		C.nox_xxx_tileSetDrawFn_481420()
	}
	if C.dword_5d4594_1193156 == 1 && C.nox_client_texturedFloors2_154960 == 0 && gameFrame()%30 == 0 && C.nox_xxx_tileCheckRedrawMB_482570(vp.C()) == 0 {
		C.dword_5d4594_1193156 = 0
		C.nox_client_texturedFloors_154956 = C.nox_client_texturedFloors2_154960
		C.nox_xxx_tileSetDrawFn_481420()
	}
	if C.dword_5d4594_1193188 != 0 {
		C.nox_xxx_tileDrawImpl_4826A0(vp.C())
		C.dword_5d4594_1193188 = 0
	} else {
		C.nox_xxx_tileDrawMB_481C20_A(vp.C(), C.int(dx))
		C.nox_xxx_tileDrawMB_481C20_B(vp.C(), C.int(dy))
	}
	if C.nox_client_texturedFloors_154956 != 0 {
		C.nox_xxx_tileDrawMB_481C20_C_textured(vp.C(), C.int(dx), C.int(dy))
	} else {
		C.nox_xxx_tileDrawMB_481C20_C_solid(vp.C(), C.int(dx), C.int(dy))
	}
}

func sub_4754F0(vp *Viewport) {
	rect := image.Rect(
		int(vp.field_4), int(vp.field_5),
		int(vp.field_4+vp.width), int(vp.field_5+vp.height)+128,
	)
	nox_drawable_list_1 = nox_drawable_list_1[:0]
	nox_drawable_list_3 = nox_drawable_list_3[:0]
	nox_drawable_list_2 = nox_drawable_list_2[:0]
	nox_drawable_list_4 = nox_drawable_list_4[:0]
	nox_backWalls = nox_backWalls[:0]
	nox_frontWalls = nox_frontWalls[:0]
	nox_xxx_forEachSprite(rect, nox_xxx_spriteAddQueue_475560_draw)
}

func nox_xxx_spriteAddQueue_475560_draw(dr *Drawable) {
	if C.nox_xxx_sprite_4756E0_drawable(dr.C()) != 0 {
		nox_drawable_list_2 = append(nox_drawable_list_2, dr)
		return
	}
	if C.nox_xxx_sprite_475740_drawable(dr.C()) != 0 {
		nox_drawable_list_3 = append(nox_drawable_list_3, dr)
		return
	}
	if C.nox_xxx_sprite_4757A0_drawable(dr.C()) != 0 {
		nox_drawable_list_4 = append(nox_drawable_list_4, dr)
		return
	}
	if C.sub_4757D0_drawable(dr.C()) != 0 {
		if C.nox_client_fadeObjects_80836 != 0 || unsafe.Pointer(dr.C()) == *memmap.PtrPtr(0x852978, 8) || C.nox_xxx_client_4984B0_drawable(dr.C()) != 0 {
			if dr.field_122 == 0 {
				if C.nox_xxx_client_4984B0_drawable(dr.C()) != 0 {
					dr.field_121 = 1
					dr.field_120 = 0
				} else {
					if dr.field_121 == 0 || (dr.flags28&0x6 != 0) {
						return
					}
					dr.field_120 = 1
				}
			}
			if (dr.field_120 != 0 || dr.field_122 != 0) && (gameFrame()-uint32(dr.field_85)) > gameFPS() {
				dr.field_120 = 0
			} else {
				nox_drawable_list_1 = append(nox_drawable_list_1, dr)
			}
		}
	}
}

const (
	nox_drawable_list_1_cap = 8192
	nox_drawable_lists_cap  = 512

	noxDrawableWallsCap = 256
)

var (
	nox_drawable_list_1 []*Drawable
	nox_drawable_list_3 []*Drawable
	nox_drawable_list_2 []*Drawable
	nox_drawable_list_4 []*Drawable

	nox_backWalls  []*Wall
	nox_frontWalls []*Wall
	nox_wallsYyy   []*Wall
)

func initDrawableLists() {
	nox_drawable_list_1 = make([]*Drawable, 0, nox_drawable_list_1_cap)
	nox_drawable_list_3 = make([]*Drawable, 0, nox_drawable_lists_cap)
	nox_drawable_list_2 = make([]*Drawable, 0, nox_drawable_lists_cap)
	nox_drawable_list_4 = make([]*Drawable, 0, nox_drawable_lists_cap)

	nox_backWalls = make([]*Wall, 0, noxDrawableWallsCap)
	nox_frontWalls = make([]*Wall, 0, noxDrawableWallsCap)
	nox_wallsYyy = make([]*Wall, 0, noxDrawableWallsCap)
}

func sub_473B30_free() {
	nox_drawable_list_1 = nil
	nox_drawable_list_3 = nil
	nox_drawable_list_2 = nil
	nox_drawable_list_4 = nil

	nox_backWalls = nil
	nox_frontWalls = nil
	nox_wallsYyy = nil
}

//export nox_xxx_drawList1096512_Append_4754C0
func nox_xxx_drawList1096512_Append_4754C0(p unsafe.Pointer) {
	nox_wallsYyy = append(nox_wallsYyy, asWall(p))
}

func nox_xxx_cliGetSpritePlayer_45A000() *Drawable {
	return asDrawable(C.nox_xxx_drawablePlayer_1046600)
}

func nox_client_queueWallsDraw(vp *Viewport, xmin, ymin int) { // nox_xxx_drawAllMB_475810_draw_C
	xmax := xmin + int(vp.width)/common.GridStep + 2
	ymax := ymin + int(vp.height)/common.GridStep + 4
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			wl := getWallAtGrid(types.Point{X: x, Y: y})
			if wl == nil {
				continue
			}
			if memmap.Uint8(0x85B3FC, 43076+12332*uintptr(wl.field1()))&4 == 0 {
				if wl.field4()&2 != 0 {
					nox_frontWalls = append(nox_frontWalls, wl)
				} else {
					nox_backWalls = append(nox_backWalls, wl)
				}
			}
		}
	}
}

func nox_client_drawBackWalls(vp *Viewport) {
	for _, v20 := range nox_backWalls {
		nox_xxx_drawWalls_473C10(vp, v20)
	}
	nox_backWalls = nox_backWalls[:0]
}

func sub_476160(a1, a2 *Drawable) bool {
	var v1 int
	if *(*int8)(a1.field(112)) >= 0 {
		v1 = a1.Pos().Y + int(*(*int16)(a1.field(104)))
	} else {
		v1 = int(C.sub_4761B0(a1.C()))
	}
	var v2 int
	if *(*int8)(a2.field(112)) >= 0 {
		v2 = a2.Pos().Y + int(*(*int16)(a2.field(104)))
	} else {
		v2 = int(C.sub_4761B0(a2.C()))
	}
	return v1-v2 < 0
}

func nox_xxx_drawAllMB_475810_draw_E(vp *Viewport) {
	sort.Slice(nox_drawable_list_1, func(i, j int) bool {
		a, b := nox_drawable_list_1[i], nox_drawable_list_1[j]
		return sub_476160(a, b)
	})
	sort.Slice(nox_wallsYyy, func(i, j int) bool {
		a, b := nox_wallsYyy[i], nox_wallsYyy[j]
		return C.sub_476080((*C.uchar)(a.C())) < C.sub_476080((*C.uchar)(b.C()))
	})
	arr1 := nox_drawable_list_1
	arr2 := nox_wallsYyy
	v41 := 0x7FFFFFFF
	if len(arr1) > 0 {
		v41 = arr1[0].Pos().Y
	}
	v21 := 0x7FFFFFFF
	if len(arr2) > 0 {
		v21 = int(C.sub_476080((*C.uchar)(arr2[0].C())))
	}
LOOP:
	for len(arr1) > 0 || len(arr2) > 0 {
		if v41 >= v21 {
			if len(arr2) != 0 {
				nox_xxx_drawWalls_473C10(vp, arr2[0])
				arr2 = arr2[1:]
				if len(arr2) != 0 {
					v21 = int(C.sub_476080((*C.uchar)(arr2[0].C())))
					continue
				}
			}
			v21 = 0x7FFFFFFF
			continue
		}
		if len(arr1) == 0 {
			v41 = 0x7FFFFFFF
			continue
		}
		dr := arr1[0]
		arr1 = arr1[1:]
		//nox_drawable_list_1_size--
		if len(arr1) > 0 {
			v41 = arr1[0].Pos().Y
		} else {
			v41 = 0x7FFFFFFF
		}
		if uint32(dr.field_27) == memmap.Uint32(0x5D4594, 1096448) && C.nox_xxx_TeamGet_418B10() != nil {
			for v25 := nox_xxx_cliGetSpritePlayer_45A000(); v25 != nil; v25 = v25.Field104() {
				if v25.CheckFlag31(30) {
					continue LOOP
				}
			}
		}
		v26 := dr.Flags28()
		if !((v26&6 == 0) || gameFrame()-uint32(dr.field_72) <= 5) {
			if v26&2 != 0 {
				v27 := dr.field_69
				if !(v27 == 9 || v27 == 10) {
					continue
				}
			} else if *memmap.PtrPtr(0x852978, 8) != unsafe.Pointer(dr.C()) {
				continue
			}
		}
		drawCreatureBackEffects(noxrend, vp, dr)
		if dr.DrawFunc(vp) == 0 {
			continue
		}
		if getEngineFlag(NOX_ENGINE_FLAG_ENABLE_SHOW_EXTENTS) {
			C.nox_thing_debug_draw(vp.C(), dr.C())
		}
		dr.field_33 = 0
		if dr.Flags70()&0x40 != 0 {
			C.nox_xxx_drawShinySpot_4C4F40(vp.C(), dr.C())
		}
		drawCreatureFrontEffects(noxrend, vp, dr)
		C.sub_495BB0(dr.C(), vp.C())
		if dr.field_120 == 0 && dr.field_122 == 0 {
			dr.field_85 = C.uint(gameFrame())
		}
		if C.sub_459DB0(dr.C()) != 0 {
			C.sub_459DD0(dr.C(), 1)
		}
		if dr.Flags28()&0x20006 != 0 {
			C.sub_49A6A0(vp.C(), dr.C())
		}
	}
	nox_drawable_list_1 = nox_drawable_list_1[:0]
}

func sub_475FE0(vp *Viewport) {
	for _, dr := range nox_drawable_list_4 {
		if C.nox_xxx_client_4984B0_drawable(dr.C()) != 0 {
			dr.field_121 = 1
			dr.DrawFunc(vp)
			if getEngineFlag(NOX_ENGINE_FLAG_ENABLE_SHOW_EXTENTS) {
				C.nox_thing_debug_draw(vp.C(), dr.C())
			}
			dr.field_33 = 0
			if dr.field_120 == 0 && dr.field_122 == 0 {
				dr.field_85 = C.uint(gameFrame())
			}
		}
	}
	nox_drawable_list_4 = nox_drawable_list_4[:0]
}

func sub_475F10(vp *Viewport) {
	for _, dr := range nox_drawable_list_3 {
		drawCreatureBackEffects(noxrend, vp, dr)
		if C.nox_xxx_client_4984B0_drawable(dr.C()) != 0 {
			dr.field_121 = 1
			dr.DrawFunc(vp)
			if dr.Flags70()&0x40 != 0 {
				C.nox_xxx_drawShinySpot_4C4F40(vp.C(), dr.C())
			}
			drawCreatureFrontEffects(noxrend, vp, dr)
			C.sub_495BB0(dr.C(), vp.C())
			if getEngineFlag(NOX_ENGINE_FLAG_ENABLE_SHOW_EXTENTS) {
				C.nox_thing_debug_draw(vp.C(), dr.C())
			}
			dr.field_33 = 0
			if dr.field_120 == 0 && dr.field_122 == 0 {
				dr.field_85 = C.uint(gameFrame())
			}
		}
	}
	nox_drawable_list_3 = nox_drawable_list_3[:0]
}

func nox_xxx_drawWalls_473C10(vp *Viewport, p *Wall) {
	C.nox_xxx_drawWalls_473C10(vp.C(), p.C())
}

func nox_client_maybeDrawFrontWalls(vp *Viewport) { // nox_client_maybeDrawFrontWalls_475810_F
	if C.nox_client_drawFrontWalls_80812 != 0 {
		for _, wl := range nox_frontWalls {
			nox_xxx_drawWalls_473C10(vp, wl)
		}
	} else {
		for _, wl := range nox_frontWalls {
			*(*byte)(wl.field(3)) = 0
			*(*byte)(wl.field(4)) &= 0xFC
		}
	}
	nox_frontWalls = nox_frontWalls[:0]
}

//export sub_4745F0
func sub_4745F0(cvp *C.nox_draw_viewport_t) {
	vp := asViewport(cvp)
	for _, dr := range nox_drawable_list_2 {
		drawCreatureBackEffects(noxrend, vp, dr)
		if C.nox_xxx_client_4984B0_drawable(dr.C()) == 0 {
			continue
		}
		dr.field_121 = 1
		C.sub_476AE0(vp.C(), dr.C())
		if dr.Flags70()&0x40 != 0 {
			C.nox_xxx_drawShinySpot_4C4F40(vp.C(), dr.C())
		}
		drawCreatureFrontEffects(noxrend, vp, dr)
		C.sub_495BB0(dr.C(), vp.C())
		if getEngineFlag(NOX_ENGINE_FLAG_ENABLE_SHOW_EXTENTS) {
			C.nox_thing_debug_draw(vp.C(), dr.C())
		}
		dr.field_33 = 0
		if dr.field_120 == 0 && dr.field_122 == 0 {
			dr.field_85 = C.uint(gameFrame())
		}
	}
	nox_drawable_list_2 = nox_drawable_list_2[:0]
}

//export nox_video_drawCircleColored_4C3270
func nox_video_drawCircleColored_4C3270(a1, a2, a3, a4 C.int) {
	noxrend.DrawCircleColored(int(a1), int(a2), int(a3), uint32(a4))
}

//export nox_video_drawCircle_4B0B90
func nox_video_drawCircle_4B0B90(a1, a2, a3 C.int) {
	noxrend.DrawCircle(int(a1), int(a2), int(a3))
}

func (r *NoxRender) DrawImageAt(img *Image, pos types.Point) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("panic drawing image %v: %v", img, r))
		}
	}()
	if C.dword_5d4594_3799452 != 0 {
		C.nox_xxx_wndDraw_49F7F0()
		C.sub_49F780(C.int(memmap.Int32(0x973F18, 52)), C.int(memmap.Int32(0x973F18, 12)))
		r.p.flag_0 = 1
	}
	r.drawImage16(img, pos)
	if C.dword_5d4594_3799452 != 0 {
		C.sub_49F860()
		C.dword_5d4594_3799452 = 0
	}
	C.dword_5d4594_3799484 = 0
	*memmap.PtrUint32(0x973F18, 120) = 0
	if memmap.Uint32(0x973F18, 68) != 0 && img != nil {
		if memmap.Uint32(0x973F18, 92) != uint32(img.field_1_0) || memmap.Uint32(0x973F18, 84) != uint32(img.field_1_1) {
			C.dword_5d4594_3799524 = 1
		}
		img.field_1_0 = uint16(memmap.Uint32(0x973F18, 92))
		img.field_1_1 = uint16(memmap.Uint32(0x973F18, 84))
	}
}

//export nox_client_drawImageAt_47D2C0
func nox_client_drawImageAt_47D2C0(img *C.nox_video_bag_image_t, x, y C.int) {
	noxrend.DrawImageAt(asImage(img), types.Point{X: int(x), Y: int(y)})
}

func sub_47D200() {
	if C.dword_5d4594_3801780 == 0 {
		panic("TODO")
		//nox_client_drawAtFunc_3799500 = sub_4C5EB0
	}
	C.dword_5d4594_3799484 = 0
	*memmap.PtrUint32(0x973F18, 52) = 0
	*memmap.PtrUint32(0x973F18, 12) = 0
	C.dword_5d4594_3799476 = 0
	C.dword_5d4594_3799452 = 0
	*memmap.PtrUint32(0x973F18, 44) = 0
	*memmap.PtrUint32(0x973F18, 104) = 0
	*memmap.PtrUint32(0x973F18, 8) = 0
	*memmap.PtrUint32(0x973F18, 36) = 0
	*memmap.PtrUint32(0x973F18, 32) = 0
	*memmap.PtrUint32(0x973F18, 100) = 0
	*memmap.PtrUint32(0x973F18, 20) = 0
	*memmap.PtrUint32(0x973F18, 120) = 0
	*memmap.PtrUint32(0x973F18, 68) = 0
	C.dword_5d4594_3799524 = 0
	C.dword_5d4594_3799468 = 0
	C.dword_5d4594_3799552 = 0
	C.dword_5d4594_3799508 = 0
}

func (r *NoxRender) drawImage16(img *Image, pos types.Point) { // nox_client_xxxDraw16_4C7440
	if img == nil {
		return
	}
	switch img.Type() & 0x3F {
	case 2, 7:
		r.draw27 = pixCopyN
		r.nox_client_drawImg_bbb_4C7860(img, pos)
	case 3, 4, 5, 6:
		r.draw5 = r.sub_4C96A0
		r.draw6 = func(dst []uint16, src []byte, op byte, val int) (_ []uint16, _ []byte) { return dst, src }
		if !r.IsAlphaEnabled() {
			if r.p.field_14 != 0 {
				r.draw5 = r.sub_4C9970
				r.draw27 = r.sub_4C86B0
				r.draw4 = r.sub_4C91C0
			} else {
				r.draw27 = r.sub_4C8D60
				if r.p.field_17 == 0 {
					r.draw27 = pixCopyN
				}
				r.draw4 = r.sub_4C8DF0
			}
		} else {
			r.draw5 = r.sub_4C97F0
			if r.p.field_14 != 0 {
				v3 := r.p.field_259
				if v3 == 255 {
					if r.p.field_16 == 0 {
						r.draw27 = r.sub_4C86B0
						r.draw4 = r.sub_4C91C0
					} else {
						r.draw27 = pixCopyN
						r.draw4 = r.sub_4C8DF0
					}
				} else if v3 == 128 {
					r.draw27 = r.pixBlend
					r.draw4 = r.sub_4C94D0
				} else {
					r.draw27 = r.sub_4C8850
					r.draw4 = r.sub_4C92F0
				}
			} else {
				v4 := r.p.field_259
				if v4 == 255 {
					r.draw27 = pixCopyN
					r.draw4 = r.sub_4C8DF0
				} else if v4 == 128 {
					r.draw27 = r.sub_4C8410
					r.draw4 = r.sub_4C9050
				} else {
					r.draw27 = r.sub_4C8130
					r.draw4 = r.sub_4C8EC0
				}
			}
		}
		r.nox_client_drawImg_aaa_4C79F0(img, pos)
	case 8:
		r.draw27 = r.pixBlendPremult
		r.nox_client_drawImg_aaa_4C79F0(img, pos)
	}
}

func pixCopyN(dst []uint16, src []byte, _ byte, n int) (_ []uint16, _ []byte) { // sub_4C80E0
	if n < 0 {
		panic("negative size")
	}
	copy16b(dst[:n], src[:n*2])
	return dst[n:], src[n*2:]
}

func drawOpC(fnc func()) drawOpFunc {
	return func(dst []uint16, src []byte, op byte, val int) (_ []uint16, _ []byte) {
		var pdst, psrc uintptr
		if len(dst) == 0 {
			C.nox_draw_sprite_dstPtr_3799540 = nil
		} else {
			C.nox_draw_sprite_dstPtr_3799540 = unsafe.Pointer(&dst[0])
			pdst = uintptr(unsafe.Pointer(&dst[0]))
		}
		if len(src) == 0 {
			C.nox_video_cur_pixdata_3799444 = nil
		} else {
			C.nox_video_cur_pixdata_3799444 = (*C.uchar)(unsafe.Pointer(&src[0]))
			psrc = uintptr(unsafe.Pointer(&src[0]))
		}

		*memmap.PtrUint32(0x973F18, 0) = uint32(op)
		*memmap.PtrInt32(0x973F18, 28) = int32(val)
		fnc()

		if p := uintptr(unsafe.Pointer(C.nox_draw_sprite_dstPtr_3799540)); p > pdst {
			dst = dst[(p-pdst)/2:]
		}
		if p := uintptr(unsafe.Pointer(C.nox_video_cur_pixdata_3799444)); p > psrc {
			src = src[p-psrc:]
		}
		return dst, src
	}
}

//export nox_video_getImagePixdata_func
func nox_video_getImagePixdata_func(img *C.nox_video_bag_image_t) unsafe.Pointer {
	data := asImage(img).Pixdata()
	if len(data) == 0 {
		return nil
	}
	return unsafe.Pointer(&data[0])
}

//export nox_video_getImagePixdata_42FB30
func nox_video_getImagePixdata_42FB30(img *C.nox_video_bag_image_t) unsafe.Pointer {
	data := asImage(img).Pixdata()
	if len(data) == 0 {
		return nil
	}
	return unsafe.Pointer(&data[0])
}

func skipPixdata(pix []byte, width int, skip int) []byte {
	for i := 0; i < skip; i++ {
		val := 0
		for j := 0; j < width; j += val {
			op := pix[0]
			val = int(pix[1])
			pix = pix[2:]
			switch op & 0xF {
			case 2, 5, 6, 7:
				pix = pix[2*val:]
				break
			case 4:
				pix = pix[val:]
			}
		}
	}
	return pix
}

func (r *NoxRender) nox_client_drawImg_aaa_4C79F0(img *Image, pos types.Point) { // nox_client_drawImg_aaa_4C79F0
	src := img.Pixdata()
	if len(src) == 0 {
		return
	}

	width := binary.LittleEndian.Uint32(src[0:])
	height := binary.LittleEndian.Uint32(src[4:])
	src = src[8:]

	offX := int32(binary.LittleEndian.Uint32(src[0:]))
	offY := int32(binary.LittleEndian.Uint32(src[4:]))
	src = src[8:]
	pos.X += int(offX)
	pos.Y += int(offY)

	src = src[1:] // unused

	if C.dword_5d4594_3799484 != 0 {
		height -= uint32(C.dword_5d4594_3799484)
		if height <= 0 {
			return
		}
		C.dword_5d4594_3799476 = C.uint(int(height) + pos.Y)
	}
	*memmap.PtrInt32(0x973F18, 92) = int32(pos.X)
	*memmap.PtrInt32(0x973F18, 84) = int32(pos.Y)
	*memmap.PtrUint32(0x973F18, 88) = width
	*memmap.PtrUint32(0x973F18, 76) = height
	if r.p.flag_0 != 0 {
		rc := types.Rect{Left: pos.X, Top: pos.Y, Right: pos.X + int(width), Bottom: pos.Y + int(height)}
		a1a, ok := nox_xxx_utilRect_49F930(rc, r.Rect())
		if !ok {
			return
		}
		if rc != a1a {
			r.nox_client_drawXxx_4C7C80(src, pos, int(width), a1a)
			return
		}
	}
	C.dword_5d4594_3799508 ^= C.uint(pos.Y & 1)
	pixbuf := r.PixBuffer()
	pitch := pixbuf.Stride
	for i := 0; i < int(height); i++ {
		dst := pixbuf.Pix[pitch*(pos.Y+i)+pos.X:]
		if C.dword_5d4594_3799552 != 0 {
			C.dword_5d4594_3799508 ^= 1
			if C.dword_5d4594_3799508 != 0 {
				if i != 0 {
					copy(dst[:width], pixbuf.Pix[pitch*(pos.Y+i-1)+pos.X:])
				}
				src = skipPixdata(src, int(width), 1)
				continue
			}
		}
		var val int
		for j := 0; j < int(width); j += val {
			op := src[0]
			val = int(src[1])
			src = src[2:]

			if op&0xF == 1 {
				dst = dst[val:]
				continue
			}
			switch op & 0xF {
			case 2, 7:
				dst, src = r.draw27(dst, src, op, val)
			case 4:
				dst, src = r.draw4(dst, src, op, val)
			case 5:
				dst, src = r.draw5(dst, src, op, val)
			case 6:
				dst, src = r.draw6(dst, src, op, val)
			default:
				panic(fmt.Errorf("invalid draw op: 0x%x, (%d,%d)", op, i, j))
			}
		}
	}
}

func (r *NoxRender) nox_client_drawXxx_4C7C80(pix []byte, pos types.Point, width int, clip types.Rect) { // nox_client_drawXxx_4C7C80
	left := clip.Left
	right := clip.Right
	dy := clip.Top - pos.Y
	height := clip.Bottom - clip.Top
	if C.dword_5d4594_3799484 != 0 {
		height -= int(C.dword_5d4594_3799484)
		if height <= 0 {
			return
		}
		C.dword_5d4594_3799476 = C.uint(height + pos.Y)
	}
	if height == 0 {
		return
	}
	ys := pos.Y
	if dy != 0 {
		ys += dy
		pix = skipPixdata(pix, width, dy)
	}
	C.dword_5d4594_3799508 ^= C.uint(ys & 1)
	pixbuf := r.PixBuffer()
	pitch := pixbuf.Stride
	for i := 0; i < height; i++ {
		yi := ys + i
		if C.dword_5d4594_3799552 != 0 {
			C.dword_5d4594_3799508 ^= 1
			if C.dword_5d4594_3799508 != 0 {
				if i != 0 {
					src := pixbuf.Pix[pitch*(yi-1)+left:]
					dst := pixbuf.Pix[pitch*(yi+0)+left:]
					w := right - left
					if w > width {
						w = width
					}
					copy(dst[:w], src[:w])
				}
				pix = skipPixdata(pix, width, 1)
				continue
			}
		}
		if width <= 0 {
			continue
		}
		row := pixbuf.Pix[pitch*yi : pitch*(yi+1)]
		var n int
		for j := 0; j < width; j += n {
			op := pix[0]
			n = int(pix[1]) // TODO: custom bag images fail here
			pix = pix[2:]

			if op&0xF == 1 {
				continue
			}

			var (
				fnc  drawOpFunc
				pmul int
			)
			switch op & 0xF {
			case 2, 7:
				fnc = r.draw27
				pmul = 2
			case 4:
				fnc = r.draw4
				pmul = 1
			case 5:
				fnc = r.draw5
				pmul = 2
			case 6:
				fnc = r.draw6
				pmul = 2
			default:
				panic(op)
			}
			xs := pos.X + j
			xe := xs + n
			xw := n
			if xe <= left || xs >= right {
				pix = pix[pmul*n:]
				continue
			}

			pix2 := pix
			if xs < left {
				d := left - xs
				xw -= d
				xs = left
				pix2 = pix2[pmul*d:]
			}
			row2 := row[xs:]
			if xe > right {
				d := xe - right
				xw -= d
			}
			_, _ = fnc(row2, pix2, op, xw)
			pix = pix[pmul*n:]
		}
	}
}

//export nox_video_drawImageAt2_4B0820
func nox_video_drawImageAt2_4B0820(a1 unsafe.Pointer, x, y C.int) {
	p := noxrend.asParticle(a1)
	p.DrawAt(types.Point{X: int(x), Y: int(y)})
}

func (r *NoxRender) nox_client_drawImg_bbb_4C7860(img *Image, pos types.Point) {
	data := img.Pixdata()
	if len(data) == 0 {
		return
	}
	width := binary.LittleEndian.Uint32(data[0:])
	height := binary.LittleEndian.Uint32(data[4:])
	data = data[8:]

	offX := int32(binary.LittleEndian.Uint32(data[0:]))
	offY := int32(binary.LittleEndian.Uint32(data[4:]))
	data = data[8:]
	pos.X += int(offX)
	pos.Y += int(offY)

	data = data[1:] // unused

	if C.dword_5d4594_3799484 != 0 {
		height -= uint32(C.dword_5d4594_3799484)
		if height <= 0 {
			return
		}
		C.dword_5d4594_3799476 = C.uint(pos.Y + int(height))
	}

	wsz := int(width)
	if r.p.flag_0 != 0 {
		rc := types.Rect{
			Left:   pos.X,
			Top:    pos.Y,
			Right:  pos.X + int(width),
			Bottom: pos.Y + int(height),
		}
		a1a, ok := nox_xxx_utilRect_49F930(rc, r.Rect())
		if !ok {
			return
		}
		v11 := a1a.Left - rc.Left
		v12 := a1a.Top - rc.Top
		wsz = a1a.Right - a1a.Left
		height = uint32(a1a.Bottom - a1a.Top)
		if a1a.Left != rc.Left || v12 != 0 {
			pos.X += v11
			data = data[int(width)*v12+2*v11:]
			pos.Y += v12
		}
	}
	xoff := pos.X
	ipitch := 2 * int(width)
	pixbuf := r.PixBuffer()
	pitch := pixbuf.Stride
	for i := 0; i < int(height); i++ {
		dst := pixbuf.Pix[pitch*(pos.Y+1)+xoff:]
		src := data[ipitch*i:]
		copy16b(dst[:wsz], src[:wsz*2])
	}
}

func copy16b(dst []uint16, src []byte) int {
	n16 := len(src) / 2
	src = src[:n16*2]
	src16 := unsafe.Slice((*uint16)(unsafe.Pointer(&src[0])), n16)
	return copy(dst, src16)
}

func (r *NoxRender) pixBlend(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C8A30
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]
	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	rmul := uint16(byte(r.p.field_26))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_24))

	for i := 0; i < sz; i++ {
		c1 := dst[i]                          // old color
		c2 := binary.LittleEndian.Uint16(src) // color to draw

		rc := byte((rmul * ((rmask & c2) << rshift)) >> 8)
		gc := byte((gmul * ((gmask & c2) >> gshift)) >> 8)
		bc := byte((bmul * ((bmask & c2) >> bshift)) >> 8)

		c3 := r.colors.R[byte(int16(bc)+(int16((bmask&c1)>>bshift)-int16(bc))/2)] |
			r.colors.G[byte(int16(gc)+(int16((gmask&c1)>>gshift)-int16(gc))/2)] |
			r.colors.B[byte(int16(rc)+(int16((rmask&c1)<<rshift)-int16(rc))/2)]

		dst[i] = c3
		src = src[2:]
	}
	return dst[sz:], src
}

func (r *NoxRender) sub_4C96A0(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C96A0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]
	const (
		rshift = 7 // -10+3
		gshift = 2 // -5+3
		bshift = 3 // -0+3

		rmask = 0x7c00
		gmask = 0x03e0
		bmask = 0x001f
	)
	for i := 0; i < sz; i++ {
		c1 := dst[i]                          // old color
		c2 := binary.LittleEndian.Uint16(src) // color to draw
		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift

		c2v := (c2 << 4) & 0xFF
		cr = r.colors.R[byte(cr+((c2v*(((c2>>8)&0xF0)-cr))>>8))]
		cg = r.colors.G[byte(cg+((c2v*(((c2>>4)&0xF0)-cg))>>8))]
		cb = r.colors.B[byte(cb+((c2v*(((c2>>0)&0xF0)-cb))>>8))]
		dst[i] = cr | cg | cb

		src = src[2:]
	}
	return dst[sz:], src
}

func (r *NoxRender) sub_4C9970(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C9970
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]
	const (
		rshift = 7 // -10+3
		gshift = 2 // -5+3
		bshift = 3 // -0+3

		rmask = 0x7c00
		gmask = 0x03e0
		bmask = 0x001f
	)

	rmul := uint16(byte(r.p.field_24))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_26))

	for i := 0; i < sz; i++ {
		c1 := dst[i]                          // old color
		c2 := binary.LittleEndian.Uint16(src) // color to draw

		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift

		c2v := (c2 << 4) & 0xFF
		cr = r.colors.R[byte(cr+((c2v*(((rmul*((c2>>8)&0xF0))>>8)-cr))>>8))]
		cg = r.colors.G[byte(cg+((c2v*(((gmul*((c2>>4)&0xF0))>>8)-cg))>>8))]
		cb = r.colors.B[byte(cb+((c2v*(((bmul*(c2&0xF0))>>8)-cb))>>8))]
		dst[i] = cr | cg | cb

		src = src[2:]
	}
	return dst[sz:], src
}

func (r *NoxRender) sub_4C97F0(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C97F0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]

	const (
		rshift = 7 // -10+3
		gshift = 2 // -5+3
		bshift = 3 // -0+3

		rmask = 0x7c00
		gmask = 0x03e0
		bmask = 0x001f
	)

	v1 := uint32(r.p.field_259)

	for i := 0; i < sz; i++ {
		v4 := dst[i]                          // old color
		v3 := binary.LittleEndian.Uint16(src) // color to draw
		v5 := byte((v1 * uint32(v3&0x0F) << 4) >> 8)

		cr := (v4 & rmask) >> rshift
		cg := (v4 & gmask) >> gshift
		cb := (v4 & bmask) << bshift

		cr = r.colors.R[byte(cr+((uint16(v5)*((v3>>8)&0xF0-cr))>>8))]
		cg = r.colors.G[byte(cg+((uint16(v5)*((v3>>4)&0xF0-cg))>>8))]
		cb = r.colors.B[byte(cb+((uint16(v5)*((v3>>0)&0xF0-cb))>>8))]

		dst[i] = cr | cg | cb

		src = src[2:]
	}
	return dst[sz:], src
}

func (r *NoxRender) sub_4C86B0(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C86B0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]
	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	rmul := uint32(r.p.field_26)
	gmul := uint32(r.p.field_25)
	bmul := uint32(r.p.field_24)

	for i := 0; i < sz; i++ {
		v2 := binary.LittleEndian.Uint16(src)
		cr := r.colors.R[byte((bmul*(uint32(bmask&v2)>>bshift))>>8)]
		cg := r.colors.G[byte((gmul*(uint32(gmask&v2)>>gshift))>>8)]
		cb := r.colors.B[byte((rmul*(uint32(rmask&v2)<<rshift))>>8)]
		c := cr | cg | cb
		dst[i] = c
		src = src[2:]
	}
	return dst[sz:], src
}

func (r *NoxRender) sub_4C91C0(dst []uint16, src []byte, op byte, sz int) (_ []uint16, _ []byte) { // sub_4C91C0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[sz:]

	rmul := uint16(byte(r.p.field_24))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_26))

	v9 := r.field66(int(op >> 4))

	rpmul := v9[6]
	gpmul := v9[7]
	bpmul := v9[8]

	for i := 0; i < sz; i++ {
		c := src[i] // color to draw
		cr := r.colors.R[byte((rmul*uint16(((rpmul*uint32(c))>>8)&0xFF))>>8)]
		cg := r.colors.G[byte((gmul*uint16(((gpmul*uint32(c))>>8)&0xFF))>>8)]
		cb := r.colors.B[byte((bmul*uint16(((bpmul*uint32(c))>>8)&0xFF))>>8)]
		dst[i] = cr | cg | cb
	}
	return dst[sz:], src[sz:]
}

func (r *NoxRender) sub_4C8DF0(dst []uint16, src []byte, op byte, sz int) (_ []uint16, _ []byte) { // sub_4C8DF0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[sz:]

	v9 := r.field66(int(op >> 4))

	rpmul := v9[6]
	gpmul := v9[7]
	bpmul := v9[8]

	for i := 0; i < sz; i++ {
		c := src[i] // color to draw
		cr := r.colors.R[byte((rpmul*uint32(c))>>8)]
		cg := r.colors.G[byte((gpmul*uint32(c))>>8)]
		cb := r.colors.B[byte((bpmul*uint32(c))>>8)]
		dst[i] = cr | cg | cb
	}
	return dst[sz:], src[sz:]
}

func (r *NoxRender) sub_4C8410(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C8410
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]

	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	for i := 0; i < sz; i++ {
		c1 := dst[i]                          // old color
		c2 := binary.LittleEndian.Uint16(src) // color to draw

		cr1 := (c1 & rmask) >> rshift
		cg1 := (c1 & gmask) >> gshift
		cb1 := (c1 & bmask) << bshift

		cr2 := (c2 & rmask) >> rshift
		cg2 := (c2 & gmask) >> gshift
		cb2 := (c2 & bmask) << bshift

		cr := r.colors.R[byte(cr1+((cr2-cr1)>>1))]
		cg := r.colors.G[byte(cg1+((cg2-cg1)>>1))]
		cb := r.colors.B[byte(cb1+((cb2-cb1)>>1))]

		dst[i] = cr | cg | cb
	}

	return dst[sz:], src
}

func (r *NoxRender) sub_4C9050(dst []uint16, src []byte, op byte, sz int) (_ []uint16, _ []byte) { // sub_4C9050
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[sz:]

	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	v5 := r.field66(int(op >> 4))

	rpmul := v5[6]
	gpmul := v5[7]
	bpmul := v5[8]

	for i := 0; i < sz; i++ {
		c1 := dst[i] // old color
		c2 := src[i] // color to draw

		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift

		cr = r.colors.R[byte(cr+((uint16(((rpmul*uint32(c2))>>8)&0xFF)-cr)>>1))]
		cg = r.colors.G[byte(cg+((uint16(((gpmul*uint32(c2))>>8)&0xFF)-cg)>>1))]
		cb = r.colors.B[byte(cb+((uint16(((bpmul*uint32(c2))>>8)&0xFF)-cb)>>1))]

		dst[i] = cr | cg | cb
	}
	return dst[sz:], src[sz:]
}

func (r *NoxRender) sub_4C8130(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C8130
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]

	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	v1 := byte(r.p.field_259)

	for i := 0; i < sz; i++ {
		c1 := dst[i]                          // old color
		c2 := binary.LittleEndian.Uint16(src) // color to draw

		cr1 := (c1 & rmask) >> rshift
		cg1 := (c1 & gmask) >> gshift
		cb1 := (c1 & bmask) << bshift

		cr2 := (c2 & rmask) >> rshift
		cg2 := (c2 & gmask) >> gshift
		cb2 := (c2 & bmask) << bshift

		cr := r.colors.R[byte(cr1+((uint16(v1)*(cr2-cr1))>>8))]
		cg := r.colors.G[byte(cg1+((uint16(v1)*(cg2-cg1))>>8))]
		cb := r.colors.B[byte(cb1+((uint16(v1)*(cb2-cb1))>>8))]

		dst[i] = cr | cg | cb

		src = src[2:]
	}

	return dst[sz:], src
}

func (r *NoxRender) sub_4C8EC0(dst []uint16, src []byte, op byte, sz int) (_ []uint16, _ []byte) { // sub_4C8EC0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[sz:]

	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	v2 := byte(r.p.field_259)

	v7 := r.field66(int(op >> 4))

	rpmul := v7[6]
	gpmul := v7[7]
	bpmul := v7[8]

	for i := 0; i < sz; i++ {
		c1 := dst[i] // old color
		c2 := src[i] // color to draw

		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift

		cr = r.colors.R[byte(cr+((uint16(v2)*(((uint16(rpmul)*uint16(c2))>>8)&0xFF-cr))>>8))]
		cg = r.colors.G[byte(cg+((uint16(v2)*(((uint16(gpmul)*uint16(c2))>>8)&0xFF-cg))>>8))]
		cb = r.colors.B[byte(cb+((uint16(v2)*(((uint16(bpmul)*uint16(c2))>>8)&0xFF-cb))>>8))]

		dst[i] = cr | cg | cb
	}

	return dst[sz:], src[sz:]
}

func (r *NoxRender) sub_4C94D0(dst []uint16, src []byte, op byte, sz int) (_ []uint16, _ []byte) { // sub_4C94D0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[sz:]

	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	rmul := uint16(byte(r.p.field_24))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_26))

	v9 := r.field66(int(op >> 4))

	rpmul := v9[6]
	gpmul := v9[7]
	bpmul := v9[8]

	for i := 0; i < sz; i++ {
		c1 := dst[i] // old color
		c2 := src[i] // color to draw

		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift

		cr = r.colors.R[byte((rmul*((cr+((uint16((rpmul*uint32(c2))>>8)&0xFF-cr)>>1))&0xFF))>>8)]
		cg = r.colors.G[byte((gmul*((cg+((uint16((gpmul*uint32(c2))>>8)&0xFF-cg)>>1))&0xFF))>>8)]
		cb = r.colors.B[byte((bmul*((cb+((uint16((bpmul*uint32(c2))>>8)&0xFF-cb)>>1))&0xFF))>>8)]

		dst[i] = cr | cg | cb
	}

	return dst[sz:], src[sz:]
}

func (r *NoxRender) sub_4C8850(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C8850
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]

	const (
		rshift = 7 // -10+3
		gshift = 2 // -5+3
		bshift = 3 // -0+3

		rmask = 0x7c00
		gmask = 0x03e0
		bmask = 0x001f
	)

	v7 := byte(r.p.field_259)

	rmul := uint16(byte(r.p.field_24))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_26))

	for i := 0; i < sz; i++ {
		c1 := dst[i]                          // old color
		c2 := binary.LittleEndian.Uint16(src) // color to draw

		cr1 := (c1 & rmask) >> rshift
		cg1 := (c1 & gmask) >> gshift
		cb1 := (c1 & bmask) << bshift

		cr2 := (c2 & rmask) >> rshift
		cg2 := (c2 & gmask) >> gshift
		cb2 := (c2 & bmask) << bshift

		cr := r.colors.R[byte(cr1+((uint16(v7)*(((rmul*cr2)>>8)-cr1))>>8))]
		cg := r.colors.G[byte(cg1+((uint16(v7)*(((gmul*cg2)>>8)-cg1))>>8))]
		cb := r.colors.B[byte(cb1+((uint16(v7)*(((bmul*cb2)>>8)-cb1))>>8))]

		dst[i] = cr | cg | cb

		src = src[2:]
	}

	return dst[sz:], src
}

func (r *NoxRender) sub_4C92F0(dst []uint16, src []byte, op byte, sz int) (_ []uint16, _ []byte) { // sub_4C92F0
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[sz:]

	const (
		rshift = 3
		gshift = 2
		bshift = 7

		rmask = 0x001f
		gmask = 0x03e0
		bmask = 0x7c00
	)

	v9 := byte(r.p.field_259)

	rmul := uint16(byte(r.p.field_24))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_26))

	v12 := r.field66(int(op >> 4))

	rpmul := v12[6]
	gpmul := v12[7]
	bpmul := v12[8]

	for i := 0; i < sz; i++ {
		c1 := dst[i] // old color
		v2 := src[i] // color to draw

		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift

		cr = r.colors.R[byte(cr+(uint16(uint32(v9)*uint32(((rmul*uint16(((rpmul*uint32(v2))>>8)&0xFF))>>8)-cr))>>8))]
		cg = r.colors.G[byte(cg+(uint16(uint32(v9)*uint32(((gmul*uint16(((gpmul*uint32(v2))>>8)&0xFF))>>8)-cg))>>8))]
		cb = r.colors.B[byte(cb+(uint16(uint32(v9)*uint32(((bmul*uint16(((bpmul*uint32(v2))>>8)&0xFF))>>8)-cb))>>8))]

		dst[i] = cr | cg | cb
	}

	return dst[sz:], src[sz:]
}

func (r *NoxRender) sub_4C8D60(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C8D60
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]

	rmul := uint16(byte(r.p.field_24))
	gmul := uint16(byte(r.p.field_25))
	bmul := uint16(byte(r.p.field_26))

	for i := 0; i < sz; i++ {
		ci := binary.LittleEndian.Uint16(src)
		var c byte
		if int(ci) < len(nox_draw_colorTablesRev_3804668) {
			c = nox_draw_colorTablesRev_3804668[ci]
		}
		cr := r.colors.R[byte((uint32(rmul)*uint32(c))>>8)]
		cg := r.colors.G[byte((uint32(gmul)*uint32(c))>>8)]
		cb := r.colors.B[byte((uint32(bmul)*uint32(c))>>8)]
		dst[i] = cr | cg | cb
		src = src[2:]
	}
	return dst[sz:], src
}

// SADD8 is a saturating 8-bit addition.
func SADD8(x, y byte) byte {
	z := uint16(x) + uint16(y)
	if z > 0xff {
		return 0xff
	}
	return byte(z)
}

func (r *NoxRender) pixBlendPremult(dst []uint16, src []byte, _ byte, sz int) (_ []uint16, _ []byte) { // sub_4C9B20
	if sz < 0 {
		panic("negative size")
	}
	_ = dst[sz:]
	_ = src[2*sz:]
	const (
		rshift = 7 // -10+3
		gshift = 2 // -5+3
		bshift = 3 // -0+3

		rmask = 0x7c00
		gmask = 0x03e0
		bmask = 0x001f
	)
	for i := 0; i < sz; i++ {
		c1 := dst[i]
		c2 := binary.LittleEndian.Uint16(src)
		cr := (c1 & rmask) >> rshift
		cg := (c1 & gmask) >> gshift
		cb := (c1 & bmask) << bshift
		crb := ((c2>>8)&0xF8 | 7) + cr
		cgb := ((c2>>3)&0xFC | 3) + cg
		rbb := ((c2<<3)&0xF8 | 7) + cb
		if crb > 0xff {
			crb = 0xff
		}
		if cgb > 0xff {
			cgb = 0xff
		}
		if rbb > 0xff {
			rbb = 0xff
		}
		c := r.colors.R[crb] | r.colors.G[cgb] | r.colors.B[rbb]
		dst[i] = c
		src = src[2:]
	}
	return dst[sz:], src
}

func nox_draw_initColorTables_434CC0() {
	mode := noxcolor.GetMode()
	if C.dword_5d4594_3801780 == 0 {
		mode = noxcolor.ModeRGBA5551
	}
	noxrend.SetColorMode(mode)
	arrR, _ := alloc.Uints16(257)
	arrG, _ := alloc.Uints16(257)
	arrB, _ := alloc.Uints16(257)
	arrR = arrR[:256]
	arrG = arrG[:256]
	arrB = arrB[:256]
	copy(arrR, noxrend.colors.R[:])
	copy(arrG, noxrend.colors.G[:])
	copy(arrB, noxrend.colors.B[:])
	C.nox_draw_colors_r_3804672 = (*C.uchar)(unsafe.Pointer(&arrR[0]))
	C.nox_draw_colors_g_3804656 = (*C.uchar)(unsafe.Pointer(&arrG[0]))
	C.nox_draw_colors_b_3804664 = (*C.uchar)(unsafe.Pointer(&arrB[0]))
}

func (r *NoxRender) SetColorMode(mode noxcolor.Mode) {
	if r.colors.mode == mode {
		return
	}
	for i := 0; i < 256; i++ {
		r.colors.R[i] = mode.RGB(byte(i), 0, 0).Color16()
		r.colors.G[i] = mode.RGB(0, byte(i), 0).Color16()
		r.colors.B[i] = mode.RGB(0, 0, byte(i)).Color16()
	}
	r.colors.mode = mode
}

var (
	drawWhiteBubbleParticle     uint32
	drawLightBlueBubbleParticle uint32
	drawRedBubbleParticle       uint32
	drawOrangeBubbleParticle    uint32
	drawYellowBubbleParticle    uint32
	drawGreenBubbleParticle     uint32
	drawWhiteSpark              uint32
	drawColorXxxLoaded1096552   bool
	drawColorXxx1096452         uint32
	drawColorXxx1096436         uint32
)

func sub_499F60(p uint32, pos types.Point, a4 int, a5, a6, a7, a8, a9 int, a10 int) {
	C.sub_499F60(C.int(p), C.int(pos.X), C.int(pos.Y), C.short(a4), C.char(a5), C.char(a6), C.char(a7), C.char(a8), C.char(a9), C.int(a10))
}

func drawCreatureBackEffects(r *NoxRender, vp *Viewport, dr *Drawable) { // nox_xxx_drawHasteAndRunParticles_4746C0
	if dr.CheckFlag31(0) && C.sub_474B40(dr.C()) == 0 {
		return
	}
	if dr.CheckFlag31(14) {
		pos := vp.toScreenPos(dr.Pos())
		r.DrawGlow(pos, memmap.Uint32(0x85B3FC, 980), 30, 31)
	}
	if dr.CheckFlag31(9) && !nox_xxx_checkGameFlagPause_413A50() { // Haste effect
		if drawWhiteBubbleParticle == 0 {
			drawWhiteBubbleParticle = nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteBubbleParticle")
			drawLightBlueBubbleParticle = nox_xxx_getTTByNameSpriteMB_44CFC0("LightBlueBubbleParticle")
		}
		pos := dr.Pos()
		v2 := 0
		if dr.Pos() != dr.Point8() {
			v2 = 2
		}
		for ; v2 > 0; v2-- {
			v22 := randomIntMinMax(3, 5)
			v18 := randomIntMinMax(3, 6)
			v14 := randomIntMinMax(2, 4)
			pos2 := types.Point{
				X: randomIntMinMax(-10, 10),
				Y: randomIntMinMax(-10, 10) + int(*(*int16)(dr.field(104))),
			}
			sub_499F60(drawWhiteBubbleParticle, pos.Add(pos2), 1, v14, v18, 0, 0, 0, v22)

			v23 := randomIntMinMax(3, 5)
			v19 := randomIntMinMax(3, 6)
			v15 := randomIntMinMax(2, 4)
			pos3 := types.Point{
				X: randomIntMinMax(-10, 10),
				Y: randomIntMinMax(-10, 10) + int(*(*int16)(dr.field(104))),
			}
			sub_499F60(drawLightBlueBubbleParticle, pos.Add(pos3), 1, v15, v19, 0, 0, 0, v23)
		}
	}
	if dr.CheckFlag31(8) && !nox_xxx_checkGameFlagPause_413A50() { // Run effect
		if drawRedBubbleParticle == 0 {
			drawRedBubbleParticle = nox_xxx_getTTByNameSpriteMB_44CFC0("RedBubbleParticle")
			drawOrangeBubbleParticle = nox_xxx_getTTByNameSpriteMB_44CFC0("OrangeBubbleParticle")
		}
		pos := dr.Pos()
		v5 := 1
		if dr.Pos() != dr.Point8() {
			v5 = 2
		}
		for ; v5 > 0; v5-- {
			v24 := randomIntMinMax(3, 5)
			v20 := randomIntMinMax(3, 6)
			v16 := randomIntMinMax(2, 4)
			pos2 := types.Point{
				X: randomIntMinMax(-10, 10),
				Y: randomIntMinMax(-10, 10) + int(*(*int16)(dr.field(104))),
			}
			sub_499F60(drawRedBubbleParticle, pos.Add(pos2), 1, v16, v20, 0, 0, 0, v24)

			v25 := randomIntMinMax(3, 5)
			v21 := randomIntMinMax(3, 6)
			v17 := randomIntMinMax(2, 4)
			pos3 := types.Point{
				X: randomIntMinMax(-10, 10),
				Y: randomIntMinMax(-10, 10) + int(*(*int16)(dr.field(104))),
			}
			sub_499F60(drawOrangeBubbleParticle, pos.Add(pos3), 1, v17, v21, 0, 0, 0, v25)
		}
	}
	// Protection effects
	if dr.CheckFlag31(17) {
		r.drawProtectEffectDefault(vp, dr.Pos(), dr, 0, 0, memmap.Uint32(0x85B3FC, 940), memmap.Uint32(0x5D4594, 2589776), true)
	}
	if dr.CheckFlag31(18) {
		r.drawProtectEffectDefault(vp, dr.Pos(), dr, 85, 1, memmap.Uint32(0x8531A0, 2572), memmap.Uint32(0x852978, 24), true)
	}
	if dr.CheckFlag31(20) {
		r.drawProtectEffectDefault(vp, dr.Pos(), dr, 170, 2, memmap.Uint32(0x85B3FC, 980), memmap.Uint32(0x5D4594, 2523948), true)
	}
	if dr.CheckFlag31(27) { // Shield effects
		switch *(*byte)(dr.field(297)) {
		case 0, 1, 2:
			C.nox_xxx_drawShield_499810(vp.C(), dr.C())
		}
	}
}

func drawCreatureFrontEffects(r *NoxRender, vp *Viewport, dr *Drawable) { // nox_xxx_drawEffectsMB_474E60
	if dr.CheckFlag31(0) && C.sub_474B40(dr.C()) == 0 {
		return
	}
	if dr.CheckFlag31(22) {
		if drawWhiteSpark == 0 {
			drawWhiteSpark = nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteSpark")
		}
		pos := dr.Pos()
		C.nox_xxx_drawEnergyBolt_499710(C.int(pos.X), C.int(pos.Y), C.short(*(*int16)(dr.field(104))), C.int(drawWhiteSpark))
	}
	if dr.CheckFlag31(3) || dr.CheckFlag31(5) || dr.CheckFlag31(29) || dr.CheckFlag31(28) {
		pos := vp.toScreenPos(dr.Pos())
		v5 := 5 - int(*(*int16)(dr.field(106))) - int(*(*int16)(dr.field(104))) - int(dr.Field25())
		v6 := *(*byte)(dr.field(112))
		pos.Y += v5
		if v6&0x4 != 0 && *(*uint32)(dr.field(276)) == 6 {
			v8 := 8 * uintptr(*(*byte)(dr.field(297)))
			pos.X += int(memmap.Int32(0x587000, 149432+v8))
			pos.Y += int(memmap.Int32(0x587000, 149436+v8))
		}
		if dr.CheckFlag31(29) {
			r.setField17(1)
			C.sub_433E40(C.int(memmap.Uint32(0x85B3FC, 980)))
		}
		r.nox_video_drawAnimatedImageOrCursorAt(asImageRefP(*memmap.PtrPtr(0x5D4594, 1096456)), pos.Add(types.Point{X: -64, Y: -64}))
		r.setField17(0)
	}
	if dr.CheckFlag31(4) && !nox_xxx_checkGameFlagPause_413A50() {
		v11 := int(*(*float32)(dr.field(48)))
		v44 := int(dr.Field25() * 0.5)
		if drawYellowBubbleParticle == 0 {
			drawYellowBubbleParticle = nox_xxx_getTTByNameSpriteMB_44CFC0("YellowBubbleParticle")
		}
		pos := dr.Pos()
		for v12 := 0; v12 < 2; v12++ {
			v40 := randomIntMinMax(3, 5)
			v36 := randomIntMinMax(3, 6)
			v34 := randomIntMinMax(1, 2)
			v32 := randomIntMinMax(0, v44)
			pos2 := types.Point{
				X: randomIntMinMax(-v11, v11),
				Y: randomIntMinMax(-v11, v11),
			}
			sub_499F60(drawYellowBubbleParticle, pos.Add(pos2), v32, v34, v36, -5, 0, 0, v40)
		}
	}
	if dr.CheckFlag31(21) && !nox_xxx_checkGameFlagPause_413A50() {
		if drawGreenBubbleParticle == 0 {
			drawGreenBubbleParticle = nox_xxx_getTTByNameSpriteMB_44CFC0("GreenBubbleParticle")
		}
		pos := dr.Pos()
		v41 := randomIntMinMax(2, 3)
		v37 := randomIntMinMax(3, 6)
		v35 := randomIntMinMax(2, 4)
		v29 := int(dr.Field25()) + 12
		v14 := int(dr.Field25())
		v33 := randomIntMinMax(v14+8, v29)
		pos2 := types.Point{
			X: randomIntMinMax(-6, 6),
			Y: randomIntMinMax(-10, 10),
		}
		sub_499F60(drawGreenBubbleParticle, pos.Add(pos2), v33, v35, v37, 1, 0, 0, v41)
	}
	if dr.CheckFlag31(13) && !nox_xxx_checkGameFlagPause_413A50() {
		if !drawColorXxxLoaded1096552 {
			drawColorXxx1096452 = noxcolor.ExtendColor16(noxcolor.RGBColor(255, 0, 255))
			drawColorXxx1096436 = noxcolor.ExtendColor16(noxcolor.RGBColor(255, 180, 255))
			drawColorXxxLoaded1096552 = true
		}
		pos := vp.toScreenPos(dr.Pos())

		for v16 := 0; v16 < 10; v16++ {
			v17 := randomIntMinMax(1, 2)
			v38 := int(*(*float32)(dr.field(48)))
			v18 := int(*(*float32)(dr.field(48)))
			v20 := int(dr.Field25())
			pos2 := pos.Add(types.Point{
				X: randomIntMinMax(-v18, v38),
				Y: randomIntMinMax(-10-v20, 0) + int(*(*int16)(dr.field(104))),
			})
			v22 := randomIntMinMax(3, 4)
			r.DrawGlow(pos2, drawColorXxx1096452, v17+v22, v17+2)
			r.SetColor2(drawColorXxx1096436)
			r.DrawPoint(pos2, v17)
		}
	}
	if dr.CheckFlag31(17) {
		r.drawProtectEffectDefault(vp, dr.Pos(), dr, 0, 0, memmap.Uint32(0x85B3FC, 940), memmap.Uint32(0x5D4594, 2589776), false)
	}
	if dr.CheckFlag31(18) {
		r.drawProtectEffectDefault(vp, dr.Pos(), dr, 85, 1, memmap.Uint32(0x8531A0, 2572), memmap.Uint32(0x852978, 24), false)
	}
	if dr.CheckFlag31(20) {
		r.drawProtectEffectDefault(vp, dr.Pos(), dr, 170, 2, memmap.Uint32(0x85B3FC, 980), memmap.Uint32(0x5D4594, 2523948), false)
	}
	if dr.CheckFlag31(26) {
		pos := vp.toScreenPos(dr.Pos())
		v23 := *(*uint32)(dr.field(276))
		v24 := -90 - int(*(*int16)(dr.field(104)))
		pos.X -= 64
		pos.Y += v24
		if v23 == 6 {
			v26 := 8 * uintptr(*(*byte)(dr.field(297)))
			pos.X += int(memmap.Int32(0x587000, 149504+v26))
			pos.Y += int(memmap.Int32(0x587000, 149508+v26))
		}
		r.SetAlphaEnabled(true)
		r.SetAlpha(0x80)
		r.nox_video_drawAnimatedImageOrCursorAt(asImageRefP(*memmap.PtrPtr(0x5D4594, 1096460)), pos)
		r.SetAlphaEnabled(false)
	}
	if dr.CheckFlag31(27) {
		switch *(*byte)(dr.field(297)) {
		default:
			C.nox_xxx_drawShield_499810(vp.C(), dr.C())
		case 0, 1, 2:
		}
	}
}
