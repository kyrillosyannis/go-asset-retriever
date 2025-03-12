create table if not exists users(
    id bigserial primary key,
    username varchar(255),
    password varchar(255)
);

insert into users(username, password) values('user', '$2a$12$bXd7Hif1bhPyWkvsrY2U5./a8mO9jAPCmKFiVue23VIz9dS1sNQmW');