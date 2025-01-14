Create Table Users(
    id int primary key Auto_INCREMENT ,
    name   varchar(255) not null,
    phone_number varchar(255) not null unique,
    password varchar(255) not null,
    email varchar(255) not null,
    Create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP



);