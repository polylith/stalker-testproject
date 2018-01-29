package fizzbuzz

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "strings"
)

func TestReverse1(t *testing.T) {
  var result string = fizzbuzz()
  splitted_result := strings.Split(result, "\n")
  assert.Equal(t, "1", splitted_result[1])
  assert.Equal(t, "2", splitted_result[2])
  assert.Equal(t, "fizz", splitted_result[3])
  assert.Equal(t, "4", splitted_result[4])
  assert.Equal(t, "buzz", splitted_result[5])
  assert.Equal(t, "fizz", splitted_result[6])
  assert.Equal(t, "7", splitted_result[7])
  assert.Equal(t, "8", splitted_result[8])
  assert.Equal(t, "fizz", splitted_result[9])
  assert.Equal(t, "buzz", splitted_result[10])
  assert.Equal(t, "11", splitted_result[11])
  assert.Equal(t, "fizz", splitted_result[12])
  assert.Equal(t, "13", splitted_result[13])
  assert.Equal(t, "14", splitted_result[14])
  assert.Equal(t, "fizzbuzz", splitted_result[15])
}
