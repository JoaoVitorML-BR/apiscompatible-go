package secure

import "golang.org/x/crypto/bcrypt"

// Cryptography password and return hashed password as a string
func Hash(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// Compare password with hash. Returns an error if the passwords do not match.
func ComparePasswordWithHash(passwordHash string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
