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
   "net/http"
"io/ioutil"
	  "runtime"
	  "github.com/jpillora/overseer"
	 // "github.com/manifoldco/promptui"
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
var server = "127.0.0.1" //"http://" + server + "/sudokuserver.php?action=stophost&nameserver="+nameserver)
var multiplayermoveold string
var nameserver string
var strBoardSolved string
var strBoardUnsolved string

var nickname1 string
var modeonline string = ""
var boolgamestarted bool = false

var PrintBoardOn bool = false

var sudokulogo string = `
―――――――――――
│ 4     6 │ 
│ 8  5  7 │ 	Sudoku - Beta 1
│    1    │ 
―――――――――――`

func StartServer() {
 currPlayer=1
 nameserver = ""
nickname1 =""
CreateBoard()
strBoardSolved = ""
strBoardUnsolved = ""
 
  

nameserver = TextPrompt("Server name: ")

nickname1 = TextPrompt("Your nickname: ")

	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			strBoardSolved = strBoardSolved + strconv.Itoa(COMPLETED[i][j]) +","
			strBoardUnsolved = strBoardUnsolved + strconv.Itoa(baseUNSOLVED[i][j]) +","
		}
	}
resp, err := http.Get("http://" + server + "/sudokuserver.php?action=stophost&nameserver="+nameserver)
   if err != nil {
      fmt.Println(err)
   }

   resp, err = http.Get("http://" + server + "/sudokuserver.php?action=starthost&nameserver="+nameserver+"&boardsolved="+strBoardSolved+"&boardunsolved="+strBoardUnsolved+"&boardsize="+strconv.Itoa(boardsize)+"&boarddiff="+strconv.Itoa(diff)+"&nicknamehost=" + nickname1)
   if err != nil {
      fmt.Println(err)
   }

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }

   sbody := string(body)
      modeonline = "server"
 
   fmt.Println("\nStarting server.. : " +sbody)

   ClearTerminal()
//printCurrSudokub(baseUNSOLVED, posx, posy)
boolgamestarted = true

   	go func() { 
	for (modeonline == "client" ||  modeonline =="server"){
	
	if (modeonline == "client" || modeonline =="server") {
	GetNewMove()
	time.Sleep(1 * time.Second)
	}
	if (modeonline != "client" && modeonline != "server") {
	//ResetVars()
	keyboard.Close()
	return } 
	} }()
   ClearTerminal()
   PrintBoardOn = true
printCurrSudokub(baseUNSOLVED, posx, posy)
return
}

func detectJoinPlayer() {
var PlayersOnline int = 1 
 fmt.Println(" - Awaiting joining player") 
for {
   resp, err := http.Get("http://" + server + "/sudokuserver.php?action=joinhost&detectJoin=1&nameserver="+nameserver)
   if err != nil {
      fmt.Println(err)
	  break
   }
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }
   sbody := string(body)
   ReceiveJoin := strings.Split(sbody, ":")
   ServerPlayers, _ := strconv.Atoi(ReceiveJoin[9])
   PlayersOnline = ServerPlayers

   if (PlayersOnline > 1) { break }
   time.Sleep(4*time.Second)
   }
   if (PlayersOnline == 2) { fmt.Println(" - Player 2 joined server") }
   return
}
func JoinServer() {

CreateBoard()
strBoardSolved = ""
strBoardUnsolved = ""
currPlayer=1
nameserver=""
fmt.Print("\n")

nameserver = TextPrompt("Join server name: ")

   resp, err := http.Get("http://" + server + "/sudokuserver.php?action=joinhost&nameserver="+nameserver)
   if err != nil {
      fmt.Println(err)
   }

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }
	//echo "BOARDSOLVED:" . $boardsolved .":";
	//echo "BOARDUNSOLVED:".$boardunsolved.":";
   sbody := string(body)
   ReceiveJoin := strings.Split(sbody, ":")
    //ReceiveJoin[1] = BOARDSOLVED
	//ReceiveJoin[3] = BOARDUNSOLVED
   BoardSolved := strings.Split(ReceiveJoin[1], ",")
   BoardUnsolved := strings.Split(ReceiveJoin[3], ",")
   boardsize, _ = strconv.Atoi(ReceiveJoin[5])
   ServerPlayers, _ := strconv.Atoi(ReceiveJoin[9])
   if (ServerPlayers != 1) { 
   fmt.Print("Server full")
   //MenuA()
   }
  FilledIN = createEmptyBoard()
   var cnterx int =0
   var cntery int =0
   for i, strSolved := range BoardSolved {
   if (len(strSolved) > 0) {
	COMPLETED[cntery][cnterx], _ = strconv.Atoi(strSolved)
	if (cnterx < boardsize) { cnterx++ } else { }
	if (cnterx > (boardsize-1)) { 
	cnterx =0
	cntery++ } else { }
	}
	fmt.Print(i)
	}
	cnterx=0
	cntery=0
   for i, strUnsolved := range BoardUnsolved {
      if (len(strUnsolved) > 0) {
	baseUNSOLVED[cntery][cnterx], _ = strconv.Atoi(strUnsolved)
	if (cnterx < boardsize) { cnterx++ } else { }
	if (cnterx > (boardsize-1)) { 
	cnterx =0
	cntery++ } else { }
	}
	fmt.Print(i)
	}
	fmt.Println("\n\nLoaded sudoku board from " + nameserver)

	 fmt.Println("\nJoining server.. : " +nameserver)
if (len(BoardSolved) <1 ) { 
fmt.Print("Error connecting..")

}


modeonline = "client"
ClearTerminal()
time.Sleep(2*time.Second)
PrintBoardOn = true
printCurrSudokub(baseUNSOLVED, posx, posy)
boolgamestarted = true
	go func() { 
	for (modeonline == "client" ||  modeonline =="server"){
	
	if (modeonline == "client" || modeonline =="server") {
	GetNewMove()
	time.Sleep(1 * time.Second)
	}
	if (modeonline != "client" && modeonline != "server") { break } 
	} }()
ReceiveKeyboard()
}
var boolNewMoverecieved bool = false
var boolquit = false
func GetNewMove() {
  resp, err := http.Get("http://" + server + "/sudokuserver.php?action=getmovement&nameserver="+nameserver)
   if err != nil {
      fmt.Println(err)
	    resp, _ = http.Get("http://" + server + "/sudokuserver.php?action=getmovement&nameserver="+nameserver)
   }

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }
   
    sbody := string(body)
	
	movement := strings.Split(sbody, ":")

	if (movement[1] != "starthost" && movement[1] != "offline") {
	if ((movement[1] + ":"+ movement[2]+ ":"+ movement[3]) != multiplayermoveold) {

	posx,_=strconv.Atoi(movement[1])
	posy,_=strconv.Atoi(movement[2])
	newvalmultiplayer,_:=strconv.Atoi(movement[3])
	fillinnumber(newvalmultiplayer)
	boolNewMoverecieved = true
	ClearTerminal()
	printCurrSudokub(baseUNSOLVED, posx, posy)
	multiplayermoveold = movement[1] + ":"+ movement[2]+ ":"+ movement[3]
	}
	}
	if (movement[1] == "offlinebb") {
	modeonline = ""
	nameserver = ""
	var winner string

if (scorePlayer1 > scorePlayer2) { winner = "Player 1 won the game!" }
if (scorePlayer2 > scorePlayer1) { winner = "Player 2 won the game!" }
if (scorePlayer2 == scorePlayer1) { winner = "Player 1 scored equal to Player 2" }
	fmt.Println("Game finished\n\n"+winner+"\n\nServer disconnected..")
	fmt.Scanln()
	ResetVars()
	main()
	}
	//echo "MOVEMENT:" . $lastmovement;


}
func MakeNewMove(posx int, posy int, newval int) {

var newmovement string = strconv.Itoa(posx) + ":" + strconv.Itoa(posy) + ":" + strconv.Itoa(newval)
if (newmovement != multiplayermoveold) { multiplayermoveold = newmovement } else {}
var player int = 0
if (modeonline =="client") { player=2 } 
if (modeonline =="server") { player=1 } 
//fmt.Println("http://" + server + "/sudokuserver.php?action=setmovement&nameserver="+nameserver+"&player="+strconv.Itoa(player)+"&makemove="+newmovement)

  resp, err := http.Get("http://" + server + "/sudokuserver.php?action=setmovement&nameserver="+nameserver+"&player="+strconv.Itoa(player)+"&makemove="+newmovement)
   if err != nil {
      fmt.Println(err)
   }

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }
   
    sbody := string(body)
	fmt.Print(sbody)
} 

func Multiplayer() {
var modeselect  string
fmt.Print("\nType H to start server\nType J to join server: \nType L for server list\nO: for offline\n" )
fmt.Scanln(&modeselect)
if (len(modeselect) > 0) {
if (modeselect == "H" || modeselect == "h") { StartServer() }
if (modeselect == "J" || modeselect == "j") { JoinServer() }
if (modeselect == "O" || modeselect == "o") { return }
if (modeselect == "L" || modeselect == "l") {
  resp, err := http.Get("http://" + server + "/sudokuserver.php?action=serverlist")
   if err != nil {
      fmt.Println(err)
   }
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }
    sbody := string(body)
	serverlist := strings.Split(sbody, ":")
	fmt.Println("Server \t\t Players")
	   for _, server := range serverlist {
	   	serverdata := strings.Split(server, "-")
	   if (len(serverdata[0])>0) { fmt.Println(serverdata[0] + "\t\t" + serverdata[1]) }
	   }
	fmt.Scanln(&modeselect)
	Multiplayer()
 }
} else {
return
}
}
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
var boolLoaded bool =false
func CreateBoard() {

baseUNSOLVED = createEmptyBoard()

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

	FilledIN = createEmptyBoard()
	FilledINPlayer1or2= createEmptyBoard()
}
func ResetVars() {
 SOLVED = make([][]int, 27)
 COMPLETED = make([][]int, 27)
 FilledIN = make([][]int, 27)
 FilledINPlayer1or2 = make([][]int, 27)
 baseUNSOLVED = make([][]int, 27)

 posx  = 0
 posy  = 0

 boardsize  = 27

 players  = 1
 currPlayer = 1
 scorePlayer1  =0
 scorePlayer2  =0
 diff  = 65
 newint  =0
 //server = "127.0.0.1"

 //modeonline = ""
 keyboard.Close()
 boolgamestarted = false
}
func Substring(str string, start, end int) string {
   return strings.TrimSpace(str[start:end])
}

func TextPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
	intreturn :=strings.TrimSpace(s)
    return intreturn
}
func DecimalPrompt(label string) int {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
	intreturn, _ :=strconv.Atoi(strings.TrimSpace(s))
    return intreturn
}

func MenuA() {
PrintBoardOn = false
var menu1 int =-1 
var menu2 int =-1
fmt.Print(modeonline)

if (modeonline == "server") {
modeonline = ""
_, err := http.Get("http://" + server + "/sudokuserver.php?action=stophost&nameserver="+nameserver)
   if err != nil {
      fmt.Println(err)
   }
ResetVars()
}
if (modeonline == "client") {
modeonline = ""
ResetVars()
}
modeonline = ""
_ = keyboard.Close()
ResetVars()
boolgamestarted = false


ClearTerminal()
modeonline = ""

boolgamestarted = false

fmt.Print(sudokulogo +"\tMain menu" +modeonline+"\n\n1:		Singleplayer\n2:		Multiplayer\n3:		Exit\n\n")
	//prompt := promptui.Select{
	//	Label: "Menu",
	//	Items: []string{"Single player", "Multi player", "Exit"},
	//}
	//
	//i,_, _ := prompt.Run()

menu1=DecimalPrompt("Menu: ")
if (menu1 == 1) {
ClearTerminal()
var size string
fmt.Print(sudokulogo + "\tSingle player\n\n1:		1 players\n2:		2 players\n3:		Return\n\n")

menuSP:=DecimalPrompt("Menu: ")
if (menuSP == 1) { players = 1 }
if (menuSP == 2) { players = 2 }
if (menuSP == 3) { MenuA() }
fmt.Print("Size S M , L: " )
fmt.Scanln(&size)
if (size == "S" || size=="s" || size == "") { boardsize = 9 }
if (size == "M" || size == "m") { boardsize = 18 }
if (size == "L" || size == "l") { boardsize = 27 }

diff=DecimalPrompt("Difficulty: ")
if (diff < 12) { diff=12 } 
if (diff > 89) { diff=89 }
fmt.Print("\nPress any key to start game.." )
modeonline = "offline"
CreateBoard()
boolgamestarted = true
PrintBoardOn = true
ReceiveKeyboard()
}
if (menu1 == 2) {
menu2=-1
ResetVars()
ClearTerminal()
nameserver=""
modeonline = ""
fmt.Print(sudokulogo + "\tMultiplayer "+modeonline+"\n\n1:		Host game\n2:		Join game\n3:		Return\n\n")
	//prompt = promptui.Select{
	//	Label: "Menu",
	///	Items: []string{"Host game", "Join game", "Return"},
	//}
	//
	//i,_, _ := prompt.Run()

menu2=DecimalPrompt("Menu: ")
if (menu2 == 1) {
ResetVars()
nameserver=""
modeonline = "server"
ClearTerminal()
diff =0
currPlayer =1
players=2
boardsize =9
var size string
fmt.Print("Sudoku Puzzle - Host online game\n" )

size=TextPrompt("Size S,M, L: ")
if (size == "S" || size=="s" || size == "") { boardsize = 9 }
if (size == "M" || size == "m") { boardsize = 18 }
if (size == "L" || size == "l") { boardsize = 27 }

diff=DecimalPrompt("Difficulty: ")
if (diff < 12) { diff=12 } 
if (diff > 89) { diff=89 }

players=2
currPlayer =1

StartServer()
detectJoinPlayer()
ReceiveKeyboard()
overseer.Restart()
} else if (menu2 == 2) {
ResetVars()
nameserver=""
modeonline = "client"
ClearTerminal()

diff =0
currPlayer =1
players=2
boardsize =9
fmt.Print("Sudoku Puzzle - Find online sudoku game\n" )
  resp, err := http.Get("http://" + server + "/sudokuserver.php?action=serverlist")
   if err != nil {
      fmt.Println(err)
   }
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
   }
    sbody := string(body)
	serverlist := strings.Split(sbody, ":")
	fmt.Println("\nServer\t\t\t\tPlayers\t\tDifficulty\tSize board")
	   for _, server := range serverlist {
	   
	   	serverdata := strings.Split(server, "-")
		var nameserverin string = serverdata[0]
		if (len(serverdata[0])>0 && serverdata[1] == "2") { continue }
	   if (len(serverdata[0])>0) { 
		   
	   	if (len(nameserverin) < 20) {
		var i int = len(nameserverin)
		fmt.Print(serverdata[0])
		for (i < 20) { 
		fmt.Print(" ")
		i++
		}
		if (serverdata[2] == "") { serverdata[2] = "?" }
		if (serverdata[3] == "") { serverdata[3] = "?" }
		} else {fmt.Print(Substring(nameserverin, 0, 20))}
	   fmt.Print("\t\t" + serverdata[1] +"\t\t"+ serverdata[2]+"\t\t"+ serverdata[3] +"\n") }
	   }
	   


JoinServer()
//ReceiveKeyboard()
overseer.Restart()

}else if (menu2 == 3) {
MenuA()
}
}
if (menu1 == 3) {
os.Exit(0)
}
return
}
func main() {
//addressint := strconv.Itoa(randInt(5000, 6000))
	overseer.Run(overseer.Config{
		Program: sudokumain,
		//Address: ":" +addressint,
		Debug: false,
	})
}
func sudokumain(state overseer.State){
if (modeonline == "") {
fmt.Print("a")
keyboard.Close()
ResetVars()
ClearTerminal()
modeonline = "menu"
 MenuA()  


	fmt.Print("b")
	}
	//if (modeonline != "" && modeonline != "menu") {  }

 


}

func ReceiveKeyboard() {
if (modeonline != "" && modeonline != "menu") {
 	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	
	} else {
	_ = keyboard.Close()
	}
var boardmax int = boardsize-1
	for (len(modeonline) >0  && modeonline != "menu"){
			keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Println(err)
	}
		event := <-keysEvents
		if event.Err != nil {
			fmt.Println(event.Err)
		}
if (modeonline == "" || modeonline == "menu") {
keyboard.Close()
 break } 
		keyevent := fmt.Sprintf("%X", event.Key)

if (string(event.Rune) == "w") { 
if (posy < 1) { posy=boardmax } else { posy=posy-1 } 
start = true
}
if (string(event.Rune) == "s") { 
if (posy > (boardmax-1)) { posy=0 } else { posy=posy+1 } 
start = true
}
if (string(event.Rune) == "a") {
 if (posx < 1) { posx=boardmax } else { posx=posx-1 } 
 start = true
 }
if (string(event.Rune) == "d") {
 if (posx > (boardmax-1)) { posx=0 } else { posx=posx+1 } 
 start = true
 }
		if keyevent == "FFED" { //UP
		start = true
			if (posy < 1) { posy=boardmax } else { posy=posy-1 }
			
				for i := posy; i > -1; i-- {
	if (baseUNSOLVED[i][posx] == 0) { 
	posy=i
	break
	}
	}
			
		}
		if keyevent == "FFEC" { //DOWN
		start = true
			if (posy > (boardmax-1)) { posy=0 } else { posy=posy+1 }
	for i := posy; i < boardsize; i++ {
	if (baseUNSOLVED[i][posx] == 0) { 
	posy=i
	break
	}
	}
		}
		if keyevent == "FFEB" { //LEFT
		start = true
			if (posx < 1) { posx=boardmax } else { posx=posx-1 }
				for i := posx; i > -1; i-- {
	if (baseUNSOLVED[posy][i] == 0) { 
	posx=i
	break
	}
	}

		}
		if keyevent == "FFEA" { //RIGHT
		start = true
			if (posx > (boardmax-1)) { posx=0 } else { posx=posx+1 }
				for i := posx; i < boardsize; i++ {
	if (baseUNSOLVED[posy][i] == 0) { 
	posx=i
	break
	}
	}
		}
		if keyevent == "1B" { //ESC
		if (modeonline == "server" && len(nameserver)>1) { http.Get("http://" + server + "/sudokuserver.php?action=stophost&nameserver="+nameserver) }
		modeonline = ""
		
		break
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
if checkIsNumber(string(event.Rune)) { intpressed, _ = strconv.Atoi(string(event.Rune)) } else {intpressed=-2}

}
}
if (modeonline == "client") {
if (currPlayer == 2) { 
	if (intpressed > -1 &&intpressed < 10) {
	MakeNewMove(posx,posy, intpressed)
	boolNewMoverecieved = false
	}
} else { 

for (boolNewMoverecieved == false) {}
}
}
if (modeonline == "server") {
if (currPlayer == 1) {
	if (intpressed > -1 &&intpressed < 10) {
	MakeNewMove(posx,posy, intpressed)
	boolNewMoverecieved = false
	}
} else { 
for (boolNewMoverecieved == false) {}
 }

}
fmt.Print("\n\n")
		if (intpressed > -1 &&intpressed < 10 && boolgamestarted == true) {
		fillinnumber(intpressed)
		intpressed=-2
		//start = false
		}
		
		if (modeonline != "" && PrintBoardOn == true) {
	
		ClearTerminal()
		printCurrSudokub(baseUNSOLVED, posx, posy) } else {
		
		}
	}
	keyboard.Close()
	modeonline = ""
	boolgamestarted = false
	return
}
func checkIsNumber(s string) bool {
    for _, c := range s {
        if c < '0' || c > '9' {
            return false
        }
    }
    return true
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
func PrintlineHorizb() string {
var lineout string = "―"
				for line := 0; line < ((boardsize/9)*30); line++ {
					lineout = lineout + "―"
				}
				lineout = lineout + "\n"
				
				return lineout
}
//var SCREENNEWVALUES bool = false

func printCurrSudokub(b [][]int, posxin int, posyin int) {
	c := exec.Command("cls")
c.Stdout = os.Stdout
c.Run()
var strScreenout string = ""
if (modeonline != ""){
if (CheckSolved(b, FilledIN, COMPLETED) == true) {
if (players == 1) {
fmt.Println("Puzzle done!\n\nScore: "+strconv.Itoa(scorePlayer1)+"\n\nPress any key to start a new puzzle..")
fmt.Scanln()
} else {
var winner string
var whowon string = ""
if (currPlayer == 1 && modeonline == "server" && scorePlayer1 > scorePlayer2) { whowon = "You won the game!\n" }
if (currPlayer == 2 && modeonline == "client" && scorePlayer2 > scorePlayer1) { whowon = "You won the game!\n" }
if (currPlayer == 1 && modeonline == "server" && scorePlayer1 < scorePlayer2) { whowon = "You lost the game!\n" }
if (currPlayer == 2 && modeonline == "client" && scorePlayer2 < scorePlayer1) { whowon = "You lost the game!\n" }
fmt.Print(whowon)
if (scorePlayer1 > scorePlayer2) { winner = "Player 1 won the game!" }
if (scorePlayer2 > scorePlayer1) { winner = "Player 2 won the game!" }
if (scorePlayer2 == scorePlayer1) { winner = "Player 1 scored equal to Player 2" }
fmt.Println("Puzzle done!\n\n"+winner+"\n\nScore player 1: "+strconv.Itoa(scorePlayer1)+"\nScore player 2: "+strconv.Itoa(scorePlayer2)+"\n\nPress any key to start a new puzzle..")
keyboard.Close()
fmt.Scanln()
if (modeonline == "client" || modeonline == "server") {
http.Get("http://" + server + "/sudokuserver.php?action=stophost&nameserver="+nameserver)
boolgamestarted = false
modeonline = ""
ResetVars()
main()
}
}



ResetVars()
scorePlayer1 = 0
scorePlayer2 = 0
//main()
boolgamestarted = false
modeonline = ""
ResetVars()
main()
}
} else {
ResetVars()
modeonline = ""
keyboard.Close()
main()
}
//fmt.Println("Player: " + strconv.Itoa(currPlayer) +"\n")
var connectedserver string
if (modeonline == "client" || modeonline == "server") {
if (currPlayer == 1 && modeonline == "server") { connectedserver = " (you)" }
if (currPlayer == 2 && modeonline == "client") { connectedserver = " (you)" }
connectedserver = connectedserver + " server: " + nameserver + "\n" 
} else { connectedserver = ""}
if (modeonline != "") {
	if (currPlayer == 1) {strScreenout = strScreenout +"Player: " + strconv.Itoa(currPlayer) + connectedserver +"\n" } else { strScreenout = strScreenout +"Player: " + strconv.Itoa(currPlayer)+ connectedserver +"\n"}
	for i := 0; i < boardsize; i++ {
		if (i==0) { 	strScreenout = strScreenout + PrintlineHorizb()} 
		for j := 0; j < boardsize; j++ {
	
		if (j==0) { strScreenout = strScreenout +"│" }
			if (FilledIN[i][j] != 0) { 
			if (posx == j && posy == i) { strScreenout = strScreenout +"["
			strScreenout = strScreenout + strconv.Itoa(FilledIN[i][j]) 
			strScreenout = strScreenout + "]" } else {
			//if (FilledINPlayer1or2[i][j] == 2) { fmt.Print(" " + strconv.Itoa(FilledIN[i][j])+ ":")  } else { fmt.Print(" " + strconv.Itoa(FilledIN[i][j])+ ".") }
			if (FilledINPlayer1or2[i][j] == 2) {strScreenout = strScreenout +" "  + strconv.Itoa(FilledIN[i][j])+ ":" } else { strScreenout = strScreenout +" " + strconv.Itoa(FilledIN[i][j])+ "." }
			}
			} else {

			if ((b[i][j]) !=0) {
			if (posx == j && posy == i) { strScreenout = strScreenout + "[" +strconv.Itoa(b[i][j])+ "]" } else {
			strScreenout = strScreenout +" " + strconv.Itoa(b[i][j])+ " " } 
			} else {
			if (posx == j && posy == i) { strScreenout = strScreenout +"[ ]" } else {
			strScreenout = strScreenout +"   " } 
			}
			}
			if (j==2 || j==5 || j == 8 || j==11 || j==14 || j == 17 || j==20 || j==23 || j == 26) { strScreenout = strScreenout +"│"}
	
		}
		strScreenout = strScreenout +"\n"
	if (i==2 || i ==5|| i==8 || i==11 || i==14 || i == 17 || i==20 || i==23 || i == 26) { strScreenout = strScreenout + PrintlineHorizb()} 
	}
	strScreenout = strScreenout +""
	fmt.Print(strScreenout)
	}
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
rand.Seed(time.Now().UTC().UnixNano())
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