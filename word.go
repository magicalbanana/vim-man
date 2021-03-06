package vimman

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Word struct {
	*Entity
	Content            string
	Speed              float64
	Direction          Direction
	CenterHorizontally bool
	Fg                 termbox.Attribute
	Bg                 termbox.Attribute
}

func (w *Word) ShouldCenterHorizontally() bool {
	return w.CenterHorizontally
}

type WordOptions struct {
	InitCallback       func()
	Fg                 termbox.Attribute
	Bg                 termbox.Attribute
	CollidesPhysically bool
	CenterHorizontally bool
	Tags               []Tag
}

func ConvertStringToCells(s string, fg termbox.Attribute, bg termbox.Attribute) []*TermBoxCell {
	var arr []*TermBoxCell

	for i := 0; i < len([]rune(s)); i++ {
		cell := &TermBoxCell{
			Cell: &termbox.Cell{
				[]rune(s)[i],
				fg,
				bg,
			},
			collidesPhysically: false,
			cellData:           TileMapCellData{}}

		arr = append(arr, cell)
	}

	return arr
}

func NewWord(s *Stage, x, y int, content string, options WordOptions) *Word {
	fg, bg, collidesPhysically, centerHorizontally := options.Fg, options.Bg, options.CollidesPhysically, options.CenterHorizontally

	cells := ConvertStringToCells(content, fg, bg)
	entityOptions := EntityOptions{2000, options.Tags, nil}
	e := NewEntity(s, x, y, len(content), 1, ' ', fg, bg, cells, collidesPhysically, entityOptions)
	return &Word{
		Entity:             e,
		Content:            content,
		Direction:          horizontal,
		CenterHorizontally: centerHorizontally,
		Fg:                 fg,
		Bg:                 bg,
	}
}

func DefaultWordOptions() WordOptions {
	return WordOptions{InitCallback: nil, Fg: typedCharacterFg, Bg: typedCharacterBg, CollidesPhysically: false, CenterHorizontally: true}
}

func NewEmptyCharacter(s *Stage, x, y int, options WordOptions) *Word {
	return NewWord(s, x, y, string(" "), options)
}

func (w *Word) Update(s *Stage, event termbox.Event, delta time.Duration) {
}
