-- =============================================
-- Author:      William Wibowo Ciptono
-- Create date: 14 Des 2022
-- Description: Initiate restaurant DB tables
-- =============================================

CREATE DATABASE DB_RESTAURANT;

CREATE TABLE IF NOT EXISTS roles (
	id INT PRIMARY KEY,
    role_name VARCHAR,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR NOT NULL,
	phone VARCHAR check (phone ~ '^[0-9]*$'),
	email VARCHAR UNIQUE NOT NULL,
	username VARCHAR UNIQUE NOT NULL check (username ~ '^[a-z0-9]+'),
	password VARCHAR NOT NULL,
	register_date date NOT NULL,
	profile_picture BYTEA,
	play_attempt INT,
	role_id INT NOT NULL,
	foreign key (role_id) references roles(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS sessions (
	id SERIAL PRIMARY KEY,
	refresh_token VARCHAR,
	user_id INT,
	foreign key (user_id) references users(id),
	expired_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS categories (
	id SERIAL primary key,
	category_name varchar,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS menus (
	id SERIAL PRIMARY KEY,
	menu_name VARCHAR,
	avg_rating numeric,
	number_of_favorites numeric,
	price numeric,
	menu_photo varchar,
	category_id INT,
	foreign key (category_id) references categories(id),
	customization json,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

alter table menus 
alter column customization
set default '[]';

CREATE TABLE IF NOT EXISTS favorited_menus (
	id SERIAL PRIMARY KEY,
	user_id INT,
	foreign key (user_id) references users(id),
	menu_id INT,
	foreign key (menu_id) references menus(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS promotions (
	id SERIAL PRIMARY KEY,
	admin_id INT,
	foreign key (admin_id) references users(id),
	name VARCHAR,
	description VARCHAR,
	promotion_photo varchar,
	discount_rate numeric,
	started_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
	expired_at TIMESTAMP null,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS promo_menus (
	id SERIAL PRIMARY KEY,
	promotion_id INT,
	foreign key (promotion_id) references promotions(id),
	menu_id INT,
	foreign key (menu_id) references menus(id),
	promotion_price numeric,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

create table if not exists carts (
	id SERIAL PRIMARY KEY,
	user_id INT,
	foreign key (user_id) references users(id),
	promotion_id INT,
	foreign key (promotion_id) references promotions(id),
	menu_id INT,
	foreign key (menu_id) references menus(id),
	quantity int,
	menu_option json,	
	is_ordered boolean default false,
	promotion_price numeric,

	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

alter table carts 
alter column menu_option
set default '[]';

CREATE TABLE IF NOT EXISTS payment_options (
	id SERIAL primary key,
	payment_name varchar,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

create table if not exists coupons (
	id SERIAL primary key,
	admin_id INT,
	foreign key (admin_id) references users(id),
	description VARCHAR,
	discount_amount numeric,
	quota_initial numeric,
	quota_left numeric,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

create table if not exists user_coupons (
	id SERIAL primary key, 
	user_id INT,
	foreign key (user_id) references users(id),	
	coupon_id INT,
	foreign key (coupon_id) references coupons(id),
	coupon_code VARCHAR,
	discount_amount numeric,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);


create table if not exists orders (
	id SERIAL PRIMARY KEY,
	user_id INT,
	foreign key (user_id) references users(id),
	order_date TIMESTAMP not null default CURRENT_TIMESTAMP,
	coupon_id INT,
	foreign key (coupon_id) references coupons(id),
	payment_option_id INT,
	foreign key (payment_option_id) references payment_options(id),
	gross_amount numeric,
	discount_amount numeric,
	net_amount numeric,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);


create table if not exists ordered_menus (
	id SERIAL PRIMARY KEY,
	order_id INT,
	foreign key (order_id) references orders(id),
	menu_id INT,
	foreign key (menu_id) references menus(id),
	promotion_id INT,
	foreign key (promotion_id) references promotions(id),
	quantity int,
	menu_option json,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

alter table ordered_menus 
alter column menu_option
set default '[]';

create table if not exists reviews (
	id SERIAL PRIMARY KEY,
	review_description varchar,
	rating numeric,
	ordered_menu_id INT,
	foreign key (ordered_menu_id) references ordered_menus(id),
	menu_id INT,
	foreign key (menu_id) references menus(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS auth_sessions (
	id SERIAL PRIMARY KEY,
    user_id int,
    foreign key (user_id) references users(id),
    refresh_token varchar,
    is_invalid bool,
    expired_at TIMESTAMP,
	
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);