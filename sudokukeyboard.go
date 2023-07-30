package main

import (
	"bufio"
	"fmt"
	"os"
		"math/rand"
		"math"
		"strconv"
		"time"
"strings"
  "os/exec"
  "github.com/eiannone/keyboard"
"github.com/fatih/color"
	  "runtime"
)
var SOLVED = make([][]int, 27)
var COMPLETED = make([][]int, 27)
var FilledIN = make([][]int, 27)
var FilledINPlayer1or2 = make([][]int, 27)
var baseUNSOLVED = make([][]int, 27)

var posx int = 0
var posy int = 0

var boardsize int = 27

var players int = 1
var currPlayer int= 1
var scorePlayer1 int =0
var scorePlayer2 int =0
var diff int = 65
var newint int =0

func runCmd(name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func ClearTerminal() {
    switch runtime.GOOS {
    case "darwin":
        runCmd("clear")
    case "linux":
        runCmd("clear")
    case "windows":
        runCmd("cmd", "/c", "cls")
    default:
        runCmd("clear")
    }
}
var start bool =false
func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	
	
fmt.Print("Sudoku Puzzle\n\nHelp:\nUse arrow keys to move\nuse e or number 1-9 keys to select or save new value\nuse f for autofind empty block\n")
fmt.Print("\nPlayers: ")
fmt.Scanln(&players)
if (players > 2) { players=2 }
if (players < 1) { players=1 }
var size string
fmt.Print("\nSize S M , L: " )
fmt.Scanln(&size)
if (size == "S" || size=="s" || size == "") { boardsize = 9 }
if (size == "M" || size == "m") { boardsize = 18 }
if (size == "L" || size == "l") { boardsize = 27 }
fmt.Print("\nDifficulty: " )
fmt.Scanln(&diff)
if (diff < 12) { diff=12 } 
if (diff > 89) { diff=89 }
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

	COMPLETED[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9}
	COMPLETED[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3}
	COMPLETED[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6}
	COMPLETED[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1}
	COMPLETED[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4}
	COMPLETED[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7}
	COMPLETED[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2}
	COMPLETED[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5}
	COMPLETED[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8}
	
	COMPLETED[9] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9}
	COMPLETED[10] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3}
	COMPLETED[11] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6}
	COMPLETED[12] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1}
	COMPLETED[13] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4}
	COMPLETED[14] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7}
	COMPLETED[15] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2}
	COMPLETED[16] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5}
	COMPLETED[17] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8}

	COMPLETED[18] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9}
	COMPLETED[19] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3}
	COMPLETED[20] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6}
	COMPLETED[21] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1}
	COMPLETED[22] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4}
	COMPLETED[23] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7}
	COMPLETED[24] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2}
	COMPLETED[25] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5}
	COMPLETED[26] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8}
	
	COMPLETED = createSolvedSudoku(COMPLETED, diff)

	

//mainmenu()

	FilledIN = createEmptyBoard()
	FilledINPlayer1or2= createEmptyBoard()


		var boardmax int = boardsize-1
	printCurrSudoku(baseUNSOLVED, posx, posy)

	scanner := bufio.NewScanner(os.Stdin)
		keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		keyevent := fmt.Sprintf("%X", event.Key)
		if keyevent == "FFED" { //UP
		start = true
			if (posy < 1) { posy=boardmax } else { posy=posy-1 }
		}
		if keyevent == "FFEC" { //UP
		start = true
			if (posy > (boardmax-1)) { posy=0 } else { posy=posy+1 }
		}
		if keyevent == "FFEB" { //UP
		start = true
			if (posx < 1) { posx=boardmax } else { posx=posx-1 }
		}
		if keyevent == "FFEA" { //UP
		start = true
			if (posx > (boardmax-1)) { posx=0 } else { posx=posx+1 }
		}
		if keyevent == "1B" { //ESC
			main()
		}

var intpressed = -2
if (string(event.Rune) == "f") {
	if (posx >= (boardsize-1)) { 
	posy=posy+1  
	posx=0 
	} else { posx=posx+1 }
	for (posx <= (boardsize-1)) {
	if (baseUNSOLVED[posy][posx] == 0 && FilledIN[posy][posx] == 0) { break }
	if (posx > (boardsize-1)) { 
	posx=0
	if (posy < (boardsize-1)) { 
	posx=0
	posy=posy+1 } else { posy =0 }
	}
	posx++
	}

} else {
if (keyevent =="0" && string(event.Rune) != "" && start ==true) {
intpressed, _ = strconv.Atoi(string(event.Rune))
}
}
		if (intpressed > -1 &&intpressed < 10) {
		fillinnumber(intpressed)
		intpressed=-2
		}
		fmt.Print("\n\n")
		ClearTerminal()
		printCurrSudoku(baseUNSOLVED, posx, posy)
	}
	
	
	for scanner.Scan() {

	textstringin:=scanner.Text()

	if _, err := strconv.Atoi(textstringin); err == nil {
   	pressednumber, _ := strconv.Atoi(textstringin)
fillinnumber(pressednumber)
}

	if (textstringin == "f") {
	if (posx >= 8) { 
	posy=posy+1  
	posx=0 
	} else { posx=posx+1 }
	for (posx < 9) {
	if (baseUNSOLVED[posy][posx] == 0 && FilledIN[posy][posx] == 0) { break }
	if (posx > 8) { 
	posx=0
	if (posy < 8) { posy=posy+1 } else { posy =0 }
	}
	posx++
	}

	}
	amountup := strings.Count(textstringin, "w")
	amountdown := strings.Count(textstringin, "s")
	amountleft := strings.Count(textstringin, "a")
	amountright := strings.Count(textstringin, "d")
	amountedit := strings.Count(textstringin, "e")

	
	for i := 0; i < amountup; i++ {
	if (posy < 1) { posy=boardmax } else { posy=posy-1 }
	}
	for i := 0; i < amountdown; i++ {
	if (posy > (boardmax-1)) { posy=0 } else { posy=posy+1 }
	}
	for i := 0; i < amountleft; i++ {
	if (posx < 1) { posx=boardmax } else { posx=posx-1 }
	}
	for i := 0; i < amountright; i++ {
	if (posx > (boardmax-1)) { posx=0 } else { posx=posx+1 }
	}
	if (amountedit > 0) {
		fmt.Println(strconv.Itoa(FilledIN[posy][posx]) + " - New value: ")

		fmt.Scanln(&newint)
	fillinnumber(newint)
	
	}

		if (scanner.Text() == "q") { os.Exit(0) } 
		fmt.Print("\n\n")
		printCurrSudoku(baseUNSOLVED, posx, posy)
	}

	if scanner.Err() != nil {
		// handle error
	}
}

func fillinnumber(newnr int) {

		if (newnr > -1 && newnr < 10) {
		if (baseUNSOLVED[posy][posx] == 0) { 
		FilledINPlayer1or2[posy][posx] = currPlayer
		FilledIN[posy][posx] = newnr 
		
		}
		if (COMPLETED[posy][posx] == FilledIN[posy][posx]) {
		if (players == 1) { 
		scorePlayer1=scorePlayer1+1 
		}
		if (players == 2) { 
		if (currPlayer == 1) {scorePlayer1=scorePlayer1+1 
		}
		if (currPlayer == 2) {
		scorePlayer2=scorePlayer2+1 }
		}
		} else {
		if (players == 1) { 
		scorePlayer1=scorePlayer1-1 
		}
		if (players == 2) { 
		if (currPlayer == 1) {scorePlayer1=scorePlayer1-1 
		}
		if (currPlayer == 2) {
		scorePlayer2=scorePlayer2-1 }
		}
		}
		if (players == 2) { 
		if (currPlayer == 1) { currPlayer =2 } else { currPlayer = 1}
		}
		}
}
func CheckSolved(b [][]int, filledinb [][]int, solvedsudoku [][]int) bool {
CurrentSoduku := createEmptyBoard()
var done bool = true
	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
	
		if (filledinb[i][j]!=0) { CurrentSoduku[i][j] = filledinb[i][j] } else {	CurrentSoduku[i][j] = b[i][j]  }
		if (CurrentSoduku[i][j] != solvedsudoku[i][j]) { done = false }
		}	
		}
		return done
}
func PrintlineHoriz() {
				for line := 0; line < ((boardsize/9)*30); line++ {
					fmt.Print("-")
				}
				fmt.Print("\n")
}
func printCurrSudoku(b [][]int, posxin int, posyin int) {
	c := exec.Command("cls")
c.Stdout = os.Stdout
c.Run()
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
//main()
main()
} 
//fmt.Println("Player: " + strconv.Itoa(currPlayer) +"\n")
blue := color.New(color.FgBlue)
boldBlue := blue.Add(color.Bold)
red := color.New(color.FgRed)
boldRed := red.Add(color.Bold)
Green := color.New(color.FgGreen)
	if (currPlayer == 1) {boldBlue.Println("Player: " + strconv.Itoa(currPlayer) +"\n") } else { boldRed.Println("Player: " + strconv.Itoa(currPlayer) +"\n")}
	for i := 0; i < boardsize; i++ {
		if (i==0) { PrintlineHoriz()} 
		for j := 0; j < boardsize; j++ {
	
		if (j==0) { fmt.Print("|") }
			if (FilledIN[i][j] != 0) { 
			if (posx == j && posy == i) { Green.Print("[") 
			fmt.Print(strconv.Itoa(FilledIN[i][j])) 
			Green.Print("]") } else {
			//if (FilledINPlayer1or2[i][j] == 2) { fmt.Print(" " + strconv.Itoa(FilledIN[i][j])+ ":")  } else { fmt.Print(" " + strconv.Itoa(FilledIN[i][j])+ ".") }
			if (FilledINPlayer1or2[i][j] == 2) { boldRed.Print(" " + strconv.Itoa(FilledIN[i][j])+ " ")  } else { boldBlue.Print(" " + strconv.Itoa(FilledIN[i][j])+ " ") }
			}
			} else {

			if ((b[i][j]) !=0) {
			if (posx == j && posy == i) { fmt.Print("[" +strconv.Itoa(b[i][j])+ "]") } else {
			fmt.Print(" " + strconv.Itoa(b[i][j])+ " ") } 
			} else {
			if (posx == j && posy == i) { Green.Print("[ ]") } else {
			fmt.Print("   ") } 
			}
			}
			if (j==2 || j==5 || j == 8 || j==11 || j==14 || j == 17 || j==20 || j==23 || j == 26) { fmt.Print("|") }
	
		}
		fmt.Println(" ")
	if (i==2 || i ==5|| i==8 || i==11 || i==14 || i == 17 || i==20 || i==23 || i == 26) { PrintlineHoriz()} 
	}
	fmt.Println(" ")
}
func printBase(b [][]int) {
	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
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
	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
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
	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
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
	base := make([][]int, 27)
	base[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9}
	base[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3}
	base[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6}
	base[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1}
	base[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4}
	base[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7}
	base[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2}
	base[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5}
	base[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8}
	
	base[9] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9}
	base[10] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3}
	base[11] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6}
	base[12] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1}
	base[13] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4}
	base[14] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7}
	base[15] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2}
	base[16] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5}
	base[17] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8}

	base[18] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9,1, 2, 3, 4, 5, 6, 7, 8, 9}
	base[19] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3,4, 5, 6, 7, 8, 9, 1, 2, 3}
	base[20] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6,7, 8, 9, 1, 2, 3, 4, 5, 6}
	base[21] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1,2, 3, 4, 5, 6, 7, 8, 9, 1}
	base[22] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4,5, 6, 7, 8, 9, 1, 2, 3, 4}
	base[23] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7,8, 9, 1, 2, 3, 4, 5, 6, 7}
	base[24] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2,3, 4, 5, 6, 7, 8, 9, 1, 2}
	base[25] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5,6, 7, 8, 9, 1, 2, 3, 4, 5}
	base[26] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8,9, 1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
				base[i][j] = 0
		}
	}
	return base
}

func swapLine(b [][]int) [][]int {
newswap := b
	randInts := [9]int{0, 3, 6, 9, 12, 15, 18, 21, 24}
	r := randInts[rand.Intn(len(randInts))]
	l1 := r + randInt(0, 3) // 0-2
	l2 := r + randInt(0, 3) // 0-2
	newswap[l1], newswap[l2] = newswap[l2], newswap[l1]
return newswap
}

func swapCol(b [][]int) [][]int {
newswap := b
	randInts := [9]int{0, 3, 6, 9, 12, 15, 18, 21, 24}
	r := randInts[rand.Intn(len(randInts))]
	c1 := r + randInt(0, 3) // 0-2
	c2 := r + randInt(0, 3) // 0-2
	for line := 0; line < boardsize; line++ {
		newswap[line][c1], newswap[line][c2] = newswap[line][c2], newswap[line][c1]
	}
	return newswap
}

func swapNumber(b [][]int) [][]int{
newswap := b
	n1 := randInt(1, 10)
	n2 := randInt(1, 10)

	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			if newswap[i][j] == n1 {
				newswap[i][j] = n2
			} else if newswap[i][j] == n2 {
				newswap[i][j] = n1
			}
		}
	}
	return newswap
}