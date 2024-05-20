# SUI ez wallet

## This SUI wallet tool has two functions:

1. Generate wallets with specified prefix or suffix address.
2. Generate wallets based on a given word.

## build

> go build -o ez-wallet main.go

## Usage

```bash
$ ./ez-wallet
NAME:
   ez-wallet - A new cli application

USAGE:
   ez-wallet [global options] command [command options]

COMMANDS:
   gen, g          gen [flag]
   create-word, c  create-word <key-words>
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

1. To generate a wallet with a specified prefix or suffix address.

```bash
$ ./ez-wallet g --prefix 66 --suffix 88
wallet address: 0x66df32c6264c436da8101551dac867bba818d0ff53d1734462cab462f1837788
private key: suiprivkey1qz**************usfamskp6997e

$ ./ez-wallet g --prefix 66
wallet address: 0x665380a3147fa8d4a3ad54daa1a6a872d5daffe929e507b7a76c36ffc8a704f8
private key: suiprivkey1qqv0fs********a4kqkawqm9svk0w6jt

$ ./ez-wallet g --suffix 88
wallet address: 0x446aa09d427e6faec955d74663c5a727b88d7ca466ffcb3187e2dafb12adea88
private key: suiprivkey1qr5xeph5rpjr***********87yrd5y4m

```

* This command is essentially generating a million wallets, selecting addresses with specified prefixes or suffixes.

```bash
$ ./ez-wallet g --prefix foo --suffix boo
Your request is quite challenging; please try again, or change the prefix or suffix.
```

2. generate a wallet based on keywords, you only need to remember the keywords to generate the wallet again without having to remember complex private keys or mnemonic phrases, making it easy to manage wallets, especially when dealing with a large number of wallets.
**Warning**: If the keyword is leaked, it means the wallet is compromised. So, try to make the keywords as long and complex as possible, just like a password.

```bash
$ ./ez-wallet c myname
create task:
wallet address: 0x9b8ec12f93f25dfa23df39b5011056c917789346d98d0d038e603541bfae6de7
private key: suiprivkey1qqcphnx7efm5f5spj9p9yjxwxfnequavsrc3gj8qrnmk43kcdmz7w85e8lx

$ ./ez-wallet c myname
create task:
wallet address: 0x9b8ec12f93f25dfa23df39b5011056c917789346d98d0d038e603541bfae6de7
private key: suiprivkey1qqcphnx7efm5f5spj9p9yjxwxfnequavsrc3gj8qrnmk43kcdmz7w85e8lx

$ ./ez-wallet c wallet-1
create task:
wallet address: 0xc8dbea3e25e83b45ebd4d58852c912334fbed7210becdf96aeba07bf601384d7
private key: suiprivkey1qqwn6cd6u8sr76xa8027ej6sz3m2ls8mqavv0ah9w6wjgsrryrskzy7mwkw

$ ./ez-wallet c wallet-2
create task:
wallet address: 0x0f260ca4985048474d459abde1d9b3f63a345a9b2d3053a578ae536fb2eded15
private key: suiprivkey1qzr4wxx0283tu0qw02lqzarnsfcjwp9t0mjchtza9jwm8wnpnsrgkxve97j
```

* The essence of this command is to use the hash of the keyword as the seed for the private key.
* This command can be slightly modified to manage a large number of addresses without the need to manage private keys or other difficult-to-remember mnemonic words or seeds.For instance, if I need 10 thousand addresses, I just need to remember the keyword plus the numbers, like xxx-wallet-1-10000.

