group "default" {
  targets = ["elephpants-api"]
}

target "elephpants-api" {
  context = "./elephpants-api"
  dockerfile = "Dockerfile"
  tags = ["scraly/elephpants-api:bake"]
}
