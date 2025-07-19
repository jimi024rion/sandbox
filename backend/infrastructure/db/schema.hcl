table "todos" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
    default = gen_random_uuid()
  }
  column "title" {
    null = false
    type = varchar(255)
  }
  column "description" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = timestamptz
    default = "now()"
  }
  column "updated_at" {
    null = false
    type = timestamptz
    default = "now()"
  }
  primary_key {
    columns = [column.id]
  }
}

schema "public" {
  comment = "A public schema"
}
