# Godoc Comments

## Definition
Godoc comments are special comments in Go programming language used to document your code. These comments are written in a specific format to be recognized by the godoc tool, which generates documentation for Go code.

Godoc comments are placed immediately before the item (e.g., package, function, variable) they are documenting. They should start with the name of the item followed by a sentence describing its purpose or usage.

## Usage
To write Godoc comments, use the following format:

```go
// ItemName is a description of the item.
//
// Additional details or usage instructions.
// This can span multiple lines.
//
// Examples can be provided to illustrate usage.
package packageName

// FunctionName is a description of the function.
//
// Parameters:
//   param1: Description of the first parameter.
//   param2: Description of the second parameter.
//
// Returns:
//   Description of the return value.
//
// Example:
//   functionName(arg1, arg2)
func functionName(param1, param2 int) string {
    // Function implementation
}
```

## Example
```go
// GreetUser generates a personalized greeting for the user.
//
// Parameters:
//   name: The name of the user.
//
// Returns:
//   A string containing the personalized greeting.
//
// Example:
//   greeting := GreetUser("John")
//   fmt.Println(greeting)
func GreetUser(name string) string {
    return "Hello, " + name + "!"
}
```

