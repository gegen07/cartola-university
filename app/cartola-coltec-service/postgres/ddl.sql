DROP TABLE IF EXISTS formation CASCADE;
CREATE TABLE formation (
    id          SERIAL PRIMARY KEY,
    goalkeeper  SMALLINT NOT NULL,
    defenders   SMALLINT NOT NULL,
    attackers   SMALLINT NOT NULL
);

DROP TABLE IF EXISTS round CASCADE;
CREATE TABLE round (
    id                  SERIAL PRIMARY KEY,
    round_begin_date    TIMESTAMP NOT NULL,
    round_finish_date   TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS scout CASCADE;
CREATE TABLE scout (
    id              SERIAL PRIMARY KEY,
    description     VARCHAR(80) NOT NULL,
    points          SMALLINT NOT NULL
);

DROP TABLE IF EXISTS team CASCADE;
CREATE TABLE team (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    nickname    VARCHAR(255) NOT NULL,
    image_url   VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS team_stats CASCADE;
CREATE TABLE team_stats (
    id              SERIAL PRIMARY KEY,
    id_team         INT REFERENCES team(id) ON DELETE CASCADE,
    victory         SMALLINT NOT NULL,
    lose            SMALLINT NOT NULL,
    draw            SMALLINT NOT NULL,
    goal_against    SMALLINT NOT NULL,
    goal_difference SMALLINT NOT NULL,
    goal_pro        SMALLINT NOT NULL
);

DROP TABLE IF EXISTS player CASCADE;
CREATE TABLE player (
    id          SERIAL PRIMARY KEY,
    id_team     INT REFERENCES team(id) ON DELETE CASCADE,
    name        VARCHAR(255) NOT NULL,
    nickname    VARCHAR(255) NOT NULL,
    photo       VARCHAR(255) NOT NULL,
    price       DECIMAL NOT NULL,
    score       DECIMAL NOT NULL,
    median      DECIMAL NOT NULL,
    num_matches SMALLINT NOT NULL,
    position    VARCHAR(255) NOT NULL,
    status      VARCHAR(255) NOT NULL

);

DROP TABLE IF EXISTS player_stats CASCADE;
CREATE TABLE player_stats (
    id              SERIAL PRIMARY KEY,
    round_id        INT REFERENCES round(id) ON DELETE CASCADE,
    player_id       INT REFERENCES player(id) ON DELETE CASCADE,
    scout_id        INT REFERENCES scout(id) ON DELETE CASCADE,
    amount_scout    SMALLINT  NOT NULL
);

DROP TABLE IF EXISTS match CASCADE;
CREATE TABLE match (
    id                      SERIAL PRIMARY KEY,
    id_home_team            INT REFERENCES team(id) ON DELETE CASCADE,
    id_visitor_team         INT REFERENCES team(id) ON DELETE CASCADE,
    scoreboard_home_team    SMALLINT NOT NULL,
    scoreboard_visitor_team SMALLINT NOT NULL,
    match_date              TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS round_matches CASCADE;
CREATE TABLE round_matches (
    id          SERIAL PRIMARY KEY,
    id_round    INT REFERENCES round(id) ON DELETE CASCADE,
    id_match    INT REFERENCES match(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS cup CASCADE;
CREATE TABLE cup (
    id  SERIAL PRIMARY KEY,
    id_winner_team     INT REFERENCES team(id) ON DELETE CASCADE,
    id_second_team     INT REFERENCES team(id) ON DELETE CASCADE,
    id_third_team      INT REFERENCES team(id) ON DELETE CASCADE,
    year               INT NOT NULL,
    begin_date         TIMESTAMP NOT NULL,
    finish_date        TIMESTAMP NOT NULL,
    name_cup           VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS shop CASCADE;
CREATE TABLE shop (
    id          SERIAL PRIMARY KEY,
    id_round    INT REFERENCES round(id) ON DELETE CASCADE,
    id_cup      INT REFERENCES cup(id) ON DELETE CASCADE,
    status      VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS players_shop CASCADE;
CREATE TABLE players_shop (
    id          SERIAL PRIMARY KEY,
    id_player   INT REFERENCES player(id) ON DELETE CASCADE,
    id_shop     INT REFERENCES shop(id) ON DELETE CASCADE
);
