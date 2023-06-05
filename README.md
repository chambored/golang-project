# Go Programming Language

The programming language I picked for this project is Go, also known as Golang. The Go version used is Go 1.16.

## Why Go?

I chose Go because it's simple, efficient, and excels in concurrent programming. Go is a statically-typed language, meaning the variable type is known at compile time.

## How Go Handles Different Programming Concepts

**Object-Oriented Programming:** Go does not have classes. Instead, Go uses struct types and interfaces to achieve similar goals. Methods can be defined on types and interfaces can be used to make code more abstract and flexible.

**File Ingestion:** Go has a standard library package called "os" that provides functions for working with files and the OS. Files can be opened, read, written, and closed easily with this package.

**Conditional Statements:** Go supports if-else and switch-case conditional statements similar to other popular languages.

**Assignment Statements:** Go uses "=" for assignment and ":=" for declaring and assigning at the same time.

**Loops:** Go uses the "for" loop for all kinds of iterations including traditional for loop, while loop, and infinite loop.

**Functions/Methods:** Functions are defined using the "func" keyword. Go supports functions with multiple return values. Go uses pass by value, which means that functions get a copy of the variable and changes to it won't affect the original one. But pointers can be passed to the variable if there's a need to modify it directly.

**Unit Testing:** Go has a built-in testing tool called "testing". It provides a simple mechanism to write and execute tests.

**Exception Handling:** Go doesn't have exceptions, it uses error values to indicate an error. Errors are returned as an extra return value from functions.

**Data Types:** Go provides a variety of data types, including unsigned integer types (uint), byte (an alias for an unsigned 8-bit integer), rune (an alias for a signed 32-bit integer), and uintptr (an integer representation of a memory address).

## Libraries Used

**fmt:** It's a standard library that provides functions for formatted I/O. I used this for printing to the console.

**sort:** It's a standard library for sorting slices and user-defined collections. I used this for sorting years in the function "countPhonesByYear".

**regexp:** This is a standard library for regular expressions. I used this for parsing various strings into numbers and other formats in the parsing functions.

**strconv:** The strconv package provides functions to convert strings to primitive types, like integers or floats. It was used in this project to convert string data extracted from the CSV file to required data types.

**testing:** This is Go's built-in testing package. It provides a set of functions and conventions for writing and executing tests. It was used to validate the logic of various functions in the codebase.

## Misc

The underscore in Go is a blank identifier. You can use it when you don't care about a variable in a context. For instance, if a function returns multiple values and you only care about some of them, you can assign the ones you don't need to _. In loops, you can use _ if you only need the index or key but not the value, or vice versa.

A special note about Go loops - the "for {}" construct loops indefinitely, similar to a "while(true)" loop in other languages like Java.

Overall, Go is a very powerful language that can handle complex tasks with simplicity and efficiency. It is highly suitable for concurrent operations and networked tasks, which make it an excellent choice for projects that need performance and scalability.
