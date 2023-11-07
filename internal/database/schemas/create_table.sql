CREATE TABLE public."URLs"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    "longURL" text NOT NULL,
    "shortURL" character varying(29) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."URLs"
    OWNER to postgres;