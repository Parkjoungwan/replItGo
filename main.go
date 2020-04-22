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

func nojam2747(){
  var N int
  fmt.Scan(&N)
  var F[46] int
  F[1]=1
  for i:=2; i<=N; i++{
    F[i]=F[i-1]+F[i-2]
  }
  fmt.Println(F[N])
}

func nojam10984(){
  var T int
  fmt.Scan(&T)
  var sum2[10] float32
  var sum1[10] float32
  var result1 float32
  var result2 float32
  for i:=0; i<T; i++{
    var N int
    fmt.Scan(&N)
    for j:=0; j<N;j++{
      var C float32
      fmt.Scan(&C)
      sum1[j]=C
      var G float32
      fmt.Scan(&G)
      sum2[j]=C*G
    }
    for k:=0; k<N; k++{
      result1+=sum1[k]
      result2+=sum2[k]
    }
    fmt.Printf("%.0f %.1f\n", result1, result2/result1)
    result1=0
    result2=0
  }
}

func nojam1149(){
  
}

func main() {
  nojam10984()
}