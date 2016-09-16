package main

import ("fmt"
)

func main() {

  cases := readCases() //Lectura de casos

  for i:=1; i<= cases;i++{
    pancakes := readPancakes()

    max := getMax(pancakes)
    counter := make([]int, max+1)


    for j:=0; j < len(pancakes); j++{
      counter[pancakes[j]]++
    }
    moves:=max;

    for k:=1; k <= max;k++{
      splits:=0
      for l:=1;l <= max;l++{
        splits += ((l - 1) / k) * counter[l]
      }
      if (splits+k) <= moves{
        moves = splits + k
      }
    }
    fmt.Printf("Case #%v: %v\n", i,moves)
  }
}

func readCases() int {
  var i int
  fmt.Scanln(&i)
  return i
}

func readPancakes() []int{
  var j int
  fmt.Scanln(&j)
  pancakesarray := make([]int, j)
  for a := 0; a < j; a++{
  var m int
  fmt.Scan(&m)
  pancakesarray[a] = m
  }
  return pancakesarray
}

func getMax(pancakes []int) int {
  max := 0
  for a := 0; a < len(pancakes); a++ {
    if pancakes[a] > max{
      max=pancakes[a]

    }
  }
  return max
}
