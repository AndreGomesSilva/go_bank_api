package bank

import (
  "fmt"
  "errors"
)

//bank
type Bank interface {
	Statement() string
}

//Customer
type Customer struct {
  Name string
  Address string
  Phone string
}

//Account
type Account struct {
  Customer
  Number int32
  Balance float64
}

//Deposit ...
func (a *Account) Deposit(amount float64) error {
  if amount <= 0 {
    return errors.New("the amount to deposit should be greater than zero")
  }

  a.Balance += amount
  return nil
}

//Withdraw ...
func (a *Account) Withdraw(amount float64) error {
  if amount <= 0 {
    return errors.New("the amount to withdraw should be greater than zero")
  }

  if amount > a.Balance {
    return errors.New("the amount to withdraw should be less than balance")
  }

  a.Balance -= amount
  return nil
}

func (a *Account) Statement() string {
  return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(amount float64, target *Account) error {
  if a.Balance < amount {
    return errors.New("the amount to transfer should be less than balance")
  }
  err := a.Withdraw(amount)
  if err != nil {
    return err
  }
  err = target.Deposit(amount)
  if err != nil {
    a.Deposit(amount)
    return err
  }
  return nil
}

func Statement(b Bank) string {
  return b.Statement()
}
