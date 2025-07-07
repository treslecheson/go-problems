# Defang an IP Address

Write a function called `convertAddress` that takes in an IP address as a string
and validates and returns a defanged IP address also as a string.
If it isn't a valid address, then return an empty string.

A defanged IP address is formatted to separate the parts of an address
so that software doesn't convert it into a link, such as in documentation.
The method to use in this problem is to replace all `.` with `[.]`.
A valid IP address is one where every octet (the 3 digits between each period)
is within 0 - 255, inclusive.

```go
convertAddress("192.168.1.9") // -> "192[.]168[.]1[.]9"
convertAddress("172.3.265.42") // -> ""
```
