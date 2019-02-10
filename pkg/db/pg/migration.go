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

var query1 = `alter table notes alter column description type varchar(255) using description::varchar(255);

alter table notes
	add text text;`

var query2 = `create table keywords
(
  id serial not null
    constraint keywords_pk
      primary key,
  name varchar(32) not null
);`

var query3 = `create table notes_has_keywords
(
	note_id integer not null
		constraint notes_has_keywords_notes_id_fk
			references notes
				on delete cascade,
	keyword_id integer not null
		constraint notes_has_keywords_keywords_id_fk
			references keywords
				on delete cascade
);
`
