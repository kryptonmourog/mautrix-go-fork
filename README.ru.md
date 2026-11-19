# mautrix-go-fork

**Русский** | [English](README.md)

Форк [mautrix-go](https://github.com/mautrix/go) — фреймворка Matrix, на котором построены мосты mautrix, — с
изменениями, которые нужны [telegram-bridge-advanced](https://github.com/kryptonmourog/telegram-bridge-advanced).

## Зачем он создан

Часть функций моста требует правок в самом bridgev2, а не только в мосте. Они собраны здесь, небольшие и близкие
к upstream, чтобы мост собирался с ними, пока в upstream нет такого же (если появится). Чтобы пользоваться мостом,
этот репозиторий не нужен: он уже внутри образов моста.

## Что он меняет (bridgev2)

- Хук перевода (`bridgev2.Translate`) для текстов, которые видит пользователь: вход, уведомления, переадресация
  комнат.
- Шифрование только личных порталов и никогда — комнат, которые мост не сможет расшифровать.
- Прямые медиа вместе с websocket-подключением к серверу, с настраиваемым именем сервера медиа.
- Мосты могут дождаться настроек медиа сервера и выбирать глубину первичной истории для каждого портала.
- Марионетка выходит из комнаты, когда её пользователь показан там своим аккаунтом (двойной вход).

## Версии

Ветка `i18n` построена на выпуске upstream. Выпуски — теги `v<версия mautrix-go>-adv.<N>` (например
`v0.33.0-adv.1`); мост подключает их строкой `replace` в `go.mod`:

```
replace maunium.net/go/mautrix => github.com/kryptonmourog/mautrix-go-fork v0.33.0-adv.1
```

## Проект

| Репозиторий | Что это |
|---|---|
| [telegram-bridge-advanced](https://github.com/kryptonmourog/telegram-bridge-advanced) | мост: код, образы, инструкция по установке, выпуски |
| **mautrix-go-fork** (вы здесь) | форк mautrix-go, с которым собирается мост |
| [wsproxy-fork](https://github.com/kryptonmourog/wsproxy-fork) | websocket-прокси для моста на отдельном сервере, с шифрованием |
| [telegram-bridge-advanced-ansible](https://github.com/kryptonmourog/telegram-bridge-advanced-ansible) | установка через matrix-docker-ansible-deploy, в том числе на отдельный сервер |

## Лицензия

MPL-2.0, как у mautrix-go.
