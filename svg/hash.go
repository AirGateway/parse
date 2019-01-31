package svg

// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT, except for adding more constants to the list and rerun go generate

// uses github.com/AirGateway/hasher
//go:generate hasher -type=Hash -file=hash.go

// Hash defines perfect hashes for a predefined list of strings
type Hash uint32

// Unique hash definitions to be used instead of strings
const (
	A                            Hash = 0x101   // a
	Alignment_Baseline           Hash = 0x2e12  // alignment-baseline
	BaseProfile                  Hash = 0xb     // baseProfile
	Baseline_Shift               Hash = 0x380e  // baseline-shift
	Buffered_Rendering           Hash = 0x5212  // buffered-rendering
	Clip                         Hash = 0x6404  // clip
	Clip_Path                    Hash = 0x6409  // clip-path
	Clip_Rule                    Hash = 0x8009  // clip-rule
	Color                        Hash = 0xd805  // color
	Color_Interpolation          Hash = 0xd813  // color-interpolation
	Color_Interpolation_Filters  Hash = 0xd81b  // color-interpolation-filters
	Color_Profile                Hash = 0x1ea0d // color-profile
	Color_Rendering              Hash = 0x2250f // color-rendering
	ContentScriptType            Hash = 0xa011  // contentScriptType
	ContentStyleType             Hash = 0xb110  // contentStyleType
	Cursor                       Hash = 0xc106  // cursor
	D                            Hash = 0x5901  // d
	Defs                         Hash = 0x35c04 // defs
	Direction                    Hash = 0x2ff09 // direction
	Display                      Hash = 0x9807  // display
	Dominant_Baseline            Hash = 0x18511 // dominant-baseline
	Enable_Background            Hash = 0x8811  // enable-background
	FeImage                      Hash = 0x14507 // feImage
	Fill                         Hash = 0xc904  // fill
	Fill_Opacity                 Hash = 0x3300c // fill-opacity
	Fill_Rule                    Hash = 0xc909  // fill-rule
	Filter                       Hash = 0xec06  // filter
	Flood_Color                  Hash = 0xd20b  // flood-color
	Flood_Opacity                Hash = 0x1050d // flood-opacity
	Font                         Hash = 0x11404 // font
	Font_Family                  Hash = 0x1140b // font-family
	Font_Size                    Hash = 0x11f09 // font-size
	Font_Size_Adjust             Hash = 0x11f10 // font-size-adjust
	Font_Stretch                 Hash = 0x1370c // font-stretch
	Font_Style                   Hash = 0x14c0a // font-style
	Font_Variant                 Hash = 0x1560c // font-variant
	Font_Weight                  Hash = 0x1620b // font-weight
	G                            Hash = 0x1601  // g
	Glyph_Orientation_Horizontal Hash = 0x1c61c // glyph-orientation-horizontal
	Glyph_Orientation_Vertical   Hash = 0x161a  // glyph-orientation-vertical
	Height                       Hash = 0x6c06  // height
	Href                         Hash = 0x14204 // href
	Image                        Hash = 0x16d05 // image
	Image_Rendering              Hash = 0x16d0f // image-rendering
	Kerning                      Hash = 0x1af07 // kerning
	Letter_Spacing               Hash = 0x90e   // letter-spacing
	Lighting_Color               Hash = 0x1e10e // lighting-color
	Line                         Hash = 0x3c04  // line
	Marker                       Hash = 0x17c06 // marker
	Marker_End                   Hash = 0x17c0a // marker-end
	Marker_Mid                   Hash = 0x1960a // marker-mid
	Marker_Start                 Hash = 0x1a00c // marker-start
	Mask                         Hash = 0x1ac04 // mask
	Metadata                     Hash = 0x1b608 // metadata
	Missing_Glyph                Hash = 0x1be0d // missing-glyph
	Opacity                      Hash = 0x10b07 // opacity
	Overflow                     Hash = 0x25508 // overflow
	Paint_Order                  Hash = 0x2a10b // paint-order
	Path                         Hash = 0x6904  // path
	Pattern                      Hash = 0x1f707 // pattern
	Pointer_Events               Hash = 0x1fe0e // pointer-events
	Points                       Hash = 0x21a06 // points
	Polygon                      Hash = 0x23407 // polygon
	Polyline                     Hash = 0x23b08 // polyline
	PreserveAspectRatio          Hash = 0x24313 // preserveAspectRatio
	Rect                         Hash = 0x30104 // rect
	Rx                           Hash = 0x4f02  // rx
	Ry                           Hash = 0xc602  // ry
	Script                       Hash = 0xf206  // script
	Shape_Rendering              Hash = 0x20b0f // shape-rendering
	Solid_Color                  Hash = 0x21f0b // solid-color
	Solid_Opacity                Hash = 0x35f0d // solid-opacity
	Stop_Color                   Hash = 0x12d0a // stop-color
	Stop_Opacity                 Hash = 0x2670c // stop-opacity
	Stroke                       Hash = 0x27306 // stroke
	Stroke_Dasharray             Hash = 0x27310 // stroke-dasharray
	Stroke_Dashoffset            Hash = 0x28311 // stroke-dashoffset
	Stroke_Linecap               Hash = 0x2940e // stroke-linecap
	Stroke_Linejoin              Hash = 0x2ac0f // stroke-linejoin
	Stroke_Miterlimit            Hash = 0x2bb11 // stroke-miterlimit
	Stroke_Opacity               Hash = 0x2cc0e // stroke-opacity
	Stroke_Width                 Hash = 0x2da0c // stroke-width
	Style                        Hash = 0x15105 // style
	Svg                          Hash = 0x2e603 // svg
	Switch                       Hash = 0x2e906 // switch
	Symbol                       Hash = 0x2ef06 // symbol
	Text_Anchor                  Hash = 0x450b  // text-anchor
	Text_Decoration              Hash = 0x710f  // text-decoration
	Text_Rendering               Hash = 0xf70e  // text-rendering
	Type                         Hash = 0x11004 // type
	Unicode_Bidi                 Hash = 0x2f50c // unicode-bidi
	Use                          Hash = 0x30803 // use
	Vector_Effect                Hash = 0x30b0d // vector-effect
	Version                      Hash = 0x31807 // version
	ViewBox                      Hash = 0x31f07 // viewBox
	Viewport_Fill                Hash = 0x3270d // viewport-fill
	Viewport_Fill_Opacity        Hash = 0x32715 // viewport-fill-opacity
	Visibility                   Hash = 0x33c0a // visibility
	White_Space                  Hash = 0x25c0b // white-space
	Width                        Hash = 0x2e105 // width
	Word_Spacing                 Hash = 0x3460c // word-spacing
	Writing_Mode                 Hash = 0x3520c // writing-mode
	X                            Hash = 0x4701  // x
	X1                           Hash = 0x5002  // x1
	X2                           Hash = 0x32502 // x2
	Xml_Space                    Hash = 0x36c09 // xml:space
	Y                            Hash = 0x1801  // y
	Y1                           Hash = 0x9e02  // y1
	Y2                           Hash = 0xc702  // y2
)

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// ToHash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := uint32(_Hash_hash0)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				goto NEXT
			}
		}
		return i
	}
NEXT:
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				return 0
			}
		}
		return i
	}
	return 0
}

const _Hash_hash0 = 0x30372d7b
const _Hash_maxLen = 28
const _Hash_text = "baseProfiletter-spacinglyph-orientation-verticalignment-base" +
	"line-shiftext-anchorx1buffered-renderingclip-patheightext-de" +
	"corationclip-rulenable-backgroundisplay1contentScriptTypecon" +
	"tentStyleTypecursory2fill-ruleflood-color-interpolation-filt" +
	"erscriptext-renderingflood-opacitypefont-familyfont-size-adj" +
	"ustop-colorfont-stretchrefeImagefont-stylefont-variantfont-w" +
	"eightimage-renderingmarker-endominant-baselinemarker-midmark" +
	"er-startmaskerningmetadatamissing-glyph-orientation-horizont" +
	"alighting-color-profilepatternpointer-eventshape-renderingpo" +
	"intsolid-color-renderingpolygonpolylinepreserveAspectRatiove" +
	"rflowhite-spacestop-opacitystroke-dasharraystroke-dashoffset" +
	"stroke-linecapaint-orderstroke-linejoinstroke-miterlimitstro" +
	"ke-opacitystroke-widthsvgswitchsymbolunicode-bidirectionusev" +
	"ector-effectversionviewBox2viewport-fill-opacityvisibilitywo" +
	"rd-spacingwriting-modefsolid-opacityxml:space"

var _Hash_table = [1 << 7]Hash{
	0x0:  0x2940e, // stroke-linecap
	0x1:  0x1140b, // font-family
	0x2:  0x23b08, // polyline
	0x3:  0x1f707, // pattern
	0x4:  0x30104, // rect
	0x5:  0x5212,  // buffered-rendering
	0x7:  0x2f50c, // unicode-bidi
	0x8:  0x450b,  // text-anchor
	0x9:  0x2bb11, // stroke-miterlimit
	0xa:  0xc909,  // fill-rule
	0xb:  0x27310, // stroke-dasharray
	0xc:  0xc904,  // fill
	0xd:  0x1af07, // kerning
	0xe:  0x2670c, // stop-opacity
	0x10: 0x1a00c, // marker-start
	0x11: 0x380e,  // baseline-shift
	0x14: 0x17c0a, // marker-end
	0x15: 0x18511, // dominant-baseline
	0x16: 0xc602,  // ry
	0x17: 0x161a,  // glyph-orientation-vertical
	0x18: 0x5002,  // x1
	0x19: 0x20b0f, // shape-rendering
	0x1a: 0x32502, // x2
	0x1b: 0x11f10, // font-size-adjust
	0x1c: 0x2250f, // color-rendering
	0x1d: 0x28311, // stroke-dashoffset
	0x1f: 0x3520c, // writing-mode
	0x20: 0x2e906, // switch
	0x21: 0xf70e,  // text-rendering
	0x22: 0x23407, // polygon
	0x23: 0x3460c, // word-spacing
	0x24: 0x21f0b, // solid-color
	0x25: 0xec06,  // filter
	0x26: 0x1801,  // y
	0x27: 0x1be0d, // missing-glyph
	0x29: 0x11404, // font
	0x2a: 0x4f02,  // rx
	0x2b: 0x9807,  // display
	0x2c: 0x2e603, // svg
	0x2d: 0x1050d, // flood-opacity
	0x2f: 0x14204, // href
	0x30: 0x6404,  // clip
	0x31: 0x3c04,  // line
	0x32: 0x1620b, // font-weight
	0x33: 0x1c61c, // glyph-orientation-horizontal
	0x34: 0x6c06,  // height
	0x35: 0x9e02,  // y1
	0x36: 0x6904,  // path
	0x37: 0x31807, // version
	0x38: 0x2ac0f, // stroke-linejoin
	0x39: 0x4701,  // x
	0x3a: 0x30803, // use
	0x3b: 0x2cc0e, // stroke-opacity
	0x3c: 0x15105, // style
	0x3d: 0x30b0d, // vector-effect
	0x3e: 0x14c0a, // font-style
	0x40: 0x16d05, // image
	0x41: 0x1e10e, // lighting-color
	0x42: 0xd813,  // color-interpolation
	0x43: 0x27306, // stroke
	0x44: 0x2ef06, // symbol
	0x47: 0x8811,  // enable-background
	0x48: 0x33c0a, // visibility
	0x49: 0x25508, // overflow
	0x4b: 0x31f07, // viewBox
	0x4c: 0x2e12,  // alignment-baseline
	0x4d: 0x5901,  // d
	0x4e: 0x1560c, // font-variant
	0x4f: 0x1ac04, // mask
	0x50: 0x21a06, // points
	0x51: 0x1b608, // metadata
	0x52: 0x710f,  // text-decoration
	0x53: 0xd81b,  // color-interpolation-filters
	0x54: 0x2ff09, // direction
	0x55: 0x6409,  // clip-path
	0x56: 0x2da0c, // stroke-width
	0x59: 0x35f0d, // solid-opacity
	0x5a: 0xd805,  // color
	0x5b: 0xd20b,  // flood-color
	0x5c: 0x1601,  // g
	0x5d: 0x2e105, // width
	0x5e: 0x1ea0d, // color-profile
	0x61: 0x35c04, // defs
	0x62: 0x1370c, // font-stretch
	0x63: 0x11004, // type
	0x64: 0x8009,  // clip-rule
	0x66: 0x24313, // preserveAspectRatio
	0x67: 0x14507, // feImage
	0x68: 0x36c09, // xml:space
	0x69: 0xc106,  // cursor
	0x6a: 0x16d0f, // image-rendering
	0x6b: 0x90e,   // letter-spacing
	0x6c: 0xf206,  // script
	0x6d: 0x12d0a, // stop-color
	0x6e: 0x101,   // a
	0x70: 0x10b07, // opacity
	0x71: 0xb110,  // contentStyleType
	0x72: 0x1fe0e, // pointer-events
	0x73: 0xb,     // baseProfile
	0x74: 0x11f09, // font-size
	0x75: 0x3270d, // viewport-fill
	0x76: 0x3300c, // fill-opacity
	0x77: 0x25c0b, // white-space
	0x79: 0x17c06, // marker
	0x7b: 0x2a10b, // paint-order
	0x7c: 0xc702,  // y2
	0x7d: 0x32715, // viewport-fill-opacity
	0x7e: 0x1960a, // marker-mid
	0x7f: 0xa011,  // contentScriptType
}
