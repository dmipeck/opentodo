-- +goose Up
CREATE TABLE "todo" (
	"id" UUID PRIMARY KEY,
	"title" VARCHAR(255) NOT NULL,
	"is_completed" BOOLEAN NOT NULL DEFAULT FALSE,
)

-- +goose Down

DROP TABLE IF EXISTS "todo";
