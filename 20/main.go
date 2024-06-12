package main

import "fmt"

func reverseChar(char rune) rune {
  if char == '[' {
    return ']'
  }

  if char == '{' {
    return '}'
  }

  return ')'
}

func isValid(s string) bool {
  if s == "" { return true }

  queue := ""

  for _, char := range s {
    // fmt.Printf("char %q\n", char)
    // fmt.Printf("queue %s\n", queue)

    if char == '{' || char == '[' || char == '(' {
      queue = string(reverseChar(char)) + queue
    } else {
      if len(queue) == 0 || char != rune(queue[0]) {
        return false
      }
      queue = queue[1:]
    }
  }


  return len(queue) == 0
}

func main() {
  function := isValid
  s := ""

  s = "{}[]()"
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", true)

  s = ""
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", true)

  s = "[]"
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", true)

  s = "({()[]})"
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", true)



  s = "[["
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", false)

  s = "([)]"
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", false)

  s = "]["
  fmt.Printf("Input    %v\n", s)
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", false)
}

