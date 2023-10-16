---
theme: "solarized"
---

<style type="text/css">
	p { text-align: left; }
	code {
		max-height: 100% !important;
	}
</style>

# Mutation Testing In Go

---

## What is Mutation Testing

* Meta-Testing (Test the efficacy of your tests)
* No maintenance, works with existing tests
* Can be used on all types of tests that integrate into `go test` ootb
  * Fuzzing
  * Unit
  * Integration
  * ...

---

## Why Mutation Testing

* Enforces high code coverage
* Tests will become less brittle
* Makes sure your tests don't just execute statements, but also test them

---

## Requirements

* High code coverage
* Low Execution Time
* No flaky tests

---

## How does it work?

* Makes mutated (changed) copies of your production code
* Runs tests against the mutated code
* The mutation tests succeeds, if the real tests fail

---

## Mutations

* Change to an expression / statement or logical unit of code
* Different types of mutations
  * arithmetic
  * conditional
  * ...

---

## Libraries

* There's no super mature tool / library yet
* Available:
  * [gtramontina/ooze](https://github.com/gtramontina/ooze) <- We use this
  * [go-gremlins/gremlins](https://github.com/go-gremlins/gremlins)
  * [avito-tech/go-mutesting](https://github.com/avito-tech/go-mutesting)

---

## Ooze

* Features
  * Integrates into the `go test` command
  * Choose which mutations to use
  * Provide custom mutations
* Bugs
  * Only works on Nix-like systems (as of now)
  * Code needs to be `gofmt`'d right now.
  * One of the mutations is broken, causing compile errors 

---

## Performance Issues

Mutation Testing is slow, execution time grows heavily with the number of mutations.
Each mutation, if all is well, causes a rerun of your whole testsuite.

Tips:
* `-failfast` - kill mutant on first failure
* `go test -timeout=...` - prevent running into default timeout (10m)
* Use `t.Parallel()`
* Reuse test data / bootstrapped dependencies

---

## Example - 100% Coverage

```go
func do(x int) {
	if x > 0 {
		sideEffectA()
	}
	sideEffectB()
}
```

```go
func Test_do(t *testing.T) {
	for _, x := range []int{math.MinInt32, -1, 0, 1, 5, math.MaxInt32} {
		t.Run(fmt.Sprint(x), func(t *testing.T) {
			do(x)
		})
	}
}
```

```shell
$> go test -covermode=atomic ./examples/...
ok  ./examples/...  0.201s  coverage: 100.0% of statements
```

---

## 100% Coverage - Mutation Test 

<div>
    <img src="/media/100_percent_output.png"/>
</div>

---

## 100% Coverage - Improvements

*Transition to VSCode*


---

## Example - Custom Mutations

*Transition to VSCode*

---

## Example - Bigger Codebase

*Transition to VSCode*

---

## End

Thanks for listening :D

This presentation can be found at:

[Bios-Marcel/presentation_go_mutationtests](https://github.com/Bios-Marcel/presentation_go_mutationtests).
