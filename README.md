# Практическое задание № 4 Борисов Денис Александрович ЭФМО-01-25
Тема: Настройка Prometheus + Grafana для метрик. Интеграция с приложением

Задачи практики:
1.	Изучить назначение метрик и их роль в сопровождении серверных приложений.
2.	Понять различие между логированием и мониторингом на основе метрик.
3.	Освоить подключение Prometheus-клиента к Go-приложению.
4.	Научиться публиковать метрики приложения через HTTP-эндпоинт /metrics.
5.	Настроить Prometheus для регулярного сбора метрик приложения.
6.	Подключить Prometheus как источник данных в Grafana.
7.	Построить базовые панели для визуализации количества запросов, ошибок и времени обработки.
8.	Отработать анализ состояния приложения по метрикам.


Выполнение практического задания.

1.	Структура проекта

<img width="293" height="491" alt="3" src="https://github.com/user-attachments/assets/44f5fcaa-3a3f-44b0-869b-bf9b084a7a8d" />

2.	Установка зависимостей Go и инструментов генерации.

Установка зависимостей

<img width="750" height="397" alt="1" src="https://github.com/user-attachments/assets/3c9a796d-716d-46c5-a845-7bde419f1bce" />

3.	Создание файла docker-compose.yml.

docker-compose.yml, предназначен для развертывания БД Postgres 

<img width="513" height="352" alt="2" src="https://github.com/user-attachments/assets/cc2d8bc7-af61-40db-a5e0-0fbda9c7f772" />

Развертывание сервисов

<img width="867" height="161" alt="3" src="https://github.com/user-attachments/assets/fb261420-3faa-4514-938b-12f31d0c0d0d" />

Рабочие контейнеры в Docker

<img width="1567" height="296" alt="4" src="https://github.com/user-attachments/assets/f93cb82a-a6f6-4c48-b9ee-4b6ff191a7be" />

Создание и заполнение БД

<img width="887" height="293" alt="5" src="https://github.com/user-attachments/assets/2f9454ed-bc20-4a32-9ee8-49e84daf3b87" />

4. Создание SSH ключа

<img width="1462" height="451" alt="6" src="https://github.com/user-attachments/assets/503929da-9919-43a1-98ae-a35445634884" />

Созданные ключи

![Uploading 7.png…]()

5. Тестирование

Запуск сервера

<img width="723" height="93" alt="9" src="https://github.com/user-attachments/assets/c3f8dbf0-d5bd-48bb-b93f-7b8cc9b3f675" />

Выполнение тестов

1. Вывод первого студенты

<img width="957" height="552" alt="11" src="https://github.com/user-attachments/assets/9bb4b668-d790-4923-b83d-993f502b488c" />

2. Ошибка несуществующий студент

<img width="953" height="607" alt="12" src="https://github.com/user-attachments/assets/275ee7e6-1231-421f-9173-eec27f4c3284" />

3. Попытка иньекции

<img width="952" height="592" alt="13" src="https://github.com/user-attachments/assets/67e3fa63-ef11-4353-a5d8-befa02a219a3" />

4. Поиск по почте студента

<img width="952" height="687" alt="14" src="https://github.com/user-attachments/assets/00706056-5af1-48cb-abaa-5326df8dd2b5" />

5. Переадресация

<img width="961" height="650" alt="15" src="https://github.com/user-attachments/assets/d5274f94-6201-47d3-91ee-75a97af21dab" />

