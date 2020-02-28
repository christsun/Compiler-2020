package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nats-io/nuid"
	_ "github.com/nats-io/nuid"
	"net/http"
)

var n = nuid.New()
var user_nuid = nuid.New()

var semanticPool []JudgePoolElement
var codegenPool []JudgePoolElement
var optimizePool []JudgePoolElement
var compilePool = make(map[string]JudgePoolElement)
var compilePoolCounter = 0
var db, _ = sql.Open("mysql", "client:password1A@tcp(127.0.0.1:3306)/compiler")
var db2, _ = sql.Open("mysql", "client:password1A@tcp(127.0.0.1:3306)/compiler")

func main() {
	db, err := sql.Open("mysql", "client:password1A@tcp(127.0.0.1:3306)/compiler")
	if err != nil {
		fmt.Printf("runtime Error: %s", err.Error())
		return
	}
	if db == nil {
		fmt.Printf("runtime Error: Database open failed.(db is nil)")
		return
	}
	defer db.Close()
	db2, err2 := sql.Open("mysql", "client:password1A@tcp(127.0.0.1:3306)/compiler")
	if err2 != nil {
		fmt.Printf("runtime Error: %s", err2.Error())
		return
	}
	if db2 == nil {
		fmt.Printf("runtime Error: Database open failed.(db is nil)")
		return
	}
	defer db2.Close()
	http.HandleFunc("/fetchRepo", getUserList)           // Get test ok!
	http.HandleFunc("/fetchTask", getTask)               // Get
	http.HandleFunc("/addUser", addUser)                 // Post test ok!
	http.HandleFunc("/requestJudge", reqJudge)           // Post
	http.HandleFunc("/addDataSemantic", addDataSemantic) // Post semantic data test ok!
	http.HandleFunc("/addDataCodegen", addDataCodegen)   // Post
	http.HandleFunc("/fetchStatus", getJudgeStatus)      // Post
	http.HandleFunc("/submitTask", submitTask)
	http.HandleFunc("/loginUser", loginUser)
	http.HandleFunc("/queryID", queryID)
	http.HandleFunc("/fetchRepoWeb", getUserListWeb)
	http.HandleFunc("/fetchBuild", fetchCompileTask)
	http.HandleFunc("/submitBuild", submitBuildTask)
	fmt.Print("Start to serve\n")
	http.ListenAndServe(":10430", nil)
}
