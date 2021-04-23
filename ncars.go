package main
import (
  "fmt"
)

func main() {

  // 1 Car

  fmt.Printf("Cost of car 1: ")
  var car1Cost float64
  fmt.Scanf("%f", &car1Cost)

  fmt.Printf("Monthly revenue of car 1: ")
  var car1MonthlyRevenue float64
  fmt.Scanf("%f", &car1MonthlyRevenue)

  fmt.Printf("Cost of car 1 is recovered after %.1f months.",
    car1Cost/car1MonthlyRevenue)

  // n Cars -> n+1 Cars

  n := 1;
  maxCars := 2;
  nCarMonthlyRevenue := car1MonthlyRevenue;
  var nextCarCost float64
  var nextCarOutsideContributions float64
  var nextCarMonthlyRevenue float64
  var anotherCar bool

  for n < maxCars {

    fmt.Printf("\n\nCost of car %d: ", n+1)
    fmt.Scanf("%f", &nextCarCost)

    fmt.Printf("Existing car revenue alone will buy car %d after %.1f months.\n\n",
      n+1, nextCarCost/nCarMonthlyRevenue)

    fmt.Printf("Outside contributions toward purchase of car %d: ", n+1)
    fmt.Scanf("%f", &nextCarOutsideContributions)

    fmt.Printf("With this, car %d can be purchased after %.1f months\n\n",
      n+1, (nextCarCost - nextCarOutsideContributions)/nCarMonthlyRevenue)

    // n+1 Cars

    fmt.Printf("Monthly revenue of car %d: ", n+1)
    fmt.Scanf("%f", &nextCarMonthlyRevenue)

    totalMonthlyRevenue := nCarMonthlyRevenue + nextCarMonthlyRevenue;

    fmt.Printf("Total monthly revenue is now %.2f\n\n",
      totalMonthlyRevenue)

    fmt.Printf("Cost of car %d is recovered after %.1f months.\n\n",
      n+1, nextCarCost/totalMonthlyRevenue)

    fmt.Printf("Buying another car? (true/false) ")
    fmt.Scanf("%t\n", &anotherCar)

    n++

    if anotherCar {
      maxCars++
      nCarMonthlyRevenue = totalMonthlyRevenue
    }

  }

  fmt.Printf("--------------------------------------------------------------\n\n")

  fmt.Printf("\n\n Revenue per year with these %d cars: %.2f",
    n, nCarMonthlyRevenue*12)

  fmt.Printf("\n Total after 2 years: %.2f", nCarMonthlyRevenue*12*2)
  fmt.Printf("\n Total after 3 years: %.2f", nCarMonthlyRevenue*12*3)
  fmt.Printf("\n Total after 4 years: %.2f", nCarMonthlyRevenue*12*4)

}
