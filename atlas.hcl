variable "envfile" {
    type    = string
    default = "./.env"
}

locals {
    envfile = {
        for line in split("\n", file(var.envfile)): split("=", line)[0] => regex("=(.*)", line)[0]
        if !startswith(line, "#") && length(split("=", line)) > 1
    }
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/infra/gorm-repositories",
    "--dialect", "postgres", // mysql | postgres | sqlite | sqlserver
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/17/dev"
  url = local.envfile["DATABASE_URL"]
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

