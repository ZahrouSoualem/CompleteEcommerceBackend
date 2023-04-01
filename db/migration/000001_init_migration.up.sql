CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "address" varchar NOT NULL DEFAULT '',
  "city" varchar NOT NULL DEFAULT '',
  "state" varchar NOT NULL DEFAULT '',
  "country" varchar NOT NULL DEFAULT '',
  "zip_code" bigint NOT NULL DEFAULT 0,
  "phone_number" bigint UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "market" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "marketname" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "catname" varchar UNIQUE NOT NULL
);

CREATE TABLE "brand" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "braname" varchar UNIQUE NOT NULL,
  "imageurl" varchar UNIQUE NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "category_id" bigint NOT NULL,
  "brand_id" bigint NOT NULL,
  "market_id" bigint NOT NULL,
  "proname" varchar NOT NULL,
  "description" varchar NOT NULL,
  "imageurl" varchar UNIQUE NOT NULL,
  "price" float NOT NULL,
  "quantity" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_updated" timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z')
);

CREATE TABLE "ordersproduct" (
  "orders_product_id" bigserial PRIMARY KEY NOT NULL,
  "orders_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "quantity" bigint NOT NULL DEFAULT 1
);

CREATE TABLE "review" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "product_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "rating" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "comment" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "review_id" bigint NOT NULL,
  "opinion" varchar NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "market" ("marketname");

CREATE INDEX ON "category" ("catname");

CREATE INDEX ON "brand" ("braname");

CREATE INDEX ON "products" ("proname");

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "ordersproduct" ("orders_id");

CREATE INDEX ON "ordersproduct" ("product_id");

CREATE UNIQUE INDEX ON "ordersproduct" ("orders_id", "product_id");

CREATE INDEX ON "review" ("product_id");

CREATE INDEX ON "review" ("user_id");

CREATE UNIQUE INDEX ON "review" ("product_id", "user_id");

CREATE INDEX ON "comment" ("review_id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("brand_id") REFERENCES "brand" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("market_id") REFERENCES "market" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "ordersproduct" ADD FOREIGN KEY ("orders_id") REFERENCES "orders" ("id");

ALTER TABLE "ordersproduct" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("review_id") REFERENCES "review" ("id");
