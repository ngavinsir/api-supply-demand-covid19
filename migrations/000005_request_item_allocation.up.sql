create table request_item_allocation
(
    id text
		constraint request_item_allocation_pk
			primary key,
	request_item_id text not null
		constraint request_item_allocation_request_items_id_fk
			references request_items
				on delete cascade,
	allocation_date timestamptz not null,
	description text
);

create unique index request_item_allocation_request_item_id_uindex
	on request_item_allocation (request_item_id);