package main

import (
  "fmt"
  "encoding/json"
  "github.com/invopop/jsonschema"
  talosTypes "github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1"
)


func main() {
  s := jsonschema.Reflect(&talosTypes.Config{})
  data, _ := json.MarshalIndent(s, "", "  ")
  fmt.Println(string(data))
}
