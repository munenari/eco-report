
# eco-report

this is a log-agent for eco-report toward elasticsearch (for kibana)

## how to run

### docker

```bash
$ vi config.toml
$ docker-compose up -d
```

### plain binary

edit `config.toml` and build binary.
then setting up your machine for running agent.

- setup binary and its config

```
/path/to/eco-report
├── config.toml
└── eco-report-agent (built by go)
```

- enable and start agent unit

```bash
$ systemctl enable /path/to/some/eco-report-agent.service
$ systemctl start /path/to/some/eco-report-agent.service
```
