package main

import "fmt"

func nojam17173() {
  var N int
  var M int
  fmt.Scan(&N)
  fmt.Scan(&M)
  var K[1001] int
  var Mul[1001] int

  for i:=0; i<M; i++{
    fmt.Scan(&K[i])
  }

  for i:=0; i<M; i++{
    for j:=0; j<=N; j++{
      if j%K[i]==0 {Mul[j] = j}
    } 
  }
  sum:=0

  for i:=1; i<=N; i++{
    sum+=Mul[i]
  }

  fmt.Println(sum)

}

func main() {
  nojam17173()
}