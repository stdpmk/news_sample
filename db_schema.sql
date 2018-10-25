
-- PostgreSQL New sample schema

create table author (
	id bigserial primary key,
	name text not NULL
);

create table news (
	id bigserial primary key,
	title text NOT NULL,
	text text NOT NULL,
	tags text[],
	date timestamp without time zone NOT NULL DEFAULT now()::timestamp without time zone,
	id_author bigint references author(id)
);

create table comment (
	id bigserial primary key,
	createdate timestamp without time zone NOT NULL DEFAULT now()::timestamp without time zone,
	updatedate timestamp without time zone NOT NULL DEFAULT now()::timestamp without time zone,
	id_author bigint references author(id),
	text text NOT NULL,
	likes int DEFAULT 0,
	dislikes int DEFAULT 0,
	id_news bigint references news(id)
);



