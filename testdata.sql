/*
delete from thread_comments;
delete from threads;
delete from thread_categories;
delete from articles;
delete from logins;
delete from users;
*/



-- ユーザーデータ
INSERT INTO users(user_id,name,email,password,created_at,updated_at) VALUES ('1','鈴木','test1@test.com','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,email,password,created_at,updated_at) VALUES ('2','田中','test2@test.com','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,email,password,created_at,updated_at) VALUES ('3','高橋','test3@test.com','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,email,password,created_at,updated_at) VALUES ('4','サトシ','test4@test.com','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,email,password,created_at,updated_at) VALUES ('5','ゴルシ','test5@test.com','pass1234',current_timestamp,current_timestamp);


-- ログインデータ
INSERT INTO logins(uuid,user_id,created_at) VALUES ('tjwaeoimaiso','1',current_timestamp);
INSERT INTO logins(uuid,user_id,created_at) VALUES ('tawletmqi23o','4',current_timestamp);
INSERT INTO logins(uuid,user_id,created_at) VALUES ('ai32qmiotqww','5',current_timestamp);

-- 記事データ
INSERT INTO articles(user_id ,title,description,created_at,updated_at) VALUES ('1','タイトル(仮)','記事(仮)',current_timestamp,current_timestamp);
INSERT INTO articles(user_id ,title,description,created_at,updated_at) VALUES ('1','タイトル(仮)2','記事(仮)2',current_timestamp,current_timestamp);
INSERT INTO articles(user_id ,title,description,created_at,updated_at) VALUES ('1','ポケモンって何だっけ...','ピカァ...',current_timestamp,current_timestamp);
INSERT INTO articles(user_id ,title,description,created_at,updated_at) VALUES ('1','人参嫌い','固すぎて食えたもんじゃないわ',current_timestamp,current_timestamp);



-- カテゴリーデータ
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('JavaScriprt',current_timestamp,current_timestamp);
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('Go',current_timestamp,current_timestamp);
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('ポケモン',current_timestamp,current_timestamp);
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('ウマ',current_timestamp,current_timestamp);

--スレッドデータ
INSERT INTO threads(title,thread_category_id,user_id,created_at,updated_at) VALUES ('JavaScriptとJavaの違い',1,1,current_timestamp,current_timestamp);
INSERT INTO threads(title,thread_category_id,user_id,created_at,updated_at) VALUES ('Goと60の違い',2,2,current_timestamp,current_timestamp);
INSERT INTO threads(title,thread_category_id,user_id,created_at,updated_at) VALUES ('ライチュウとデデンネの違い',3,4,current_timestamp,current_timestamp);
INSERT INTO threads(title,thread_category_id,user_id,created_at,updated_at) VALUES ('馬とUMAの違い',4,5,current_timestamp,current_timestamp);


--コメントデータ
INSERT INTO thread_comments(thread_id,text,session_id,created_at,updated_at) VALUES (1,'テストコメント1','asetawqmaoiu',current_timestamp,current_timestamp);
INSERT INTO thread_comments(thread_id,text,user_id,created_at,updated_at) VALUES (2,'テストコメント2','2',current_timestamp,current_timestamp);
INSERT INTO thread_comments(thread_id,text,session_id,created_at,updated_at) VALUES (3,'テストコメント3','asetawqmaoiu',current_timestamp,current_timestamp);
INSERT INTO thread_comments(thread_id,text,user_id,created_at,updated_at) VALUES (4,'テストコメント4','3',current_timestamp,current_timestamp);