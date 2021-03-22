CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"username" varchar(64) NOT NULL UNIQUE,
	"email" varchar(100) NOT NULL UNIQUE,
	"fullname" TEXT,
	"bio" TEXT,
	"admin" BOOLEAN NOT NULL DEFAULT FALSE,
	"created_at" DATE NOT NULL DEFAULT (now()),
	"updated_at" DATE NOT NULL DEFAULT (now()),
	"deleted_at" DATE,
	CONSTRAINT "user_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "password" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"password" TEXT NOT NULL,
	"created_at" DATE NOT NULL DEFAULT (now()),
	"updated_at" DATE NOT NULL DEFAULT (now()),
	"deleted_at" DATE,
	CONSTRAINT "password_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "post" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"content" TEXT NOT NULL,
	"created_at" DATE NOT NULL DEFAULT (now()),
	"updated_at" DATE NOT NULL DEFAULT (now()),
	"deleted_at" DATE,
	CONSTRAINT "post_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "post" ADD CONSTRAINT "post_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

