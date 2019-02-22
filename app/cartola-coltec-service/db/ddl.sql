CREATE TABLE formation (
    id          SERIAL,
    goalkeeper  SMALLINT,
    defenders   SMALLINT,
    attackers   SMALLINT  
);

CREATE TABLE scout (
    id              SERIAL,
    description     VARCHAR(80),
    points          SMALLINT,
);

CREATE TABLE player_stats (
    id              SERIAL,
    round_id        SMALLINT,
    player_id       SMALLINT,
    scout_id        SMALLINT,
    amount_scout    SMALLINT
);

CREATE TABLE status (
    id              SERIAL,
    description     VARCHAR(80)
);

CREATE TABLE team (
    id          SERIAL,
    name        VARCHAR(80),
    nickname    VARCHAR(255),
    image_url   VARCHAR(255)
);

CREATE TABLE team_stats (
    id              SERIAL,
    id_team         SMALLINT,
    victory         SMALLINT,
    lose            SMALLINT,
    draw            SMALLINT,
    goal_against    SMALLINT,
    goal_difference SMALLINT,
    goal_pro        SMALLINT
);

