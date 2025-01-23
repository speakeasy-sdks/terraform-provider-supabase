terraform {
  required_providers {
    supabase = {
      source  = "speakeasy/supabase"
      version = "0.0.5"
    }
  }
}

provider "supabase" {
  bearer = "sbp_b836c69bdd19c67f1e80153870908b8e4cb12ebc"
}

resource "supabase_project" "my_project" {
  db_pass         = "Barry0325!"
  kps_enabled     = false
  name            = "Demo Project"
  organization_id = "dspubgjvweempctumvtc"
  plan            = "free"
  postgres_engine = "15"
  region          = "us-east-1"
  template_url    = "https://github.com/supabase/supabase/tree/master/examples/slack-clone/nextjs-slack-clone"
}
