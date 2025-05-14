# 1) Convert Temperature

Write a function `toFahrenheit(celsius: number): number` that accepts a Celsius temperature as a `number` and returns the converted Fahrenheit temperature rounded to one decimal place, if applicable. If the result is an integer don't worry about trying to force empty decimals
(`212` is fine, you don't need to make sure it's explicitly `212.0`).

Celsius to Fahrenheit:
`F = C * (9/5) + 32`

#### Rounding:

There is a `.toFixed(x)` method that will allow you to round to `x` decimal places. You may use this if you wish,
but know that it returns a `string` and you will still need to convert it to a `number`.

Alternatively, you can use this method:

If you want to round 3.14159 to 2 decimal places, multiply the value by `10^x` where `x` is the number of decimal places.
Then use `Math.round()` to round the resulting value to the nearest whole number and divide by the factor again to get
the right decimals.

```typescript
const val = 3.14159;
const factor = 10 ** 2;
const rounded = Math.round(val * factor) / factor;

// 3.14159 * 100 = 314.159
// Math.round(314.159) = 314
// 314 / 100 = 3.14
```
