#ifndef NOX_PORT_GAME1_2
#define NOX_PORT_GAME1_2

#include "defs.h"

int nox_xxx_wallMath_427F30(int2* a1, int* a2);
int sub_428170(void* a1, int4* a2);
int nox_xxx_pointInRect_4281F0(int2* a1, int4* a2);
int sub_428220(float2* a1, float4* a2);
void nox_shape_box_calc(nox_shape* s);
char* sub_4282D0(char* a1, int a2);
void* sub_4282F0(int a1, int a2, size_t a3);
unsigned int sub_428540(int a1, char* a2, int a3);
char* sub_4285C0(short* a1);
int sub_428810(int a1, int a2);
int sub_428890(short* a1);
void* sub_4289D0(void** a1);
int nox_server_mapRWObjectTOC_428B30();
int nox_server_mapRWPolygons_428CD0(int a1);
int nox_server_mapRWAmbientData_429200();
int nox_server_mapRWWindowWalls_4292C0(uint32_t* a1);
void sub_429450(uint8_t* a1, uint32_t* a2);
void sub_4294B0(uint8_t* a1, uint32_t* a2);
void nox_xxx_wallBreakableCounterClear_429520();
int nox_server_mapRWDestructableWalls_429530(uint32_t* a1);
void nox_xxx_wall_4296E0(uint8_t* a1, uint32_t* a2);
void sub_429740(uint8_t* a1, uint32_t* a2);
void nox_xxx_wallSecretCounterClear_4297B0();
int nox_server_mapRWSecretWalls_4297C0(uint32_t* a1);
void sub_429A00(uint8_t* a1, uint32_t* a2);
void sub_429A60(int a1, uint32_t* a2);
int nox_server_mapRWWallMap_429B20(uint32_t* a1);
int sub_42A0F0(int a1);
int sub_42A150(short a1, uint32_t* a2);
int sub_42A650(unsigned char* a1);
unsigned char nox_xxx_wall_42A6C0(unsigned char a1, unsigned char a2);
int nox_server_mapRWMapInfo_42A6E0();
uint16_t* sub_42A8B0(uint8_t* a1, int* a2);
int sub_42A970(uint8_t* a1, uint8_t* a2, int* a3);
double sub_42AAA0(int* a1);
void sub_42ABF0(int a1, int a2, int a3);
uint8_t* sub_42AC50(uint8_t* a1, size_t* a2);
uint16_t* sub_42ADA0(int a1, int a2, short a3, unsigned int* a4);
uint16_t* sub_42B810(short* a1, unsigned int* a2);
int sub_42BCE0(uint32_t* this, char* a2, char a3);
int sub_42BD50(uint32_t* this, char* a2, char a3);
int sub_42BDC0(uint32_t* this, char* a2, char a3);
int sub_42BE30(uint32_t* this, char* a2, const char* a3);
int sub_42BEA0(uint32_t* this, char* a2, const void* a3, unsigned short a4);
void sub_42BFB0();
void sub_42BFE0();
int nox_xxx_objectTOCgetTT_42C2B0(unsigned short a1);
unsigned short sub_42C2E0(int a1);
unsigned short sub_42C300();
void sub_42C310(int a1, unsigned short a2);
void sub_42C330(uint32_t* this);
int sub_42C360(uint32_t* this, int a2);
uint16_t* sub_42C480(uint32_t* this, unsigned int* a2);
int sub_42C770(void** this);
int sub_42C7F0(int this, char* a2, char a3);
int sub_42C820(int this, char* a2, char a3);
int sub_42C8B0(int this, char* a2, char a3);
int sub_42C8E0(int this, char* a2, const char* a3);
int sub_42C910(int this, char* a2, const void* a3, unsigned short a4);
char* sub_42C9A0(void** this, char* a2, char a3);
char* sub_42CA00(void** this, char* a2, char a3);
char* sub_42CB20(void** this, char* a2, char a3);
char* sub_42CB80(void** this, char* a2, const char* a3);
void* sub_42CBF0(void** this, char* a2, const void* a3, unsigned short a4);
uint16_t sub_42CC70(int this);
short sub_42CCE0(uint16_t* this);
void nox_xxx_clientTalk_42E7B0(nox_drawable* a1p);
void nox_xxx_clientCollideOrUse_42E810(nox_drawable* a1p);
void nox_xxx_clientTrade_42E850(nox_drawable* a1p);
int sub_42EB90(int a1);
int sub_42EBA0();
void sub_42EDC0();
nox_video_bag_image_t* nox_xxx_readImgMB_42FAA0(int known_idx, char a2, char* a3);
void* nox_video_getImagePixdata_42FB30(nox_video_bag_image_t* img);
nox_things_imageRef_t* nox_xxx_gLoadAnim_42FA20(char* a1);
nox_video_bag_image_t* nox_xxx_gLoadImg_42F970(char* name);
int sub_430AA0(int a1);
int nox_client_mousePriKey_430AF0();
int nox_xxx_cursor_430B00();
void nox_client_setMousePos_430B10(int x, int y);
int sub_430B50(int a1, int a2, int a3, int a4);
void sub_430C30_set_video_max(int w, int h);
void nox_xxx_screenGetSize_430C50_get_video_max(int* w, int* h);
int nox_audio_initall(int a3);
int sub_4311F0();
void sub_431270();
void sub_431290();
int sub_431370();
void sub_431380();
void nox_client_resetScreenParticles_431510();
nox_screenParticle* nox_client_newScreenParticle_431540(int a1, int a2, int a3, int a4, int a5, int a6, char a7,
														char a8, char a9, char a10);
void nox_client_addScreenParticle_431680(nox_screenParticle* p);
void sub_4316C0(nox_screenParticle* p);
void sub_431700(uint64_t* a1);
void nox_client_screenParticlesDraw_431720(nox_draw_viewport_t* rdr);
char* nox_xxx_getHostInfoPtr_431770();
char* nox_xxx_copyServerIPAndPort_431790(char* a1);
void sub_434080(int a1);
void nox_xxx_drawSetTextColor_434390(int a1);
void nox_xxx_drawSetColor_4343E0(int a1);
void nox_client_drawSetColor_434460(int a1);
void nox_client_drawEnableAlpha_434560(int a1);
void nox_client_drawSetAlpha_434580(unsigned char a1);
void sub_4345F0(int a1);
void nox_xxx_draw_434600(int a1);
void sub_434990(int a1, int a2, int a3);
void sub_4349C0(uint32_t* a1);
void sub_435040();
void sub_435120(void* a1, void* a2);
void sub_435150(uint8_t* a1, char* a2);
void nox_draw_splitColor_435280(short a1, uint8_t* a2, uint8_t* a3, uint8_t* a4);
long long nox_xxx_initTime_435570();
int sub_435590();
void sub_4355B0(int a1);
void nox_xxx_cliUpdateCameraPos_435600(int a1, int a2);
void nox_xxx_getSomeCoods_435670(int2* a1);
uint32_t* sub_435690(uint32_t* a1);
int nox_xxx_gameGetPlayState_4356B0();
bool nox_client_drawable_testBuff_4356C0(nox_drawable* dr, char a2);
void nox_xxx_spriteLoadError_4356E0();
wchar_t* sub_435700(wchar_t* a1, int a2);
void nox_client_setServerConnectAddr_435720(char* addr);
int nox_xxx_cliToggleObsWindow_4357A0();
int sub_435F60();
int sub_436550();
void sub_437100();
nox_draw_viewport_t* nox_draw_getViewport_437250();
void sub_437260();
void sub_437290();
int nox_xxx_playerAnimCheck_4372B0();
int nox_xxx_clientIsObserver_4372E0();
int sub_437320(int a1);
void sub_4375C0(int a1);
int sub_437860(int a1, int a2);
void sub_4379C0();
int sub_438330();
int sub_438370();
int sub_438480();
int nox_client_joinGame_438A90();
int sub_438C80(int a1, int a2);
int sub_438DD0(unsigned int a1, unsigned int a2);
int sub_438E30(uint32_t* a1, int a2);
int sub_438EF0(uint32_t* a1, int a2, unsigned int a3, int a4);
int sub_439050(int a1, unsigned int a2, int* a3, unsigned int a4);
uint32_t* sub_439450(int a1, int a2, uint32_t* a3);
char* sub_439CC0(int a1, char* a2);
int sub_439D00(int* a1, int a2, unsigned int a3, int a4);
int sub_439D90(unsigned int a1, unsigned int a2);
int sub_43A920();
int nox_client_guiXxx_43A9D0();
char* sub_43AA70();
int sub_43AF30();
int sub_43AF40();
int sub_43AF80();
int sub_43AF90(int a1);
void nox_client_setConnError_43AFA0(int a1);
unsigned int nox_client_getServerAddr_43B300();
int nox_client_getServerPort_43B320();
int sub_43B340();
int nox_xxx_cliDrawConnectedLoop_43B360();
int sub_43B460();
int sub_43B490();
void nox_xxx_serverHost_43B4D0();

#endif // NOX_PORT_GAME1_2
