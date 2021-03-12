# Theory
## 1.0 main package
- main is entry point
- package and function
- func not function
- main is for compile

## 1.1 packages and imports
- fmt: package for formatting
- import if I want to export function,
function have to start uppercase
- create something folder and we can export SayHello becuase it's start with uppercase
```go
package main

import (
	"fmt"

	"github.com/serin0837/learngo/something"
)

func main() {
	fmt.Println("hello world")
	something.SayHello()
}
```

## 1.2 variables and constants
const : const
var : let 
`const name string = "nico"`

var name string = "nico"
	name = "lynn"
	fmt.Println(name)

`name := "nico"` (same as var)
go is going to find type for me 
and I can not change type

shorthand only inside function and work with variable 

## 1.3 function part one 

- type : stirng, boolean and manynumbers

```go
func multiply(a, b) {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}

```
not working like this 

```go

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
```
working now 
add int

shorthand
```go
func multiply(a, b int) int {
	return a * b
}
```

- function can return multiple values
```go

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
}
```

```go

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)
}
```

- with `_` go will ignore value of uppername and withoud `_` would not work 

- go can take multiple arguments
```go
func repeatMe(words ...string) {
	fmt.Println(words)
}
//[nico lynn dal marl flynn]
func main() {
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
```

- **I dont have to put return value's type if it's not returning anything**


## 1.4 function part two 
- go have naked return 
```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("nico")
	fmt.Println(totalLength, up)
}
```

becuase we alrady mention what we want to return with variable 
and return nothing

but we can also do return length, uppercase but we don't have to 

nico does not use that often 

- defer 
defer is executed when function is finished
 really cool!!!!

```go
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I am done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("nico")
	fmt.Println(totalLength, up)
}

// I am done
//4 NICO
```
## 1.5 loop
- for, range
foreach in js
```go

func superAdd(numbers ...int) int {
	for number := range numbers {
		fmt.Println(number)
	}
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}

//0 1 2 3 4 5 -> index
```

```go

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
//index and number together
```
```go
func superAdd(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
//123456
```
 we can use range only inside for 

 ```go
 func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}//21
```

## 1.6 if/ else
- no need () or :
```go
func canIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(canIDrink(16))
}
```
also can get rid of else block 

- can create variable right in the moment when creating if block (variable expression)
```go
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
```
same as below code, 
if you but variable only in the if block, then it's good for peopler to think that you only use this variable only in the if block!!
```go
func canIDrink(age int) bool {
	koreanAge := age + 2
	if koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
```

## 1.6 switch
almost like js
```go

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
```

```go
func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
```
can also do variable expression
```go

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
```

👇from here only thing for GO!!
## 1.8 pointers
- low-level programming - get into the memory memory address

```go
func main() {
	a := 2
	b := a // b copy the value of a
	a = 10 // so even though we update a, b does not care
	fmt.Println(a, b)
}
```

```go
func main() {
	a := 2
	b := 5
	fmt.Println(&a, &b)
}
// 0xc0000ae008 0xc0000ae010
```

I want to point to the address not copy the value of a then 
I need to do `b := &a`

```go
func main() {
	a := 2
	b := &a
	fmt.Println(&a, b)
}
//0xc000198000 0xc000198000// same address
```

`*` is see through, look through
```go
func main() {
	a := 2
	b := &a
	fmt.Println(*b)
}//2
```

```go
func main() {
	a := 2
	b := &a
	fmt.Println(b)
}//0xc0000160b0
```
this is how you can make your program fast because we don't want to copy sometime 

```go
func main() {
	a := 2
	b := &a
	a = 10
	fmt.Println(*b)
}// 10 because b is connected to a now 
```
I can update a with b 
```go
func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
}
```

## 1.9 Arryas and Slice
- Arrya in GO is quite different with other laguages
becuase we have to specity the length of array 
```go
func main() {
	names := [5]string{"nico", "lynn", "dal"}
	fmt.Println(names)
}
```
```go
func main() {
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "serin"
	names[4] = "serin"
	fmt.Println(names)
}
```

- sometime we want to limit array, sometime we want unlimited item in 
that is slice, slice in go is array without length 
```go
func main() {
	names := []string{"nico", "lynn", "dal"}

	fmt.Println(names)
}
```
- but we can not add more itmes like `names[4] ="lala"`
- then we have to use append method
- return new slice!!!!
```go
func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}
```

## 1.10 Maps
- a bit like js and python's object not 100%
- `map[key type]value type {stuff}`
```go
func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
}
//map[age:12 name:nico]
```
- iteration
```go

func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}

}
//name nico
//age 12
```

```go
func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for _, value := range nico {
		fmt.Println(value)
	}

}// nico
//12
```

## 1.11 Structs
- structs like object and they are more flexible than map 
- what is if we want 
{"name":"serin", "age":12}

- for this we have structs
```go
type person struct {
	name    string
	age     int
	favFood []string // slice of string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico)
}
```
it does not that good cuase it's not clear

```go
type person struct {
	name    string
	age     int
	favFood []string // slice of string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}
```
above is better

go does not have class
go does not have constructor method 
we have to implement constructor ourself

a lot of things come fromt structs so it's really important



# 2.Firt Project

## 2.0 
had error about module so I had to type `go mod init github.com/serin0837/learngo` in terminal and working



// need to start with B uppercase cause we want to export and want to use in main.go
// and I have to add comment!! vscode give me small error to put comment about export struct
//owner and balance need to start with capital as well becuase it need to be public


if we change balance to lowercase, people can not change money balalnce

so change back to private
and we have to make constructor
and change back the struct to private

```go
type account struct {
	owner   string
	balance int
}
```
to demonstrate constructor we back to Account,
```go
type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}
```

after change folder name everything is not working 

```go

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)	
}
//&{nico 0}
```

now we can not change account.balabce = 1000 

## 2.1 -2.2 Methods part One and Two
- have to create method
- add (a Account) in between of func and func name
```go
func (a Account) Deposit(amount int){
}
```
- (a Account) is receiver
- a can be anything
- convention is lowercase of struct

```go
func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
} // 0 now so it's not actually change the value of account balance
```
- because before it made a copy of account
a is 
- Balance method is okay to copy becuase we just want value of balance
- Deposit method is not okay to copy becuase we want to change the value of balance
- so *Account (pointer receiver)
- and now 10

- withdraw method with pointer receiver

- if we don't have money we should not able to withdraw money

- need to write error ourself
- so go does not know how to show error it's just not working with withdraw function 
- so we have to write function
```go
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
```
- Fatalln - println and os.Exit (kill program)
- create ChangeOwner, Owner method
- go can call method automatically

- string method will be called without I type
```go

func(a Account) String() string{
	return "whatever you want "
}
```

```go
fmt.Print(account)
//whatever you want 
```

```go
func(a Account) String() string{
	return fmt.Sprint(a.Owner(), "'s account.\nHas:", a.Balance())
}
//nico's account.
//Has:10
```

code of main.go for account project
```go
package main

import ("fmt"

"github.com/serin0837/learngo/accounts"

"log"
)


func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Print(account)
	fmt.Println(account.Balance())
	fmt.Println(account.Balance(),account.Owner())

	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	
}

```

# 3. Secound project Dictionary 
- Dictionary is alias(가명) for map 
- create Dictionary
```go
package mydict

//Dictionary type
type Dictionary map[string]string

```

```go
package main

import ("fmt"

"github.com/serin0837/learngo/mydict"
)


func main() {
	//create dictionary 
	dictionary := mydict.Dictionary{}
	//add something in dictioanry 
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
}

```

- we can use method to a type not only to struct

- we can add methoes like add, delete etc
- we can do like this and can find word but
```go
func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary["first"])
}
//First word 
```

- we can add method

```go
var errNotFound = errors.New("Not found")
//methods // retrun string and error 
//Search for word
func (d Dictionary) Search(word string) (string, error){
	//map have two value// exist is boolean
	value, exists := d[word]	
	if exists{
		//nil (error is nil no error)//we have to return two values
		return value, nil
	}
	//return two values, first one is nothing cause no value
	return "",errNotFound
}
```
```go
package main

import (
"fmt"
"github.com/serin0837/learngo/mydict"
)


func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(definition)
	}
}
```

- add add method
```go
func (d Dictionary) Add(word, def string) error{
	//if there is word already that would throw error
	//we don't need definition
	_, err := d.Search(word)
	if err == errNotFound{
		d[word] = def
	}else if err == nil {
		return errWordExists
	}
}
```
- can write in switch as well
```go
switch err {
case errNotFound:
	d[word] =def
case nil:
	return errWOrdExists
}
``` 
```go
func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	err := dictionary.Add("hello", "Greeting")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(dictionary)
}
```

```go

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	err := dictionary.Add("hello", "Greeting")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(dictionary)
	definition, err2 := dictionary.Search("hello")
	if err2 != nil{
		fmt.Println(err2)
	}else{
		fmt.Println(definition)
	}	
}

```

- add update and delete methods
```go

func(d Dictionary) Update(word, def string)err{
	//search the word exist first
	_, err := d.Search(word)
	switch err{
	//if there is no word
	case errNotFound:
		return errCantUpdate
	//word exist so we can update word
	case nil:
		d[word] = def
	}
	return nil
}
```

```go

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil{
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word) 
}

```

- delete function
```go
func(d Dictionary) Delete (word string){
//delete function is built in Go
	delete(d, word)
}
```

```go

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	fmt.Println(dictionary)
	dictionary.Delete(baseWord)
	fmt.Println(dictionary)
}
```