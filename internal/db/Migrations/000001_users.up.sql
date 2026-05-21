create extension if not exists "pgcrypto";

create table if not exists Users(
	UserID serial primary key,
	uuid uuid not null default gen_random_uuid(),
	Name varchar(50) not null,
	Email varchar(100) unique not null,
	create_at timestamptz not null default now()
);