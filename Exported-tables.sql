--
-- PostgreSQL database dump
--

-- Dumped from database version 16.8 (Ubuntu 16.8-0ubuntu0.24.04.1)
-- Dumped by pg_dump version 16.8

-- Started on 2025-05-11 13:21:09 UTC

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
-- TOC entry 218 (class 1259 OID 22536)
-- Name: beehive; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.beehive (
    beehive_id integer NOT NULL,
    beehive_number character varying DEFAULT 1 NOT NULL,
    beekeeper_id integer NOT NULL,
    beetype_id integer NOT NULL,
    plant_id integer NOT NULL,
    location_id integer NOT NULL,
    beehive_status character varying DEFAULT 'Active'::character varying NOT NULL,
    beehivestartdate timestamp without time zone NOT NULL,
    beehive_description text
);


ALTER TABLE public.beehive OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 22535)
-- Name: beehive_beehive_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.beehive_beehive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.beehive_beehive_id_seq OWNER TO postgres;

--
-- TOC entry 3567 (class 0 OID 0)
-- Dependencies: 217
-- Name: beehive_beehive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.beehive_beehive_id_seq OWNED BY public.beehive.beehive_id;


--
-- TOC entry 216 (class 1259 OID 22527)
-- Name: beekeeper; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.beekeeper (
    beekeeper_id integer NOT NULL,
    land_area character varying,
    location_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.beekeeper OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 22526)
-- Name: beekeeper_beekeeper_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.beekeeper_beekeeper_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.beekeeper_beekeeper_id_seq OWNER TO postgres;

--
-- TOC entry 3570 (class 0 OID 0)
-- Dependencies: 215
-- Name: beekeeper_beekeeper_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.beekeeper_beekeeper_id_seq OWNED BY public.beekeeper.beekeeper_id;


--
-- TOC entry 238 (class 1259 OID 22624)
-- Name: webuser; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.webuser (
    user_id integer NOT NULL,
    email character varying NOT NULL,
    role_id integer,
    createdat timestamp without time zone DEFAULT now(),
    name character varying NOT NULL,
    birth_date timestamp without time zone,
    phone bigint,
    username character varying NOT NULL
);


ALTER TABLE public.webuser OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 23028)
-- Name: beekeeper_with_name; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW public.beekeeper_with_name AS
 SELECT b.beekeeper_id,
    u.name
   FROM (public.beekeeper b
     JOIN public.webuser u ON ((b.user_id = u.user_id)));


ALTER VIEW public.beekeeper_with_name OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 22554)
-- Name: beetype; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.beetype (
    beetype_id integer NOT NULL,
    beetype_name character varying NOT NULL,
    beetype_description text
);


ALTER TABLE public.beetype OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 22553)
-- Name: beetype_beetype_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.beetype_beetype_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.beetype_beetype_id_seq OWNER TO postgres;

--
-- TOC entry 3575 (class 0 OID 0)
-- Dependencies: 221
-- Name: beetype_beetype_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.beetype_beetype_id_seq OWNED BY public.beetype.beetype_id;


--
-- TOC entry 232 (class 1259 OID 22599)
-- Name: customer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer (
    customer_id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    phone character varying,
    address text,
    createdat timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.customer OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 22598)
-- Name: customer_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customer_customer_id_seq OWNER TO postgres;

--
-- TOC entry 3578 (class 0 OID 0)
-- Dependencies: 231
-- Name: customer_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customer_customer_id_seq OWNED BY public.customer.customer_id;


--
-- TOC entry 220 (class 1259 OID 22545)
-- Name: geolocation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.geolocation (
    location_id integer NOT NULL,
    province_name character varying NOT NULL,
    district_name character varying NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    location_name character varying NOT NULL
);


ALTER TABLE public.geolocation OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 22544)
-- Name: geolocation_location_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.geolocation_location_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.geolocation_location_id_seq OWNER TO postgres;

--
-- TOC entry 3581 (class 0 OID 0)
-- Dependencies: 219
-- Name: geolocation_location_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.geolocation_location_id_seq OWNED BY public.geolocation.location_id;


--
-- TOC entry 228 (class 1259 OID 22581)
-- Name: harvestlog; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.harvestlog (
    harvest_id integer NOT NULL,
    beehive_id integer NOT NULL,
    harvestdate timestamp without time zone NOT NULL,
    production numeric NOT NULL,
    production_note text
);


ALTER TABLE public.harvestlog OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 22580)
-- Name: harvestlog_harvest_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.harvestlog_harvest_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.harvestlog_harvest_id_seq OWNER TO postgres;

--
-- TOC entry 3584 (class 0 OID 0)
-- Dependencies: 227
-- Name: harvestlog_harvest_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.harvestlog_harvest_id_seq OWNED BY public.harvestlog.harvest_id;


--
-- TOC entry 230 (class 1259 OID 22590)
-- Name: honeystock; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.honeystock (
    stock_id integer NOT NULL,
    beehive_id integer NOT NULL,
    quantity numeric NOT NULL,
    harvest_id integer NOT NULL,
    stock_date timestamp without time zone NOT NULL,
    is_sold boolean DEFAULT false NOT NULL,
    description character varying
);


ALTER TABLE public.honeystock OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 22589)
-- Name: honeystock_stock_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.honeystock_stock_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.honeystock_stock_id_seq OWNER TO postgres;

--
-- TOC entry 3587 (class 0 OID 0)
-- Dependencies: 229
-- Name: honeystock_stock_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.honeystock_stock_id_seq OWNED BY public.honeystock.stock_id;


--
-- TOC entry 236 (class 1259 OID 22617)
-- Name: orderitem; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orderitem (
    orderitem_id integer NOT NULL,
    order_id integer NOT NULL,
    price double precision NOT NULL,
    stock_id integer NOT NULL
);


ALTER TABLE public.orderitem OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 22616)
-- Name: orderitem_orderitem_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orderitem_orderitem_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orderitem_orderitem_id_seq OWNER TO postgres;

--
-- TOC entry 3590 (class 0 OID 0)
-- Dependencies: 235
-- Name: orderitem_orderitem_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orderitem_orderitem_id_seq OWNED BY public.orderitem.orderitem_id;


--
-- TOC entry 234 (class 1259 OID 22608)
-- Name: orderlist; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orderlist (
    order_id integer NOT NULL,
    customer_id integer NOT NULL,
    order_date timestamp without time zone NOT NULL,
    status character varying DEFAULT 'Pending'::character varying NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.orderlist OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 22607)
-- Name: orderlist_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orderlist_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orderlist_order_id_seq OWNER TO postgres;

--
-- TOC entry 3593 (class 0 OID 0)
-- Dependencies: 233
-- Name: orderlist_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orderlist_order_id_seq OWNED BY public.orderlist.order_id;


--
-- TOC entry 224 (class 1259 OID 22563)
-- Name: planttype; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.planttype (
    plant_id integer NOT NULL,
    plant_name character varying NOT NULL,
    anti_oxidant character varying,
    pollen character varying,
    plant_description text,
    flowering_season character varying,
    climatezone character varying
);


ALTER TABLE public.planttype OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 22562)
-- Name: planttype_plant_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.planttype_plant_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.planttype_plant_id_seq OWNER TO postgres;

--
-- TOC entry 3596 (class 0 OID 0)
-- Dependencies: 223
-- Name: planttype_plant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.planttype_plant_id_seq OWNED BY public.planttype.plant_id;


--
-- TOC entry 226 (class 1259 OID 22572)
-- Name: queenbee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.queenbee (
    queenbee_id integer NOT NULL,
    beehive_id integer NOT NULL,
    startdate timestamp without time zone NOT NULL,
    enddate timestamp without time zone NOT NULL,
    origin character varying
);


ALTER TABLE public.queenbee OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 22571)
-- Name: queenbee_queenbee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.queenbee_queenbee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.queenbee_queenbee_id_seq OWNER TO postgres;

--
-- TOC entry 3599 (class 0 OID 0)
-- Dependencies: 225
-- Name: queenbee_queenbee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.queenbee_queenbee_id_seq OWNED BY public.queenbee.queenbee_id;


--
-- TOC entry 240 (class 1259 OID 22633)
-- Name: webrole; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.webrole (
    role_id integer NOT NULL,
    role_name character varying NOT NULL,
    role_description character varying
);


ALTER TABLE public.webrole OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 22632)
-- Name: staffrole_role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.staffrole_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.staffrole_role_id_seq OWNER TO postgres;

--
-- TOC entry 3602 (class 0 OID 0)
-- Dependencies: 239
-- Name: staffrole_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.staffrole_role_id_seq OWNED BY public.webrole.role_id;


--
-- TOC entry 237 (class 1259 OID 22623)
-- Name: webuser_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.webuser_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.webuser_user_id_seq OWNER TO postgres;

--
-- TOC entry 3604 (class 0 OID 0)
-- Dependencies: 237
-- Name: webuser_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.webuser_user_id_seq OWNED BY public.webuser.user_id;


--
-- TOC entry 3300 (class 2604 OID 22539)
-- Name: beehive beehive_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive ALTER COLUMN beehive_id SET DEFAULT nextval('public.beehive_beehive_id_seq'::regclass);


--
-- TOC entry 3299 (class 2604 OID 22530)
-- Name: beekeeper beekeeper_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beekeeper ALTER COLUMN beekeeper_id SET DEFAULT nextval('public.beekeeper_beekeeper_id_seq'::regclass);


--
-- TOC entry 3304 (class 2604 OID 22557)
-- Name: beetype beetype_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beetype ALTER COLUMN beetype_id SET DEFAULT nextval('public.beetype_beetype_id_seq'::regclass);


--
-- TOC entry 3310 (class 2604 OID 22602)
-- Name: customer customer_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer ALTER COLUMN customer_id SET DEFAULT nextval('public.customer_customer_id_seq'::regclass);


--
-- TOC entry 3303 (class 2604 OID 22548)
-- Name: geolocation location_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.geolocation ALTER COLUMN location_id SET DEFAULT nextval('public.geolocation_location_id_seq'::regclass);


--
-- TOC entry 3307 (class 2604 OID 22584)
-- Name: harvestlog harvest_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.harvestlog ALTER COLUMN harvest_id SET DEFAULT nextval('public.harvestlog_harvest_id_seq'::regclass);


--
-- TOC entry 3308 (class 2604 OID 22593)
-- Name: honeystock stock_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.honeystock ALTER COLUMN stock_id SET DEFAULT nextval('public.honeystock_stock_id_seq'::regclass);


--
-- TOC entry 3314 (class 2604 OID 22620)
-- Name: orderitem orderitem_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderitem ALTER COLUMN orderitem_id SET DEFAULT nextval('public.orderitem_orderitem_id_seq'::regclass);


--
-- TOC entry 3312 (class 2604 OID 22611)
-- Name: orderlist order_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderlist ALTER COLUMN order_id SET DEFAULT nextval('public.orderlist_order_id_seq'::regclass);


--
-- TOC entry 3305 (class 2604 OID 22566)
-- Name: planttype plant_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.planttype ALTER COLUMN plant_id SET DEFAULT nextval('public.planttype_plant_id_seq'::regclass);


--
-- TOC entry 3306 (class 2604 OID 22575)
-- Name: queenbee queenbee_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queenbee ALTER COLUMN queenbee_id SET DEFAULT nextval('public.queenbee_queenbee_id_seq'::regclass);


--
-- TOC entry 3317 (class 2604 OID 22636)
-- Name: webrole role_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webrole ALTER COLUMN role_id SET DEFAULT nextval('public.staffrole_role_id_seq'::regclass);


--
-- TOC entry 3315 (class 2604 OID 22627)
-- Name: webuser user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webuser ALTER COLUMN user_id SET DEFAULT nextval('public.webuser_user_id_seq'::regclass);


--
-- TOC entry 3357 (class 2606 OID 23216)
-- Name: queenbee Beehive; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queenbee
    ADD CONSTRAINT "Beehive" UNIQUE (beehive_id);


--
-- TOC entry 3335 (class 2606 OID 23117)
-- Name: beehive Beehive U; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive
    ADD CONSTRAINT "Beehive U" UNIQUE (beehive_id);


--
-- TOC entry 3329 (class 2606 OID 23115)
-- Name: beekeeper Beekeeper ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beekeeper
    ADD CONSTRAINT "Beekeeper ID" UNIQUE (beekeeper_id, user_id);


--
-- TOC entry 3345 (class 2606 OID 23113)
-- Name: beetype Beetype ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beetype
    ADD CONSTRAINT "Beetype ID" UNIQUE (beetype_name, beetype_id);


--
-- TOC entry 3347 (class 2606 OID 23230)
-- Name: beetype Beetype Name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beetype
    ADD CONSTRAINT "Beetype Name" UNIQUE (beetype_name);


--
-- TOC entry 3371 (class 2606 OID 23226)
-- Name: customer Customer email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT "Customer email" UNIQUE (email);


--
-- TOC entry 3373 (class 2606 OID 23109)
-- Name: customer Customer iD; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT "Customer iD" UNIQUE (email, phone, customer_id);


--
-- TOC entry 3375 (class 2606 OID 23228)
-- Name: customer Customer phone; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT "Customer phone" UNIQUE (phone);


--
-- TOC entry 3318 (class 2606 OID 23399)
-- Name: queenbee Dates; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.queenbee
    ADD CONSTRAINT "Dates" CHECK ((startdate <= enddate)) NOT VALID;


--
-- TOC entry 3339 (class 2606 OID 23107)
-- Name: geolocation Geolocation ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.geolocation
    ADD CONSTRAINT "Geolocation ID" UNIQUE (location_id, location_name);


--
-- TOC entry 3341 (class 2606 OID 23224)
-- Name: geolocation Geolocation name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.geolocation
    ADD CONSTRAINT "Geolocation name" UNIQUE (location_name);


--
-- TOC entry 3367 (class 2606 OID 23103)
-- Name: honeystock Honey Stock U; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.honeystock
    ADD CONSTRAINT "Honey Stock U" UNIQUE (stock_id);


--
-- TOC entry 3383 (class 2606 OID 23101)
-- Name: orderitem Orderitem ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderitem
    ADD CONSTRAINT "Orderitem ID" UNIQUE (stock_id, orderitem_id);


--
-- TOC entry 3379 (class 2606 OID 23099)
-- Name: orderlist Orderlist U; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderlist
    ADD CONSTRAINT "Orderlist U" UNIQUE (order_id);


--
-- TOC entry 3351 (class 2606 OID 23218)
-- Name: planttype Plant; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.planttype
    ADD CONSTRAINT "Plant" UNIQUE (plant_name);


--
-- TOC entry 3353 (class 2606 OID 23097)
-- Name: planttype Plant ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.planttype
    ADD CONSTRAINT "Plant ID" UNIQUE (plant_id, plant_name);


--
-- TOC entry 3326 (class 2606 OID 23403)
-- Name: orderitem Price validate; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.orderitem
    ADD CONSTRAINT "Price validate" CHECK ((price > (0)::double precision)) NOT VALID;


--
-- TOC entry 3359 (class 2606 OID 23095)
-- Name: queenbee Queenbee ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queenbee
    ADD CONSTRAINT "Queenbee ID" UNIQUE (queenbee_id, beehive_id);


--
-- TOC entry 3397 (class 2606 OID 23093)
-- Name: webrole Role; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webrole
    ADD CONSTRAINT "Role" UNIQUE (role_name, role_id);


--
-- TOC entry 3385 (class 2606 OID 23220)
-- Name: orderitem Stock_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderitem
    ADD CONSTRAINT "Stock_id" UNIQUE (stock_id);


--
-- TOC entry 3331 (class 2606 OID 23234)
-- Name: beekeeper User ID; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beekeeper
    ADD CONSTRAINT "User ID" UNIQUE (user_id);


--
-- TOC entry 3337 (class 2606 OID 22543)
-- Name: beehive beehive_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive
    ADD CONSTRAINT beehive_pkey PRIMARY KEY (beehive_id);


--
-- TOC entry 3333 (class 2606 OID 22534)
-- Name: beekeeper beekeeper_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beekeeper
    ADD CONSTRAINT beekeeper_pkey PRIMARY KEY (beekeeper_id);


--
-- TOC entry 3349 (class 2606 OID 22561)
-- Name: beetype beetype_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beetype
    ADD CONSTRAINT beetype_pkey PRIMARY KEY (beetype_id);


--
-- TOC entry 3327 (class 2606 OID 23409)
-- Name: webuser birthdate validation; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.webuser
    ADD CONSTRAINT "birthdate validation" CHECK (((birth_date IS NULL) OR (birth_date <= (CURRENT_DATE - '18 years'::interval)))) NOT VALID;


--
-- TOC entry 3323 (class 2606 OID 23408)
-- Name: customer createdat validate; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.customer
    ADD CONSTRAINT "createdat validate" CHECK ((createdat <= now())) NOT VALID;


--
-- TOC entry 3377 (class 2606 OID 22606)
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);


--
-- TOC entry 3389 (class 2606 OID 23210)
-- Name: webuser email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webuser
    ADD CONSTRAINT email UNIQUE (email);


--
-- TOC entry 3343 (class 2606 OID 22552)
-- Name: geolocation geolocation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.geolocation
    ADD CONSTRAINT geolocation_pkey PRIMARY KEY (location_id);


--
-- TOC entry 3363 (class 2606 OID 23222)
-- Name: harvestlog harvest_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.harvestlog
    ADD CONSTRAINT harvest_id UNIQUE (harvest_id);


--
-- TOC entry 3319 (class 2606 OID 23406)
-- Name: harvestlog harvestdate validate; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.harvestlog
    ADD CONSTRAINT "harvestdate validate" CHECK ((harvestdate <= now())) NOT VALID;


--
-- TOC entry 3365 (class 2606 OID 22588)
-- Name: harvestlog harvestlog_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.harvestlog
    ADD CONSTRAINT harvestlog_pkey PRIMARY KEY (harvest_id);


--
-- TOC entry 3369 (class 2606 OID 22597)
-- Name: honeystock honeystock_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.honeystock
    ADD CONSTRAINT honeystock_pkey PRIMARY KEY (stock_id);


--
-- TOC entry 3324 (class 2606 OID 23401)
-- Name: orderlist order_date validation; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.orderlist
    ADD CONSTRAINT "order_date validation" CHECK ((order_date <= now())) NOT VALID;


--
-- TOC entry 3387 (class 2606 OID 22622)
-- Name: orderitem orderitem_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderitem
    ADD CONSTRAINT orderitem_pkey PRIMARY KEY (orderitem_id);


--
-- TOC entry 3381 (class 2606 OID 22615)
-- Name: orderlist orderlist_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderlist
    ADD CONSTRAINT orderlist_pkey PRIMARY KEY (order_id);


--
-- TOC entry 3391 (class 2606 OID 23212)
-- Name: webuser phone; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webuser
    ADD CONSTRAINT phone UNIQUE (phone);


--
-- TOC entry 3355 (class 2606 OID 22570)
-- Name: planttype planttype_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.planttype
    ADD CONSTRAINT planttype_pkey PRIMARY KEY (plant_id);


--
-- TOC entry 3320 (class 2606 OID 23407)
-- Name: harvestlog production validate; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.harvestlog
    ADD CONSTRAINT "production validate" CHECK ((production > (0)::numeric)) NOT VALID;


--
-- TOC entry 3321 (class 2606 OID 23404)
-- Name: honeystock quantity validate; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.honeystock
    ADD CONSTRAINT "quantity validate" CHECK ((quantity > (0)::numeric)) NOT VALID;


--
-- TOC entry 3361 (class 2606 OID 22579)
-- Name: queenbee queenbee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queenbee
    ADD CONSTRAINT queenbee_pkey PRIMARY KEY (queenbee_id);


--
-- TOC entry 3399 (class 2606 OID 23214)
-- Name: webrole role_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webrole
    ADD CONSTRAINT role_name UNIQUE (role_name);


--
-- TOC entry 3401 (class 2606 OID 22640)
-- Name: webrole staffrole_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webrole
    ADD CONSTRAINT staffrole_pkey PRIMARY KEY (role_id);


--
-- TOC entry 3325 (class 2606 OID 23402)
-- Name: orderlist status validation; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.orderlist
    ADD CONSTRAINT "status validation" CHECK (((status)::text = ANY ((ARRAY['Pending'::character varying, 'Delivered'::character varying])::text[]))) NOT VALID;


--
-- TOC entry 3322 (class 2606 OID 23405)
-- Name: honeystock stock_date validate; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.honeystock
    ADD CONSTRAINT "stock_date validate" CHECK ((stock_date <= now())) NOT VALID;


--
-- TOC entry 3393 (class 2606 OID 23208)
-- Name: webuser user_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webuser
    ADD CONSTRAINT user_id UNIQUE (user_id);


--
-- TOC entry 3395 (class 2606 OID 22631)
-- Name: webuser webuser_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webuser
    ADD CONSTRAINT webuser_pkey PRIMARY KEY (user_id);


--
-- TOC entry 3404 (class 2606 OID 22651)
-- Name: beehive beehive_beekeeper_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive
    ADD CONSTRAINT beehive_beekeeper_id_fkey FOREIGN KEY (beekeeper_id) REFERENCES public.beekeeper(beekeeper_id);


--
-- TOC entry 3405 (class 2606 OID 22656)
-- Name: beehive beehive_beetype_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive
    ADD CONSTRAINT beehive_beetype_id_fkey FOREIGN KEY (beetype_id) REFERENCES public.beetype(beetype_id);


--
-- TOC entry 3406 (class 2606 OID 22666)
-- Name: beehive beehive_location_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive
    ADD CONSTRAINT beehive_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.geolocation(location_id);


--
-- TOC entry 3407 (class 2606 OID 22661)
-- Name: beehive beehive_plant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beehive
    ADD CONSTRAINT beehive_plant_id_fkey FOREIGN KEY (plant_id) REFERENCES public.planttype(plant_id);


--
-- TOC entry 3402 (class 2606 OID 22641)
-- Name: beekeeper beekeeper_location_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beekeeper
    ADD CONSTRAINT beekeeper_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.geolocation(location_id);


--
-- TOC entry 3403 (class 2606 OID 22646)
-- Name: beekeeper beekeeper_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.beekeeper
    ADD CONSTRAINT beekeeper_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.webuser(user_id);


--
-- TOC entry 3409 (class 2606 OID 22676)
-- Name: harvestlog harvestlog_beehive_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.harvestlog
    ADD CONSTRAINT harvestlog_beehive_id_fkey FOREIGN KEY (beehive_id) REFERENCES public.beehive(beehive_id);


--
-- TOC entry 3410 (class 2606 OID 22681)
-- Name: honeystock honeystock_beehive_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.honeystock
    ADD CONSTRAINT honeystock_beehive_id_fkey FOREIGN KEY (beehive_id) REFERENCES public.beehive(beehive_id);


--
-- TOC entry 3411 (class 2606 OID 22686)
-- Name: honeystock honeystock_harvest_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.honeystock
    ADD CONSTRAINT honeystock_harvest_id_fkey FOREIGN KEY (harvest_id) REFERENCES public.harvestlog(harvest_id);


--
-- TOC entry 3414 (class 2606 OID 22701)
-- Name: orderitem orderitem_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderitem
    ADD CONSTRAINT orderitem_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orderlist(order_id);


--
-- TOC entry 3415 (class 2606 OID 23016)
-- Name: orderitem orderitem_stock_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderitem
    ADD CONSTRAINT orderitem_stock_id_fkey FOREIGN KEY (stock_id) REFERENCES public.honeystock(stock_id);


--
-- TOC entry 3412 (class 2606 OID 22696)
-- Name: orderlist orderlist_customer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderlist
    ADD CONSTRAINT orderlist_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customer(customer_id);


--
-- TOC entry 3413 (class 2606 OID 23021)
-- Name: orderlist orderlist_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orderlist
    ADD CONSTRAINT orderlist_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.webuser(user_id);


--
-- TOC entry 3408 (class 2606 OID 23120)
-- Name: queenbee queenbee_beehive_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queenbee
    ADD CONSTRAINT queenbee_beehive_id_fkey FOREIGN KEY (beehive_id) REFERENCES public.beehive(beehive_id) ON DELETE CASCADE;


--
-- TOC entry 3416 (class 2606 OID 22711)
-- Name: webuser webuser_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.webuser
    ADD CONSTRAINT webuser_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.webrole(role_id);


--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 218
-- Name: TABLE beehive; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.beehive TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.beehive TO api_user;


--
-- TOC entry 3568 (class 0 OID 0)
-- Dependencies: 217
-- Name: SEQUENCE beehive_beehive_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.beehive_beehive_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.beehive_beehive_id_seq TO api_user;


--
-- TOC entry 3569 (class 0 OID 0)
-- Dependencies: 216
-- Name: TABLE beekeeper; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.beekeeper TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.beekeeper TO api_user;


--
-- TOC entry 3571 (class 0 OID 0)
-- Dependencies: 215
-- Name: SEQUENCE beekeeper_beekeeper_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.beekeeper_beekeeper_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.beekeeper_beekeeper_id_seq TO api_user;


--
-- TOC entry 3572 (class 0 OID 0)
-- Dependencies: 238
-- Name: TABLE webuser; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.webuser TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.webuser TO api_user;


--
-- TOC entry 3573 (class 0 OID 0)
-- Dependencies: 241
-- Name: TABLE beekeeper_with_name; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE public.beekeeper_with_name TO api_user;
GRANT SELECT ON TABLE public.beekeeper_with_name TO web_anon;


--
-- TOC entry 3574 (class 0 OID 0)
-- Dependencies: 222
-- Name: TABLE beetype; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.beetype TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.beetype TO api_user;


--
-- TOC entry 3576 (class 0 OID 0)
-- Dependencies: 221
-- Name: SEQUENCE beetype_beetype_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.beetype_beetype_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.beetype_beetype_id_seq TO api_user;


--
-- TOC entry 3577 (class 0 OID 0)
-- Dependencies: 232
-- Name: TABLE customer; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.customer TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.customer TO api_user;


--
-- TOC entry 3579 (class 0 OID 0)
-- Dependencies: 231
-- Name: SEQUENCE customer_customer_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.customer_customer_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.customer_customer_id_seq TO api_user;


--
-- TOC entry 3580 (class 0 OID 0)
-- Dependencies: 220
-- Name: TABLE geolocation; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE public.geolocation TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.geolocation TO api_user;


--
-- TOC entry 3582 (class 0 OID 0)
-- Dependencies: 219
-- Name: SEQUENCE geolocation_location_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.geolocation_location_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.geolocation_location_id_seq TO api_user;


--
-- TOC entry 3583 (class 0 OID 0)
-- Dependencies: 228
-- Name: TABLE harvestlog; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.harvestlog TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.harvestlog TO api_user;


--
-- TOC entry 3585 (class 0 OID 0)
-- Dependencies: 227
-- Name: SEQUENCE harvestlog_harvest_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.harvestlog_harvest_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.harvestlog_harvest_id_seq TO api_user;


--
-- TOC entry 3586 (class 0 OID 0)
-- Dependencies: 230
-- Name: TABLE honeystock; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.honeystock TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.honeystock TO api_user;


--
-- TOC entry 3588 (class 0 OID 0)
-- Dependencies: 229
-- Name: SEQUENCE honeystock_stock_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.honeystock_stock_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.honeystock_stock_id_seq TO api_user;


--
-- TOC entry 3589 (class 0 OID 0)
-- Dependencies: 236
-- Name: TABLE orderitem; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.orderitem TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.orderitem TO api_user;


--
-- TOC entry 3591 (class 0 OID 0)
-- Dependencies: 235
-- Name: SEQUENCE orderitem_orderitem_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.orderitem_orderitem_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.orderitem_orderitem_id_seq TO api_user;


--
-- TOC entry 3592 (class 0 OID 0)
-- Dependencies: 234
-- Name: TABLE orderlist; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.orderlist TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.orderlist TO api_user;


--
-- TOC entry 3594 (class 0 OID 0)
-- Dependencies: 233
-- Name: SEQUENCE orderlist_order_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.orderlist_order_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.orderlist_order_id_seq TO api_user;


--
-- TOC entry 3595 (class 0 OID 0)
-- Dependencies: 224
-- Name: TABLE planttype; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.planttype TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.planttype TO api_user;


--
-- TOC entry 3597 (class 0 OID 0)
-- Dependencies: 223
-- Name: SEQUENCE planttype_plant_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.planttype_plant_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.planttype_plant_id_seq TO api_user;


--
-- TOC entry 3598 (class 0 OID 0)
-- Dependencies: 226
-- Name: TABLE queenbee; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.queenbee TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.queenbee TO api_user;


--
-- TOC entry 3600 (class 0 OID 0)
-- Dependencies: 225
-- Name: SEQUENCE queenbee_queenbee_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.queenbee_queenbee_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.queenbee_queenbee_id_seq TO api_user;


--
-- TOC entry 3601 (class 0 OID 0)
-- Dependencies: 240
-- Name: TABLE webrole; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE public.webrole TO web_anon;
GRANT SELECT,INSERT,UPDATE ON TABLE public.webrole TO api_user;


--
-- TOC entry 3603 (class 0 OID 0)
-- Dependencies: 239
-- Name: SEQUENCE staffrole_role_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.staffrole_role_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.staffrole_role_id_seq TO api_user;


--
-- TOC entry 3605 (class 0 OID 0)
-- Dependencies: 237
-- Name: SEQUENCE webuser_user_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.webuser_user_id_seq TO web_anon;
GRANT SELECT,USAGE ON SEQUENCE public.webuser_user_id_seq TO api_user;


-- Completed on 2025-05-11 13:21:09 UTC

--
-- PostgreSQL database dump complete
--

