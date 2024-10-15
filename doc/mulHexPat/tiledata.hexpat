#pragma pattern_limit 0x30A800

struct StaticTileData {
	u64 flags;
	u8 weight;
	u8 layer;
	u32 count;
	u16 animId;
	u16 hue;
	u16 lightIndex;
	u8 heigt;
	char name[20] [[color("A0A000")]];
} [[color("0000A0")]];

struct StaticTileGroup {
	u32 unknown;
	StaticTileData blocks[32];
};

struct LandTileData {
	u64 flags;
	u16 textureId;
	char name[20] [[color("A0A000")]];
	
} [[color("00A000")]];

struct LandTileGroup {
	u32 unknown;
	LandTileData blocks[32];
} [[color("00FF00")]];

struct TileData {
	LandTileGroup groups[512];
	StaticTileGroup statics[while($ < 0x30A800)];
//	StaticTileGroup statics[512];
};

TileData tileData @ 0;
