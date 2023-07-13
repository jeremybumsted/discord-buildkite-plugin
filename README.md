# Discord Notification Plugin

A Buildkite plugin written in Go to send job and build notifications formatted for [Discord](https://discord.com/) using Discord [Webhooks](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks).

This plugin works similarly to the `notify` [attribute](https://buildkite.com/docs/pipelines/notifications#slack-channel-and-direct-messages) in a command step.

You can create a webhook integration in your Discord guild following the steps [here](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks).

The plugin has been tested and built using **go 1.20.3**.

###⚠️ **Note**: this plugin requires the `polyglot-hooks` [experiment](https://github.com/buildkite/agent/blob/main/EXPERIMENTS.md#polyglot-hooks) to be enabled on your agents.

## Usage

Add the following to your `pipeline.yml`:

```yaml
    steps:
        command: ./build.sh
        plugins:
            discord-notifications#v0.1.0:
                message: "From Buildkite with Love"
                webhook-url: DISCORD_WEBHOOK_URL
```

The plugin builds a `pre-exit` hook in the /hooks directory, and when executed, will check the exit status of the job's command hook, and send a formatted message to the configured Discord Webhook URL based on the result.


## Examples
Below are a list of exmaples that you can use to implement the plugin in your pipelines.

### Example 


## Options

### `message` (optional, string)
The message you want to send in the notification

### `webhook-env` (optional, string)
The environment variable that contains the value of the Discord Webhook URL. Default: `DISCORD_WEBHOOK_URL`.
It is not recommended to pass secure


## Development
### Running the tests
The tests are written using Go's built-in testing package, as well as using packages from [stretchr/testify](https://github.com/stretchr/testify).

Tests can be run from the command line:

```shell
go mod tidy

go test -v ./...
```


## Contributing

We welcome all contributions to improve the plugin! To contribute, please follow these guidelines:

- Fork the repository and create a new branch.
- Make your changes and ensure that the tests pass.
- Write clear and concise commit messages.
- Submit a pull request.

By contributing, you agree to license your contributions under the LICENSE file of this repository.

## License
MIT (see [LICENSE](LICENSE.MD))
