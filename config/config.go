package config

import (
  "flag"
  "fmt"
  "io/ioutil"
  "github.com/pelletier/go-toml"
)

type Config struct {
  *flag.FlagSet
}

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
    if subtree, isTree := value.(*toml.TomTree); isTree {
      err := c.loadTomlTree(subtree, fullPath)
      if err != nill {
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
    }
    err := c.Set(fullPath, fmt.Sprintf("%v", value))
    if err != nil {
      return err
    }
  }
  return nil
}

func NewConfig(name string, errorHandling flag.ErrorHandling) *Config {
  return &Config {
    FlagSet: flag.NewFlagSet(name, errorHandling),
  }
}

// Global Config

var globalConfig = NewConfig(os.Args[0], flag.ExitOnError)

func Parse(path string) error {
  return globalConfig.Parse(path)
}