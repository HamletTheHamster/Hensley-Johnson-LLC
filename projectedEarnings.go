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

  fmt.Printf("Cost of car 1 is recovered after %.1f months.\n\n",
    car1Cost/car1MonthlyRevenue)

  // n Cars -> n+1 Cars

  n := 1;
  maxCars := 2;
  nCarMonthlyRevenue := car1MonthlyRevenue;

  for n < maxCars {

    fmt.Printf("Cost of car %d: ", n+1)
    var nextCarCost float64
    fmt.Scanf("%f", &nextCarCost)

    fmt.Printf("Car %d revenue alone will buy car %d after %.1f months.\n\n",
      n, n+1, nextCarCost/nCarMonthlyRevenue)

    fmt.Printf("Outside contributions toward purchase of car %d: ", n+1)
    var nextCarOutsideContributions float64
    fmt.Scanf("%f", &nextCarOutsideContributions)

    fmt.Printf("With this, car %d can be purchased after %.1f months\n\n",
      n+1, (nextCarCost - nextCarOutsideContributions)/nCarMonthlyRevenue)

    // n+1 Cars

    fmt.Printf("Monthly revenue of car %d: ", n+1)
    var nextCarMonthlyRevenue float64
    fmt.Scanf("%f", &nextCarMonthlyRevenue)

    totalMonthlyRevenue := nCarMonthlyRevenue + nextCarMonthlyRevenue;

    fmt.Printf("Total monthly revenue is now %.2f\n\n",
      totalMonthlyRevenue)

    fmt.Printf("Cost of car %d is recovered after %.1f months.\n\n",
      n+1, nextCarCost/totalMonthlyRevenue)

    fmt.Printf("Buying another car? (true/false) ")
    var anotherCar bool
    fmt.Scanf("%t", anotherCar)

    n++
    if anotherCar {
      maxCars++
    }
  }

}
