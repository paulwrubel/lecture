# Lecture

Lecture is a lecture-style natural speech esoteric programming language. It is without any known productive use or value.

Lecture is actually two separate things in a single repository: A well defined grammar for the Lecture language, and a tool called `lecture` which can interpret and transpile lecture programs.

**[Click here](using-the-lecture-tool.md) to view documentation on compiling and running the `lecture` binary, which can interpret valid Lecture files.**

## Table of Contents

- [Lecture](#lecture)
  - [Table of Contents](#table-of-contents)
  - [Language Features](#language-features)
    - [Syntax](#syntax)
      - [Entrypoint](#entrypoint)
      - [Comments](#comments)
      - [Assignment](#assignment)
      - [Reassignment](#reassignment)
      - [Types](#types)
        - [Number](#number)
        - [String](#string)
      - [Printing](#printing)
      - [Function calls](#function-calls)
      - [Operators](#operators)
      - [If Statements](#if-statements)
      - [Functions](#functions)
      - [General](#general)
        - [Capitalization](#capitalization)
  - [Style](#style)
    - [Indentation](#indentation)
    - [Variable names](#variable-names)
    - [Function Phrasing](#function-phrasing)
  - [Example Programs](#example-programs)
    - [Hello World](#hello-world)
    - [Fibonacci](#fibonacci)
  - [A Note on Valid Lecture](#a-note-on-valid-lecture)

## Language Features

The language grammar is defined using ANTLR4. The grammar file, `Lecture.g4` defines the grammar in [EBNF](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form).

### Syntax

Lecture is modelled after english speech. All tokens and keywords in the language are all-lowercase common words.

The following sections are written to introduce language feature incrementally, in the order you might reach to use them when writing a simple program.

#### Entrypoint

To start a lecture program, you can define the bounds of the main function that will execute using the following phrases:
```
okay, hear me out.
...
i rest my case.
```

Statements inside these phrases are executed automatically and exclusively. It is invalid Lecture to write statements outside of this bounds, unless they are in another function.

All statements in Lecture end in a single period ('.'), and whitespace, beside that which naturally separates words and tokens, is completely ignored (untokenized), which allows indentation for organization.

#### Comments

To add a comment, use the following phrase:
```
by the way, COMMENT_TEXT.
```

A comment statement always starts with `by the way, ` and always ends with `.`. All characters between them are added as a literal comment to the generated code.

#### Assignment

To assign a variable, use the following phrase:
```
let's say myVariable is literally 5.
```

This will assign a literal, `5`, to a variable called `myVariable`.

#### Reassignment

For a variable which has already been assigned, simply prepend the assignment phrase with `now`, as follows:
```
now let's say myVariable is literally 6.
```

#### Types

Currently, lecture supports two types, number and string. A number is always a 64-bit integer.

##### Number

A 64-bit integer. It can be defined using the keyword `literally`:
```
let's say twelve is literally 12.
```

##### String

A basic string. It can be defined using the keywords `quote, ` and `, unquote`:
```
let's say myString is quote, just a normal average string here, unquote.
```

#### Printing

To print a literal or a variable, use the phrase `then we have`, as follows:
```
let's save myVariable is literally 5.
by the way, the following line will print out variable.
then we have myVariable.
```

#### Function calls

Function calls are special. They can only be used in assignment and reassignment statements. They also cannot be used with operators, since I ran into [ambiguity issues](https://en.wikipedia.org/wiki/Serial_comma#Ambiguity) when defining the grammar. Once the result is saved to a variable, it can be used like all other variables.

To call a function, use the phrase `the result of FUNCTION_NAME, using ARG1 and ARG2 and ... and ARGX.`:
```
let's say myResult is the result of FindingMyResult, using literally 1 and literally 2.
```

[See here](#functions) for defining functions.

#### Operators

Lecture allows addition and subtraction with numbers and variables:
```
let's say seven is literally 7.
let's say five is literally 5.

let's say twelve is seven plus five.
let's say two is seven minus five.
```

These operators, in most cases, can be combined and used in all places where you would normaly reference some user token, e.g. assigning a variable, printing, passing arguments, returning, etc.

#### If Statements

Lecture supports `if` statements, `else if` statements, and `else` statements:
```
by the way, this is an `if (...) { ... }` block.
let's say five is literally 5.
if five is literally 5, here's what we need to do.
    then we have quote, five is 5, unquote.
now that we've done that, we can move on.
```

The phrases `where's what we need to do.` and `now that we've done that, we can move on` defined the bounds of the if statement.

Here are examples of an `if else` block and an `if else if else` block:
```
by the way, this is an `if (...) { ... } else { ... }` block.
let's say five is literally 5.
if five is literally 5, here's what we need to do.
    then we have quote, five is 5, unquote.
otherwise, here's what we need to do.
    then we have quote, five is not 5, unquote.
    then we have five.
now that we've done that, we can move on.

by the way, this is an `if (...) { ... } else if (...) { ... }` block.
let's say five is literally 6.
if five is literally 5, here's what we need to do.
    then we have quote, five is 5, unquote.
otherwise, if five is literally 6, here's what we need to do.
    then we have quote, five is 6, unquote.
now that we've done that, we can move on.

by the way, this is an `if (...) { ... } else if (...) { ... } else { ... }` block.
let's say five is literally 7.
if five is literally 5, here's what we need to do.
    then we have quote, five is 5, unquote.
otherwise, if five is literally 6, here's what we need to do.
    then we have quote, five is 6, unquote.
otherwise, here's what we need to do.
    then we have quote, five is not 5 or 6, unquote.
    then we have five.
now that we've done that, we can move on.
```

#### Functions

To define a function, use the (very verbose) phrase `we can use a process known as FUNCTION_NAME, which needs a TYPE called PARAMETER_NAME, to produce a TYPE.
it proceeds as follows.
`, as follows:
```
we can use a process known as AddingTwoNumbersTogether, which needs a number called a and a number called b, to produce a number.
it proceeds as follows.
    finally, we get a plus b.
```

It is a requirement that functions return a value, and that they return it on the final line. Return statements use the phrase `finally, we get`, as seen above.

Functions can be defined anywhere on the top level of a file. Nesting functions is not supported. It is not required to define a function before calling it.

#### General

##### Capitalization

Outside of user-defined variables, functions, and string literals, all token and phrases are **required** to be all lowercase.

## Style

In general, a few style rules are suggested to both make the source easier to read for control flow (e.g. indentation) and easier to read grammatically (e.g. intelligible sentences):

### Indentation

Indent inside of blocks, such as if statements and functions. Whitespace of all forms outside of tokens is ignore (except for newlines in comments), so indentation is always options (again, except for comments)

### Variable names

Variables should be named in camelCase. Functions should be named in PascalCase.

### Function Phrasing

Due to the way [function calls](#function-calls) work, it is better to phrase function names in [continuous or progressive tense](https://en.wikipedia.org/wiki/Continuous_and_progressive_aspects) (e.g. `AddingTwoNumbers`) instead of simple tense (`AddTwoNumber`). This is in opposition to most other languages.

The following is an example of the two tenses, to demonstrate the fludity of continuous/progressive tense:
```
by the way, this function is named in a SIMPLE tense.
by the way, this should NOT be used.
let's say myResult is the result of AddTwoNumbers, using literally 2 and literally 2.

by the way, this function is named in a CONTINUOUS tense.
by the way, this should be used.
let's say myResult is the result of AddingTwoNumbers, using literally 2 and literally 2.
```

## Example Programs

### Hello World

[See the file here](./examples/helloworld.ltr)

lecture text:
```
okay, hear me out. 
    let's say helloWorld is quote, Hello, World!, unquote.
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

[See the file here](./examples/fibonacci.ltr)

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

## A Note on Valid Lecture

Is it very possible (arguably _easy_) to write valid a valid Lecture file which produces invalid Golang. As Golang is the only target language for traspilation, this means you can write _grammatically correct Lecture which can never be ran_.

An easy example of this is trying to add a number to a string. There is nothing invalid, grammar-wise, about the following program:
```
okay, hear me out.
    let's say myNumber is literally 5.
    let's say myString is quote, just an average string, unquote.

    by the way, the following line is absolutely not valid Golang.
    let's say myTotallyNotValidGolangVariable is myNumber plus myString.
    then we have myTotallyNotValidGolangVariable.
i rest my case.
```

Despite this, the Golang output will be invalid:
```go
package main

import "fmt"

func main() {
        myNumber := int64(5)
        myString := "just an average string"
        // the following line is absolutely not valid Golang
        myTotallyNotValidGolangVariable := myNumber + myString
        fmt.Println(myTotallyNotValidGolangVariable)
}
```

This problem isn't that important in practice, since it's not exclusively the parser's job to catch syntax errors, and no tools exist to catch lecture errors before compilation anyways. This means, to the end user, it wouldn't really matter if the error was grammatical or otherwise, they will surface at the same time: during build.

Nonetheless, this is something to watch out for when writing Lecture.