# Go Basic Types

- **bool** → Boolean type, values are `true` or `false`.
- **string** → Immutable sequence of characters (UTF-8 encoded text).

## Integers

- **int** → Signed integer, size depends on system (32 or 64 bits).
- **int8** → Signed 8-bit integer (−128 to 127).
- **int16** → Signed 16-bit integer (−32,768 to 32,767).
- **int32** → Signed 32-bit integer (−2,147,483,648 to 2,147,483,647).
- **int64** → Signed 64-bit integer (−9e18 to 9e18).

- **uint** → Unsigned integer, size depends on system (32 or 64 bits).
- **uint8** → Unsigned 8-bit integer (0 to 255).
- **uint16** → Unsigned 16-bit integer (0 to 65,535).
- **uint32** → Unsigned 32-bit integer (0 to 4e9).
- **uint64** → Unsigned 64-bit integer (0 to 18e18).
- **uintptr** → Unsigned integer large enough to store pointer values (used in low-level programming).

## Aliases

- **byte** → Alias for `uint8`, represents raw data or ASCII characters.
- **rune** → Alias for `int32`, represents a Unicode code point (Unicode-safe character).

## Floating Point

- **float32** → 32-bit floating-point number (single precision).
- **float64** → 64-bit floating-point number (double precision, default).

## Complex Numbers

- **complex64** → Complex number with `float32` real and imaginary parts.
- **complex128** → Complex number with `float64` real and imaginary parts.
