# whichnumber
**whichnumber** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite-hq/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/zale144/whichnumber@latest! | sudo bash
```
`zale144/whichnumber` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Use

### Tx

#### Update Params

     whichnumberd tx whichnumber update-params 600 600 11 121 --from bob

#### Create a new Game

     whichnumberd tx whichnumber new-game 42 11stake 121stake --from bob

#### Commit a number

     whichnumberd tx whichnumber commit-number 1 38 --from bob

#### Reveal a number
    
     whichnumberd tx whichnumber reveal-number 1 38 [salt] --from bob

### Query

#### List all games

    whichnumberd query whichnumber list-games

#### Show a game
    
    whichnumberd query whichnumber show-game 1

#### Show params

    whichnumberd query whichnumber params

#### Show system info
    
    whichnumberd query whichnumber show-system-info

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
