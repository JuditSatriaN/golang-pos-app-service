CREATE TABLE public.products
(
    product_id VARCHAR NOT NULL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    unit VARCHAR(10) NOT NULL DEFAULT '',
    barcode BIGINT NOT NULL DEFAULT 0,
    create_time timestamp not null default now(),
    update_time timestamp
);

CREATE UNIQUE INDEX products_barcode_idx ON products (barcode);