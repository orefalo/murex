package readline

import (
	"fmt"
	"strconv"
	"strings"
)

func (rl *instance) tabCompletion() {
	if rl.TabCompleter == nil {
		return
	}

	rl.tcPrefix, rl.tcSuggestions = rl.TabCompleter(rl.line, rl.pos)
	if len(rl.tcSuggestions) == 0 {
		return
	}

	if len(rl.tcSuggestions) == 1 {
		if len(rl.tcSuggestions[0]) == 0 || rl.tcSuggestions[0] == " " || rl.tcSuggestions[0] == "\t" {
			return
		}
		insert([]byte(rl.tcSuggestions[0]))
		return
	}

	rl.initTabGrid()
	rl.renderSuggestions()
}

func (rl *instance) initTabGrid() {
	rl.getTermWidth()

	tcMaxLength := 1
	for i := range rl.tcSuggestions {
		if len(rl.tcPrefix+rl.tcSuggestions[i]) > tcMaxLength {
			tcMaxLength = len([]rune(rl.tcPrefix + rl.tcSuggestions[i]))
		}
	}

	rl.modeTabGrid = true
	rl.tcPosX = 1
	rl.tcPosY = 1
	rl.tcMaxX = termWidth / (tcMaxLength + 2)
	rl.tcMaxY = rl.MaxTabCompleterRows
}

func (rl *instance) moveTabHighlight(x, y int) {
	rl.tcPosX += x
	rl.tcPosY += y

	if rl.tcPosX < 1 {
		rl.tcPosX = rl.tcMaxX
		rl.tcPosY--
	}

	if rl.tcPosX > rl.tcMaxX {
		rl.tcPosX = 1
		rl.tcPosY++
	}

	if rl.tcPosY < 1 {
		rl.tcPosY = rl.tcUsedY
	}

	if rl.tcPosY > rl.tcUsedY {
		rl.tcPosY = 1
	}

	if rl.tcPosY == rl.tcUsedY && (rl.tcMaxX*(rl.tcPosY-1))+rl.tcPosX > len(rl.tcSuggestions) {
		if x < 0 {
			rl.tcPosX = len(rl.tcSuggestions) - (rl.tcMaxX * (rl.tcPosY - 1))
		}

		if x > 0 {
			rl.tcPosX = 1
			rl.tcPosY = 1
		}

		if y < 0 {
			rl.tcPosY--
		}

		if y > 0 {
			rl.tcPosY = 1
		}
	}

	rl.renderSuggestions()
}

func (rl *interface)renderSuggestions() {
	newlines := strings.Repeat("\r\n", rl.hintY+1)
	//fmt.Print("\r\n")
	//moveCursorDown(hintY )
	fmt.Print(newlines)

	cellWidth := strconv.Itoa((rl.termWidth / rl.tcMaxX) - 2)
	x := 0
	y := 1

	for i := range rl.tcSuggestions {
		x++
		if x > rl.tcMaxX {
			x = 1
			y++
			if y > rl.tcMaxY {
				y--
				break
			} else {
				fmt.Print("\r\n")
			}
		}

		if x == rl.tcPosX && y == rl.tcPosY {
			fmt.Print(seqBgWhite + seqFgBlack)
		}
		fmt.Printf(" %-"+cellWidth+"s %s", rl.tcPrefix+rl.tcSuggestions[i], seqReset)
	}

	rl.tcUsedY = y
	moveCursorUp(y + rl.hintY)
	moveCursorBackwards(termWidth)
	moveCursorForwards(rl.promptLen + rl.pos)
}

func (rl *instance)clearTabSuggestions() {
	move := rl.termWidth * rl.tcUsedY
	blank := strings.Repeat(" ", move)

	// It's ugly but required as we don't know the absolute position of the cursor
	fmt.Print("\r\n")
	moveCursorDown(rl.hintY)
	fmt.Print(blank)
	moveCursorBackwards(rl.termWidth)
	moveCursorUp(rl.hintY + rl.tcUsedY)
	moveCursorForwards(rl.promptLen + rl.pos)

	rl.modeTabGrid = false
}
