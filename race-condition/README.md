# Race Condition

## Getting Started

Go includes a [built-in data race dectector](https://golang.org/doc/articles/race_detector) that we can use:

```bash
go run -race main.go
```

It will output the stack trace as well as the go routines involved. For example:

```
balance before deposit 10
balance before withdraw 10
==================
WARNING: DATA RACE
Write at 0x000000657300 by goroutine 7:
  main.deposit()
      /home/runner/golang-exercises/race-condition/main.go:14 +0xe1

Previous read at 0x000000657300 by goroutine 8:
  main.withdraw()
      /home/runner/golang-exercises/race-condition/main.go:19 +0x41

Goroutine 7 (running) created at:
  main.main()
      /home/runner/golang-exercises/race-condition/main.go:28 +0x61

Goroutine 8 (running) created at:
  main.main()
      /home/runner/golang-exercises/race-condition/main.go:29 +0x79
==================
balance after deposit: 11
==================
WARNING: DATA RACE
Write at 0x000000657300 by goroutine 8:
  main.withdraw()
      /home/runner/golang-exercises/race-condition/main.go:22 +0xe1

Previous write at 0x000000657300 by goroutine 7:
  main.deposit()
      /home/runner/golang-exercises/race-condition/main.go:14 +0xe1

Goroutine 8 (running) created at:
  main.main()
      /home/runner/golang-exercises/race-condition/main.go:29 +0x79

Goroutine 7 (running) created at:
  main.main()
      /home/runner/golang-exercises/race-condition/main.go:28 +0x61
==================
balance after withdraw: 0
Found 2 data race(s)
exit status 66
```

## Race Condition
What is the race condition here?

We are simulating a case where a bank account doing a deposit and withdraw concurrently.

One of the outputs from this program is:

```
balance before withdraw 10
balance after withdraw: 0
balance before deposit 10
balance after deposit: 11
```

The `withdraw` function deducts 10 from the balance and the `deposit` function adds 1 to the balance. In the above case, we originally inteded the output to be:
* "balance before deposit" should be 0
* "balance after deposit" should be 1

What's causing the race condition?

Races occur due to **interleaving** and **communication**.

Interleaving means that the order of execution **between** the concurrent task is non-deterministic.

Here, `withdraw` and `deposit` are communicating through the global variable `balance`.

Therefore, race conditions occur in this program leading to random outputs and causing bugs to happen.