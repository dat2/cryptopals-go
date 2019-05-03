package cryptopals

import "testing"

func TestEncodeBase64(t *testing.T) {
    inputHexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    inputBytes := DecodeHexString(inputHexString)
    actualBase64 := EncodeBase64(inputBytes)
    expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    if actualBase64 != expectedBase64 {
        t.Fatalf("%s != %s", actualBase64, expectedBase64)
    }
}
