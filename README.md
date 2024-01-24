# Anagram testing function

Strings are defined inside the main function.

The first string test is not an anagram. This is due to the capitalization. I did not normalize all characters to lowercase.

The second string test is a valid anagram.

`go run anagram.go`

```
$ go run anagram.go
String A:[This is a test!] length:15  String B:[is This test! A] length:15
Strings are unfortunately not anagrams!
String A:[This is a test!] length:15  String B:[This is test! a] length:15
Strings are anagrams!
```
