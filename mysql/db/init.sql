use production_db;

create table users(
    id int auto_increment not null primary key,
    name varchar(128) not null,
    email varchar(128) not null,
    password varchar(128) not null,
    token varchar(128) not null
);

create table weights(
    id int not null primary key,
    degree varchar(128) not null
);

create table tasks(
    id int auto_increment not null primary key,
    title varchar(128) not null,
    deadline_date date,
    deadline_time time,
    description varchar(128),
    weight_id int,
    isAchieve boolean not null,
    foreign key (weight_id) references weights(id) on delete cascade
);

create table user_and_task_links(
    user_id int not null,
    task_id int not null,
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (task_id) references tasks(id) on delete cascade
);

insert into 
    weights(id,degree) 
values
    (1,"ぬるい");

insert into 
    weights(id,degree) 
values
    (2,"ふつう");

insert into 
    weights(id,degree) 
values
    (3,"えぐい");
