// Test the app.
package main

import (
	"finance/finances"
	// // "app/umath"
	"fmt"
	// "math"
)

// type S struct{}
// func (s *S) addr() { fmt.Printf("%p\n", s) }

func main() {

	// var a, b S
	// a.addr()
	// b.addr()

	// if len(os.Args) > 2 && strings.EqualFold(os.Args[1], "ordinary") {
	//   if os.Args[2] == "interest" && len(os.Args) == 8 {
	//     _, err := strconv.Atoi(os.Args[4])
	//     if err != nil {
	//       panic(err)
	//     }
	//     var si finances.SimpleInterest
	//     var interest = (&si).OrdinaryInterest(100, 0.04, finances.Monthly, 1, finances.Months)
	//     fmt.Printf("interest = $%.2f (Ordinary Interest)\n", interest)
	//   }
	// }

	/***
	var si finances.SimpleInterest
	var interest = (&si).OrdinaryInterest(100, 0.04, finances.Monthly, 1, finances.Months)
	fmt.Printf("interest = $%.2f (Ordinary Interest)\n", interest)
	rate := si.OrdinaryRate(100, interest, finances.Monthly, 1, finances.Months)
	fmt.Printf("rate = %.2f%% (Ordinary Interest)\n", rate*100.0)
	principal := si.OrdinaryPrincipal(interest, rate, finances.Monthly, 1, finances.Months)
	fmt.Printf("principal = $%.2f (Ordinary Interest)\n", principal)
	time := si.OrdinaryTime(principal, interest, rate, finances.Monthly, finances.Months)
	fmt.Printf("time period = %.2f (Ordinary Interest)\n\n", time)
	//
	interest = (&si).OrdinaryInterest(100, 0.04, finances.Annually, 1, finances.Months)
	fmt.Printf("interest = $%.2f (Ordinary Interest)\n", interest)
	rate = si.OrdinaryRate(100, interest, finances.Annually, 1, finances.Months)
	fmt.Printf("rate = %.2f%% (Ordinary Interest)\n", rate*100.0)
	principal = si.OrdinaryPrincipal(interest, rate, finances.Annually, 1, finances.Months)
	fmt.Printf("principal = $%.2f (Ordinary Interest)\n", principal)
	time = si.OrdinaryTime(principal, interest, rate, finances.Annually, finances.Months)
	fmt.Printf("time period = %.2f (Ordinary Interest)\n\n", time)
	//
	interest = (&si).OrdinaryInterest(10000, 0.09, finances.Annually, 153, finances.Days)
	fmt.Printf("interest = $%.2f (Ordinary Interest)\n", interest)
	rate = si.OrdinaryRate(10000, interest, finances.Annually, 153, finances.Days)
	fmt.Printf("rate = %.2f%% (Ordinary Interest)\n", rate*100.0)
	principal = si.OrdinaryPrincipal(interest, rate, finances.Annually, 153, finances.Days)
	fmt.Printf("principal = $%.2f (Ordinary Interest)\n", principal)
	time = si.OrdinaryTime(principal, interest, rate, finances.Annually, finances.Years) //Days
	fmt.Printf("time period = %.4f (Ordinary Interest)\n\n", time)
	***/
	/***
	  var si finances.SimpleInterest
	  var interest = (&si).BankersInterest(100, 0.04, finances.Monthly, 1, finances.Months)
	  fmt.Printf("interest = $%.2f (Banker's Interest)\n", interest)
	  rate := si.BankersRate(100, interest, finances.Monthly, 1, finances.Months)
	  fmt.Printf("rate = %.2f%% (Banker's Interest)\n", rate * 100.0)
	  principal := si.BankersPrincipal(interest, rate, finances.Monthly, 1, finances.Months)
	  fmt.Printf("principal = $%.2f (Banker's Interest)\n", principal)
	  time := si.BankersTime(principal, interest, rate, finances.Monthly, finances.Months)
	  fmt.Printf("time period = %.2f (Banker's Interest)\n\n", time)
	  //
	  interest = (&si).BankersInterest(100, 0.04, finances.Annually, 1, finances.Months)
	  fmt.Printf("interest = $%.2f (Banker's Interest)\n", interest)
	  rate = si.BankersRate(100, interest, finances.Annually, 1, finances.Months)
	  fmt.Printf("rate = %.2f%% (Banker's Interest)\n", rate * 100.0)
	  principal = si.BankersPrincipal(interest, rate, finances.Annually, 1, finances.Months)
	  fmt.Printf("principal = $%.2f (Banker's Interest)\n", principal)
	  time = si.BankersTime(principal, interest, rate, finances.Annually, finances.Months)
	  fmt.Printf("time period = %.2f (Banker's Interest)\n\n", time)
	  //
	  interest = (&si).BankersInterest(10000, 0.09, finances.Annually, 153, finances.Days)
	  fmt.Printf("interest = $%.2f (Banker's Interest)\n", interest)
	  rate = si.BankersRate(10000, interest, finances.Annually, 153, finances.Days)
	  fmt.Printf("rate = %.2f%% (Banker's Interest)\n", rate * 100.0)
	  principal = si.BankersPrincipal(interest, rate, finances.Annually, 153, finances.Days)
	  fmt.Printf("principal = $%.2f (Banker's Interest)\n", principal)
	  time = si.BankersTime(principal, interest, rate, finances.Annually, finances.Years)  //Days
	  fmt.Printf("time period = %.4f (Banker's Interest)\n\n", time)
	  ***/
	/***
	  var si finances.SimpleInterest
	  var interest = (&si).AccurateInterest(100, 0.04, finances.Monthly, 1, finances.Months)
	  fmt.Printf("interest = $%.2f (Accurate Interest)\n", interest)
	  rate := si.AccurateRate(100, interest, finances.Monthly, 1, finances.Months)
	  fmt.Printf("rate = %.2f%% (Accurate Interest)\n", rate * 100.0)
	  principal := si.AccuratePrincipal(interest, rate, finances.Monthly, 1, finances.Months)
	  fmt.Printf("principal = $%.2f (Accurate Interest)\n", principal)
	  time := si.AccurateTime(principal, interest, rate, finances.Monthly, finances.Months)
	  fmt.Printf("time period = %.2f (Accurate Interest)\n\n", time)
	  //
	  interest = (&si).AccurateInterest(100, 0.04, finances.Annually, 1, finances.Months)
	  fmt.Printf("interest = $%.2f (Accurate Interest)\n", interest)
	  rate = si.AccurateRate(100, interest, finances.Annually, 1, finances.Months)
	  fmt.Printf("rate = %.2f%% (Accurate Interest)\n", rate * 100.0)
	  principal = si.AccuratePrincipal(interest, rate, finances.Annually, 1, finances.Months)
	  fmt.Printf("principal = $%.2f (Accurate Interest)\n", principal)
	  time = si.AccurateTime(principal, interest, rate, finances.Annually, finances.Months)
	  fmt.Printf("time period = %.2f (Accurate Interest)\n\n", time)
	  //
	  interest = (&si).AccurateInterest(10000, 0.09, finances.Annually, 153, finances.Days)
	  fmt.Printf("interest = $%.2f (Accurate Interest)\n", interest)
	  rate = si.AccurateRate(10000, interest, finances.Annually, 153, finances.Days)
	  fmt.Printf("rate = %.2f%% (Accurate Interest)\n", rate * 100.0)
	  principal = si.AccuratePrincipal(interest, rate, finances.Annually, 153, finances.Days)
	  fmt.Printf("principal = $%.2f (Accurate Interest)\n", principal)
	  time = si.AccurateTime(principal, interest, rate, finances.Annually, finances.Days)  //Years
	  fmt.Printf("time period = %.4f (Accurate Interest)\n\n", time)
	  ***/

  var m finances.Miscellaneous
  var real = (&m).RealInterestRate(0.045, 0.065)
  fmt.Printf("Real Interest Rate = %.2f%%\n", real * 100)

  // fmt.Println("eps = ", math.Nextafter(1.0, 2.0) - 1.0)
  fmt.Println("Hello, world.")
}