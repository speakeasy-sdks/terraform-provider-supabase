resource "supabase_function" "my_function" {
  body               = "...my_body..."
  compute_multiplier = 2.57
  entrypoint_path    = "...my_entrypoint_path..."
  import_map         = false
  import_map_path    = "...my_import_map_path..."
  name               = "...my_name..."
  ref                = "...my_ref..."
  slug               = "...my_slug..."
  verify_jwt         = true
}