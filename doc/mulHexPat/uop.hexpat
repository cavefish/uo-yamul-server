struct MapData{
	u16 color;
	u8 altitude;
} [[color("808000")]];


struct TableEntry {
	u64 offset;
	u32 headerLength;
	u32 size;
	u32 decompressedSize;
	u64 filenameHash;
	u32 hash;
	u16 compression;
	MapData data[100] @ (offset + headerLength);
} [[color("00FF00")]];

struct Table {
	u32 entries;
	u64 next;
	TableEntry values[10];
} [[color("0000FF")]];

struct Header {
	u32 header;
	u32 version;
	u32 timestamp;
	u64 nextTable;
	Table table @ nextTable;
} [[color("FF0000")]] ;

Header header @ 0;
