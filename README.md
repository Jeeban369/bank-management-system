# Bank Account Management System

A simple **Bank Account Management System** implemented in Go. This program enables users to create and manage bank accounts, deposit and withdraw money, check balances, and view transaction histories.

---

## Features

1. **Create Account**: Open a new account with a name and initial deposit.
2. **Deposit Money**: Add money to the account.
3. **Withdraw Money**: Deduct money from the account while ensuring no overdrafts.
4. **Check Balance**: View the current account balance.
5. **Transaction History**: Track all deposits and withdrawals.

---

## Code Overview

### Transaction Struct
The `Transaction` struct represents a single transaction in the account with the following fields:
- `Amount` (float64): The transaction amount.
- `Type` (string): The type of transaction (Deposit or Withdrawal).
- `Timestamp` (time.Time): The time the transaction occurred.

### Account Struct
The `Account` struct contains:
- `Name` (string): The account holder’s name.
- `Balance` (float64): The current balance in the account.
- `Transactions` ([]Transaction): A slice to store the transaction history.

### Methods

#### `newAccount`
Creates a new bank account with a name and initial deposit.
```go
func newAccount(name string, initialDeposit float64) *Account
```

#### `deposit`
Adds money to the account.
```go
func (acc *Account) deposit(amount float64)
```

#### `withdraw`
Deducts money from the account, ensuring sufficient balance.
```go
func (acc *Account) withdraw(amount float64)
```

#### `checkBalance`
Displays the current balance.
```go
func (acc *Account) checkBalance()
```

#### `transactionHistory`
Displays a list of all transactions.
```go
func (acc *Account) transactionHistory()
```

---

## How to Run

1. **Install Go**: Ensure Go is installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

2. **Clone the Repository**: Save the code in a `.go` file, for example, `bank.go`.

3. **Run the Program**:
```bash
go run bank.go
```

---

## Example Output

```
Account created for John Doe with initial balance: ₹5000.00
₹1500.00 deposited. New balance: ₹6500.00
₹2000.00 withdrawn. New balance: ₹4500.00
Current balance: ₹4500.00
Transaction History:
Deposit - ₹5000.00 (14 Dec 2024 12:00:00)
Deposit - ₹1500.00 (14 Dec 2024 12:05:00)
Withdrawal - ₹2000.00 (14 Dec 2024 12:10:00)
Insufficient balance for withdrawal.
```

---

## Learning Highlights
- Struct-based design for encapsulating data and functionality.
- Method-based operations for seamless user interaction.
- Handling edge cases like overdrafts and invalid inputs.
- Tracking transaction history using slices.



