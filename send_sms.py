#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
import json
import requests
import sys


def send_sms(recipient, message, url, username, password):
    """
    Отправляет SMS-сообщение через API
    
    Args:
        recipient: номер телефона получателя
        message: текст сообщения
        url: URL API для отправки SMS
        username: имя пользователя для авторизации
        password: пароль для авторизации
    
    Returns:
        bool: True в случае успеха, False в случае ошибки
    """
    # Формируем данные запроса
    sms_data = {
        "recipient": recipient,
        "message": message
    }
    
    # Отправляем запрос с Basic Auth
    try:
        response = requests.post(
            url, 
            json=sms_data,
            auth=(username, password),
            headers={"Content-Type": "application/json"}
        )
        
        # Проверяем статус ответа
        if response.status_code == 200:
            print("SMS успешно отправлено!")
            print(f"Ответ сервера: {response.text}")
            return True
        else:
            print(f"Ошибка при отправке SMS. Код: {response.status_code}, Ответ: {response.text}")
            return False
            
    except Exception as e:
        print(f"Ошибка при отправке запроса: {e}")
        return False


def main():
    # Создаем парсер аргументов командной строки
    parser = argparse.ArgumentParser(description="Отправка SMS через API")
    
    # Добавляем параметры
    parser.add_argument("-recipient", required=True, 
                        help="Номер телефона получателя в формате +XXXXXXXXXXX")
    parser.add_argument("-message", required=True, 
                        help="Текст сообщения")
    parser.add_argument("-url", default="http://192.168.1.1/api/sendsms",
                        help="URL API сервера для отправки SMS")
    parser.add_argument("-user", default="root",
                        help="Имя пользователя для авторизации")
    parser.add_argument("-pass", dest="password", default="root",
                        help="Пароль для авторизации")
    
    # Парсим аргументы
    args = parser.parse_args()
    
    # Отправляем SMS
    success = send_sms(
        args.recipient, 
        args.message, 
        args.url, 
        args.user, 
        args.password
    )
    
    # Выходим с соответствующим кодом
    sys.exit(0 if success else 1)


if __name__ == "__main__":
    main() 