package types

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitBytes = "1234567890"
const specialCharBytes = "~!@#$%^&*()_+-={}[]|;:',<.?/?"

// Marker: =======================================
// Marker: Basic string functions
// Marker: =======================================

/*
Returns Random Alphabetic string of length n
 */
func GetRandomAlphaString(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/*
Returns a random numeric string (consisting of digits 0-9) of length n
 */
func GetRandomNumericString(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = digitBytes[rand.Intn(len(digitBytes))]
	}
	return string(b)
}

/*
Returns a random string of length n containing special characters
 */
func GetRandomSpecialCharacterString(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = specialCharBytes[rand.Intn(len(specialCharBytes))]
	}
	return string(b)
}

/*
Returns a random string of length n which may contain alphabets, digits and special characters
 */
func GetRandomString(n int) string {
	s := letterBytes + digitBytes + specialCharBytes
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UTC().UnixNano())
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}

// Marker: =======================================
// Marker: Extension type and methods
// Marker: =======================================

type String struct {
	strValue string
}

func NewString(str string) *String {
	return &String{strValue: str}
}

func (s *String) Value() string {
	return s.strValue
}

func (s *String) ValueInBytes() []byte {
	return []byte(s.strValue)
}

func (s *String) SafeForPassword() bool {
	if s.Length() < 8 {
		return false
	}
	matchedCaps, errCaps := regexp.Match(`[A-Z]`, []byte(s.strValue))
	matchedSmall, errSmall := regexp.Match(`[a-z]`, []byte(s.strValue))
	matchedNumeric, errNumeric := regexp.Match(`[0-9]`, []byte(s.strValue))
	if matchedCaps == false || errCaps != nil ||
		matchedSmall == false || errSmall != nil ||
		matchedNumeric == false || errNumeric != nil {
		return false
	}

	return true
}

func (s *String) ValidMobileNumber() bool {
	matched, err := regexp.MatchString(`^\+[\d]+$`, s.strValue)
	if matched == true && err == nil {
		return true
	}
	return false
}

func (s *String) ValidIndianMobileNumber() bool {
	matched, err := regexp.MatchString(`^\+91[9876]+[\d]{9}$`, s.strValue)
	if matched == true && err == nil {
		return true
	}
	return false
}

func (s *String) ValidEmail() bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(s.strValue)
}

// True, if the truncated value is blank string
func (s *String) IsEmpty() bool {
	str := s.strValue

	strTrimmed := strings.TrimSpace(str)
	if strTrimmed == "" {
		return true
	}
	return false
}

/*

Checks if the String starts with another string

The method takes a second boolean value which when set to true will compare
in a case-insensitive manner.

	NewString("Vaibhav").StartsWith("") // true
	NewString("Vaibhav").StartsWith("Vai") // true
	NewString("Vaibhav").StartsWith("vai") // false
	NewString("Vaibhav").StartsWith("Vai", true) // true
	NewString("Vaibhav").StartsWith("Vaibhav Kaushal") // false

*/
func (s *String) StartsWith(subString string, caseInsensitive ...bool) bool {
	if len(s.strValue) < len(subString) {
		// A string cannot start with another string larger than its length
		return false
	}

	origStringToCompare := s.strValue
	subStringToCompare := subString
	if len(caseInsensitive) > 0 && caseInsensitive[0] == true {
		origStringToCompare = strings.ToLower(origStringToCompare)
		subStringToCompare = strings.ToLower(subStringToCompare)
	}

	if origStringToCompare[0:len(subString)] == subStringToCompare {
		return true
	}

	return false
}

/*

Checks if the String ends in another string

The method takes a second boolean value which when set to true will compare
in a case-insensitive manner.

	NewString("Vaibhav").EndsWith("") // true
	NewString("Vaibhav").EndsWith("v") // true
	NewString("Vaibhav").EndsWith("V") // false
	NewString("Vaibhav").EndsWith("V", true) // true
	NewString("Vaibhav").EndsWith("Vaibhav Kaushal") // false
	NewString("Vaibhav").EndsWith("av") // true

*/
func (s *String) EndsWith(subString string, caseInsensitive ...interface{}) bool {
	if len(s.strValue) < len(subString) {
		// A string cannot end in another string larger than its length
		return false
	}

	origStringToCompare := s.strValue
	subStringToCompare := subString
	if len(caseInsensitive) > 0 && caseInsensitive[0] == true {
		origStringToCompare = strings.ToLower(origStringToCompare)
		subStringToCompare = strings.ToLower(subStringToCompare)
	}

	if origStringToCompare[len(origStringToCompare)-len(subStringToCompare):] == subStringToCompare {
		return true
	}

	return false
}

func (s *String) Trim() *String {
	s.strValue = strings.TrimSpace(s.strValue)
	return s
}

//func (s *String) ToPasswordHash() string {
//	return GetPasswordHash(s.strValue)
//}

func (s *String) Length() int {
	return len(s.strValue)
}

func (s *String) IsNumeric() bool {
	if _, err := strconv.Atoi(s.strValue); err == nil {
		return true
	}

	return false
}

func (s *String) ToUpperCase() *String {
	s.strValue = strings.ToUpper(s.strValue)
	return s
}

func (s *String) ToLowerCase() *String {
	s.strValue = strings.ToLower(s.strValue)
	return s
}
