create table todo (
  todo_id serial primary key,
  title varchar(30) not null,
  completed boolean not null default false
);

insert into todo (title, completed) values ('Buy groceries', false);
