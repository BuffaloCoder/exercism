package account

import "sync"

// Account manages balance
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open opens a new account with a given balance. If balance is negative, no
// account is created and nil is returned
func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}

	return &Account{balance: amt, closed: false}
}

// Close closes the account and returns a balance, if available
func (acct *Account) Close() (int64, bool) {
	acct.Lock()
	defer acct.Unlock()

	// Don't close if already closed
	if acct.closed {
		return 0, false
	}

	payout := acct.balance
	acct.closed = true
	acct.balance = 0

	return payout, true
}

// Balance returns the remaining balance
func (acct *Account) Balance() (int64, bool) {
	acct.RLock()
	defer acct.RUnlock()
	if acct.closed {
		return 0, false
	}
	return acct.balance, true
}

// Deposit adds & withdraws to/from the account
func (acct *Account) Deposit(amount int64) (int64, bool) {
	acct.Lock()
	defer acct.Unlock()
	if acct.closed {
		return 0, false
	}
	if acct.balance+amount < 0 {
		return acct.balance, false
	}
	acct.balance += amount
	return acct.balance, true
}
