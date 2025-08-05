env "local" {
  src = "ent://ent/schema"
  dev = "mysql://user:password@localhost:3307/dbname_dev"
  url = "mysql://user:password@localhost:3306/dbname"
  migration {
    dir = "file://migrations"
  }
}

env "dev" {
  src = "ent://ent/schema"
  dev = "mysql://user:password@localhost:3307/dbname_dev"
  url = "mysql://user:password@localhost:3306/dbname"
  migration {
    dir = "file://migrations"
  }
}

env "prod" {
  src = "ent://ent/schema"
  dev = "mysql://user:password@localhost:3307/dbname_dev"
  url = "mysql://user:password@localhost:3306/dbname"
  migration {
    dir = "file://migrations"
  }
} 