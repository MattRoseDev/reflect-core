CREATE TABLE "user" (
	"id" uuid NOT NULL,
	"username" varchar(64) NOT NULL UNIQUE,
	"email" varchar(100) NOT NULL UNIQUE,
	"fullname" TEXT,
	"bio" TEXT,
	"admin" BOOLEAN,
	"deleted" BOOLEAN,
	CONSTRAINT "user_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "password" (
	"id" serial NOT NULL,
	"userId" uuid NOT NULL,
	"password" TEXT NOT NULL,
	"deleted" BOOLEAN,
	CONSTRAINT "password_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "post" (
	"id" serial NOT NULL,
	"userId" uuid NOT NULL,
	"content" TEXT NOT NULL,
	"deleted" BOOLEAN,
	CONSTRAINT "post_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("userId") REFERENCES "user"("id");

ALTER TABLE "post" ADD CONSTRAINT "post_fk0" FOREIGN KEY ("userId") REFERENCES "user"("id");

