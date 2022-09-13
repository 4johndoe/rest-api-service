DROP SEQUENCE IF EXISTS accounts_id_seq;
CREATE SEQUENCE accounts_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "accounts"
(
    account_id      INT DEFAULT nextval('accounts_id_seq') NOT NULL,
    customer_id     INT,
    opening_date    TIMESTAMP NOT NULL DEFAULT now(),
    account_type    VARCHAR,
    amount          FLOAT,
    status          VARCHAR DEFAULT '1'
);
