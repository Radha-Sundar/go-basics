# GO Data structure : Map in GO
- Maps are key-value storage with fast lookups.
- The zero value of map is nil.
- Maps are a reference to the hash table.
- The data-type of key must be of same type and data-type of values must be of same type but data-type of key and value can be different.

## Declaration and Initialization
- In GO, Maps can be created in two ways , 
    - Initialization without make() will allow to initialize **values**.
    - Initialization with make function will allow to initialize **capacity**.
    - Note : For an empty map with capacity 0, both ways are same
### I) Simple Initialization 
###### Syntax :  <span style="color: #954121">map </span>[<span style="color: #aa7ee1"> key-type </span>]<span style="color: #aa7ee1"> val-type </span>{}
###### Sample-1 : <span style="color: #954121">map</span>[<span style="color: #aa7ee1">string</span>]<span style="color: #aa7ee1">int</span>{}      
###### Sample-2 : <span style="color: #954121">map</span>[<span style="color: #aa7ee1">string</span>]<span style="color: #aa7ee1">int</span>{"age": 32}
```
package main

import "fmt"

func main() {
	detail := map[string]int{"age": 32}
	fmt.Println(detail["age"])
}

Output:
32

```
###
### II) Initialization using make function
###### Syntax :  <span style="color: #954121">make </span>(<span style="color: #954121"> map </span>[<span style="color: #aa7ee1"> key-type </span>]<span style="color: #aa7ee1"> val-type </span>)
###### Sample-1 : <span style="color: #954121">make</span>(<span style="color: #954121">map</span>[<span style="color: #aa7ee1">string</span>]<span style="color: #aa7ee1">int</span>)
###### Sample-2 : <span style="color: #954121">make</span>(<span style="color: #954121">map</span>[<span style="color: #aa7ee1">string</span>]<span style="color: #aa7ee1">float64</span>, 100)
###### Note: 100 is the capacity of Map ,
###### 100 is an estimate of the initial capacity of a map, It pre-allocates memory for at least 100 entries, which results in increased performance.
```
package main

import "fmt"

func main() {
	temperature := make(map[string]float64, 100)
	temperature["Maya"] = 97.4
	temperature["Nitya"] = 98.6
	fmt.Println(temperature["Maya"])
}

OUtput:
97.4
```

## Adding values to empty map will give runtime error
```
package main

import "fmt"

func main() {
	var emptyMap map[string]int
	fmt.Println(emptyMap)
	if emptyMap == nil {
		fmt.Println("emptyMap is nil")
	}

	// Attempting to add keys to a nil map will result in a runtime error
	// emptyMap["Amala"] = 32
}

Output:
map[]
emptyMap is nil

```


## Iteration / Print
### Get by Key &nbsp;&nbsp;<span style="color: #aa7ee1"> map-name</span>[<span style="color: #aa7ee1"> key-name</span>] :
###### Sample-1: <span style="color: #aa7ee1"> marks</span>[<span style="color: #aa7ee1">"Sandy"</span>]
```
package main

import "fmt"

func main() {
        var marks = map[string]int{"Sandy": 450}
	fmt.Println(marks["Sandy"])
}

Output:

450

```
### Get all Key-Value Pairs &nbsp;&nbsp;<span style="color: #954121">fmt.Println(<span style="color: #aa7ee1"> map-name</span>)</span>:
###### Sample-1: <span style="color: #954121">fmt.Println(<span style="color: #aa7ee1">studentList</span>)</span>

```
package main

import "fmt"

func main() {
	studentList := map[int]string{
		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	fmt.Println(studentList)
}

Output:

map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Gabby]

```
###### Sample-2: Iteration with for loop

```
package main

import "fmt"

func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	for student := range studentList {
		fmt.Println(student, "=>", studentList[student])
	}

}

Output:

102 => Bobby
103 => Daniel
104 => Eve
105 => Gabby
101 => Amala
```
###### Sample-2: Iteration with range

```
package main

import "fmt"

func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	for key, value := range studentList {
		fmt.Printf("studentList[%v] = %v\n", key, value)
	}
}

Output:

studentList[101] = Amala
studentList[102] = Bobby
studentList[103] = Daniel
studentList[104] = Eve
studentList[105] = Gabby
```

## Add key and value to Map
###### Syntax :  <span style="color: #aa7ee1"> map-name </span>[<span style="color: #aa7ee1"> key-name </span>] =<span style="color: #aa7ee1"> key-value </span>
###### Sample-1 : <span style="color: #aa7ee1"> studentList</span>[<span style="color: #aa7ee1">106</span>]=<span style="color: #aa7ee1">"Issac"</span>
```
package main

import "fmt"

func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	fmt.Println("Before-studentList ", studentList)
	studentList[106] = "Issac"
	
	fmt.Println("Before-studentList ", studentList)

}

Output:

Before-studentList  map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Gabby]
Before-studentList  map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Gabby 106:Issac]

```

## Remove key/value from Map
###### Syntax :  <span style="color: #954121"> delete </span>(<span style="color: #aa7ee1"> map_name , key-name </span>)
###### Sample-1 : <span style="color: #954121"> delete</span>(<span style="color: #aa7ee1">studentList, 104</span>)

```
package main

import "fmt"

func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	fmt.Println("Before-studentList ", studentList)
	delete(studentList, 104)

	fmt.Println("Before-studentList ", studentList)

}

Output:

Before-studentList  map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Gabby]
Before-studentList  map[101:Amala 102:Bobby 103:Daniel 105:Gabby]

```

## Replace key/value in Map
#### Syntax :  <span style="color: #aa7ee1"> map-name </span>[<span style="color: #aa7ee1"> key-name </span>] =<span style="color: #aa7ee1"> new-key-value </span>
#### Sample-1 : <span style="color: #aa7ee1"> studentList</span>[<span style="color: #aa7ee1">105</span>]=<span style="color: #aa7ee1">"Issac"</span>

```
package main

import "fmt"

func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	fmt.Println("Before-studentList ", studentList)
	studentList[105] = "Issac"

	fmt.Println("Before-studentList ", studentList)

}

Output:

Before-studentList  map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Gabby]
Before-studentList  map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Issac]

```
## Built-in
- <span style="color: #72e0d1">len</span> returns the number of key/value pairs when called on a map.
###### Syntax :  <span style="color: #72e0d1">len</span> (<span style="color: #aa7ee1"> key-name </span>)
###### Sample-1 : <span style="color: #72e0d1">len</span>(<span style="color: #aa7ee1">studentList</span>)

```
package main

import "fmt"

func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	fmt.Println("length:", len(studentList))
}

```
- <span style="color: #72e0d1">delete</span> removes key/value pairs from a map.

## Search
- Value in the map is accessed like below with true/false boolean
### Check for zero value
```
package main

import "fmt"

func main() {
	marks := map[string]int{

		"Amala":  410,
		"Bobby":  440,
		"Daniel": 380,
		"Eve":    490,
		"Gabby":  350,
	}

	f, _ := marks["Eve"]
	fmt.Println("Eve:", f)

	g, _ := marks["Eves"]
	fmt.Println("Eves:", g)

}

Output:

Eve: 490
Eves: 0

if the value is zero value of value-type , then it means no match found

```
### Use second return value directly in an if statement
```
package main

import "fmt"
func main() {
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}

	if g, found := studentList[101]; found {
		fmt.Println(g)
	}

}

Output:

Amala

```

## Remove all items in Map
```
package main

import "fmt"

func main() {
	// Method - I
	marks := map[string]int{

		"Amala":  410,
		"Bobby":  440,
		"Daniel": 380,
		"Eve":    490,
		"Gabby":  350,
	}
	fmt.Println("marks Before Truncate:", marks)
	for k := range marks {
		delete(marks, k)
	}
	fmt.Println("marks After Truncate:", marks)

	// Method - II
	studentList := map[int]string{

		101: "Amala",
		102: "Bobby",
		103: "Daniel",
		104: "Eve",
		105: "Gabby",
	}
	fmt.Println("")
	fmt.Println("studentList Before Truncate:", studentList)
	studentList = make(map[int]string)
	fmt.Println("studentList After Truncate:", studentList)
}

Output:

marks Before Truncate: map[Amala:410 Bobby:440 Daniel:380 Eve:490 Gabby:350]
marks After Truncate: map[]

studentList Before Truncate: map[101:Amala 102:Bobby 103:Daniel 104:Eve 105:Gabby]
studentList After Truncate: map[]

```

## Sorting Map by Keys
```
package main

import (
	"fmt"
	"sort"
)

func main() {
	marks := map[string]int{
		"Amala":  410,
		"Eve":    490,
		"Daniel": 380,
		"Bobby":  440,
		"Gabby":  350,
	}
	fmt.Println("Before Sorting:")
	for i, _ := range marks {
		fmt.Println(i, marks[i])
	}
	keys := make([]string, 0, len(marks))

	for k := range marks {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println("\nAfter Sorting:")
	for _, k := range keys {
		fmt.Println(k, marks[k])
	}

}

Output:

Before Sorting:
Amala 410
Eve 490
Daniel 380
Bobby 440
Gabby 350

After Sorting:
Amala 410
Bobby 440
Daniel 380
Eve 490
Gabby 350


```

## Sorting the Map by Values

```
package main

import (
	"fmt"
	"sort"
)

func main() {
	marks := map[string]int{
		"Amala":  410,
		"Eve":    490,
		"Daniel": 380,
		"Bobby":  440,
		"Gabby":  350,
	}
	fmt.Println("Before Sorting:")
	for i, _ := range marks {
		fmt.Println(marks[i])
	}
	values := make([]int, 0, len(marks))

	for _, v := range marks {
		values = append(values, v)
	}

	sort.Ints(values)
	fmt.Println("\nAfter Sorting:")
	for _, v := range values {
		fmt.Println(v)
	}

}

Output:

Before Sorting:
380
440
350
410
490

After Sorting:
350
380
410
440
490

```

## Merging maps
- To merge two different map , their keys and values data-type should match
```
package main

import (
	"fmt"
)

func main() {
	TwoAMarks := map[string]int{
		"Amala":  410,
		"Eve":    490,
		"Daniel": 380,
		"Bobby":  440,
		"Gabby":  350,
	}

	fmt.Println("Before Merging:", TwoAMarks)
	TwoBMarks := map[string]int{
		"Kamala": 410,
		"Chaya":  490,
		"Maya":   380,
	}

	for k, v := range TwoBMarks {
		TwoAMarks[k] = v
	}
	fmt.Println("After Merging:", TwoAMarks)

}

Output:

Before Merging: map[Amala:410 Bobby:440 Daniel:380 Eve:490 Gabby:350]
After Merging: map[Amala:410 Bobby:440 Chaya:490 Daniel:380 Eve:490 Gabby:350 Kamala:410 Maya:380]


```

## Map is a reference type
- Go map is a reference type. It means that when we assign a reference to a new variable or pass a map to a function, the reference to the map is copied.
```
package main

import (
	"fmt"
)

func main() {
	TwoAMarks := map[string]int{
		"Amala":  410,
		"Eve":    490,
		"Daniel": 380,
		"Bobby":  440,
		"Gabby":  350,
	}

	TwoBMarks := TwoAMarks
	fmt.Println("Before Changing Reference:", TwoAMarks)

	TwoBMarks["Maya"] = 470

	fmt.Println("After Changing Reference:", TwoAMarks)
}

Output:

Before Changing Reference: map[Amala:410 Bobby:440 Daniel:380 Eve:490 Gabby:350]
After Changing Reference: map[Amala:410 Bobby:440 Daniel:380 Eve:490 Gabby:350 Maya:470]

```
