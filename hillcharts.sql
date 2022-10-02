--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: frame_scopes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.frame_scopes (
    id integer NOT NULL,
    title character varying(50) NOT NULL,
    "position" numeric(5,2) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone,
    frame_id integer NOT NULL,
    scope_id integer NOT NULL
);


ALTER TABLE public.frame_scopes OWNER TO postgres;

--
-- Name: frames; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.frames (
    id integer NOT NULL,
    hillchart_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.frames OWNER TO postgres;

--
-- Name: hillcharts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.hillcharts (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.hillcharts OWNER TO postgres;

--
-- Name: hillcharts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.hillcharts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.hillcharts_id_seq OWNER TO postgres;

--
-- Name: hillcharts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.hillcharts_id_seq OWNED BY public.hillcharts.id;


--
-- Name: scopes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scopes (
    id integer NOT NULL,
    color character varying(7),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone,
    hillchart_id integer NOT NULL
);


ALTER TABLE public.scopes OWNER TO postgres;

--
-- Name: scopes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scopes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scopes_id_seq OWNER TO postgres;

--
-- Name: scopes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scopes_id_seq OWNED BY public.frame_scopes.id;


--
-- Name: scopes_id_seq1; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scopes_id_seq1
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scopes_id_seq1 OWNER TO postgres;

--
-- Name: scopes_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scopes_id_seq1 OWNED BY public.scopes.id;


--
-- Name: statuses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.statuses_id_seq OWNER TO postgres;

--
-- Name: statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.statuses_id_seq OWNED BY public.frames.id;


--
-- Name: frame_scopes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frame_scopes ALTER COLUMN id SET DEFAULT nextval('public.scopes_id_seq'::regclass);


--
-- Name: frames id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frames ALTER COLUMN id SET DEFAULT nextval('public.statuses_id_seq'::regclass);


--
-- Name: hillcharts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hillcharts ALTER COLUMN id SET DEFAULT nextval('public.hillcharts_id_seq'::regclass);


--
-- Name: scopes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scopes ALTER COLUMN id SET DEFAULT nextval('public.scopes_id_seq1'::regclass);


--
-- Name: frame_scopes frame_id_scope_id_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frame_scopes
    ADD CONSTRAINT frame_id_scope_id_unique UNIQUE (frame_id, scope_id);


--
-- Name: hillcharts hillcharts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hillcharts
    ADD CONSTRAINT hillcharts_pkey PRIMARY KEY (id);


--
-- Name: frame_scopes scopes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frame_scopes
    ADD CONSTRAINT scopes_pkey PRIMARY KEY (id);


--
-- Name: scopes scopes_pkey1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scopes
    ADD CONSTRAINT scopes_pkey1 PRIMARY KEY (id);


--
-- Name: frames statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frames
    ADD CONSTRAINT statuses_pkey PRIMARY KEY (id);


--
-- Name: frames fk_hillchart; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frames
    ADD CONSTRAINT fk_hillchart FOREIGN KEY (hillchart_id) REFERENCES public.hillcharts(id);


--
-- Name: frame_scopes frame_scopes_scope_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frame_scopes
    ADD CONSTRAINT frame_scopes_scope_id_fkey FOREIGN KEY (scope_id) REFERENCES public.scopes(id);


--
-- Name: frame_scopes scopes_frame_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.frame_scopes
    ADD CONSTRAINT scopes_frame_id_fkey FOREIGN KEY (frame_id) REFERENCES public.frames(id);


--
-- Name: scopes scopes_hillchart_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scopes
    ADD CONSTRAINT scopes_hillchart_id_fkey FOREIGN KEY (hillchart_id) REFERENCES public.hillcharts(id);


--
-- PostgreSQL database dump complete
--

