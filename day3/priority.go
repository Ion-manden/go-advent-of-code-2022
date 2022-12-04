package day3

import "strings"

func getItemPriority(item string) int {
  chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  return strings.Index(chars, item) + 1
}
