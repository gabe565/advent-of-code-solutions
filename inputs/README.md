# Inputs

This directory contains inputs used for automated CI tests. Since Advent of Code [does not allow inputs to be committed publicly](https://adventofcode.com/about#faq_copying), they are encrypted using [SOPS](https://github.com/getsops/sops).

# Usage

Helper scripts are provided to assist with encrypting and decrypting inputs:
- [`encrypt-inputs.sh`](../hack/encrypt-inputs.sh)
- [`decrypt-inputs.sh`](../hack/decrypt-inputs.sh)
 
> [!NOTE]  
> You will not be able to decrypt these inputs because the SOPS AGE key is private. These scripts are used solely for me to add inputs and decrypt them during CI.
