-- Creates table actions if doesn't exist
CREATE TABLE IF NOT EXISTS stocktrader.actions (
    uid character varying NOT null,
    symbol character varying (50),
    company_name character varying (50),
    action_type character varying (50),
    current_value double precision,
    added_timestamp timestamp DEFAULT CURRENT_TIMESTAMP ,
    is_deleted boolean DEFAULT false,
    CONSTRAINT actions_pkey PRIMARY KEY (uid)
);