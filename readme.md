# Whichnumber
**Whichnumber** is a simple game to guess a number. It is a simple example of a Cosmos SDK blockchain application. 
The flow of the game is the following:

- A game is first created. The creator sets the entry fee and the reward. The reward amount is transferred from the creator's account to the game's account.
- Then during the commit phase, players can commit a number. The number commit is a hash of the guessed number and a salt. The entry fee is transferred from the player's account to the game's account.
- Each player can only commit once, and the creator cannot participate.
- In the reveal phase, each player reveals the number they guessed by providing the guessed number and salt. 
- If the number is within the minimum allowed range of the secret number, the player wins.
- There can be multiple winners. In that case the prize is split among the winners, proportionally to their proximity to the correct number.

## Features

- Commit and reveal scheme
- A CLI to interact with the chain
- A Makefile to facilitate the interaction with the chain
- Unit tests for the game logic
- Manually consuming gas where there are potential for DoS attacks
- Emitted relevant events for the game lifecycle, facilitating indexing 

## Things to improve

- More unit tests, and add integration tests
- Better logging and error handling
- Make the code easier to test by breaking down larger functions
- Find more ways to defend against malicious actors
- There are probably more edge cases to consider, and likely more bugs to find

## Get started

```
ignite chain serve
```

## Use

### Tx

#### Update Params

     make update-params

#### Create a new Game

     make new-game

#### Commit a number

     make commit-number

#### Reveal a number
    
    make reveal-number

### Query

#### List all games

    make list-games

#### Show a game
    
    make show-game

#### Show params

    make show-params

#### Show system info
    
    make show-system
