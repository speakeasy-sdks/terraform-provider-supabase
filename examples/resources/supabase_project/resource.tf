resource "supabase_project" "my_project" {
  db_pass               = "...my_db_pass..."
  desired_instance_size = "2xlarge"
  name                  = "...my_name..."
  organization_id       = "...my_organization_id..."
  postgres_engine       = "15"
  ref                   = "...my_ref..."
  region                = "us-east-1"
  release_channel       = "withdrawn"
  template_url          = "https://github.com/supabase/supabase/tree/master/examples/slack-clone/nextjs-slack-clone"
}