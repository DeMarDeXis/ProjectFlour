CREATE TABLE IF NOT EXISTS type_materials
(
    id                  serial       not null unique,
    name_type_material  varchar(255) not null,
    percent_of_marriage float not null
);

CREATE TABLE IF NOT EXISTS type_products (
    id serial       not null unique,
    name_type_product  varchar(255) not null,
    coefficient float not null
);

CREATE TABLE IF NOT EXISTS products (
    id serial not null unique,
    type_product_id int references type_products(id) on delete cascade not null,
    name_product varchar(255) not null,
    article varchar(255) not null,
    min_price_for_partner float not null
);

CREATE TABLE IF NOT EXISTS organizations_type (
    id serial not null unique,
    name_of_type_organization varchar(10) not null
);

CREATE TABLE IF NOT EXISTS partners (
    id serial not null unique,
    type_organization int references organizations_type(id) not null,
    name_partner varchar(255) not null,
    full_name_boss varchar(255) not null,
    email varchar(255) not null,
    phone_number varchar(255) not null,
    legal_address varchar(255) not null,
    inn varchar(26) not null,
    rate int not null
);

CREATE TABLE IF NOT EXISTS partner_products (
    id serial not null unique,
    partner_id int references partners(id) not null,
    product_id int references products(id) not null,
    quantity_of_product int not null,
    date_of_sale date not null,
    price_for_partner float
);

CREATE TABLE IF NOT EXISTS users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);