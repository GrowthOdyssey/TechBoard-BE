-- 一旦仮でusersテーブルのみ
-- 値も仮
drop table users;

create table users (
  id         serial primary key,
  name       varchar(255),
  email      varchar(255) not null unique
);