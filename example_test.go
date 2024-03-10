package env_test

import (
	"fmt"
	"log"
	"os"

	"github.com/garaekz/go-env"
)

type Config struct {
	Host     string
	Port     int
	Password string `env:",secret"`
}

func Example_one() {
	_ = os.Setenv("APP_HOST", "127.0.0.1")
	_ = os.Setenv("APP_PORT", "8080")

	var cfg Config
	if err := env.Load(&cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg.Host)
	fmt.Println(cfg.Port)
	// Output:
	// 127.0.0.1
	// 8080
}

func Example_two() {
	_ = os.Setenv("API_HOST", "127.0.0.1")
	_ = os.Setenv("API_PORT", "8080")
	_ = os.Setenv("API_PASSWORD", "test")

	var cfg Config
	loader := env.New("API_", log.Printf)
	if err := loader.Load(&cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg.Host)
	fmt.Println(cfg.Port)
	fmt.Println(cfg.Password)
	// Output:
	// 127.0.0.1
	// 8080
	// test
}

type Embedded struct {
	URL  string
	Port int
}

type Config2 struct {
	Nested Embedded  `prefix:"NESTED_"`
	Other  *Embedded `prefix:"OTHER_"` // pointer to struct
}

func Example_three() {
	_ = os.Setenv("APP_NESTED_URL", "http://example.com")
	_ = os.Setenv("APP_NESTED_PORT", "8080")
	_ = os.Setenv("APP_OTHER_URL", "http://other.com")
	_ = os.Setenv("APP_OTHER_PORT", "9090")

	var cfg Config2

	if err := env.Load(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Nested.URL)
	fmt.Println(cfg.Nested.Port)
	fmt.Println(cfg.Other.URL)
	fmt.Println(cfg.Other.Port)

	// Output:
	// http://example.com
	// 8080
	// http://other.com
	// 9090
}
