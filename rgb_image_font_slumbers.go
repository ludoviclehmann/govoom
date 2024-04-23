package govoom

var FontSlumbers = PixelFont{
	FontName:    "Slumbers",
	Description: "Slim (3x5) font with mathematical symbols and uppercase letters",
	FixedWidth:  true,
	FontSpacing: 1,
	LineHeight:  5,
	Glyphs: map[rune]PixelGlyph{
		// Numbers
		'0': {Pixels: [][]byte{{0, 1, 0}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {0, 1, 0}}},
		'1': {Pixels: [][]byte{{0, 1, 0}, {1, 1, 0}, {0, 1, 0}, {0, 1, 0}, {1, 1, 1}}},
		'2': {Pixels: [][]byte{{1, 1, 0}, {0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {1, 1, 1}}},
		'3': {Pixels: [][]byte{{1, 1, 0}, {0, 0, 1}, {0, 1, 0}, {0, 0, 1}, {1, 1, 0}}},
		'4': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {1, 1, 1}, {0, 0, 1}, {0, 0, 1}}},
		'5': {Pixels: [][]byte{{1, 1, 1}, {1, 0, 0}, {1, 1, 0}, {0, 0, 1}, {1, 1, 0}}},
		'6': {Pixels: [][]byte{{0, 1, 1}, {1, 0, 0}, {1, 1, 0}, {1, 0, 1}, {0, 1, 0}}},
		'7': {Pixels: [][]byte{{1, 1, 1}, {0, 0, 1}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}}},
		'8': {Pixels: [][]byte{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}, {1, 0, 1}, {0, 1, 0}}},
		'9': {Pixels: [][]byte{{0, 1, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}, {1, 1, 0}}},
		// Uppercase Letters
		'A': {Pixels: [][]byte{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}, {1, 0, 1}, {1, 0, 1}}},
		'B': {Pixels: [][]byte{{1, 1, 0}, {1, 0, 1}, {1, 1, 0}, {1, 0, 1}, {1, 1, 0}}},
		'C': {Pixels: [][]byte{{0, 1, 1}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {0, 1, 1}}},
		'D': {Pixels: [][]byte{{1, 1, 0}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 1, 0}}},
		'E': {Pixels: [][]byte{{1, 1, 1}, {1, 0, 0}, {1, 1, 0}, {1, 0, 0}, {1, 1, 1}}},
		'F': {Pixels: [][]byte{{1, 1, 1}, {1, 0, 0}, {1, 1, 0}, {1, 0, 0}, {1, 0, 0}}},
		'G': {Pixels: [][]byte{{0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 0, 1}, {0, 1, 1}}},
		'H': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {1, 1, 1}, {1, 0, 1}, {1, 0, 1}}},
		'I': {Pixels: [][]byte{{1, 1, 1}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {1, 1, 1}}},
		'J': {Pixels: [][]byte{{1, 1, 1}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {1, 1, 0}}},
		'K': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {1, 1, 0}, {1, 0, 1}, {1, 0, 1}}},
		'L': {Pixels: [][]byte{{1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 0, 0}, {1, 1, 1}}},
		'M': {Pixels: [][]byte{{1, 0, 1}, {1, 1, 1}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}}},
		'N': {Pixels: [][]byte{{1, 1, 0}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}}},
		'O': {Pixels: [][]byte{{1, 1, 1}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 1, 1}}},
		'P': {Pixels: [][]byte{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}, {1, 0, 0}, {1, 0, 0}}},
		'Q': {Pixels: [][]byte{{0, 1, 0}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {0, 1, 1}}},
		'R': {Pixels: [][]byte{{1, 1, 0}, {1, 0, 1}, {1, 1, 0}, {1, 0, 1}, {1, 0, 1}}},
		'S': {Pixels: [][]byte{{0, 1, 1}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {1, 1, 0}}},
		'T': {Pixels: [][]byte{{1, 1, 1}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}}},
		'U': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 1, 1}}},
		'V': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {0, 1, 0}}},
		'W': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {1, 0, 1}, {1, 1, 1}, {1, 0, 1}}},
		'X': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {0, 1, 0}, {1, 0, 1}, {1, 0, 1}}},
		'Y': {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}}},
		'Z': {Pixels: [][]byte{{1, 1, 1}, {0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {1, 1, 1}}},
		// Punctuation and special chars from the ascii table
		' ':  {Pixels: [][]byte{{0, 0}}},
		'!':  {Pixels: [][]byte{{1}, {1}, {1}, {0}, {1}}},
		'"':  {Pixels: [][]byte{{1, 0, 1}, {1, 0, 1}}},
		'%':  {Pixels: [][]byte{{1, 0, 1}, {0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {1, 0, 1}}},
		'\'': {Pixels: [][]byte{{1}, {1}}},
		'(':  {Pixels: [][]byte{{0, 1}, {1, 0}, {1, 0}, {1, 0}, {0, 1}}},
		')':  {Pixels: [][]byte{{1, 0}, {0, 1}, {0, 1}, {0, 1}, {1, 0}}},
		'*':  {Pixels: [][]byte{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}},
		'+':  {Pixels: [][]byte{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, Offset: 1},
		',':  {Pixels: [][]byte{{1}, {1}}, Offset: 4},
		'-':  {Pixels: [][]byte{{1, 1, 1}}, Offset: 2},
		'.':  {Pixels: [][]byte{{1}}, Offset: 4},
		'/':  {Pixels: [][]byte{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}, Offset: 1},
		':':  {Pixels: [][]byte{{1}, {0}, {1}}, Offset: 2},
		';':  {Pixels: [][]byte{{1}, {0}, {1}, {1}}, Offset: 2},
		'<':  {Pixels: [][]byte{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
		'=':  {Pixels: [][]byte{{1, 1, 1}, {0, 0, 0}, {1, 1, 1}}, Offset: 1},
		'>':  {Pixels: [][]byte{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {0, 1, 0}, {1, 0, 0}}},
		'?':  {Pixels: [][]byte{{1, 1, 0}, {0, 0, 1}, {0, 1, 0}, {0, 0, 0}, {0, 1, 0}}},
		'[':  {Pixels: [][]byte{{1, 1}, {1, 0}, {1, 0}, {1, 0}, {1, 1}}},
		'\\': {Pixels: [][]byte{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, Offset: 1},
		']':  {Pixels: [][]byte{{1, 1}, {0, 1}, {0, 1}, {0, 1}, {1, 1}}},
		'^':  {Pixels: [][]byte{{0, 1, 0}, {1, 0, 1}}},
		'_':  {Pixels: [][]byte{{1, 1, 1}}, Offset: 4},
		'`':  {Pixels: [][]byte{{1, 0}, {0, 1}}},
		'{':  {Pixels: [][]byte{{0, 0, 1}, {0, 1, 0}, {1, 1, 0}, {0, 1, 0}, {0, 0, 1}}},
		'|':  {Pixels: [][]byte{{1}, {1}, {1}, {1}, {1}}},
		'}':  {Pixels: [][]byte{{1, 0, 0}, {0, 1, 0}, {0, 1, 1}, {0, 1, 0}, {1, 0, 0}}},
		// Others
		'k': {Pixels: [][]byte{{1, 0, 0}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 0, 1}}},
	},
}