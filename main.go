package main

//1. Описать user файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
//2. Реализовать gRPC сервис на основе user файла на Go
//3. Для хранения данных использовать PostgreSQL
//4. на запрос получения списка пользователей данные будут кешироваться в redis на минуту и брать из редиса
//5. При добавлении пользователя делать лог в clickHouse
//6. Добавление логов в clickHouse делать через очередь Kafka
