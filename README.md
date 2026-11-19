# mautrix-go-fork

**English** | [Русский](README.ru.md)

A fork of [mautrix-go](https://github.com/mautrix/go), the Matrix framework that mautrix bridges are built on,
with the changes [telegram-bridge-advanced](https://github.com/kryptonmourog/telegram-bridge-advanced) needs.

## Why it exists

Some features of the bridge need changes in bridgev2 itself, not only in the bridge. They live here, kept small
and close to upstream, so the bridge can be built against them until (and unless) upstream has an equivalent.
You don't need this repository to use the bridge: its images already contain it.

## What it changes (bridgev2)

- A translation hook (`bridgev2.Translate`) for user-facing text: login prompts, notices, tombstones.
- Encryption of personal (private chat) portals only, and never of rooms the bridge can't decrypt.
- Direct media served alongside the websocket appservice, with a configurable server identity.
- Network connectors can wait for the homeserver's media config and choose the initial backfill limit per
  portal.
- A ghost leaves the room once its user is shown there by their own account (double puppeting).

## Versions

The `i18n` branch is based on an upstream release. Releases are tagged `v<mautrix-go version>-adv.<N>`
(e.g. `v0.33.0-adv.1`) and used by the bridge with a `replace` in its `go.mod`:

```
replace maunium.net/go/mautrix => github.com/kryptonmourog/mautrix-go-fork v0.33.0-adv.1
```

## The project

| Repository | What it is |
|---|---|
| [telegram-bridge-advanced](https://github.com/kryptonmourog/telegram-bridge-advanced) | the bridge itself: code, images, installation guide, releases |
| **mautrix-go-fork** (you are here) | the fork of mautrix-go the bridge is built with |
| [wsproxy-fork](https://github.com/kryptonmourog/wsproxy-fork) | the websocket proxy for a bridge on a separate server, with encryption |
| [telegram-bridge-advanced-ansible](https://github.com/kryptonmourog/telegram-bridge-advanced-ansible) | installation with matrix-docker-ansible-deploy, including on a separate server |

## License

MPL-2.0, like mautrix-go.
