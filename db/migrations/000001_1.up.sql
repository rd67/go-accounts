DROP TABLE IF EXISTS "accounts";
DROP SEQUENCE IF EXISTS accounts_id_seq;
CREATE SEQUENCE accounts_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."accounts" (
    "id" bigint DEFAULT nextval('accounts_id_seq') NOT NULL,
    "name" character varying(255) NOT NULL,
    "balance" bigint NOT NULL,
    "currency" character(3) DEFAULT 'INR' NOT NULL,
    "isDeleted" boolean DEFAULT false NOT NULL,
    "createdAt" timestamp DEFAULT now() NOT NULL,
    "updatedAt" timestamp DEFAULT now() NOT NULL,
    CONSTRAINT "accounts_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

TRUNCATE "accounts";

DROP TABLE IF EXISTS "entries";
DROP SEQUENCE IF EXISTS entries_id_seq;
CREATE SEQUENCE entries_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."entries" (
    "id" bigint DEFAULT nextval('entries_id_seq') NOT NULL,
    "account_id" bigint NOT NULL,
    "amount" bigint NOT NULL,
    "currency" character(3) DEFAULT 'INR' NOT NULL,
    "exchange_rate" integer DEFAULT '1' NOT NULL,
    "isDeleted" boolean DEFAULT false NOT NULL,
    "createdAt" timestamptz DEFAULT now() NOT NULL,
    "updatedAt" timestamptz DEFAULT now() NOT NULL,
    CONSTRAINT "entries_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "entries_account_id" ON "public"."entries" USING btree ("account_id");


DROP TABLE IF EXISTS "transfers";
DROP SEQUENCE IF EXISTS transfers_id_seq;
CREATE SEQUENCE transfers_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."transfers" (
    "id" bigint DEFAULT nextval('transfers_id_seq') NOT NULL,
    "sender_account_id" bigint NOT NULL,
    "receiver_account_id" bigint NOT NULL,
    "amount" bigint NOT NULL,
    "createdAt" timestamptz DEFAULT now() NOT NULL,
    "updatedAt" timestamp DEFAULT now() NOT NULL,
    "isDeleted" boolean DEFAULT false NOT NULL,
    CONSTRAINT "transfers_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "transfers_receiver_account_id" ON "public"."transfers" USING btree ("receiver_account_id");

CREATE INDEX "transfers_sender_account_id" ON "public"."transfers" USING btree ("sender_account_id");

CREATE INDEX "transfers_sender_account_id_receiver_account_id" ON "public"."transfers" USING btree ("sender_account_id", "receiver_account_id");


ALTER TABLE ONLY "public"."entries" ADD CONSTRAINT "entries_account_id_fkey" FOREIGN KEY (account_id) REFERENCES accounts(id) NOT DEFERRABLE;
