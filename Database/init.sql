CREATE TABLE JUGADORES(
    nombre text NOT NULL,
    contra text NOT NULL,
 -- perfil bytea NOT NULL,
    descrp text NOT NULL,
    codigo text,
    PRIMARY KEY (codigo)
);

CREATE TABLE AMISTAD(
    estado integer NOT NULL,
    usr1 text REFERENCES JUGADORES (codigo),
    usr2 text REFERENCES JUGADORES (codigo),
    PRIMARY KEY (usr1, usr2)
);

CREATE TABLE CARTAS(
    numero integer,
    palo text,
 -- foto bytea NOT NULL,
    PRIMARY KEY (numero, palo)
);



