# GO Data structure : Slice in GO
- A slice is a flexible and extensible data structure
- Like arrays, slices are also used to store multiple values of the same type in a single variable.
- Like Array , Slice have index value and length but the size of the slice can be resized as it is designed as variable-length sequence
- A slice can grow and shrink within the bounds of the underlying array.
- Slices have a capacity and length property ,
  - Length: The length is the total number of elements present in the array,
    - Slicing beyond its length extends the slice
  - Capacity: The capacity represents the maximum size upto which it can expand,
    - Slicing beyond its capacity causes a panic
- Slice is a reference to an underlying Array, Slice abstracts and manipulates an underlying Array
- Zero value of slice is nil, here capacity and the length of this slice is 0 also known as Nil Slice.

## Declaration and Initialization

These are the ways for creating Slice

### 1) Using the []type{values} format
###### Syntax : []type{value1, value2, value3, ...value n}
###### type can be string,int,float,struct,map
###### Sample-2 : []string{"Amala","Bobby","Charley"}
```
package main

import (
	"fmt"
)

func main() {

	names := []string{"Amala", "Bobby", "Charley"}
	fmt.Println("Slice:", names)
	fmt.Printf("Len: %v \tCap: %v\n", len(names), cap(names))
}

Output:

Slice: [Amala Bobby Charley]
Len: 3 	Cap: 3

```

### 2) Create a slice from an array
###### Syntax : array-name[start:end]
###### Sample-1 : names[2:4]
```
package main

import (
	"fmt"
)

func main() {
	nameArray := []string{"Amala", "Bobby", "Charley", "Don", "Eve"}
	nameSlice := nameArray[2:4]
	fmt.Println("Slice:", nameSlice)
}

Output:

Slice: [Charley Don]

```

### 3) Using the make() function
###### Syntax : make([]type, length, capacity)
###### Sample-1 : make([]int, 5, 10)
```
package main

import (
	"fmt"
)

func main() {
	nameSlice := make([]string, 5, 10)
	nameSlice[0] = "Amala"
	nameSlice[1] = "Bobby"
	
	// duplicate values are allowed in slice
	nameSlice[2] = "Bobby"
	fmt.Println("Slice:", nameSlice)
}

Output:

Slice: [Amala Bobby Bobby  ]

```

### 4) Using new keyword
###### Syntax :  new([capacity]type)[start:end]
###### Sample-1 :  new([50]int)[0:10]
```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var sl = new([50]int)[0:10]
	// to get type of the variable
	fmt.Println(reflect.ValueOf(sl).Kind())

	fmt.Printf("Len: %v \tCap: %v\n", len(sl), cap(sl))
	fmt.Println(sl)
	sl[0] = 23
	fmt.Println(sl)
}

Output:

slice
Len: 10 	Cap: 50
[0 0 0 0 0 0 0 0 0 0]
[23 0 0 0 0 0 0 0 0 0]

```

## Iteration / Print
### Using classic for loop
```
package main

import "fmt"

func main() {
	numbers := []int{2, 3, 5, 7, 11}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}

Output:
2
3
5
7
11

```

### Using range in for loop
```
package main

import "fmt"

func main() {

	names := []string{"Amala", "Bobby", "Charlie", "Belinda", "MIke", "Eve"}
	//Sample-1
	for i, name := range names {
		fmt.Println(i, "-", name)
	}
	
	fmt.Println("")
	//Sample-2
	for _, name := range names {
		fmt.Println(name)
	}

}

Output:

0 - Amala
1 - Bobby
2 - Charlie
3 - Belinda
4 - MIke
5 - Eve

Amala
Bobby
Charlie
Belinda
MIke
Eve

```

## Nil slice vs Empty slice:
### Nil slice:
var s []string,
####where length and capacity is 0 and the underlying array is Nil
### Empty slice:

s := []string{}
####where length and capacity is 0 and underlying array is with length 0
####Empty slice can also be created using make() function



## Add key and value to slice
- We use append() function to add elements to a slice.
###### Syntax :  <span style="color: #aa7ee1"> append </span>(<span style="color: #aa7ee1"> slice-name </span>, <span style="color: #aa7ee1"> value1 </span>, <span style="color: #aa7ee1"> value2 </span>)
###### Sample-1 : <span style="color: #aa7ee1"> append </span>(<span style="color: #aa7ee1"> prime </span>, <span style="color: #aa7ee1"> 17 </span>, <span style="color: #aa7ee1"> 19 </span>)
```
package main

import "fmt"

func main() {
	prime := []int{2, 3, 5, 7}
	prime = append(prime, 13, 17)
	fmt.Println("Prime Numbers:", prime)
}

Output:

Prime Numbers: [2 3 5 7 13 17]

```

## Remove value from slice
###### Syntax :  <span style="color: #aa7ee1"> append </span>(<span style="color: #aa7ee1"> slice-name[start-index] , slice-name[end-index] ... ) </span>
###### Syntax :  <span style="color: #aa7ee1"> append</span>(<span style="color: #aa7ee1">prime[:1],prime[2:]...) </span>

```
package main

import (
	"fmt"
)

func main() {
	prime := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Before Deletion:", prime)
	prime = append(prime[:1], prime[2:]...) //deletes 2nd element (3)
	fmt.Println("After Deletion:", prime)
}

Output:

Before Deletion: [2 3 5 7 11 13]
After Deletion: [2 5 7 11 13]

```

## Replace value in slice
#### Syntax :  <span style="color: #aa7ee1"> slice-name </span>[<span style="color: #aa7ee1"> index </span>] =<span style="color: #aa7ee1"> new-value </span>
#### Sample-1 : <span style="color: #aa7ee1"> evenslice</span>[<span style="color: #aa7ee1">2</span>]=<span style="color: #aa7ee1">6</span>

```
package main

import "fmt"

func main() {

	evenarray := [6]int{2, 4, 5, 7, 9, 12}
	evenslice := evenarray[:]

	fmt.Println("Before Update: ", evenslice)

	evenslice[2] = 6
	evenslice[3] = 8
	evenslice[4] = 10

	fmt.Println("After Update: ", evenslice)
}

Output:

Before Update:  [2 4 5 7 9 12]
After Update:  [2 4 6 8 10 12]

```

## Built-in functions
- These are the built-in functions used for Slice,
  - make(): helps to allocate and initialize the slice
  ```
  package main
  import "fmt"
  func main() {

	   var numbers = make([]int, 3, 5)

	   numbers[0] = 2
	   numbers[1] = 4
	   numbers[2] = 6

	   fmt.Println("Numbers: ", numbers)
  }
  
  Output:
  Numbers:  [2 4 6]
  ```
  - cap(): returns capacity of the slice
  ```
  package main
  import "fmt"
  func main() {
	   var numbers = make([]int, 3, 5)

	   numbers[0] = 2
	   numbers[1] = 4
	   numbers[2] = 6

	   fmt.Println("Numbers: ", numbers)
	   fmt.Println("Capacity: ", cap(numbers))
  }
  
  Output:
  Numbers:  [2 4 6]
  Capacity:  5
  ```
  - len(): returns length of the slice
  ```
  package main
  import "fmt"
  func main() {
    var numbers = make([]int, 3, 5)

	   numbers[0] = 2
	   numbers[1] = 4
	   numbers[2] = 6

	   fmt.Println("Numbers: ", numbers)
	   fmt.Println("Length: ", len(numbers))
  }
  
  Output:
  Numbers:  [2 4 6]
  Length:  3
  ```
  - append(): appends an element at the end of the slice
  ```
  package main
  import "fmt"
  func main() {
    var first []string

	   fmt.Println("Before Append:", first)
	   first = append(first, "Bobby")
	   first = append(first, "Charlie")
	   first = append(first, "Mike")
	   fmt.Println("After Append:", first)
  }
  
  Output:
  
  Before Append: []
  After Append: [Bobby Charlie Mike]
  
  ```
  - copy(): copies elements from source slice to destination
  ```
  package main
  import "fmt"
  func main() {
      first := []string{"Antony", "Bobby", "Charlie"}
      second := make([]string, len(first))

	     copy(second, first)

	     fmt.Printf("First slice: %[1]v, address: %[1]p\n", first)
	     fmt.Printf("Copied slice: %[1]v, address: %[1]p\n", second)
  }
  
  Output:
  First slice: [Antony Bobby Charlie], address: 0xc0000121b0
  Copied slice: [Antony Bobby Charlie], address: 0xc0000121e0
  // Address will be different
  ```
## Compare two slices in GO
- We can compare two slices by two ways,
  - custom method
  - reflect
### Custom method
```
package main

import "fmt"

func checkStringEquality(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func main() {
	fruits := []string{"lychee", "dragon"}
	exotic := []string{"lychee", "dragon"}
	fmt.Printf("slices equal: %t\n", checkStringEquality(fruits, exotic))

	exotic = append(exotic, "blueberry")
	fmt.Printf("slices equal: %t\n", checkStringEquality(fruits, exotic))
}

Output:

slices equal: true
slices equal: false

```
### Reflect
```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fruits := []string{"lychee", "dragon"}
	exotic := []string{"lychee", "dragon"}
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(fruits, exotic))

	exotic = append(exotic, "blueberry")
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(fruits, exotic))
}

Output:

slices equal: true
slices equal: false

```
## Search
###### Linear search
```
package main

import (
	"fmt"
)

func main() {

	fruits := []string{"apple", "banana", "kiwi", "orange"}
	fmt.Println("Fruits :", fruits)
	fmt.Println("Search (kiwi) in fruits: ", Contains(fruits, "kiwi"))
	fmt.Println("Search (pineapple) in fruits: ", Contains(fruits, "pineapple"))

}
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}


Output:

Fruits : [apple banana kiwi orange]
Search (kiwi) in fruits:  true
Search (pineapple) in fruits:  false

```
## Sorting slices using sort
```
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

	fmt.Println("Before sorting (int):", slice)
	sort.Ints(slice)
	fmt.Println("After sorting (int):", slice)

	sliceStr := []string{"Zen", "Charlie", "Annie", "Mike"}
	fmt.Println("\nBefore sorting (string):", sliceStr)
	sort.Strings(sliceStr)
	fmt.Println("After sorting (string):", sliceStr)
}

Output:

Before sorting (int): [45 67 23 90 33 21 56 78 89]
After sorting (int): [21 23 33 45 56 67 78 89 90]

Before sorting (string): [Zen Charlie Annie Mike]
After sorting (string): [Annie Charlie Mike Zen]

```

## Merging slices

```
package main

import "fmt"

func main() {
	odd := []int{1, 2, 3, 4, 5}
	even := []int{6, 7, 8, 9, 10}
	prime := []int{11, 13, 17, 19, 23}

	numbers := []int{}
	numbers = append(odd, even...) // Can't concatenate more than 2 slice at once
	fmt.Printf("numbers after merging (odd,even): %v\n", numbers)
	fmt.Printf("capacity: %v\n", cap(numbers))
	numbers = append(numbers, prime...) // concatenate numbers with prime

	fmt.Printf("\nnumbers after merging (prime): %v\n", numbers)
	fmt.Printf("capacity: %v\n", cap(numbers))
}

Output:

numbers after merging (odd,even): [1 2 3 4 5 6 7 8 9 10]
capacity: 10

numbers after merging (prime): [1 2 3 4 5 6 7 8 9 10 11 13 17 19 23]
capacity: 20
```

## Slice is a reference type
- When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same underlying array. Hence when a slice is passed to a function as parameter, changes made inside the function are reflected outside the function too.
```
package main

import (
	"fmt"
)

func multiply(numbers []int) {
	for i := range numbers {
		numbers[i] *= 10
	}

}
func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	multiply(nos)
	fmt.Println("slice after function call", nos)
}

Output:

slice before function call [8 7 6]
slice after function call [80 70 60]

```


## Slicing
```
package main

import "fmt"

func main() {
	var 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 = []string{"Lotus", "Jasmine", "Rose", "Sunflower", "Lavender"}

	fmt.Printf("Flowers                                    : %v\n", flowers)
	fmt.Printf("First Value [:1]                           : %v\n", flowers[:1])
	fmt.Printf("First two Values [:2]                      : %v\n", flowers[:2])
	fmt.Printf("Values from start-index to end-index [1:3] : %v\n", flowers[1:3])
	fmt.Printf("Values from start-index to end [3:]        : %v\n", flowers[3:])
	fmt.Printf("Values from start-index to end-index [0:3] : %v\n", flowers[0:3])
	fmt.Printf("Last Value                                 : %v\n", flowers[len(flowers)-1])

	fmt.Println("All Values using [:]                       :", flowers[:])
	fmt.Println("All Values using [0:]                      :", flowers[0:])
	fmt.Println("All Values using [0:len(flowers)]        :", flowers[0:len(flowers)])
}

Output:

Flowers                                    : [Lotus Jasmine Rose Sunflower Lavender]
First Value [:1]                           : [Lotus]
First two Values [:2]                      : [Lotus Jasmine]
Values from start-index to end-index [1:3] : [Jasmine Rose]
Values from start-index to end [3:]        : [Sunflower Lavender]
Values from start-index to end-index [0:3] : [Lotus Jasmine Rose]
Last Value                                 : Lavender
All Values using [:]                       : [Lotus Jasmine Rose Sunflower Lavender]
All Values using [0:]                      : [Lotus Jasmine Rose Sunflower Lavender]
All Values using [0:len(flowers)]        : [Lotus Jasmine Rose Sunflower Lavender]

```

