package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/fnv"

	"github.com/cespare/xxhash/v2"
	"golang.org/x/crypto/bcrypt"
)

func GenerateMD5Hash(concatenatedString string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(concatenatedString)))
}

func CheckMDSignIsValid(sign string, operatorCode string, requestTime string, action string, secretKey string) bool {
	temp := operatorCode + requestTime + action + secretKey

	refSign := GenerateMD5Hash(temp)

	return refSign == sign
}

// hashPassword creates a bcrypt hash of the password.
func HashPassword(password string) (string, error) {

	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is the cost for hashing
	return string(bytes), err
}

// checkPasswordHash compares a hashed password with its possible plaintext equivalent.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateFNV32aHashString generates a 32-bit FNV-1a hash of the input string.
//
// The hash is returned as an 8-character hexadecimal string (0-9, a-f).
func GenerateFNV32aHashString(s string) string {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum32())
}

// GenerateSHA1HashString generates a SHA-1 hash of the input string.
//
// The hash is returned as a 40-character hexadecimal string (0-9, a-f).
func GenerateSHA1HashString(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// GenerateSHA256HashString generates a SHA-256 hash of the input string.
// The hash is returned as a 64-character hexadecimal string (0-9, a-f).
func GenerateSHA256HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// 無加密, 速度快, 碰撞低
func GenerateXXHash64String(s string) string {
	return fmt.Sprintf("%x", xxhash.Sum64String(s))
}

// HMACSignSHA256 generates HMAC-SHA256 signature for the given data and secret key.
// Returns the signature as a hexadecimal string.
func HMACSignSHA256(data, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// HMACSignSHA1 generates HMAC-SHA1 signature for the given data and secret key.
// Returns the signature as a hexadecimal string.
func HMACSignSHA1(data, secretKey string) string {
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMACSignSHA256 verifies if the given signature matches the HMAC-SHA256 signature
// of the data with the secret key.
func VerifyHMACSignSHA256(data, secretKey, signature string) bool {
	expectedSignature := HMACSignSHA256(data, secretKey)
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
