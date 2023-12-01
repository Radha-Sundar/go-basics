# Go Templating Language

## Overview

Go's templating language is a lightweight and powerful template engine designed for embedding dynamic content within HTML and other text-based documents. It is widely used in Go web applications for rendering dynamic web pages with ease.

## Features

- **Simple Syntax:** The templating language uses a straightforward syntax that allows developers to embed dynamic content seamlessly.

- **Control Structures:** Supports control structures like conditionals and loops for dynamic content generation.

- **Automatic HTML Escaping:** Helps prevent common security issues by automatically escaping content to prevent Cross-Site Scripting (XSS) attacks.

- **Integration with Other Packages:** Easily integrates with other Go packages and libraries, making it versatile for various use cases.

## Usage

1. **Import the `html/template` package:**

   ```go
   import "html/template"

### 1.Create a template:
Define your template using the template.New() function.
```go
tmpl, err := template.New("index").Parse("Hello, {{.Name}}!")
if err != nil {
    // handle error
}
```

### 2.Prepare data:
Create a data structure to pass dynamic values into the template.
```go
data := struct {
    Name string
}{
    Name: "World",
}
```

### 3.Execute the template:
Apply the data to the template using the Execute method.
```go
err = tmpl.Execute(os.Stdout, data)
if err != nil {
    // handle error
}
```

## Example
```go
package main

import (
	"html/template"
	"os"
)

func main() {
	// Define a simple template
	tmpl, err := template.New("greeting").Parse("Hello, {{.Name}}!")
	if err != nil {
		panic(err)
	}

	// Prepare data
	data := struct {
		Name string
	}{
		Name: "Gopher",
	}

	// Execute the template
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
```

