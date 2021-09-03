# SimpleCA

Have you ever been working with a technology and needed TLS certificates quickly? 

Perhaps you wanted to set up a PKI infrastructure for testing, but were either intimidated by `openssl`, or wished 
there was a {quicker, easier} way.

SimpleCA is here to solve that need. 

## Purpose

SimpleCA is meant to be used when you need quick access to a certificate authority of your own generation. 
With a few simple commands, a (simple) CA is initialized on your workstation and you can begin
issuing certificates. 

## Quickstart

Initialize the default CA

```
$ simpleca ca init
```

Issue your first certificate

```
$ simpleca cert sign my-first-cert --dns localhost --ip 127.0.0.1
```

View that certificate

```
$ simpleca cert view my-first-cert
```

## Advanced

When issuing certificates (either CA or server), support for the following arguments is available:

```
Flags:
      --country string               2-letter country code (default "AA")
      --dns strings                  dns subject alternative name
      --email strings                email subject alternative name
      --expire-in string             duration of cert validity (default "1 year")
      --ip ipSlice                   ip address subject alternative name (default [])
      --locality string              locality
      --organization string          organization name (default "SimpleCA Ltd.")
      --organizational-unit string   organizational unit (default "SimpleCA Security")
      --passphrase string            passphrase for generated certificate (default "changeme")
      --state string                 state or province name (default "Relaxation")
      --uri strings                  uri subject alternative name
```

#### A note about expiration:

When signing a certificate, you can specify `--expire-in`. This uses a very simple syntax:

`[number] [period]`

For example, you can specify things like `1 year` or `10 months` or `15 days`.

## Contribution

File issues, PRs welcome, etc.