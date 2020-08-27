use production_db;

create table user_parameters(
    id int auto_increment not null primary key,
    hp int not null,
    updated_datetime datetime not null
);

create table users(
    id int auto_increment not null primary key,
    name varchar(128) not null,
    email varchar(128) not null,
    password varchar(128) not null,
    token varchar(128) not null,
    param_id int not null,
    foreign key (param_id) references user_parameters(id) on delete cascade
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
    registered_datetime datetime not null,
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

insert into
    user_parameters(hp,updated_datetime)
values
    (1000000,Now());

insert into
    users(id,name,email,password,token,param_id)
values
    (1,"Hoge","hoge@hoge.jp","4c716d4cf211c7b7d2f3233c941771ad0507ea5bacf93b492766aa41ae9f720d","56f91b5f3668c470912be72ea6cbb0567cfdc0e6ab2266505f3f4b30bab989c6",1);

insert into
    tasks(title,deadline_date,deadline_time,description,weight_id,isAchieve,registered_datetime)
values
    ("睡眠","2020-08-28","21:00:00","ねる",1,false,Now());

insert into
    user_and_task_links(user_id,task_id)
values
    (1,1);

