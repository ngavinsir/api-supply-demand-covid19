alter table password_reset_requests
	add new_password text default '' not null;