create table if not exists items
(
	id text not null
		constraint items_pk
			primary key,
	name text not null
);

create unique index if not exists items_name_uindex
	on items (name);

create table if not exists units
(
	id text not null
		constraint units_pk
			primary key,
	name text not null
);

create unique index if not exists units_name_uindex
	on units (name);

create table if not exists stocks
(
	id text not null
		constraint stocks_pk
			primary key,
	item_id text not null
		constraint stocks_items_id_fk
			references items
				on delete restrict,
	unit_id text not null
		constraint stocks_units_id_fk
			references units
				on delete restrict,
	quantity numeric(12,2) not null
);

create table if not exists users
(
	id text not null
		constraint users_pk
			primary key,
	email text not null,
	password text not null,
	name text not null,
	contact_person text,
	contact_number text,
	role text not null
);

create unique index if not exists users_email_uindex
	on users (email);

create table if not exists donations
(
	id text not null
		constraint donations_pk
			primary key,
	date timestamp with time zone not null,
	is_accepted boolean default false not null,
	is_donated boolean default false not null,
	donator_id text not null
		constraint donations_users_id_fk
			references users
				on delete restrict
);

create table if not exists donation_items
(
	id text not null
		constraint donation_items_pk
			primary key,
	donation_id text not null
		constraint donation_items_donations_id_fk
			references donations
				on delete cascade,
	item_id text not null
		constraint donation_items_items_id_fk
			references items
				on delete restrict,
	unit_id text not null
		constraint donation_items_units_id_fk
			references units
				on delete restrict,
	quantity numeric(12,2) not null
);

create table if not exists requests
(
	id text not null
		constraint requests_pk
			primary key,
	date timestamp with time zone not null,
	is_fulfilled boolean default false not null,
	donation_applicant_id text not null
		constraint requests_users_id_fk
			references users
				on delete restrict
);

create table if not exists request_items
(
	id text not null
		constraint request_items_pk
			primary key,
	item_id text not null
		constraint request_items_items_id_fk
			references items
				on delete restrict,
	unit_id text not null
		constraint request_items_units_id_fk
			references units
				on delete restrict,
	quantity numeric(12,2) not null,
	request_id text not null
		constraint request_items_requests_id_fk
			references requests
				on delete cascade
);

create table if not exists allocations
(
	id text not null
		constraint allocations_pk
			primary key,
	request_id text not null
		constraint allocations_requests_id_fk
			references requests
				on delete restrict,
	date timestamp with time zone not null,
	photo_url text,
	admin_id text not null
		constraint allocations_users_id_fk
			references users
				on delete restrict
);

create table if not exists allocation_items
(
	id text not null
		constraint allocation_items_pk
			primary key,
	allocation_id text not null
		constraint allocation_items_allocations_id_fk
			references allocations
				on delete cascade,
	item_id text not null
		constraint allocation_items_items_id_fk
			references items
				on delete restrict,
	unit_id text not null
		constraint allocation_items_units_id_fk
			references units
				on delete restrict,
	quantity numeric(12,2) not null
);

create table if not exists password_reset_requests
(
	id text not null
		constraint password_reset_requests_pk
			primary key,
	user_id text not null
		constraint password_reset_requests_users_id_fk
			references users
				on delete cascade,
	new_password text not null,
	date timestamp with time zone
);