-- Таблица пользователей телеграма, которые хотя бы раз начинали игру
CREATE TABLE public.t_user (
	user_id int8 NOT NULL,
	username varchar NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	lang varchar NULL,
	creation_date timestamp NOT NULL
);
CREATE UNIQUE INDEX t_user_user_id_idx ON public.t_user (user_id);

-- таблица различных типов местностей
CREATE TABLE public.map_terrain (
	id int NOT NULL GENERATED ALWAYS AS IDENTITY,
	"type" varchar NULL,
	name varchar NULL
);
CREATE INDEX map_terrain_id_idx ON public.map_terrain (id);

--таблица чанков с координатами
CREATE TABLE public.map_chunk (
	id int NOT NULL,
	x int NOT NULL,
	y int NOT NULL
);
CREATE INDEX map_chunk_id_idx ON public.map_chunk (id);

--таблица местностей для каждого чанка
CREATE TABLE public.map_chunk_terrains (
	chunk_id int NOT NULL,
	terrain_id int NOT NULL,
	"size" int NOT NULL
);

--таблица рек для чанка
CREATE TABLE public.map_chunk_rivers (
	chunk_id int NOT NULL,
	"from" varchar NOT NULL,
	"to" varchar NOT NULL,
	"size" int NOT NULL,
	is_bridge boolean NOT NULL DEFAULT false
);