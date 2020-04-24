package main

import (
    "bufio"
    "os"
    "fmt"
)
var wr = bufio.NewWriter(os.Stdout)

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
func min(a int,b int) int{
  if a>b {
      return b
  }else{return a}
}

func nojam1149(){
  var N int
  var dp [1000][3] int
  var arr [1000][3] int
  fmt.Scan(&N)
  for i:=0; i<N;i++{
    for j:=0;j<3;j++{
      fmt.Scan(&arr[i][j])
    }
  }
  dp[0][0]=arr[0][0]
  dp[0][1]=arr[0][1]
  dp[0][2]=arr[0][2]
  for i:=1;i<N;i++{
    dp[i][0]= min(dp[i-1][1],dp[i-1][2])+arr[i][0]
    dp[i][1]= min(dp[i-1][0],dp[i-1][2])+arr[i][1]
    dp[i][2]= min(dp[i-1][0],dp[i-1][1])+arr[i][2]
  }
  result:=2147438647
  for i:=0;i<3;i++{
    if dp[N-1][i]<result{result=dp[N-1][i]}
  }
  fmt.Println(result)
}

func nojam10995(){
  var N int
  fmt.Scan(&N)
  for i:=0;i<N;i++{
    if i%2==0{
      for j:=0;j<N;j++{
        fmt.Print("*")
        fmt.Print(" ")
      }
    } else{
      for k:=0;k<N;k++{
        fmt.Print(" ")
        fmt.Print("*")
      }
    }
    fmt.Println()
  }
}
func nojam16433case1(N int){
  for i:=0;i<N;i++{
    if i%2==0{
      for j:=0;j<N;j++{
        if j%2==0{
          fmt.Print(".")
        }else{
          fmt.Print("v")
        }
      }
    }else{
      for k:=0;k<N;k++{
        if k%2==0{
          fmt.Print("v")
        }else{
          fmt.Print(".")
        }
      }
    }
    fmt.Println()
  }
}
func nojam16433case2(N int){
  for i:=0;i<N;i++{
    if i%2==0{
      for j:=0;j<N;j++{
        if j%2==0{
          fmt.Print("v")
        }else{
          fmt.Print(".")
        }
      }
    }else{
      for k:=0;k<N;k++{
        if k%2==0{
          fmt.Print(".")
        }else{
          fmt.Print("v")
        }
      }
    }
    fmt.Println()
  }
}
func nojam16433(){
  var N int
  var R int
  var C int
  fmt.Scan(&N, &R, &C)
  if R%2==0{
    if C%2==0{
      nojam16433case2(N)
    }else{
      nojam16433case1(N)
    }
  }else{
    if C%2==0{
      nojam16433case1(N)
    }else{
      nojam16433case2(N)
    }
  }
}
func nojam11650(){
  
}

func nojam15649(){
    var N int
    var M int
  
    fmt.Scan(&N, &M)
    var Arr [] int = make([]int,N)
    var Select [] int = make([]int, N)
    var check [] bool = make([]bool,N) 
    for i:=0;i<N;i++{
      Arr[i]=i+1
    }
    nojam15649DFS(Arr,Select,check,0,M,N)
    wr.Flush()
}

func nojam15649DFS(Arr[] int,Select[] int,check[] bool,Cnt int, M int, N int){
    if(Cnt==M){
      for i:=0;i<M;i++{
        wr.WriteByte(byte(Select[i]) + '0')
        wr.WriteByte(' ')
      }
      wr.WriteByte('\n')
      return
    }
    for i:=0;i<N;i++{
      if(check[i]==true){continue}
      check[i] = true;
      Select[Cnt]=Arr[i];
      nojam15649DFS(Arr,Select,check,Cnt+1,M,N);
      check[i] = false;
    }
}
func nojam15649(){
  
}

func main() {
  nojam15649()
}