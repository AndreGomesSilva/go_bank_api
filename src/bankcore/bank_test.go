package bank

import (
	"testing"
)

func TestAccount(t *testing.T){
  account := Account{
    Customer: Customer{
      Name: "John",
      Address: "Los Angeles, California",
      Phone: "123-456-7890",
    },
    Number: 1001,
    Balance: 0,
  }

  if account.Name == "" {
    t.Error("can't create an Account object")
  }
}

func TestDeposit(t *testing.T) {
  account := Account{
    Customer: Customer{
      Name: "John",
      Address: "Los Angeles, California",
      Phone: "123-456-7890",
    },
    Number: 1001,
    Balance: 0,
  }

    account.Deposit(10)

    if account.Balance != 10 {
      t.Error("balance should be 10")
    }
}

func TestDepositInvalid(t *testing.T) {
  account := Account{
    Customer: Customer{
      Name: "John",
      Address: "Los Angeles, California",
      Phone: "123-456-7890",
    },
    Number: 1001,
    Balance: 0,
  }

  if err := account.Deposit(-10); err == nil {
    t.Error("should throw an error")
  }
}

func TestWithdraw(t *testing.T){
 account := Account{
   Customer: Customer{
     Name: "John",
     Address: "Los Angeles, California",
     Phone: "123-456-7890",
   },
   Number: 1001,
   Balance: 0,
 }
 account.Deposit(10)
 account.Withdraw(10)

 if account.Balance != 0 {
   t.Error("balance should be 0")
 }
}

func TestStatement(t *testing.T) {
 account := Account{
   Customer: Customer{
     Name: "John",
     Address: "Los Angeles, California",
     Phone: "123-456-7890",
   },
   Number: 1001,
   Balance: 0,
 }
 account.Deposit(100)
 statement := account.Statement()
 if statement != "1001 - John - 100" {
   t.Error("statement should be 1001 - John - 100.00")
 }
}

func TestTransfer(t *testing.T){
  Accounts := map[float64]*Account{
    1001:{
      Customer: Customer{
        Name: "John",
        Address: "Los Angeles, California",
        Phone: "123-456-7890",
      },
      Number: 1001,
      Balance: 0,
    },
    1002:{
      Customer: Customer{
        Name: "Jane",
        Address: "Los Angeles, California",
        Phone: "123-456-7890",
      },
      Number: 1002,
      Balance: 0,
    },
  }

  Accounts[1001].Deposit(100)
  Accounts[1002].Deposit(100)
  Accounts[1001].Transfer(100, Accounts[1002])
  if Accounts[1002].Balance != 200 {
    t.Error("balance should be 200")
  }
}
