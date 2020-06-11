package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  rand.Seed(time.Now().UnixNano())
  
  isHeistOn := true

  if eludedGuards := rand.Intn(100); eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good Job remember, this is the fist step.")
  } 
  else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time")
  }   
  openedVault := rand.Intn(100)
  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("Amount Stolen:", amtStolen, "$")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("Vault can't be opened.")
  }
  leftSafely := rand.Intn(5)
  if isHeistOn == false {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Mission failed backup code 84!")
      case 1:
        isHeistOn = false
        fmt.Println("Mission failed backup code 74!")
      case 2:
        isHeistOn = false
        fmt.Println("Mission failed backup code 64!")
      case 3:
        isHeistOn = false
        fmt.Println("Mission failed backup code 54!")
      default:
        fmt.Println("Start the getaway car!")
        }
  }  
}
