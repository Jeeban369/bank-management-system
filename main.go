package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	Amount    float64
	Type      string
	Timestamp time.Time
}

type Account struct {
	Name         string
	Balance      float64
	Transactions []Transaction
}

// create account
func newAccount(name string, initialDeposit float64) *Account{
	account := &Account{
		Name:    name,
		Balance: initialDeposit,
		Transactions: []Transaction{
			{Amount: initialDeposit, Type: "Deposit", Timestamp: time.Now()},
		},
	}
	fmt.Printf("Account created for %s with initial balance: %.2f\n", name, initialDeposit)
	return account
}

//depodit money
func (acc *Account) deposit(amount float64){
	if amount<0{
		fmt.Println("deposite amount must be greater  than zero.")
		return
	}
	acc.Balance += amount
	acc.Transactions = append(acc.Transactions, Transaction{Amount: amount, Type: "Deposit", Timestamp: time.Now()})
	fmt.Printf("₹%.2f ddeposited. New balance: ₹%.2f\n", amount, acc.Balance)
}


// Method to withdraw money
func (acc *Account) withdraw(amount float64) {
	if amount > acc.Balance {
		fmt.Println("Insufficient balance for withdrawal.")
		return
	}
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be greater than zero.")
		return
	}
	acc.Balance -= amount
	acc.Transactions = append(acc.Transactions, Transaction{Amount: amount, Type: "Withdrawal", Timestamp: time.Now()})
	fmt.Printf("₹%.2f withdrawn. New balance: ₹%.2f\n", amount, acc.Balance)
}

// Method to check the account balance
func (acc *Account) checkBalance() {
	fmt.Printf("Current balance: ₹%.2f\n", acc.Balance)
}

// Method to view transaction history
func (acc *Account) transactionHistory() {
	if len(acc.Transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}
	fmt.Println("Transaction History:")
	for _, tx := range acc.Transactions {
		fmt.Printf("%s - ₹%.2f (%s)\n", tx.Type, tx.Amount, tx.Timestamp.Format("02 Jan 2006 15:04:05"))
	}
}

func main() {
	// Create a new account
	account := newAccount("John Doe", 5000)

	// Perform some transactions
	account.deposit(1500)
	account.withdraw(2000)
	account.checkBalance()

	// View transaction history
	account.transactionHistory()

	// Try withdrawing more than the balance
	account.withdraw(5000)
}