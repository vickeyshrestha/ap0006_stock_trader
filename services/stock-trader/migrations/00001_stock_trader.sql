-- +goose Up
-- Create schema if not exists
CREATE SCHEMA IF NOT EXISTS stocktrader;

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

-- +goose Down
-- Drops table actions and schema if exists
DROP TABLE IF EXISTS stocktrader.actions;
DROP SCHEMA IF EXISTS stocktrader;
