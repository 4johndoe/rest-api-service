DROP SEQUENCE IF EXISTS customers_id_seq;
CREATE SEQUENCE customers_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "customers"
(
    customer_id     INT DEFAULT nextval('customers_id_seq') NOT NULL,
    name            VARCHAR NOT NULL,
    date_of_birth   TIMESTAMP NOT NULL,
    city            VARCHAR,
    zipcode         VARCHAR,
    status          VARCHAR DEFAULT '1',
    updated_at      TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO customers VALUES
    (1, 'Steve', '1978-12-15', 'Kyiv', '113445', 1),
    (2, 'Mike', '1978-12-15', 'Dnipro', '113445', 1),
    (3, 'Jack', '1978-12-15', 'Kharkiv', '113445', 1),
    (4, 'Billy', '1978-12-15', 'Uman', '113445', 1);