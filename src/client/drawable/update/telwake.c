#include "telwake.h"

#include "../../../proto.h"

//----- (004CD8D0) --------------------------------------------------------
int __cdecl nox_xxx_updDrawTeleportWake_4CD8D0(int a1, int a2) {
	int v2; // esi
	int v3; // eax
	int v4; // eax
	int v5; // esi

	if (!*getMemU32Ptr(0x5D4594, 1522980))
		*getMemU32Ptr(0x5D4594, 1522980) = nox_xxx_getTTByNameSpriteMB_44CFC0((CHAR*)getMemAt(0x587000, 189808));
	v2 = *(_DWORD*)(a2 + 12) + nox_xxx_randGetMinMax_415FF0(-5, 5, "C:\\NoxPost\\src\\client\\Drawable\\Update\\telwake.c", 32);
	v3 = nox_xxx_randGetMinMax_415FF0(-5, 5, "C:\\NoxPost\\src\\client\\Drawable\\Update\\telwake.c", 33);
	v4 = nox_xxx_spriteLoadAdd_45A360_drawable(*getMemIntPtr(0x5D4594, 1522980), v2, *(_DWORD*)(a2 + 16) + v3);
	v5 = v4;
	if (v4) {
		*(_DWORD*)(v4 + 432) = *(_DWORD*)(v4 + 12) << 12;
		*(_DWORD*)(v4 + 436) = *(_DWORD*)(v4 + 16) << 12;
		*(_BYTE*)(v4 + 299) = nox_xxx_randGetMinMax_415FF0(0, 255, "C:\\NoxPost\\src\\client\\Drawable\\Update\\telwake.c", 43);
		*(_DWORD*)(v5 + 440) = nox_xxx_randGetMinMax_415FF0(1, 100, "C:\\NoxPost\\src\\client\\Drawable\\Update\\telwake.c", 46);
		*(_DWORD*)(v5 + 448) = *getMemU32Ptr(0x5D4594, 2598000) +
							   nox_xxx_randGetMinMax_415FF0(10, 32, "C:\\NoxPost\\src\\client\\Drawable\\Update\\telwake.c", 49);
		*(_DWORD*)(v5 + 444) = *getMemU32Ptr(0x5D4594, 2598000);
		*(_WORD*)(v5 + 104) = 0;
		*(_BYTE*)(v5 + 296) = nox_xxx_randGetMinMax_415FF0(3, 8, "C:\\NoxPost\\src\\client\\Drawable\\Update\\telwake.c", 55);
		nox_xxx_sprite_45A110_drawable((_DWORD*)v5);
	}
	return 1;
}
