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

