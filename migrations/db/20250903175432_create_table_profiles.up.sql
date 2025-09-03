CREATE TABLE "profiles" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "preference" varchar,
  "weightUnit" varchar,
  "heightUnit" varchar,
  "weight" int,
  "height" int,
  "name" varchar,
  "imageUri" varchar
);


ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");