# GO Data structure : Struct in GO
- A struct in GO is a user-defined type that allows user to group/combine items of different type into single type.
- A struct is termed as lightweight class, that does not support inheritance but supports composition.
- The data-type can be  built-in and user-defined types. 
- Structs can improve modularity and allows user to create and pass complex data structures around the system.
- Structs are mutable.

## Declaration 
- The declaration contains the keyword **type** and **struct**
- The data-type of the field are mandatory in struct
### Declaration named struct
```
type struct-name struct {
	field1-name  string
	field2-name  int64
}
```
### Declaration anonymous struct:
```
type struct-name struct {
	string
	int64
}
```

## Initialization

### Initializing struct with field names
```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{
		name:          "Charlie",
		accountnumber: 3005353680,
		balance:       10000,
	}

	fmt.Println("User:", user1)
}

Output;
User: {Charlie 3005353680 10000}

```
### Initializing struct without specifying field names
```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{"Charlie", 3005353680, 10000}
	fmt.Println("User:", user1)
}

Output;
User: {Charlie 3005353680 10000}

```
### Initializing using new keyword
```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := new(User)
	user1.name = "Charlie"
	user1.accountnumber = 3005353680
	user1.balance = 10000

	fmt.Println("User:", user1)
}

Output:
User: &{Charlie 3005353680 10000}

```
### Initializing using (&) operator:
```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	var user1 = &User{"Charlie", 3005353680, 10000}
	fmt.Println(user1)

}

Output:
&{Charlie 3005353680 10000}

```
### Initializing anonymous struct
```
package main

import "fmt"

func main() {
	user1 := struct {
		name          string
		accountnumber int64
		balance       int
	}{
		name:          "Charlie",
		accountnumber: 3005353680,
		balance:       10000,
	}
	fmt.Println("User:", user1)

}

Output:
User: {Charlie 3005353680 10000}

```

### Initializing zero valued struct
- If we initialize with empty braces it will create zero valued struct

```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{}
	fmt.Println("User:", user1)
}

Output:
User: { 0 0}

```

## Nested structs
- Struct can contain field which in turn is a struct, these structs are called Nested structs.
```
package main

import (
	"fmt"
)

type Marks struct {
	language1 int
	language2 int
	maths     int
	science   int
	social    int
}

type Student struct {
	name  string
	class string
	marks Marks
}

func main() {
	student1 := Student{
		name:  "Charlie",
		class: "2d",
		marks: Marks{
			language1: 96,
			language2: 95,
			maths:     98,
			science:   97,
			social:    98,
		},
	}

	fmt.Println("student1:", student1)
	fmt.Println("\nName:", student1.name)
	fmt.Println("maths:", student1.marks.maths)
}

Output:
student1: {Charlie 2d {96 95 98 97 98}}

Name: Charlie
maths: 98
```

## Go struct as a value type 
- Go struct is value type by default.
- When we assign a struct variable to another struct variable, a new copy of the struct is created.
- when we pass a struct to another function, the function receives a new copy of the struct.
- Changing the fields of the new struct does not affect the original struct.
```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{"Charlie", 3005353680, 10000}
	fmt.Println("Name before modification :", user1.name)

	user2 := user1
	user2.name = "Richard Charlie"

	fmt.Println("Name after modification:", user1.name)

}

Output:
Name before modification : Charlie
Name after modification: Charlie
Note: value didn't change after modification

```

## Go struct as a reference type
- we can pass value by reference using (&) operator, so that the changes will get updated 
```
package main

import "fmt"

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{"Charlie", 3005353680, 10000}

	user2 := &user1
	fmt.Println("Name before modification :", user1.name)

	user2.name = "Richard Charlie"

	fmt.Println("Name after modification:", user1.name)

}

Output:
Name before modification : Charlie
Name after modification: Richard Charlie

```

## Get type of struct:
- The reflect package is used to get the underlying type of a struct.
```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{"Charlie", 3005353680, 10000}
	fmt.Println("User :", user1)
	fmt.Println("Type:", reflect.ValueOf(user1).Kind())
}

Output:

User : {Charlie 3005353680 10000}
Type: struct

```

## Compare structs with (==) operator:
- Structs of the same type can be compared using comparison operator.
```
package main

import (
	"fmt"
)

type User struct {
	name          string
	accountnumber int64
	balance       int
}

func main() {

	user1 := User{"Charlie", 3005353680, 10000}
	user2 := User{"Charlie", 3005353680, 10000}
	user3 := User{"Charles", 3005323233, 5000}

	fmt.Println("User1==user2 :", user1 == user2)
	fmt.Println("user1==user3 :", user1 == user3)
}

Output:
User1==user2 : true
user1==user3 : false
```
## Go struct Anonymous fields
- The fields available in the anonymous struct are called anonymous fields.
- Anonymous fields does not have an explicit name, by default the name of an anonymous field is the name of its type.
```
package main

import (
	"fmt"
)

type User struct {
	string
	int64
	int
}

func main() {

	user1 := User{
		string: "Charlie",
		int64:  3005353680,
		int:    10000,
	}
	fmt.Println("user:", user1)
	fmt.Println("balance:", user1.int)
}

Output:

user: {Charlie 3005353680 10000}
balance: 10000

```
## Go struct Promoted fields
- Fields that belong to an anonymous struct field in a struct are called promoted fields
```
type marks struct {
	language1 int
	language2 int
	maths     int
	science   int
	social    int
}

type Student struct {
	name  string
	class string
	marks
}

- here language1,language2,maths,science and social are called promoted fields.
```
- Example:
```
package main

import (
	"fmt"
)

type marks struct {
	language1 int
	language2 int
	maths     int
	science   int
	social    int
}

type Student struct {
	name  string
	class string
	marks
}

func main() {
	student1 := Student{
		name:  "Charlie",
		class: "2d",
		marks: marks{
			language1: 96,
			language2: 95,
			maths:     98,
			science:   97,
			social:    98,
		},
	}

	fmt.Println("student1:", student1)
	fmt.Println("\nName:", student1.name)
	fmt.Println("maths:", student1.marks.maths)
}

Output:
student1: {Charlie 2d {96 95 98 97 98}}

Name: Charlie
maths: 98
```
## Go fields as functions
- we can give user defined function as the field in struct
```
package main

import (
	"fmt"
)

type Total func(marks []int) int

type Student struct {
	name  string
	class string
	marks []int
	total Total
}

func main() {
	student1 := Student{
		name:  "Charlie",
		class: "2d",
		marks: []int{96, 95, 98, 97, 98},
		total: func(marks []int) int {
			tot := 0
			for i := range marks {
				tot += marks[i]
			}
			return tot
		},
	}

	fmt.Println("student1:", student1)
	fmt.Println("total:", student1.total(student1.marks))
}

Output:
student1: {Charlie 2d [96 95 98 97 98] 0x47dcc0}
total: 484

```
## Access individual fields of struct
- We can access individual fields of the struct using (.) operator
```
package main

import (
	"fmt"
)

type Total func(marks []int) int

type Student struct {
	name  string
	class string
	marks []int
}

func main() {
	student1 := Student{
		name:  "Charlie",
		class: "2d",
		marks: []int{96, 95, 98, 97, 98},
	}

	fmt.Println("student1:", student1)
	fmt.Println("Name:", student1.name)
	fmt.Println("Marks:", student1.marks)
}

Output:

student1: {Charlie 2d [96 95 98 97 98]}
Name: Charlie
Marks: [96 95 98 97 98]

```

