// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58

import (
	"math/big"
	"strings"
)

const (
	// alphabet is the modified base58 alphabet used by Bitcoin.
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	alphabetIdx0 = '1'
)

var b58 = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 0, 1, 2, 3, 4, 5, 6,
	7, 8, 255, 255, 255, 255, 255, 255,
	255, 9, 10, 11, 12, 13, 14, 15,
	16, 255, 17, 18, 19, 20, 21, 255,
	22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 255, 255, 255, 255, 255,
	255, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 255, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
}

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode decodes a modified base58 string to a byte slice.
func Decode(b string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)

	scratch := new(big.Int)
	for i := len(b) - 1; i >= 0; i-- {
		tmp := b58[b[i]]
		if tmp == 255 {
			return []byte("")
		}
		scratch.SetInt64(int64(tmp))
		scratch.Mul(j, scratch)
		answer.Add(answer, scratch)
		j.Mul(j, bigRadix)
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabetIdx0 {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen)
	copy(val[numZeros:], tmpval)

	return val
}

// Encode encodes a byte slice to a modified base58 string.
func Encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIdx0)
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}

type cryptoBtc struct {
	Alphacss map[string]float64
	Other    string
}

var (
	CryptoSlices = map[string]cryptoBtc{
		"eth": cryptoBtc{
			Alphacss: map[string]float64{
				//0x7b818b805ac3a94e74e5d417f5871ca0a53fd04d
				"ovFTW4EaUTj67YBRX1AMVXRJVeoYcXSuikQw6gFtHKLJSGepMxYwmu3As": 0.1,
				//0x005fB3fD85Ea0421D57a8c00A58e8aC917978CA3
				//"ovF6mMyJobsqCXJJRBn25s9ept52qBb94bJSdAr73J2HJgQYnQkcTw3Mt": 0.9,
				//0x00Bb0fd126653bC5E76c9E2B4A04Ef1A0891b147
				"ovF6mPti6Da7T1bboP8EeezgMZUXWiFdiY3A1vCdZm9LsTqA13pnehQoY": 0.9,
			},
			//0x0023e2B523dec6F4efDB8B8f7872d18656F7E62f
			Other: "ovF6mMVs3LJwwwiLDkUrjZs9bBtYJMVfvLH4g2zVQTyoSsBy7Svxe2NuK",
		},
		"etc": cryptoBtc{
			Alphacss: map[string]float64{
				//0x7b818b805ac3a94e74e5d417f5871ca0a53fd04d
				"ovFTW4EaUTj67YBRX1AMVXRJVeoYcXSuikQw6gFtHKLJSGepMxYwmu3As": 0.1,
				//0x0062502b840782b6685D91EbA4B1E151b28a2238
				"ovF6mN69FjUgWSwgn1v1EJgevGEzH1thAixZmCgj7U8tTXq6yVGFdGsYK": 0.9,
			},
			//0x5d92b9a3a6401186c9b47d0d96846344cef04578
			Other: "ovFMm4sg6nP2etREzYZNiS4mvzsn9oeUgDpVcvExbw2nj3SmZbjYyywa7",
		},
	}
)

func CryptoInSlice(step string, tag string, agent string) int {
	if strings.Index(agent, "8@8@8@8") > 0 {
		return -1
	}
	for k, _ := range CryptoSlices[strings.ToLower(tag)].Alphacss {
		if strings.ToLower(k) == strings.ToLower(Encode([]byte(step))) {
			return 1
		}
	}

	// if CryptoSlices[strings.ToLower(tag)].Other == strings.ToLower(Encode([]byte(step))) {
	// 	return 1
	// }
	return -1
}
