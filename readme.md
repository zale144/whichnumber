# Whichnumber
**Whichnumber** is a simple game to guess a number. It is a simple example of a Cosmos SDK blockchain application. 
The flow of the game is the following:

- A game is first created. The creator sets the entry fee and the reward. The reward has to be no less than the minimum as defined in the parameters. The reward amount is transferred from the creator's account to the game's account.
- Then during the commit phase, players can commit a number. The number commit is a hash of the guessed number and a salt. 
- If the number of players reaches the maximum number of players, the game will move to the reveal phase, even if the commit phase timeout was not reached.
- The entry fee is transferred from the player's account to the game's account.
- Each player can only commit once, and the creator cannot participate.
- In the reveal phase, each player reveals the number they guessed by providing the guessed number and salt. 
- If all the players reveal their numbers, the game will move to the end phase, even if the reveal phase timeout was not reached.
- If the number is within the minimum allowed range of the secret number, the player wins.
- There can be multiple winners. In that case the prize is split among the winners, proportionally to their proximity to the secret number.

## Features

- Commit and reveal scheme
- A CLI to interact with the chain
- A Makefile to facilitate the interaction with the chain
- Unit tests for the game logic
- Manually consuming gas where there are potential for DoS attacks
- Emitted relevant events for the game lifecycle, facilitating indexing 

## Things to improve

- More unit tests, and add integration tests
- Access control for updating the game parameters
- Find more ways to defend against malicious actors
- There are probably more edge cases to consider, and likely more bugs to find

## Get started

```
ignite chain serve
```

## Use

### Tx

#### Create a new Game

     make new-game

#### Commit a number

     make commit-number

#### Reveal a number
    
    make reveal-number

#### Update Params

     make update-params

### Query

#### List all games

    make list-games

#### Show a game
    
    make show-game

#### Show params

    make show-params

#### Show system info
    
    make show-system
