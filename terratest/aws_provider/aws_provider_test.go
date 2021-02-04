package aws_provider

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
  )

func TestTempProvider(t *testing.T) {
  input := &TempProviderInput{
    path: "../../module/provider.tf",
    region: "eu-west-1",
    profile: "account-1",
  }
  want := `provider "aws" {
  region = "eu-west-1"
  profile = "account-1"
}`
  file, provider := ProviderFile(input)
  defer os.Remove(input.path)
  assert.Equal(t, file, input.path)
  assert.Equal(t, provider, want)
}
