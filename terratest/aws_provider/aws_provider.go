package aws_provider

import (
  "fmt"
  "os"
  "log"
)

type TempProviderInput struct {
  path        string
  region      string
  version     string
  profile     string
  role        string
  roleSession string
}

// Creates a provider.tf file
func ProviderFile(input *TempProviderInput) (string, string) {
  if (input.path == "") {
    input.path = "../provider.tf"
  }
  file, err := os.Create(input.path)
  if err != nil {
  	log.Fatal(err)
  }
  defer file.Close()
  provider := ProviderBlock(input)
  _, err = file.WriteString(provider)
  if err != nil {
  	log.Fatal(err)
  }
  return input.path, provider
}

// Creates a Provider block
func ProviderBlock(input *TempProviderInput) string {
  provider := "provider \"aws\" {\n"
  provider += ProviderLine("region", input.region)
  provider += ProviderLine("version", input.version)
  provider += ProviderLine("profile", input.profile)
  provider += "}"
  return provider
}

// Writes a provider line
func ProviderLine(argument, input string) string {
  if (input != "") {
    return fmt.Sprint("  ", argument, " = \"", input, "\"\n")
  }
  return ""
}
