package calculate_sum

import(
  "strconv"

  "fmt"
)

func calculateSum(number string, offset int64, channel ...chan float64) float64 {

  var sum float64

  length, _ := strconv.ParseInt(number, 10, 32)
  length = int64(length)

  for i := offset + 1; i <= length; i++ {
    sum += 1.0 / float64(i)
  }

  if len(channel) > 0 {
    channel[0] <- sum
  }

  return sum

}

func Calculate(number string) string {

  return fmt.Sprintf("%v", calculateSum(number, 0))

}

func CalculateRoutine(number string) string {

  var half_length1, half_length2 int64

  if s, err := strconv.ParseInt(number, 10, 32); err == nil {
    half_length1 = int64(s / 3)
    half_length2 = int64(2 * s / 3)
  }

  channel := make(chan float64)

  go calculateSum(fmt.Sprintf("%v", half_length1), 0, channel)
  go calculateSum(fmt.Sprintf("%v", half_length2), half_length1, channel)
  go calculateSum(number, half_length2, channel)

  a1, a2, a3 := <-channel, <-channel, <-channel

  return fmt.Sprintf("%v", a1+a2+a3)

}
