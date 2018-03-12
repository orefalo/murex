package readline

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
)

const (
	modeNormal = iota
	modeTabCompletion
)

var (
	tcPrefix      string
	tcSuggestions []string
	tcGrid        [][]string
	tcPosX        int
	tcPosY        int
	tcMaxX        int
	tcMaxY        int
	tcMaxLength   int
	termWidth     int
)

func tabCompletion() {
	if TabCompleter == nil {
		return
	}

	if mode == modeTabCompletion {
		moveTabHighlight(1, 0)
		return
	}

	tcPrefix, tcSuggestions = TabCompleter(line, pos)
	if len(tcSuggestions) == 0 {
		return
	}

	if len(tcSuggestions) == 1 {
		if len(tcSuggestions[0]) == 0 || tcSuggestions[0] == " " || tcSuggestions[0] == "\t" {
			return
		}
		insert([]byte(tcSuggestions[0]))
		return
	}

	initTabGrid()
	renderSuggestions()
}

func initTabGrid() {
	var err error
	fd := int(os.Stdout.Fd())
	termWidth, _, err = terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	tcMaxLength := 1
	for i := range tcSuggestions {
		if len(tcPrefix+tcSuggestions[i]) > tcMaxLength {
			tcMaxLength = len(tcPrefix + tcSuggestions[i])
		}
	}

	mode = modeTabCompletion
	tcPosX = 1
	tcPosY = 1
	tcMaxX = termWidth / (tcMaxLength + 2)
	tcMaxY = 4
}

func moveTabHighlight(x, y int) {
	//switch x {
	//case -1:
	//}
	tcPosX += x
	tcPosY += y

	if tcPosX < 1 {
		tcPosX = tcMaxX
		tcPosY--
	}

	if tcPosX > tcMaxX {
		tcPosX = 1
		tcPosY++
	}

	if tcPosY < 1 {
		tcPosY = tcMaxY
	}

	if tcPosY > tcMaxY {
		tcPosY = 1
	}
	renderSuggestions()
}

func renderSuggestions() {
	fmt.Print("\r\n")

	cellWidth := strconv.Itoa((termWidth / tcMaxX) - 2)
	x := 0
	y := 1

	for i := range tcSuggestions {
		x++
		if x > tcMaxX {
			x = 1
			y++
			if y > tcMaxY {
				y--
				break
			} else {
				fmt.Print("\r\n")
			}
		}

		if x == tcPosX && y == tcPosY {
			fmt.Print(seqBgWhite + seqFgBlack)
		}
		fmt.Printf(" %-"+cellWidth+"s %s", tcPrefix+tcSuggestions[i], seqReset)
	}

	//fmt.Print(seqPosRestore)
	moveCursorUp(y)
	moveCursorBackwards(termWidth)
	moveCursorForwards(len(Prompt) + pos)
}

func clearTabSuggestions() {
	move := termWidth * tcMaxY
	blank := strings.Repeat(" ", move)

	fmt.Print(seqPosSave + "\r\n" + blank + seqPosRestore)
	mode = modeNormal
}
