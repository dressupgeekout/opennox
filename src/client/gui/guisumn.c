#include "guisumn.h"

#include "../../proto.h"
#include "gamewin/gamewin.h"

extern _DWORD dword_5d4594_1320992;
extern _DWORD dword_5d4594_1321204;
extern _DWORD dword_5d4594_1321040;
extern int nox_win_width;
extern int nox_win_height;

//----- (004C1D80) --------------------------------------------------------
int sub_4C1D80() {
	wchar_t* v0;         // eax
	_DWORD* v1;          // esi
	char* v2;            // eax
	_DWORD* v3;          // eax
	_DWORD* v4;          // eax
	int v5;              // edx
	unsigned __int8* v6; // eax

	*(_DWORD*)&byte_5D4594[1321004] = 0;
	*(_DWORD*)&byte_5D4594[1321000] = -145;
	*(_DWORD*)&byte_5D4594[1320988] = nox_win_width - 95;
	dword_5d4594_1320992 = -145;
	*(_DWORD*)&byte_5D4594[1321032] = nox_window_new(0, 8, nox_win_width - 95, -145, 87, 115, 0);
	nox_window_set_all_funcs(*(_DWORD**)&byte_5D4594[1321032], sub_4C2BD0, sub_4C24A0, 0);
	*(_DWORD*)&byte_5D4594[1321036] = nox_window_new(*(int*)&byte_5D4594[1321032], 136, 5, 38, 76, 76, 0);
	nox_window_set_all_funcs(*(_DWORD**)&byte_5D4594[1321036], sub_4C2B10, sub_4C1FE0, sub_4C2C20);
	v0 = loadString_sub_40F1D0((char*)&byte_587000[184592], 0, "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c", 818);
	sub_46B000((wchar_t*)(*(_DWORD*)&byte_5D4594[1321036] + 36), v0);
	*(_DWORD*)&byte_5D4594[1320996] = sub_42F970("CreatureCageBottom");
	v1 = nox_window_new(*(int*)&byte_5D4594[1321036], 160, 0, 0, 1, 1, 0);
	v2 = sub_42F970("CreatureCageTop");
	sub_46AE60((int)v1, (int)v2);
	sub_46AE40((int)v1, -5, -38);
	v3 = nox_window_new(*(int*)&byte_5D4594[1321032], 8, 19, 0, 48, 39, 0);
	nox_window_set_all_funcs(v3, sub_4C2BE0, sub_4C24A0, 0);
	*(_DWORD*)&byte_5D4594[1321008] = sub_42F970("CreatureCageHuntButtonLit");
	*(_DWORD*)&byte_5D4594[1321012] = sub_42F970("CreatureCageHuntButton");
	*(_DWORD*)&byte_5D4594[1321016] = sub_42F970("CreatureCageGuardButtonLit");
	*(_DWORD*)&byte_5D4594[1321020] = sub_42F970("CreatureCageGuardButton");
	*(_DWORD*)&byte_5D4594[1321024] = sub_42F970("CreatureCageEscortButtonLit");
	*(_DWORD*)&byte_5D4594[1321028] = sub_42F970("CreatureCageEscortButton");
	v4 = nox_window_new(0, 168, *(_DWORD*)&byte_5D4594[1320988] + 27, dword_5d4594_1320992 + 12, 34, 34,
			    0);
	dword_5d4594_1321040 = v4;
	v5 = v4[11];
	BYTE1(v5) |= 1u;
	v4[11] = v5;
	nox_window_set_all_funcs(*(_DWORD**)&dword_5d4594_1321040, sub_4C24B0, 0, sub_4C2CE0);
	sub_46AE60(*(int*)&dword_5d4594_1321040, *(int*)&byte_5D4594[1321028]);
	sub_46AEC0(*(int*)&dword_5d4594_1321040, *(int*)&byte_5D4594[1321024]);
	sub_46AEA0(*(int*)&dword_5d4594_1321040, *(int*)&byte_5D4594[1321024]);
	sub_46AE40(*(int*)&dword_5d4594_1321040, -27, -12);
	byte_5D4594[1321200] = 0;
	nox_window_set_hidden(*(int*)&byte_5D4594[1321032], 1);
	nox_window_set_hidden(*(int*)&dword_5d4594_1321040, 1);
	v6 = &byte_5D4594[1321060];
	do {
		*(_DWORD*)v6 = 0;
		v6 += 32;
	} while ((int)v6 < (int)&byte_5D4594[1321188]);
	sub_4C2BF0();
	*(_DWORD*)&byte_5D4594[1321044] = 0;
	dword_5d4594_1321204 = 0;
	*(_DWORD*)&byte_5D4594[1321196] = 0;
	return 1;
}

//----- (004C2560) --------------------------------------------------------
_DWORD* __cdecl sub_4C2560(int2* a1) {
	char** v1;            // esi
	unsigned __int16* v2; // eax
	int v3;               // esi
	bool v4;              // sf
	int v5;               // ecx
	int v6;               // edi
	int v7;               // eax
	int v8;               // edi
	int i;                // ebp
	_DWORD* v10;          // ebx
	int v12;              // [esp+10h] [ebp-8h]
	int v13;              // [esp+14h] [ebp-4h]

	*(_DWORD*)&byte_587000[184452] = 0;
	v1 = (char**)&byte_587000[184344];
	do {
		if (v1 != (char**)&byte_587000[184352]) {
			v2 = loadString_sub_40F1D0(*v1, 0, "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c", 588);
			sub_43F840(0, v2, &v12, &v13, nox_win_width);
			if (*(_DWORD*)&byte_587000[184452] < v12)
				*(_DWORD*)&byte_587000[184452] = v12;
		}
		++v1;
	} while ((int)v1 < (int)&byte_587000[184368]);
	*(_DWORD*)&byte_587000[184452] += 8;
	v3 = sub_43F320(0) + 2;
	v5 = a1->field_0 - *(_DWORD*)&byte_587000[184452] / 2;
	v4 = v5 < 0;
	v12 = a1->field_0 - *(_DWORD*)&byte_587000[184452] / 2;
	v6 = 5 * v3 + 12;
	if (v4) {
		v5 = 0;
	} else {
		if (*(_DWORD*)&byte_587000[184452] + v5 < nox_win_width)
			goto LABEL_11;
		v5 = nox_win_width - *(_DWORD*)&byte_587000[184452] - 1;
	}
	v12 = v5;
LABEL_11:
	v7 = a1->field_4 - v6 / 2;
	v13 = v7;
	if (v7 < 0) {
		v7 = 0;
	LABEL_15:
		v13 = v7;
		goto LABEL_16;
	}
	if (v6 + v7 >= nox_win_height) {
		v7 = nox_win_width - v6 - 1;
		goto LABEL_15;
	}
LABEL_16:
	*(_DWORD*)&byte_5D4594[1321044] = nox_window_new(0, 40, v5, v7, *(int*)&byte_587000[184452], 5 * v3 + 12, 0);
	nox_window_set_all_funcs(*(_DWORD**)&byte_5D4594[1321044], 0, sub_4C26F0, 0);
	sub_46A8C0(*(int*)&byte_5D4594[1321044]);
	v8 = 0;
	for (i = 0; i < 6; ++i) {
		if (i != 2) {
			v10 = nox_window_new(*(int*)&byte_5D4594[1321044], 8, 0, v8, *(int*)&byte_587000[184452],
					     v3 + 1, 0);
			nox_window_set_all_funcs(v10, sub_4C2A60, sub_4C27F0, 0);
			v10[8] = i;
			v8 += v3 + 2;
		}
	}
	return sub_452D80(791, 100);
}

//----- (004C27F0) --------------------------------------------------------
int __cdecl sub_4C27F0(_DWORD* a1) {
	int result;           // eax
	unsigned __int16* v2; // edi
	int2* v3;             // ebp
	int v4;               // esi
	int v5;               // ebx
	int v6;               // edx
	int v7;               // ebx
	int v8;               // [esp-18h] [ebp-24h]
	__int16* v9;          // [esp-14h] [ebp-20h]
	int v10;              // [esp+0h] [ebp-Ch]
	int v11;              // [esp+4h] [ebp-8h]
	int v12;              // [esp+8h] [ebp-4h]

	if (!*(_DWORD*)&byte_5D4594[1321208])
		*(_DWORD*)&byte_5D4594[1321208] = sub_4E3AA0((CHAR*)&byte_587000[184840]);
	if (dword_5d4594_1321204 || (result = 1, a1[8] != 1)) {
		v2 = loadString_sub_40F1D0(*(char**)&byte_587000[4 * a1[8] + 184344], 0,
					   "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c", 446);
		nox_client_wndGetPosition_46AA60(a1, &v11, &v10);
		sub_43F840(0, v2, &v12, 0, 0);
		v3 = nox_client_getMousePos_4309F0();
		sub_43F320(0);
		v4 = (*(_DWORD*)&byte_587000[184452] - v12) / 2 + 1;
		if (sub_46AAB0(a1, v3->field_0, v3->field_4)) {
			sub_4C2A00(v11 + v4, v10 + 3, *(int*)&byte_5D4594[2589772], *(int*)&byte_5D4594[2650656],
				   (__int16*)v2);
			if (a1[8] != *(_DWORD*)&byte_587000[184552]) {
				*(_DWORD*)&byte_587000[184552] = a1[8];
				sub_452D80(920, 100);
				return 1;
			}
			return 1;
		}
		if (dword_5d4594_1321204) {
			if (sub_4C2DD0(*(int*)&dword_5d4594_1321204)) {
				sub_4C2A00(v11 + v4, v10 + 3, *(int*)&byte_5D4594[2523948],
					   *(int*)&byte_5D4594[2650656], (__int16*)v2);
				return 1;
			}
			v5 = a1[8];
			if (v5 != 4 && v5 != 5) {
				sub_4C2A00(v11 + v4, v10 + 3, *(int*)&byte_5D4594[2523948],
					   *(int*)&byte_5D4594[2650656], (__int16*)v2);
				return 1;
			}
			v6 = *(_DWORD*)&byte_5D4594[2650660];
			v9 = (__int16*)v2;
			v8 = *(_DWORD*)&byte_5D4594[2650656];
		} else {
			v7 = a1[8];
			if (v7 != 4 && v7 != 5) {
				sub_4C2A00(v11 + v4, v10 + 3, *(int*)&byte_5D4594[2650684],
					   *(int*)&byte_5D4594[2650656], (__int16*)v2);
				return 1;
			}
			v9 = (__int16*)v2;
			if (!sub_4C2E00()) {
				sub_4C2A00(v11 + v4, v10 + 3, *(int*)&byte_5D4594[2650660],
					   *(int*)&byte_5D4594[2650656], (__int16*)v2);
				return 1;
			}
			v6 = *(_DWORD*)&byte_5D4594[2650684];
			v8 = *(_DWORD*)&byte_5D4594[2650656];
		}
		sub_4C2A00(v11 + v4, v10 + 3, v6, v8, v9);
		return 1;
	}
	return result;
}

//----- (004C2CE0) --------------------------------------------------------
int sub_4C2CE0() {
	wchar_t* v0; // eax
	wchar_t* v2; // eax
	wchar_t* v3; // eax

	switch (*(_DWORD*)&byte_587000[184448]) {
	case 3:
		v3 = loadString_sub_40F1D0((char*)&byte_587000[184980], 0, "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c",
					   785);
		sub_4776B0(v3);
		break;
	case 4:
		v2 = loadString_sub_40F1D0((char*)&byte_587000[184932], 0, "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c",
					   782);
		sub_4776B0(v2);
		return 1;
	case 5:
		v0 = loadString_sub_40F1D0((char*)&byte_587000[185028], 0, "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c",
					   788);
		sub_4776B0(v0);
		return 1;
	}
	return 1;
}
