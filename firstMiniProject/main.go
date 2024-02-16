package main

import "fmt"

func main()  {

  //Print Line
  fmt.Println("Hello, Trenton!")

  //Focusing on strings
  var name string = "lured"
  var nametwo = "luigi"
  var nameThree string
  nameThree = "Trenton"
  message := fmt.Sprintf("%s, %s, %s", name , nametwo, nameThree)
  fmt.Println(message)

  //Focusing on Numbers
  var numOne int = 20
  var numTwo = 30
  numThree := 40
  var numFour int8 = 114
  fmt.Printf("numOne: %d, numTwo: %d, numThree: %d, numFourIntType: %d \n", numOne, numTwo, numThree, numFour)
  var unsignedOne uint8 = 34
  var unsignedTwo uint16 = 145
  fmt.Printf("unsignedOne: %d, unsignedTwo: %d \n", unsignedOne, unsignedTwo)
  var floatOne float32 = 5.12
  var floatTwo float64 = 668787.32
  fmt.Printf("floatOne: %f, floatTwo: %0.1f \n", floatOne, floatTwo)

  //Arrays(fixed length)
  var ages [3]int = [3]int{2, 4, 5}
  var nums = [4]int{1,5,8}
  fmt.Println(ages, len(ages))
  fmt.Println(nums, len(nums))
  names:= [4]string{"trenton", "cole", "test", "nice"}
  fmt.Println(names, len(names))

  //Slices
  var scores = []int{100, 50, 60}
  fmt.Println(scores, len(scores))
  //append returns copy
  scores = append(scores, 123)
  fmt.Println(scores, len(scores))

  //Slice Ranges
  //[x,y) inclusive first not second
  rangeOne := names[1:3]
  fmt.Println(rangeOne)
  rangeTwo := names [1:]
  fmt.Println(rangeTwo)
  rangeThree := names[:3]
  fmt.Println(rangeThree)

}
