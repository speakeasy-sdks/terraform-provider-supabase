terraform {
  required_providers {
    supabase = {
      source  = "speakeasy/supabase"
      version = "0.0.1"
    }
  }
}

resource "random_string" "db_pass" {
  length           = 16
}

provider "supabase" {
  bearer = file("../test.key")
}

resource "supabase_project" "example" {
  name = "example"
  db_pass = random_string.db_pass.result
  organization_id = file("../test.orgid")
  region = "eu-west-2"
  plan = "free"
}

resource "supabase_function" "example_func" {
  body = file("./geoip-echo.ts")
  name = "geoip"
  ref = supabase_project.example.id
  slug = "geoip"
}