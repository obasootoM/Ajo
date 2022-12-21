CREATE TABLE "entrys" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "ammount" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

COMMENT ON COLUMN "entrys"."ammount" IS 'must be positive';