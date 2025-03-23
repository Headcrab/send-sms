# SMS отправка

[![Go Version](https://img.shields.io/github/go-mod/go-version/Headcrab/send-sms)](https://go.dev)
[![License](https://img.shields.io/github/license/Headcrab/send-sms)](LICENSE)

Простое решение для отправки SMS сообщений через API. Включает две реализации:

- Go-версия
- Python-версия

## Установка и зависимости

### Go-версия

```bash
# Клонировать репозиторий
git clone [URL репозитория]
cd send-sms3

# Собрать приложение
go build
```

### Python-версия

```bash
# Установить зависимости
pip install requests
```

## Использование

### Go-версия

```bash
# Базовое использование
./send-sms3 -recipient "+7xxxxxxxxxx" -message "Тестовое сообщение"

# Полный набор параметров
./send-sms3 -recipient "+7xxxxxxxxxx" -message "Тестовое сообщение" -url "http://192.168.1.1/api/sendsms" -user "root" -pass "root"
```

### Python-версия

```bash
# Базовое использование
python send_sms.py -recipient "+7xxxxxxxxxx" -message "Тестовое сообщение"

# Полный набор параметров
python send_sms.py -recipient "+7xxxxxxxxxx" -message "Тестовое сообщение" -url "http://192.168.1.1/api/sendsms" -user "root" -pass "root"
```

## Параметры

Обе версии используют одинаковые параметры:

- `-recipient` - Номер телефона получателя в формате `+XXXXXXXXXXX` (обязательный)
- `-message` - Текст сообщения (обязательный)
- `-url` - URL API для отправки SMS (по умолчанию "http://192.168.1.1/api/sendsms")
- `-user` - Имя пользователя для Basic Auth (по умолчанию "root")
- `-pass` - Пароль для Basic Auth (по умолчанию "root")

## Примеры

### Go-версия

```bash
# Отправить SMS с переносом строки в сообщении
./send-sms3 -recipient "+7xxxxxxxxxx" -message "Первая строка\nВторая строка"

# Отправить SMS с другими учетными данными
./send-sms3 -recipient "+7xxxxxxxxxx" -message "Тест" -user "admin" -pass "password"
```

### Python-версия

```bash
# Отправить SMS с переносом строки в сообщении
python send_sms.py -recipient "+7xxxxxxxxxx" -message "Первая строка\nВторая строка"

# Отправить SMS с другими учетными данными
python send_sms.py -recipient "+7xxxxxxxxxx" -message "Тест" -user "admin" -pass "password"
```

## Аналог curl команды

Обе версии эквивалентны следующей curl команде:

```bash
curl -X POST http://root:root@192.168.1.1/api/sendsms -d '{"recipient": "+7xxxxxxxxxx", "message": "test\n"}'
```

## Лицензия

MIT License. См. файл [LICENSE](LICENSE) для подробностей.

## Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для ваших изменений
3. Внесите изменения и создайте pull request

## Контакты

Создайте issue в репозитории для сообщения о проблемах или предложений по улучшению.

## Спасибо

- [@Headcrab](https://github.com/Headcrab)
