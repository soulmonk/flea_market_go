package pg

var initTable = `create table notes
(
	id serial,
	title varchar(64),
	description text,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

create unique index notes_id_uindex
	on notes (id);

alter table notes
	add constraint notes_pk
		primary key (id);

`
