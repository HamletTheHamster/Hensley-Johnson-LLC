package main
import (
  "fmt"
)

func main() {

  fmt.Printf("\n\n")
  fmt.Printf("  -----------------\n")
  fmt.Printf(" | nCar Calculator |        by Joel\n")
  fmt.Printf("  -----------------\n")

  // Definitions

  var car1Cost,
      car1MonthlyRevenue,
      nextCarCost,
      nextCarMonthlyRevenue,
      nextCarOutsideContributions,
      addMonths float64
  var anotherCar bool

  // 1 Car

  fmt.Printf("\n\nCost of car 1: ")
  fmt.Scanf("%f", &car1Cost)

  fmt.Printf("Monthly revenue of car 1: ")
  fmt.Scanf("%f", &car1MonthlyRevenue)

  monthsFromNow := car1Cost/car1MonthlyRevenue;

  fmt.Printf("Cost of car 1 is recovered after %.1f months, ",
    monthsFromNow)
  fmt.Printf("or %.1f months from now.", monthsFromNow)

  // n Cars -> n+1 Cars

  n := 1;
  maxCars := 2;
  nCarMonthlyRevenue := car1MonthlyRevenue;

  for n < maxCars {

    fmt.Printf("\n\nCost of car %d: ", n+1)
    fmt.Scanf("%f", &nextCarCost)

    fmt.Printf("Existing car revenue alone will buy car %d after %.1f months, ",
      n+1, nextCarCost/nCarMonthlyRevenue)

    fmt.Printf("or %.1f months from now.",
      monthsFromNow + nextCarCost/nCarMonthlyRevenue)

    fmt.Printf("\n\nOutside contributions toward purchase of car %d, if any: ",
      n+1)
    fmt.Scanf("%f", &nextCarOutsideContributions)

    addMonths = (nextCarCost - nextCarOutsideContributions)/nCarMonthlyRevenue

    if nextCarOutsideContributions != 0 {
      fmt.Printf("With this, car %d can be purchased after %.1f months, ",
        n+1, addMonths)

      monthsFromNow -= nextCarCost/nCarMonthlyRevenue
      monthsFromNow += addMonths

      fmt.Printf("or %.1f months from now.", monthsFromNow)

    } else {
      monthsFromNow += nextCarCost/nCarMonthlyRevenue
    }

    // n+1 Cars

    fmt.Printf("\n\nMonthly revenue of car %d: ", n+1)
    fmt.Scanf("%f", &nextCarMonthlyRevenue)

    totalMonthlyRevenue := nCarMonthlyRevenue + nextCarMonthlyRevenue;

    fmt.Printf("Total monthly revenue is now %.2f\n\n",
      totalMonthlyRevenue)

    fmt.Printf("Cost of car %d is recovered after %.1f months, ",
      n+1, nextCarCost/totalMonthlyRevenue)

    monthsFromNow += nextCarCost/totalMonthlyRevenue

    fmt.Printf("or %.1f months from now.", monthsFromNow)

    fmt.Printf("\n\nBuying another car? (true/false) ")
    fmt.Scanf("%t\n", &anotherCar)

    n++

    if anotherCar {
      maxCars++
      nCarMonthlyRevenue = totalMonthlyRevenue
    }
  }

  fmt.Printf("----------------------------------------------------------------")

  yearlyRevenue := nCarMonthlyRevenue*12

  fmt.Printf("\n\n\n\nRevenue per year with these %d cars: %.2f",
    n, yearlyRevenue)

  fmt.Printf("\nTotal after 2 years: %.2f", yearlyRevenue*2)
  fmt.Printf("\nTotal after 3 years: %.2f", yearlyRevenue*3)
  fmt.Printf("\nTotal after 4 years: %.2f", yearlyRevenue*4)

  fmt.Printf("\n\n")

}
