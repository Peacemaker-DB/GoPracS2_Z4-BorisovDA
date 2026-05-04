<img width="562" height="458" alt="5" src="https://github.com/user-attachments/assets/b78faaef-6f45-4bd4-86e1-b420b3b816dc" /># Практическое задание № 4 Борисов Денис Александрович ЭФМО-01-25
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

docker-compose.yml, предназначен для развертывания Prometheus и	Grafana 

<img width="562" height="458" alt="5" src="https://github.com/user-attachments/assets/5a4b0e07-5b57-4128-ab08-6388f9fa3e9f" />

Развертывание сервисов

<img width="716" height="326" alt="6" src="https://github.com/user-attachments/assets/c935e97a-fa93-41c1-ba51-d77ea16e77fd" />

Рабочие контейнеры в Docker

<img width="1542" height="383" alt="7" src="https://github.com/user-attachments/assets/0c85b657-865c-4d91-b2a9-b56fc19971a2" />

Открытые сервиси

<img width="1918" height="973" alt="14" src="https://github.com/user-attachments/assets/d031ffa8-a972-4745-b464-e7e68ce7f5b6" />

4. Тестирование

Запуск сервера

<img width="502" height="52" alt="8" src="https://github.com/user-attachments/assets/f1037f6d-41af-42c6-a4d6-7acd9f424b57" />

Выполнение тестов

1. Вывод первого студенты

<img width="957" height="682" alt="10" src="https://github.com/user-attachments/assets/2ecd5be5-fc1d-4033-b03e-c14972d88de1" />

2. Вывод второго студенты

<img width="956" height="667" alt="11" src="https://github.com/user-attachments/assets/5bb9035b-cd8b-4b4e-95ae-afe2062eabd0" />

3. Ошибка несуществующий студент

<img width="955" height="597" alt="12" src="https://github.com/user-attachments/assets/39510ad3-e184-48cf-aa6f-7df739ce6cf0" />

4. Сборка метрик

<img width="951" height="837" alt="13" src="https://github.com/user-attachments/assets/ad59e4e4-a21c-42b1-9739-54c06c557e20" />

Собранные метрики в Prometheus

<img width="947" height="767" alt="15" src="https://github.com/user-attachments/assets/53c71bb7-1f55-4693-b020-c733af10bd7a" />

5. Вывод граффиков метрик

А)Активные запросы

<img width="931" height="846" alt="Активные запросы" src="https://github.com/user-attachments/assets/934b0faf-e6df-4849-81ba-bb6c73affd47" />

Б)Запросы по студентам

<img width="932" height="851" alt="Запросы по студентам" src="https://github.com/user-attachments/assets/2d1206fd-6ea5-4449-a586-00b75f9b8168" />

В)Кол-во ошибок

<img width="923" height="847" alt="кол-во ошибок" src="https://github.com/user-attachments/assets/2ea92655-48e5-4cea-9a0e-30df11649a56" />

Г)Количество запросов по маршрутам

<img width="931" height="861" alt="Количество запросов по маршрутам" src="https://github.com/user-attachments/assets/7c1b4772-f890-41e9-a21a-9ff4bed6891b" />


Д)Ошибки по маршрутам

<img width="927" height="842" alt="Ошибки по маршрутам" src="https://github.com/user-attachments/assets/cb65c9d7-5046-4030-bd2b-47c631eb9f48" />


Е)Ошибки по статусам

<img width="927" height="852" alt="Ошибки по статусам" src="https://github.com/user-attachments/assets/a8b498d8-fe00-4681-bb2f-ac8274b35627" />


Ж)Среднее время обработки HTTP-запросов

<img width="932" height="877" alt="Среднее время обработки HTTP-запросов" src="https://github.com/user-attachments/assets/74bc72dc-c113-445e-9f7a-3d33c5fff778" />


З)Среднее время обработки маршрута студентов

<img width="925" height="856" alt="Среднее время обработки маршрута студентов" src="https://github.com/user-attachments/assets/dc7e04a0-d87c-4987-b0bd-3298b2dd640e" />


