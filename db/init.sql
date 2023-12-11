--
-- PostgreSQL database dump
--

-- Dumped from database version 15.5 (Homebrew)
-- Dumped by pg_dump version 16.0

-- Started on 2023-12-11 11:11:43 +03

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
-- TOC entry 214 (class 1259 OID 16424)
-- Name: customers; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.customers (
    id integer NOT NULL,
    login character varying(20) NOT NULL,
    password character varying(300) NOT NULL,
    email character varying(30),
    role character varying(10) NOT NULL
);


ALTER TABLE public.customers OWNER TO root;

--
-- TOC entry 217 (class 1259 OID 16449)
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.customers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 216 (class 1259 OID 16434)
-- Name: customers_products; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.customers_products (
    id integer NOT NULL,
    product_id integer NOT NULL,
    customer_id integer NOT NULL
);


ALTER TABLE public.customers_products OWNER TO root;

--
-- TOC entry 219 (class 1259 OID 16451)
-- Name: customers_products_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.customers_products ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.customers_products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 215 (class 1259 OID 16429)
-- Name: products; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    "desc" character varying(2000),
    type character varying(20),
    price real,
    image character varying(200)
);


ALTER TABLE public.products OWNER TO root;

--
-- TOC entry 218 (class 1259 OID 16450)
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

ALTER TABLE public.products ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3621 (class 0 OID 16424)
-- Dependencies: 214
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.customers (id, login, password, email, role) FROM stdin;
2				
3				
4	zaher	123		admin
5	test	$2a$10$YdySRnNHUxWRyO9mwPr7j.CiTbl7mHnUmdIphwqhwtzP2012OH3Ku		user
\.


--
-- TOC entry 3623 (class 0 OID 16434)
-- Dependencies: 216
-- Data for Name: customers_products; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.customers_products (id, product_id, customer_id) FROM stdin;
16	11	5
17	13	5
18	11	4
19	13	5
20	12	5
21	11	5
\.


--
-- TOC entry 3622 (class 0 OID 16429)
-- Dependencies: 215
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.products (id, name, "desc", type, price, image) FROM stdin;
11	AM One\n	Данный товар является частью проекта Lamoda planet - специального раздела нашего каталога, где мы собрали экологичные, этичные, инклюзивные и благотворительные товары.\n \n Товар произведен в стране присутствия Lamoda, что позволяет нам оптимизировать выбросы СО2 при доставке. Покупая этот товар, вы вносите свой вклад в сокращение углеродного следа и поддерживаете развитие локальных фабрик.	Платье	5300	pl.jpg
12	Золотой песок	Данный товар является частью проекта Lamoda planet - специального раздела нашего каталога, где мы собрали экологичные, этичные, инклюзивные и благотворительные товары. Товар произведен в стране присутствия Lamoda, что позволяет поддерживать развитие локальных фабрик, мастеров и швей. Покупая этот товар, вы также вносите свой вклад в сокращение углеродного следа.	Платье	7200	gold_sand.jpeg
13	Marichuell	Данный товар является частью проекта Lamoda planet - специального раздела нашего каталога, где мы собрали экологичные, этичные, инклюзивные и благотворительные товары.\n \n Товар произведен в стране присутствия Lamoda, что позволяет нам оптимизировать выбросы СО2 при доставке. Покупая этот товар, вы вносите свой вклад в сокращение углеродного следа и поддерживаете развитие локальных фабрик.	Платье	6800	marichuel.jpeg
\.


--
-- TOC entry 3632 (class 0 OID 0)
-- Dependencies: 217
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.customers_id_seq', 5, true);


--
-- TOC entry 3633 (class 0 OID 0)
-- Dependencies: 219
-- Name: customers_products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.customers_products_id_seq', 21, true);


--
-- TOC entry 3634 (class 0 OID 0)
-- Dependencies: 218
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.products_id_seq', 13, true);


--
-- TOC entry 3472 (class 2606 OID 16428)
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- TOC entry 3476 (class 2606 OID 16438)
-- Name: customers_products id; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.customers_products
    ADD CONSTRAINT id PRIMARY KEY (id);


--
-- TOC entry 3474 (class 2606 OID 16433)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 3477 (class 2606 OID 16439)
-- Name: customers_products customer_id; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.customers_products
    ADD CONSTRAINT customer_id FOREIGN KEY (customer_id) REFERENCES public.customers(id);


--
-- TOC entry 3478 (class 2606 OID 16444)
-- Name: customers_products product_id; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.customers_products
    ADD CONSTRAINT product_id FOREIGN KEY (product_id) REFERENCES public.products(id);


-- Completed on 2023-12-11 11:11:43 +03

--
-- PostgreSQL database dump complete
--

