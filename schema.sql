drop table thread_comments;
drop table threads;
drop table thread_categories;
drop table articles;
drop table logins;
drop table users;

-- ユーザー
create table users (
  user_id     varchar(20) primary key,
  name        varchar(100) not null,
  password    varchar(100) not null,
  avatar_id   varchar(20),
  created_at  timestamp not null,
  updated_at  timestamp not null
);

-- ログイン
create table logins (
  uuid        varchar(64) primary key,
  user_id     varchar(20) references users(user_id),
  created_at  timestamp not null
);

-- 記事
create table articles (
  id          serial primary key,
  user_id     varchar(20) references users(user_id),
  title       varchar(100) not null,
  content varchar(2000) not null,
  created_at  timestamp not null,
  updated_at  timestamp not null
);

-- スレッドのカテゴリ
create table thread_categories (
  id          serial primary key,
  name        varchar(100) not null,
  created_at  timestamp not null,
  updated_at  timestamp not null
);

-- スレッド
create table threads (
  id                  serial primary key,
  user_id             varchar(20) references users(user_id),
  thread_category_id  integer references thread_categories(id),
  title               varchar(100) not null,
  created_at          timestamp not null,
  updated_at          timestamp not null
);

-- スレッドのコメント
create table thread_comments (
  id          serial primary key,
  user_id     varchar(20),
  thread_id   integer references threads(id),
  session_id  varchar(100),
  text        varchar(2000) not null,
  created_at  timestamp not null,
  updated_at  timestamp not null
);
