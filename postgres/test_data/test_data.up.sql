INSERT INTO users (id, name, email, password_hash)
VALUES ('e5f1c9af-fa8a-4a58-9909-d887ddf7e961', 'Dennis Ritchie', 'ritchie@gmail.com', '$2y$12$37RmyCtr/AaSDX8D90slseB1jFwgP3SgOI5g5wtdk5FOpYCMvqhha');

INSERT INTO users (id, name, email, password_hash)
VALUES ('d1451907-e1ec-4291-ab14-a9a314b56b6a', 'Guido van Rossum', 'rossum@gmail.com', '$2y$12$LL6gUGE/8GuNxFKmpFowAOP/JaYPbS9ksvxzu6NdXup58E49UhxG6');

INSERT INTO users (id, name, email, password_hash)
VALUES ('0e38a4bd-87a0-447f-93fd-b904c9f7f303', 'Brendan Eich', 'eich@gmail.com', '$2y$12$MOTaU8rec4WQrzg5p4OLVO9brmharzlkygaCaxLy.K3of4wKlRd5m');

/*----------------------------------------------------------------*/

INSERT INTO articles (id, title, content, thumbnail_url, author)
VALUES ('82ba242e-e853-499f-8873-f271c53aca6a', 'Post about awsomeness of Go', 'Go is awsome.', 'http://www.example.com/path/to/photo0.jpg', '0e38a4bd-87a0-447f-93fd-b904c9f7f303');

INSERT INTO articles (id, title, content, thumbnail_url, author, created_at)
VALUES ('c3eec2ac-0fd5-41ce-829a-6f3dd74cd102', 'Very important article about programming', 'Very important contetnt of article.', 'http://www.example.com/path/to/photo1.jpg', 'd1451907-e1ec-4291-ab14-a9a314b56b6a', '01/01/2021 01:01:01');

INSERT INTO articles (id, title, content, thumbnail_url, author)
VALUES ('d50f5d60-6f59-4605-96b8-a96b9e9b17ea', 'Lorem ipsum article', 'Lorem ipsum dolor sit amet.', 'http://www.example.com/path/to/photo2.jpg', 'd1451907-e1ec-4291-ab14-a9a314b56b6a');
