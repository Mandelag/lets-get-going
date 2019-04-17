package dotenv

import (
	"os"
	"strings"
)

// Load environment from .env file, containing line separated, env variables.
// Each env variables are '=' separated.
// Only ascii encoding are supported.
//
// Example:
//   APP_SERVER=10.0.0.2
//   DB_SERVER=10.0.0.3
func LoadEnv(envFile string) error {
	file, err := os.Open(envFile)
	defer file.Close()

	env := make([]byte, 0, 128)
	buf := make([]byte, 512)
	rcount := 1024
	for rcount > 0 {
		rcount, _ = file.Read(buf)
		env = append(env, buf[:rcount]...)
	}
	envString := string(env)

	res := strings.FieldsFunc(envString, func (r rune) bool {
		switch string(r) {
		case "\n":
			return true
		case "\r":
			return true
		case "\r\n":
			return true
		default:
			return false
		}
		// return false
	})

	for i:=0; i<len(res); i++ {
		val := res[i]
		token := strings.Split(val, "=")
		key, value := token[0], token[1]
		_ = os.Setenv(string(key), string(value))
	}
	
	return err
}