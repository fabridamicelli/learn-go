package main

func SumArray5(nums [5]int) int {
	out := 0
	for _, v := range nums {
		out += v
	}
	return out
}

func Sum(nums []int) int {
	f := func(x, y int) int { return x + y }
	return Reduce(nums, 0, f)
}

// func Reduce[T any](values []T, initial T, f func(x, y T) T) T {
func Reduce[A, B any](values []A, initial B, f func(x B, y A) B) B {
	var final = initial
	for _, v := range values {
		final = f(final, v)
	}
	return final
}

func SumAll(numsToSum ...[]int) []int {
	out := make([]int, len(numsToSum))
	for i, arr := range numsToSum {

		out[i] = Sum(arr)
	}

	return out
}

func SumAllTails(numsToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(x[1:]))

	}
	return Reduce(numsToSum, []int{}, sumTail)
}

type Account struct {
	Name    string
	Balance float64
}
type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(tx []Transaction, name string) float64 {
	adjustBalance := func(balance float64, t Transaction) float64 {
		if t.From == name {
			return balance - t.Sum
		}
		if t.To == name {
			return balance + t.Sum
		}
		return balance
	}
	return Reduce(tx, 0.0, adjustBalance)
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}
