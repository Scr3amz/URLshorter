CREATE TABLE public."urls"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    "longurl" text NOT NULL,
    "shorturl" character varying(29) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public."urls"
    OWNER to postgres;