package cryptopals

import (
    "strings"
    "fmt"
)

func decodeAsciiToInt(c byte) byte {
    switch {
    case byte('a') <= c && c <= byte('f'):
        return 10 + (c - byte('a'))
    case byte('0') <= c && c <= byte('9'):
        return c - byte('0')
    default:
        return 0
    }
}

func DecodeHexString(input string) []byte {
    result := make([]byte, len(input) / 2)
    for i := 0; i < len(input) - 1; i += 2 {
        firstNybble, secondNybble := decodeAsciiToInt(input[i]), decodeAsciiToInt(input[i+1])
        result[i/2] = (firstNybble << 4) | secondNybble
    }
    return result
}

func encodeByteToBase64Char(b byte) rune {
    switch {
    case b < 26:
        return rune(byte('A') + b)
    case b < 52:
        return rune(byte('a') + (b - 26))
    case b < 62:
        return rune(byte('0') + (b - 52))
    case b == 62:
        return '+'
    case b == 63:
        return '/'
    default:
        panic(fmt.Sprintf("invalid byte: %v", b))
    }
}


func EncodeBase64(x []byte) string {
    var result strings.Builder
    for i := 0; i < len(x); i += 3 {
        first, second, third, fourth := (x[i] & 0xFC) >> 2, ((x[i] & 0x03) << 4) | ((x[i+1] & 0xF0) >> 4), ((x[i+1] & 0x0F) << 2) | ((x[i+2] & 0xC0) >> 6), (x[i+2] & 0x3F)
        result.WriteRune(encodeByteToBase64Char(first))
        result.WriteRune(encodeByteToBase64Char(second))
        result.WriteRune(encodeByteToBase64Char(third))
        result.WriteRune(encodeByteToBase64Char(fourth))
    }
    return result.String()
}
