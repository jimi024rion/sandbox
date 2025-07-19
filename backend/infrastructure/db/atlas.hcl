data "env" "src" {
  url = "docker://postgres/15/todo_db?search_path=public"
}

env {
  name = "local"
  url  = "postgres://user:password@localhost:5432/todo_db?sslmode=disable"
  dev  = data.env.src.url

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ .Name }}"
    }
  }
}
