CREATE TABLE public.cocktails (
    id uuid,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    data_url text,
    base_spirit text DEFAULT 'unset'::text NOT NULL,
    cocktail_type text DEFAULT 'unset'::text NOT NULL,
    name text NOT NULL,
    img_name text,
    type text,
    is_new boolean DEFAULT false NOT NULL,
    is_mocktail BOOLEAN NOT NULL DEFAULT FALSE,
    is_available BOOLEAN NOT NULL DEFAULT TRUE
);


CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now() NOT NULL
);



CREATE TABLE public.ingredients (
    id uuid NOT NULL,
    name text NOT NULL,
    abv real NOT NULL,
    is_available boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone
);



CREATE TABLE public.orders (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone,
    cocktail_id uuid NOT NULL,
    ordered_by text NOT NULL,
    canceled_at timestamp without time zone,
    finished boolean DEFAULT false
);


CREATE TABLE public.recipies (
    cocktail_id uuid NOT NULL,
    ingredient_id uuid NOT NULL,
    amount integer NOT NULL,
    unit text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone
);


CREATE TABLE public.users(
	id UUID UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	modified_at TIMESTAMP NOT NULL,
	last_seen_at TIMESTAMP NOT NULL
);
