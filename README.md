# Lecture

Lecture is a lecture-style natural speech esoteric programming language. It is without any known productive use.

## Examples

### Hello World

lecture text:
```
okay, hear me out. 
    let's say helloWorld is literally quote, Hello, World!, unquote.
    then we have helloWorld.
i rest my case.
```

transpiled to golang:
```go
package main

import "fmt"

func main() {
	helloWorld := "Hello, World!"
	fmt.Println(helloWorld)
}
```

### Fibonacci

lecture text:
```
we can use a process known as FindingTheNthFibonacciNumber, which needs a number called n, to produce a number.
it proceeds as follows.
    let's say theNthFibonacciNumber is literally 0.

    if n is literally 0, here's what we need to do.
        now let's say theNthFibonacciNumber is literally 0.
    otherwise, if n is literally 1, here's what we need to do.
        now let's say theNthFibonacciNumber is literally 1.
    otherwise, here's what we need to do.
        let's say theNthMinusOneFibonacciNumber is the result of FindingTheNthFibonacciNumber, using n minus literally 1.
        let's say theNthMinusTwoFibonacciNumber is the result of FindingTheNthFibonacciNumber, using n minus literally 2.
        now let's say theNthFibonacciNumber is theNthMinusOneFibonacciNumber plus theNthMinusTwoFibonacciNumber.
    now that we've done that, we can move on.

    finally, we get theNthFibonacciNumber.
    
okay, hear me out. 
    let's say theNineteenthFibonacciNumber is the result of FindingTheNthFibonacciNumber, using literally 19.
    then we have theNineteenthFibonacciNumber.
i rest my case.
```

transpiled to golang:
```go
package main

import "fmt"

func FindingTheNthFibonacciNumber(n int64) int64 {
	theNthFibonacciNumber := int64(0)
	if n == int64(0) {
		theNthFibonacciNumber = int64(0)
	} else if n == int64(1) {
		theNthFibonacciNumber = int64(1)
	} else {
		theNthMinusOneFibonacciNumber := FindingTheNthFibonacciNumber(n - int64(1))
		theNthMinusTwoFibonacciNumber := FindingTheNthFibonacciNumber(n - int64(2))
		theNthFibonacciNumber = theNthMinusOneFibonacciNumber + theNthMinusTwoFibonacciNumber
	}
	return theNthFibonacciNumber
}
func main() {
	theNineteenthFibonacciNumber := FindingTheNthFibonacciNumber(int64(19))
	fmt.Println(theNineteenthFibonacciNumber)
}

```
