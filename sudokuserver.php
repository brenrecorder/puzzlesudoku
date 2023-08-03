<?php
$sqlNewTableSudokuHosts = "CREATE TABLE SudokuHost (
id INTEGER AUTO_INCREMENT PRIMARY KEY,
ServerName TEXT,
Players INTEGER,
BoardSolved TEXT,
BoardUnSolved TEXT,
BoardSize TEXT,
BoardDiff TEXT,
Action TEXT,
NicknameP1 TEXT,
NicknameP2 TEXT,
Password TEXT
)";

if (file_exists("sudokuserver.db")) { //CREATE DATABASE
	$db = new SQLite3('sudokuserver.db');
} else {
	$db = new SQLite3('sudokuserver.db');
	$db->querySingle($sqlNewTableSudokuHosts);

	//$db->exec("INSERT INTO CoinSecure(Address, Amount, AmountRGB) VALUES('TotalMarket', '0', 'XXXXXX')");	
	echo "Tables and file sudoku created..";
}

if (!empty($_GET['action'])) { $action = $_GET['action']; } else { $action = ""; }

if ($action == "starthost") {
if (!empty($_GET['nameserver'])) { $nameserver = $_GET['nameserver']; } else { $nameserver = ""; }
$players =1;
if (!empty($_GET['boardsolved'])) { $boardsolved = $_GET['boardsolved']; } else { $boardsolved = ""; }
if (!empty($_GET['boardunsolved'])) { $boardunsolved = $_GET['boardunsolved']; } else { $boardunsolved = ""; }
if (!empty($_GET['boardsolved'])) { $boardsize = $_GET['boardsize']; } else { $boardsize = "9"; }
if (!empty($_GET['boardsolved'])) { $boardsize = $_GET['boardsize']; } else { $boardsize = "9"; }
if (!empty($_GET['boarddiff'])) { $boarddiff = $_GET['boarddiff']; } else { $boarddiff = "0"; }
if (!empty($_GET['password'])) { $password = $_GET['password']; } else { $password = ""; }
if (!empty($_GET['nicknamehost'])) { $nicknamehost = $_GET['nicknamehost']; } else { $nicknamehost = ""; }
$db->exec("INSERT INTO SudokuHost(ServerName, Players, BoardSolved, BoardUnSolved, BoardSize, BoardDiff, Action, NicknameP1, NicknameP2, Password) VALUES('".$nameserver."', '".$players."', '".$boardsolved."', '".$boardunsolved."', '".$boardsize."', '".$boarddiff."', 'starthost', '".$nicknamehost."', '', '".$password."')");
}
if ($action == "clear") {
	if (!empty($_GET['nameserver'])) { 
	$nameserver = $_GET['nameserver'];
	if ($nameserver="allservers") {	$db->exec("DELETE FROM SudokuHost"); } else { $db->exec("DELETE FROM SudokuHost WHERE ServerName=='".$nameserver."'");	 }
	} else { $nameserver = ""; }
}
if ($action == "stophost") {
	if (!empty($_GET['nameserver'])) { $nameserver = $_GET['nameserver']; } else { $nameserver = ""; }
	$db->exec("DELETE FROM SudokuHost WHERE ServerName=='".$nameserver."'");
	echo "STOPPEDHOST:" . $nameserver;
}
if ($action == "joinhost") {
	if (!empty($_GET['nameserver'])) { $nameserver = $_GET['nameserver']; } else { $nameserver = ""; }
	if (!empty($_GET['detectJoin'])) { $detectJoin = true; } else { $detectJoin = false; }

	if ($detectjoin == false) {
	$boardsolved = $db->querySingle("SELECT BoardSolved as bsolved FROM SudokuHost WHERE ServerName=='".$nameserver."'");
	$boardunsolved = $db->querySingle("SELECT BoardUnSolved as bunsolved FROM SudokuHost WHERE ServerName=='".$nameserver."'");
	$boardsize = $db->querySingle("SELECT BoardSize as bsize FROM SudokuHost WHERE ServerName=='".$nameserver."'");
	$boarddiff = $db->querySingle("SELECT BoardDiff as bsize FROM SudokuHost WHERE ServerName=='".$nameserver."'");
	} else {
	$boardsolved = "1";
	$boardunsolved  = "1";
	$boardsize = "1";
	$boarddiff  = "1";
	}
	$players = $db->querySingle("SELECT Players as bsize FROM SudokuHost WHERE ServerName=='".$nameserver."'");
	echo "BOARDSOLVED:" . $boardsolved .":";
	echo "BOARDUNSOLVED:".$boardunsolved.":";
	echo "BOARDSIZE:".$boardsize.":";
	echo "BOARDDIFF:".$boarddiff.":";
	echo "PLAYERS:".$players.":";
	
	if ($players == "1" && $detectJoin == false) { $db->exec("UPDATE SudokuHost SET Players='2' WHERE ServerName='".$nameserver."'"); }
}
if ($action == "joinhost") {
	
}
if ($action == "serverlist") {
$query="SELECT DISTINCT ServerName, Players, BoardDiff, BoardSize FROM SudokuHost";
$result=$db->query($query);
while($row= $result->fetchArray()){
$size = "S";
	if ($row['BoardSize'] == "9") { $size = "S"; }
	if ($row['BoardSize'] == "18") { $size = "M"; }
	if ($row['BoardSize'] == "27") { $size = "L"; }
	echo $row['ServerName'] . "-" . $row['Players'] . "-" . $row['BoardDiff']. "-" . $size .  ":";
}
}

if ($action == "delete") {
		if (!empty($_GET['nameserver'])) { $nameserver = $_GET['nameserver']; } else { $nameserver = ""; }
}
if ($action == "getmovement") {
	if (!empty($_GET['nameserver'])) { $nameserver = $_GET['nameserver']; } else { $nameserver = ""; }
	$lastmovement = $db->querySingle("SELECT Action as lastmove FROM SudokuHost WHERE ServerName='".$nameserver."'");
	if (strlen($lastmovement) > 0) {
	echo "MOVEMENT:" . $lastmovement;
	} else { echo "MOVEMENT:offline"; }
}

if ($action == "setmovement") {
if (!empty($_GET['nameserver'])) { $nameserver = $_GET['nameserver']; } else { $nameserver = ""; }
if (!empty($_GET['player'])) { $player = $_GET['player']; } else { $player = 1; }
if (!empty($_GET['makemove'])) { $makemove = $_GET['makemove']; } else { $makemove = 1; }

$db->exec("UPDATE SudokuHost SET Action='".$makemove."' WHERE ServerName=='".$nameserver."'");
echo "NEWMOVE:" .$player.":". $makemove;
}
?>