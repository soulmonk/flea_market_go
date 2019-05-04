create table notes
(
    id          serial not null
        constraint notes_pk
            primary key,
    title       varchar(64),
    description varchar(255),
    text        text,
    created_at  timestamp default now(),
    updated_at  timestamp default now()
);

create table keywords
(
    id   serial      not null
        constraint keywords_pk
            primary key,
    name varchar(32) not null
);

create table notes_has_keywords
(
    note_id    integer not null
        constraint notes_has_keywords_notes_id_fk
            references notes
            on delete cascade,
    keyword_id integer not null
        constraint notes_has_keywords_keywords_id_fk
            references keywords
            on delete cascade
);