/*
delete from thread_comments;
delete from threads;
delete from thread_categories;
delete from articles;
delete from logins;
delete from users;
*/



-- ユーザーデータ
INSERT INTO users(user_id,name,password,created_at,updated_at) VALUES ('1','鈴木','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,password,created_at,updated_at) VALUES ('2','田中','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,password,created_at,updated_at) VALUES ('3','高橋','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,password,created_at,updated_at) VALUES ('4','サトシ','pass1234',current_timestamp,current_timestamp);
INSERT INTO users(user_id,name,password,created_at,updated_at) VALUES ('5','ゴルシ','pass1234',current_timestamp,current_timestamp);


-- ログインデータ
INSERT INTO logins(uuid,user_id,created_at) VALUES ('tjwaeoimaiso','1',current_timestamp);
INSERT INTO logins(uuid,user_id,created_at) VALUES ('tawletmqi23o','4',current_timestamp);
INSERT INTO logins(uuid,user_id,created_at) VALUES ('ai32qmiotqww','5',current_timestamp);

-- 記事データ
INSERT INTO articles(id,user_id ,title,content,created_at,updated_at) VALUES ('awertasdfgzxcvb12345','1','タイトル(仮)','記事(仮)',current_timestamp,current_timestamp);
INSERT INTO articles(id,user_id ,title,content,created_at,updated_at) VALUES ('qwertasdfgzxcvb12345','2','タイトル(仮)2','記事(仮)2',current_timestamp,current_timestamp);
INSERT INTO articles(id,user_id ,title,content,created_at,updated_at) VALUES ('wwertasdfgzxcvb12345','3','ポケモンって何だっけ...','ピカァ...',current_timestamp,current_timestamp);
INSERT INTO articles(id,user_id ,title,content,created_at,updated_at) VALUES ('iwertasdfgzxcvb12345','4','人参嫌い','固すぎて食えたもんじゃないわ',current_timestamp,current_timestamp);



-- カテゴリーデータ
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('JavaScriprt',current_timestamp,current_timestamp);
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('Go',current_timestamp,current_timestamp);
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('ポケモン',current_timestamp,current_timestamp);
INSERT INTO thread_categories(name,created_at,updated_at) VALUES ('ウマ',current_timestamp,current_timestamp);

--スレッドデータ
INSERT INTO threads(user_id,thread_category_id,title,created_at,updated_at) VALUES (1,1,'JavaScriptとJavaの違い',current_timestamp,current_timestamp);
INSERT INTO threads(user_id,thread_category_id,title,created_at,updated_at) VALUES (2,2,'Goと60の違い',current_timestamp,current_timestamp);
INSERT INTO threads(user_id,thread_category_id,title,created_at,updated_at) VALUES (4,3,'ライチュウとデデンネの違い',current_timestamp,current_timestamp);
INSERT INTO threads(user_id,thread_category_id,title,created_at,updated_at) VALUES (5,4,'馬とUMAの違い',current_timestamp,current_timestamp);


--コメントデータ
INSERT INTO thread_comments(thread_id,text,session_id,created_at,updated_at) VALUES (1,'テストコメント1','asetawqmaoiu',current_timestamp,current_timestamp);
INSERT INTO thread_comments(thread_id,text,user_id,created_at,updated_at) VALUES (2,'テストコメント2','2',current_timestamp,current_timestamp);
INSERT INTO thread_comments(thread_id,text,session_id,created_at,updated_at) VALUES (3,'テストコメント3','asetawqmaoiu',current_timestamp,current_timestamp);
INSERT INTO thread_comments(thread_id,text,user_id,created_at,updated_at) VALUES (4,'テストコメント4','3',current_timestamp,current_timestamp);