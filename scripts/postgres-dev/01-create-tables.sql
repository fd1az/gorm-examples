CREATE TABLE "users" (
    "id" serial UNIQUE NOT NULL,
    "name" text NOT NULL,
    "email" text,
    "created_at" timestamp with time zone default NOW(),
    "updated_at" timestamp with time zone default NOW(),
    "deleted_at" timestamp with time zone default NULL
);

CREATE UNIQUE INDEX idx_users_email ON users(email);

CREATE TABLE "products" (
    "id" serial UNIQUE NOT NULL,
    "title" text NOT NULL,
    "description" text,
    "price" numeric NOT NULL DEFAULT 0.0,
    "created_at" timestamp with time zone default NOW(),
    "updated_at" timestamp with time zone default NOW(),
    "deleted_at" timestamp with time zone default NULL
);

CREATE UNIQUE INDEX idx_products_title ON products(title);

CREATE TABLE "orders" (
    "id" serial UNIQUE NOT NULL,
    "user_id" serial constraint fk_orders_users references users (id),
    "total" numeric NOT NULL DEFAULT 0.0,
    "status" text,
    "created_at" timestamp with time zone default NOW(),
    "updated_at" timestamp with time zone default NOW(),
    "deleted_at" timestamp with time zone default NULL
);

CREATE TABLE "orders_products" (
    "id" serial NOT NULL,
    "orders_id" serial NOT NULL constraint fk_orders_products_orders references orders (id),
    "product_id" serial NOT NULL constraint fk_orders_products_products references products (id),
    "unit_price" numeric NOT NULL DEFAULT 0.0,
    "quantity" numeric NOT NULL
);

[{
    {1 2021-10-28 16:17:56.252189 -0300 -03 2021-10-28 16:17:56.252189 -0300 -03 
    {0001-01-01 00:00:00 +0000 UTC false}
    } user1
     some1@email.com
      { false}}