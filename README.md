# Cartola Coltec
An app with idea of cartola (https://globoesporte.globo.com/cartola-fc/) but targeting colleges and high school.

# Backend 
- 2 roles
    - Geral Administrator
    - Users

# Geral Admin
It will be the master admin that verify the teams admin and so on!

## Actions
- Add the Teams signed up
- Update the match day/time
- Add the description of the game
    - All Players
        - Goals (plus)
        - Shoots
            - Defended Shoot (plus)
            - Missed Shoot (minus)
            - On the Bar (plus)
        - Assists (plus)
        - Skill Moves (plus)
        - Tackles (plus)
        - Missed Passes (minus)
        - Fouls
            - Suffered (plus)
            - committed (minus)
        - Cards 
            - Yellow Cards (minus)
            - Red Cards (minus)
        - Missed Shoot-Out (minus)
    - Defenders
        - Game without goal conceded (plus)
    - Goalkeepers
        - Goalkeepers Defenses (plus)
        - Shoot-Out Defended (plus)
        - Goals Conceded (plus)
# User

## Actions
- Choose 7 players
    - 2 at the bench
        - whatever that user wants
    - 5 lineup players
        - 1 goalkeeper and 4 players
- Choose formations
- Create private leagues
    - Send invitation
    - Receive invitation
- To participate the league

# Structs

## Formations
- 4 defenders / 0 attackers
- 3 defenders / 1 attackers
- 2 defenders / 2 attackers
- 1 defenders / 3 attackers

## Positions
- Goalkeeper
- Defender
- Attacker

## Status
- Probably
- Injured
- Doubt
- Suspended
- Nothing

## Shop

### Shop All Informations
- Array
    - Players
        - Name
        - nickname searched (lowercase)
        - nickname
        - photo (facebook API)
        - player id
        - rodada id
        - team id
        - status id
        - score num
        - price num
        - median
        - matches_num
        - Scout Stats (Struct)

### Shop Stats
- Current Round
- Shop Status
- Cup/Year
- Ready Teams
- Closing
    - Day
    - Month
    - Year
    - Hour
    - Minute
    - Timestamp

### Shop Highlights
- Array
    - Player
        - Id
        - Name
        - Nicname
        - Photo
    - number of teams with
    - team 
    - club shield
    - position
    - position abbreviation

## Calendar
- Array
    - Round
        - Id
        - Start (Date)
        - End (Date)

## Matches
- Array
    - Match
        - Id
        - team1_id
        - Team2-id
        - Team1 Stats
            - Array
                - Victory
                - Draw
                - Losts
        - Team2 Stats
            - Array
                - Victory
                - Draw
                - Losts
        - Date
        - Team1 Scoreboard
        - Team2 Scoreboard

## Round Matches
- Array
    - Matches

## Teams
- Array
    - Id (Key)
        - Id
        - Name
        - Abbreviation
        - Shield img

## Highlight Team
- Median Money 
- Median Points
- Best User
    - User Team Id
    - User Team Id
    - Formation Id
    - Facebook Id
    - Facebook Photo
    - User Team Name
    - Username
    - User Team name searched (lowercase)
