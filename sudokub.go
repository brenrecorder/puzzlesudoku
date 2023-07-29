package main

import (
	"fmt"
	"math/rand"
	"time"
"strconv"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
"os"
"math"
)



var text = tview.NewTextView().
    SetTextColor(tcell.ColorGreen).
    SetText("(q) to quit")
	
	var testb string
var SOLVED = make([][]int, 9)
var COMPLETED = make([][]int, 9)
var FilledIN = make([][]int, 9)
var FilledINPlayer1or2 = make([][]int, 9)
var baseUNSOLVED = make([][]int, 9)
var players int = 1
var currPlayer int= 1
var scorePlayer1 int =0
var scorePlayer2 int =0

	var diff int = 15
	
func mainmenu() {
app := tview.NewApplication()

		table := tview.NewTable().
		SetBorders(false)

		table.SetCell(0, 0,
		tview.NewTableCell("  Sudoku  ").

		SetAlign(tview.AlignCenter)).
		SetBorders(false)
	
	
	
	buttonSinglePlayer := tview.NewButton("Single player").SetSelectedFunc(func() {
		players =1
		app.Stop()
		
		return
	})
	button2Players := tview.NewButton("2 players").SetSelectedFunc(func() {
		players =2
		app.Stop()
		return
	})
		buttonExit := tview.NewButton("Exit game").SetSelectedFunc(func() {
		players =1
		app.Stop()
		os.Exit(0)
		return
	})
	
	var formDiff = tview.NewForm()
		
	   formDiff.AddInputField("Difficulty:  ", strconv.Itoa(diff), 10, nil, func(newvaluediff string) {
	      diff, _ = strconv.Atoi(newvaluediff)
		  if (diff>92) { diff = 92 }
		  if (diff<12) { diff = 12 }
	   
	   })
		gridbuttons := tview.NewGrid().
		AddItem(table, 0, 0, 1, 1, 0, 0, true).
		AddItem(formDiff, 1, 0, 1, 1, 0, 0, true).
		AddItem(buttonSinglePlayer, 2, 0, 1, 1, 0, 0, true). 
		AddItem(button2Players, 3, 0, 1, 1, 0, 0, true).
		AddItem(buttonExit, 4, 0, 1, 1, 0, 0, true)
		
	grid := tview.NewGrid().
		SetColumns(35, 35, 5, -1, -2).
		SetMinSize(1, 1).
		SetRows(1, 0).
		SetBorders(false).
		AddItem(gridbuttons, 2, 0, 1, 1, 0, 0, true)

		
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
var boolbooted bool = false
func main() {

fmt.Print("Sudoku Puzzle\n\nHelp:\nUse arrow keys to move\nuse enter key to select or save new value\n\n")

    

baseUNSOLVED = createEmptyBoard()
	rand.Seed(time.Now().UTC().UnixNano())

	base := make([][]int, 9)
	base[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	base[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	base[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	base[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	base[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	base[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	base[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	base[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	base[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8}

	COMPLETED[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	COMPLETED[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	COMPLETED[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	COMPLETED[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	COMPLETED[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	COMPLETED[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	COMPLETED[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	COMPLETED[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	COMPLETED[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8}
	

	COMPLETED = createSolvedSudoku(COMPLETED, diff)

	


mainmenu()

	FilledIN = createEmptyBoard()
	FilledINPlayer1or2= createEmptyBoard()

	
	//printBase(base)

	
	printBaseC(baseUNSOLVED)
	
	
	//solve(base)
	//printBase(base)
}

var posx int
var posy int
var newvar int

func CheckSolved(b [][]int, filledinb [][]int, solvedsudoku [][]int) bool {
CurrentSoduku := createEmptyBoard()

var done bool = true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
	
		if (filledinb[i][j]!=0) { CurrentSoduku[i][j] = filledinb[i][j] } else {	CurrentSoduku[i][j] = b[i][j]  }
		if (CurrentSoduku[i][j] != solvedsudoku[i][j]) { done = false }
		}
		
		}


		return done

}
func printBaseC(b [][]int) {
if (CheckSolved(b, FilledIN, COMPLETED) == true) {
if (players == 1) {
fmt.Println("Puzzle done!\n\nScore: "+strconv.Itoa(scorePlayer1)+"\n\nPress any key to start a new puzzle..")
} else {
var winner string
if (scorePlayer1 > scorePlayer2) { winner = "Player 1 won the game!" }
if (scorePlayer2 > scorePlayer1) { winner = "Player 2 won the game!" }
if (scorePlayer2 == scorePlayer1) { winner = "Player 1 scored equal to Player 2" }
fmt.Println("Puzzle done!\n\n"+winner+"\n\nScore player 1: "+strconv.Itoa(scorePlayer1)+"\nScore player 2: "+strconv.Itoa(scorePlayer2)+"\n\nPress any key to start a new puzzle..")
}
fmt.Scanln()
scorePlayer1 = 0
scorePlayer2 = 0
main()
} 

var intVar int =0
app := tview.NewApplication()
colorb := tcell.ColorBlue
		table := tview.NewTable().
		SetBorders(false)

		
		s, _ := tcell.NewScreen()
		s.EnableMouse()
var posvert int =0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
		//valbutton := strconv.Itoa(b[i][j])
	    //if (i == 0) { formA.AddButton(valbutton, func() { 

		//fmt.Println(strconv.Itoa(j)) }) }
		//if (i == 1) { formB.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
       //if (i == 2) { formC.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
       //if (i == 3) { formD.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
	   //if (i == 4) { formE.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
	   //if (i == 5) { formF.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
		//if (i == 6) { formG.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
		//if (i == 7) { formH.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
	    //if (i == 8) { formI.AddButton(valbutton, func() { fmt.Println(valbutton) }) }
		if (FilledIN[i][j] != b[i][j] && FilledIN[i][j] != 0) { 

				table.SetCell(i, j,
				tview.NewTableCell(strconv.Itoa(FilledIN[i][j])).
					SetTextColor(colorb).
					SetBackgroundColor(tcell.ColorSilver).
					SetAttributes(tcell.AttrBold).
					SetExpansion(5).
					SetAlign(tview.AlignCenter)).
					SetBorders(false)
	
				
		} else {
			var valuecellblock string
			
			if (b[i][j] ==0) { valuecellblock = "" } else {valuecellblock = strconv.Itoa(b[i][j]) }
		
					table.SetCell(i, j,
				tview.NewTableCell(valuecellblock).
				SetBackgroundColor(tcell.ColorSilver).
					SetTextColor(tcell.ColorBlack).
					SetAttributes(tcell.AttrBold).
					SetExpansion(5).
					SetAlign(tview.AlignCenter))
					}
		}
		posvert++
		fmt.Println("\r\n")
	}
	
		buttonNewGame := tview.NewButton("New game").SetSelectedFunc(func() {
		app.Stop()
		main()
	})
	buttonNewGame.SetBorder(false).SetRect(0, 0, 3, 3)
	buttonQuit := tview.NewButton("Quit").SetSelectedFunc(func() {
		app.Stop()
		main()
	})
	buttonQuit.SetBorder(false).SetRect(0, 0, 3, 3)
		gridbuttons := tview.NewGrid().
		AddItem(buttonNewGame, 2, 0, 1, 1, 0, 0, true). 
		AddItem(buttonQuit, 2, 1, 1, 1, 0, 0, true)
		
		var colorplayer = tcell.ColorWhite
		if (currPlayer == 1 && players == 2) { colorplayer = tcell.ColorBlue}
		if (currPlayer == 2 && players == 2) { colorplayer = tcell.ColorRed}
		Players := tview.NewTable().
		SetBorders(false)
		Players.SetCell(0, 0,
		tview.NewTableCell("Player:  "+ strconv.Itoa(currPlayer)).
		SetTextColor(colorplayer).
		SetAlign(tview.AlignLeft)).
		SetBorders(false)
	
		table.Select(posx, posy).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
						var intin int
			fmt.Scanln(&intin)
			if (intin > 9) { intin = 9 }
			FilledIN[posx][posy] = intin
			FilledINPlayer1or2[posx][posy] = currPlayer
			app.Stop()
			printBaseC(b)
			table.SetSelectable(true, true)
		}
		if key == tcell.KeyUp  {
			FilledIN[posx][posy] = newvar +1
			FilledINPlayer1or2[posx][posy] = currPlayer
	printBaseC(b)
		}
	}).SetSelectedFunc(func(row int, column int) {
	if (b[row][column] == 0 || FilledIN[row][column] != 0) {
	
	posx = row
	posy = column
	newvar = FilledIN[row][column]
	var form = tview.NewForm()
		
	   form.AddInputField("Edit " + strconv.Itoa(row)+ ":" + strconv.Itoa(column), "", 10, nil, func(newvalue string) {
	 
	   intVar, _ = strconv.Atoi(newvalue)
	     	if (intVar > 9) { intVar =9 }
		FilledIN[row][column] = intVar
		FilledINPlayer1or2[row][column] = currPlayer
		if (COMPLETED[row][column] == FilledIN[row][column]) {
		if (players == 1) { 
		scorePlayer1=scorePlayer1+1 
		table.GetCell(row, column).SetTextColor(tcell.ColorBlue) 
		}
		if (players == 2) { 
		if (currPlayer == 1) {scorePlayer1=scorePlayer1+1 
		table.GetCell(row, column).SetTextColor(tcell.ColorBlue) 
		}
		if (currPlayer == 2) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed) 
		scorePlayer2=scorePlayer2+1 }
		}
		}
		if (players == 2) { 
		if (currPlayer == 1) { currPlayer =2 } else { currPlayer = 1}
		}
		
    })
	form.SetTitle("Enter new number")
	form.AddButton("Save", func() {
app.Stop()
	printBaseC(b)
	})
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
	if (intVar > 9) { intVar =9 }
	if intVar >0 {
app.Stop()
	printBaseC(b)
		return nil
		}
	}
	return event
})



		gridb := tview.NewGrid().
		SetColumns(35, 35, 5, 15, -2).
		SetMinSize(1, 1).
		
		SetRows(2, 0).
		SetBorders(false).
		AddItem(table, 2, 0, 1, 1, 0, 0, true).
		AddItem(form, 3, 0, 1, 1, 0, 0, true).
		AddItem(Players, 1, 0, 1, 1, 0, 0, true)
		

		if err := app.SetRoot(gridb, true).SetFocus(form).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
		fmt.Println(strconv.Itoa(row))
		
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		}
		//table.SetSelectable(false, false)
	})
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

		//if (FilledIN[i][j] == COMPLETED[i][j]) {  table.GetCell(i, j).SetTextColor(tcell.ColorGreen) } else {
				if (FilledIN[i][j] != 0) { 
				if (FilledINPlayer1or2[i][j]==1) { table.GetCell(i, j).SetTextColor(tcell.ColorBlue) } else { table.GetCell(i, j).SetTextColor(tcell.ColorRed) }
				}
		//}
		}
		}
	fmt.Println(" ")
	table.SetSelectable(true, true)

	
	grid := tview.NewGrid().
		SetColumns(35, 35, 5, -1, -2).
		SetMinSize(1, 1).
		SetRows(3, 0).
		SetBorders(false).
	
		AddItem(table, 2, 0, 1, 1, 0, 0, true).
	AddItem(Players, 1, 0, 1, 1, 0, 0, true).
AddItem(gridbuttons, 3, 0, 1, 1, 0, 0, true)
	
		
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
func printBase(b [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(strconv.Itoa(b[i][j])+ " ")
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

/* Randomly swap rows and columns within the 3x3 borders
 * and swap individual numbers globally. */
func createSolvedSudoku(b [][]int, difficulty int) [][]int{
newSudoku := b
	for i := 0; i < 100; i++ {
		myrand := randInt(0, 100)
		if myrand < 33 {
			newSudoku = swapLine(newSudoku)
		} else if myrand < 66 {
			newSudoku = swapCol(newSudoku)
		} else if myrand < 100 {
			newSudoku = swapNumber(newSudoku)
		}
	}

	r1 := randInt(0, 100) // 0-99
	r2 := randInt(100-difficulty, (100-difficulty)+10) // 50-79
	var cnter float64 = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if r1 > r2 {
				baseUNSOLVED[i][j] = 0
				cnter++
			} else {
			baseUNSOLVED[i][j] = newSudoku[i][j]
			}
			r1 = randInt(0, 100)
		}
	}

    if  (math.Ceil(cnter/2) != (cnter/2) || cnter <1) {
	var booldone bool = false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
		if (baseUNSOLVED[i][j] != 0 && booldone == false) { 
		baseUNSOLVED[i][j] = 0 
		booldone = true
		}
		}
		}
    } 
	return newSudoku
}

/* Erase between 50% and 80% of the numbers to create
 * the unsolved grid with a random difficulty. */
func createUnsolvedSudoku(sudokuin [][]int, percmin int, percmax int) [][]int {

	r1 := randInt(0, 100) // 0-99
	r2 := randInt(percmin, percmax) // 50-79
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if r1 < r2 {
				sudokuin[i][j] = 0
			}
			r1 = randInt(0, 100)
		}
	}
	return sudokuin
}

func createEmptyBoard() [][]int{
	base := make([][]int, 9)
	base[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	base[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	base[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	base[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	base[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	base[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	base[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	base[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	base[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
				base[i][j] = 0
		}
	}
	return base
}

func swapLine(b [][]int) [][]int {
newswap := b
	randInts := [3]int{0, 3, 6}
	r := randInts[rand.Intn(len(randInts))]
	l1 := r + randInt(0, 3) // 0-2
	l2 := r + randInt(0, 3) // 0-2
	newswap[l1], newswap[l2] = newswap[l2], newswap[l1]
return newswap
}

func swapCol(b [][]int) [][]int {
newswap := b
	randInts := [3]int{0, 3, 6}
	r := randInts[rand.Intn(len(randInts))]
	c1 := r + randInt(0, 3) // 0-2
	c2 := r + randInt(0, 3) // 0-2
	for line := 0; line < 9; line++ {
		newswap[line][c1], newswap[line][c2] = newswap[line][c2], newswap[line][c1]
	}
	return newswap
}

func swapNumber(b [][]int) [][]int{
newswap := b
	n1 := randInt(1, 10)
	n2 := randInt(1, 10)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if newswap[i][j] == n1 {
				newswap[i][j] = n2
			} else if newswap[i][j] == n2 {
				newswap[i][j] = n1
			}
		}
	}
	return newswap
}