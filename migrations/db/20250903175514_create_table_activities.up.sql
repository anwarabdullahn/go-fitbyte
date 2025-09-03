CREATE TABLE "activities" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "activityType" varchar,
  "doneAt" timestamp,
  "durationInMinutes" int,
  "caloriesBurned" int
);

ALTER TABLE "activities" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");