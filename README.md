# Additional Notes

## Doubly linked list excercise
[Solution link](https://go.dev/play/p/Y4BTyoZv-AW)

## Data Stream: Rolling mean
[Solution link](https://go.dev/play/p/RZM8zndpyn2)

## Dynamic programming: Affordable path (using matrix)
[Solution link](https://go.dev/play/p/ByvSFqfVu1w)

## Interface and method inheritance

```go
package main

import "fmt"

type Err struct{
  Val string
}

type Str struct{
  Val string
}

func (e *Str) String() string{
  return fmt.Sprintf("Value is: %s", e.Val)
}

func (e *Err) String() string{
  return fmt.Sprintf("Value is: %s", e.Val)
}

func (e *Err) Error() string{
  return fmt.Sprintf("Error is thrown: %s", e.Val)
}


func errChecker(val string) error{
  return  &Err{val}
}


func strChecker(val string) fmt.Stringer{
  return  &Str{val}
}

func main(){

  err := errChecker("Hi")
  fmt.Println(err)
  fmt.Println(&Err{"Hello"}) //shows "Error is thrown: Hello" instead of "Valuse is: Hello"
  fmt.Println(&Str{"HI!"})

}
```

In Go, if a type implements the error interface, which has the Error() method with the signature Error() string, then it automatically satisfies the Stringer interface, which has the String() string method. The String() method is used by the fmt package when printing values with %v or Println. If both Error() and String() methods are present, the Error() method takes precedence.

In your Err type, you have both String() and Error() methods. When you print an instance of Err using fmt.Println(&Err{"HI!"}), the Error() method will be called, as it satisfies the error interface. If you want the String() method to be used, you can explicitly call it:

```go
fmt.Println((&Err{"HI!"}).String())
```
