package main

import (
  "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/msft/bank"
)

var accounts = map[float64]*bank.Account{}

func main(){
  accounts[1001] = &bank.Account{
    Customer: bank.Customer{
      Name: "John",
      Address: "Los Angeles, California", 
      Phone: "123-456-7890",
    },
    Number: 1001,
  } 
  accounts[1002] = &bank.Account{
    Customer: bank.Customer{
      Name: "Jane",
      Address: "Los Angeles, California",
      Phone: "123-456-7790",
    },
    Number: 1002,
  }

  http.HandleFunc("/statement", statement)
  http.HandleFunc("/deposit", deposit)
  http.HandleFunc("/withdraw", withdraw)
  http.HandleFunc("/transfer", transfer)
  http.ListenAndServe("localhost:8000", nil)
}

//statement
func statement(w http.ResponseWriter, req *http.Request) {
  numberqs := req.URL.Query().Get("number")

  if numberqs == "" {
    fmt.Fprintf(w, "Account number is missing!")
    return
  }

  if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
    fmt.Fprintf(w, "Invalid account number!")
  }else{
    account, ok := accounts[number]
    if !ok {
      fmt.Fprintf(w, "Account with number %v can't be found", number)
    }else{
      json.NewEncoder(w).Encode(bank.Statement(account))
    }
  }
}

//deposit
func deposit(w http.ResponseWriter, req *http.Request){
  numberqs := req.URL.Query().Get("number")
  amountqs := req.URL.Query().Get("amount")

  if numberqs == "" {
    fmt.Fprintf(w, "Account number is missing!")
    return
  }
  
  if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
    fmt.Fprintf(w, "Invalid account number!")
  } else if amount, err := strconv.ParseFloat(amountqs,  64); err != nil {
    fmt.Fprintf(w, "Invalid amount number!")
  }else{
    account, ok := accounts[number]
    if !ok {
      fmt.Fprintf(w, "Account with number %v can't be found", number)
    }else{
      err := account.Deposit(amount)
      if err != nil {
        fmt.Fprintf(w, "%v", err)
      } else {
        fmt.Fprintf(w, account.Statement())
      }
    }
  }
}

//Withdraw
func withdraw(w http.ResponseWriter, req *http.Request){
  numberqs := req.URL.Query().Get("number")
  amountqs := req.URL.Query().Get("amount")
  
  if numberqs == "" {
    fmt.Fprintf(w, "Account number is missing!")
    return
  }

  if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
    fmt.Fprintf(w, "Invalid account number!")
  }else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
    fmt.Fprintf(w, "Invalid amount number!")
  }else {
    account, ok := accounts[number]
    if !ok {
      fmt.Fprintf(w, "Account with number %v can't be found", number)
    }else {
      err := account.Withdraw(amount)
      if err != nil {
        fmt.Fprintf(w, "%v", err)
      } else {
        fmt.Fprintf(w, account.Statement())
      }
    }
  }
}

func transfer(w http.ResponseWriter, req *http.Request) {
  numberqs := req.URL.Query().Get("number") 
  amountqs := req.URL.Query().Get("amount") 
  targetqs := req.URL.Query().Get("target")

  if numberqs == "" {
    fmt.Fprintf(w, "Account number is missing!")
    return
  }
  if targetqs == "" {
    fmt.Fprintf(w, "Target account number is missing!")
    return
  }

  if number, err := strconv.ParseFloat(numberqs, 64); err !=  nil{
    fmt.Fprintf(w, "Invalid account number!")
  }else if target, err := strconv.ParseFloat(targetqs, 64); err != nil {
    fmt.Fprintf(w, "Invalid target account number!") 
  }else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
    fmt.Fprintf(w, "Invalid amount number!")
  }else {
    account, ok := accounts[number]
    if !ok {
      fmt.Fprintf(w, "Account with number %v can't be found", number)
    } else {
      targetAccount, ok := accounts[target]
      if !ok {
        fmt.Fprintf(w, "Target account with number %v can't be found", target)
      } else {
        err := account.Transfer(amount, targetAccount)
        if err != nil {
          fmt.Fprintf(w, "%v", err)
        } else {
          fmt.Fprintf(w, account.Statement())
          fmt.Fprintf(w, targetAccount.Statement())
        }
      }
    }
  }
}
