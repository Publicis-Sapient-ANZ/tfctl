#!/bin/bash
export GPG_TTY=$(tty)
export PATH=$HOME/.local/bin:/usr/local/go/bin:$PATH
export SSL_CERT_FILE='/etc/ssl/certs/ca-certificates.crt'
export REQUESTS_CA_BUNDLE='/etc/ssl/certs/ca-certificates.crt'

alias ll='ls -alF'
alias gs='git status'
alias ga='git add .'
alias gc='git commit'
alias gp='git push'
alias ggg='git add . && git commit && git push'