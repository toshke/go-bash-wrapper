## Go bash wrapper

- Forwards CLI arguments to `/bin/bash`
- Fails if there is `GO_SH_WRAPPER` variable defined as recursion
  control variable

## But, why?

Amazon ssm-agent by default runs with 'sh' command and such does
not load full bash shell, and does not process user home-directory
startup files like `~/.bashrc`

Replace standard `/bin/sh` with `sh` file produce by running `make build`

[Relevant discussion](https://github.com/aws/amazon-ssm-agent/pull/171)
