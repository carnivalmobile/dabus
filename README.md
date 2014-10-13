# Dabus

Dabus is a systemd dbus notification observer written in Go, bundled with a Hipchat notifier.

If your systemd services stop, start or restart, you can receive a Hipchat notification.

This is useful for monitoring services on CoreOS machines.

It is configured through a YAML file, and you can run it like this:

`dabus config.yaml`

A [sample configuration file is available](sample_config.yaml) as is a [sample services file](dabus.service).

## Project Setup

To get started with dabus, run `go get github.com/carnivalmobile/dabus`

Dabus depends on *systemd* being present on your system.

If systemd is not installed, you can run an appropriate Linux in a virtual machine to test it out.

## Testing

_How do I run the project's automated tests?_

1. `go test`

## Deploying

Grab the [latest binary from the available releases](https://github.com/carnivalmobile/dabus/releases).

## Troubleshooting

> - `dial unix /var/run/dbus/system_bus_socket: no such file or directory`

Check that your operating system has systemd configured correctly and that the socket is present.

## Contributing changes

Once you've made your great commits:

1. Fork Dabus
2. Create a topic branch - git checkout -b my_branch
3. Push to your branch - git push origin my_branch
4. Open a Pull Request to discuss your changes

That's it!

## License

See [LICENSE](LICENSE)
