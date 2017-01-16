package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/pelletier/go-toml"
)

// Config defines a configuration file
type Config struct {
	*flag.FlagSet
}

// BoolVar defines a bool config with a given name and default value for a Config.
// The argument p points to a bool variable in which to store the value of the config.
func (c *Config) BoolVar(p *bool, name string, value bool) {
	c.FlagSet.BoolVar(p, name, value, "")
}

// Bool defines a bool config variable with a given name and default value for
// a Config.
func (c *Config) Bool(name string, value bool) *bool {
	return c.FlagSet.Bool(name, value, "")
}

// IntVar defines a int config with a given name and default value for a Config.
// The argument p points to a int variable in which to store the value of the config.
func (c *Config) IntVar(p *int, name string, value int) {
	c.FlagSet.IntVar(p, name, value, "")
}

// Int defines a int config variable with a given name and default value for a
// Config.
func (c *Config) Int(name string, value int) *int {
	return c.FlagSet.Int(name, value, "")
}

// Int64Var defines a int64 config with a given name and default value for a Config.
// The argument p points to a int64 variable in which to store the value of the config.
func (c *Config) Int64Var(p *int64, name string, value int64) {
	c.FlagSet.Int64Var(p, name, value, "")
}

// Int64 defines a int64 config variable with a given name and default value
// for a Config.
func (c *Config) Int64(name string, value int64) *int64 {
	return c.FlagSet.Int64(name, value, "")
}

// UintVar defines a uint config with a given name and default value for a Config.
// The argument p points to a uint variable in which to store the value of the config.
func (c *Config) UintVar(p *uint, name string, value uint) {
	c.FlagSet.UintVar(p, name, value, "")
}

// Uint defines a uint config variable with a given name and default value for
// a Config.
func (c *Config) Uint(name string, value uint) *uint {
	return c.FlagSet.Uint(name, value, "")
}

// Uint64Var defines a uint64 config with a given name and default value for a Config.
// The argument p points to a uint64 variable in which to store the value of the config.
func (c *Config) Uint64Var(p *uint64, name string, value uint64) {
	c.FlagSet.Uint64Var(p, name, value, "")
}

// Uint64 defines a uint64 config variable with a given name and default value
// for a Config.
func (c *Config) Uint64(name string, value uint64) *uint64 {
	return c.FlagSet.Uint64(name, value, "")
}

// StringVar defines a string config with a given name and default value for a Config.
// The argument p points to a string variable in which to store the value of the config.
func (c *Config) StringVar(p *string, name string, value string) {
	c.FlagSet.StringVar(p, name, value, "")
}

// String defines a string config variable with a given name and default value
// for a Config.
func (c *Config) String(name string, value string) *string {
	return c.FlagSet.String(name, value, "")
}

// Float64Var defines a float64 config with a given name and default value for a Config.
// The argument p points to a float64 variable in which to store the value of the config.
func (c *Config) Float64Var(p *float64, name string, value float64) {
	c.FlagSet.Float64Var(p, name, value, "")
}

// Float64 defines a float64 config variable with a given name and default
// value for a Config.
func (c *Config) Float64(name string, value float64) *float64 {
	return c.FlagSet.Float64(name, value, "")
}

// DurationVar defines a time.Duration config with a given name and default value for a Config.
// The argument p points to a time.Duration variable in which to store the value of the config.
func (c *Config) DurationVar(p *time.Duration, name string, value time.Duration) {
	c.FlagSet.DurationVar(p, name, value, "")
}

// Duration defines a time.Duration config variable with a given name and
// default value.
func (c *Config) Duration(name string, value time.Duration) *time.Duration {
	return c.FlagSet.Duration(name, value, "")
}

// Parse a file by path to populate configuration
func (c *Config) Parse(path string) error {
	data, err := ioutil.ReadFile(path)
	switch {
	case err != nil && len(path) != 0:
		return err
	case err != nil && len(path) == 0:
		data = []byte{}
	}
	return c.ParseBytes(data)
}

// ParseBytes parses a series of bytes to populate configuration
func (c *Config) ParseBytes(data []byte) error {
	tree, err := toml.Load(string(data))
	if err != nil {
		return err
	}
	err = c.loadTomlTree(tree, []string{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) loadTomlTree(tree *toml.TomlTree, path []string) error {
	for _, key := range tree.Keys() {
		fullPath := append(path, key)
		value := tree.Get(key)
		if subtree, isTree := value.(*toml.TomlTree); isTree {
			err := c.loadTomlTree(subtree, fullPath)
			if err != nil {
				return err
			}
		} else {
			fullPath := strings.Join(append(path, key), "-")
			fullPath = strings.Replace(fullPath, "_", "-", -1)
			switch v := value.(type) {
			case []interface{}:
				var items []string
				for _, item := range v {
					items = append(items, fmt.Sprintf("%v", item))
				}
				value = strings.Join(items, ",")
			}
			err := c.Set(fullPath, fmt.Sprintf("%v", value))
			fmt.Printf("Read (%v): %v\n", fullPath, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewConfig constructs a new configuration
func NewConfig(name string, errorHandling flag.ErrorHandling) *Config {
	return &Config{
		FlagSet: flag.NewFlagSet(name, errorHandling),
	}
}

// Global Config

var globalConfig = NewConfig(os.Args[0], flag.ExitOnError)

// Parse takes a file path and parses it into the global configuration
func Parse(path string) error {
	return globalConfig.Parse(path)
}

// BoolVar defines a bool config with a given name and default value.
// The argument p points to a bool variable in which to store the value of the config.
func BoolVar(p *bool, name string, value bool) {
	globalConfig.BoolVar(p, name, value)
}

// Bool defines a bool config variable with a given name and default value.
func Bool(name string, value bool) *bool {
	return globalConfig.Bool(name, value)
}

// IntVar defines a int config with a given name and default value.
// The argument p points to a int variable in which to store the value of the config.
func IntVar(p *int, name string, value int) {
	globalConfig.IntVar(p, name, value)
}

// Int defines a int config variable with a given name and default value.
func Int(name string, value int) *int {
	return globalConfig.Int(name, value)
}

// Int64Var defines a int64 config with a given name and default value.
// The argument p points to a int64 variable in which to store the value of the config.
func Int64Var(p *int64, name string, value int64) {
	globalConfig.Int64Var(p, name, value)
}

// Int64 defines a int64 config variable with a given name and default value.
func Int64(name string, value int64) *int64 {
	return globalConfig.Int64(name, value)
}

// UintVar defines a uint config with a given name and default value.
// The argument p points to a uint variable in which to store the value of the config.
func UintVar(p *uint, name string, value uint) {
	globalConfig.UintVar(p, name, value)
}

// Uint defines a uint config variable with a given name and default value.
func Uint(name string, value uint) *uint {
	return globalConfig.Uint(name, value)
}

// Uint64Var defines a uint64 config with a given name and default value.
// The argument p points to a uint64 variable in which to store the value of the config.
func Uint64Var(p *uint64, name string, value uint64) {
	globalConfig.Uint64Var(p, name, value)
}

// Uint64 defines a uint64 config variable with a given name and default value.
func Uint64(name string, value uint64) *uint64 {
	return globalConfig.Uint64(name, value)
}

// StringVar defines a string config with a given name and default value.
// The argument p points to a string variable in which to store the value of the config.
func StringVar(p *string, name string, value string) {
	globalConfig.StringVar(p, name, value)
}

// String defines a string config variable with a given name and default value.
func String(name string, value string) *string {
	return globalConfig.String(name, value)
}

// Float64Var defines a float64 config with a given name and default value.
// The argument p points to a float64 variable in which to store the value of the config.
func Float64Var(p *float64, name string, value float64) {
	globalConfig.Float64Var(p, name, value)
}

// Float64 defines a float64 config variable with a given name and default
// value.
func Float64(name string, value float64) *float64 {
	return globalConfig.Float64(name, value)
}

// DurationVar defines a time.Duration config with a given name and default value.
// The argument p points to a time.Duration variable in which to store the value of the config.
func DurationVar(p *time.Duration, name string, value time.Duration) {
	globalConfig.DurationVar(p, name, value)
}

// Duration defines a time.Duration config variable with a given name and
// default value.
func Duration(name string, value time.Duration) *time.Duration {
	return globalConfig.Duration(name, value)
}
