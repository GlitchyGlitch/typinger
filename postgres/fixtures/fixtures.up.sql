INSERT INTO users (id, name, email, password_hash)
VALUES ('e5f1c9af-fa8a-4a58-9909-d887ddf7e961', 'First User', 'first@example.com', '$2y$12$n.qINbpxS3taWVFFCcj0FuvKnHGnQvVWBjof0EM1h2.eYhYel5AxC'); /* password: first */

INSERT INTO users (id, name, email, password_hash, created_at)
VALUES ('d1451907-e1ec-4291-ab14-a9a314b56b6a', 'Second User', 'second@example.com', '$2y$12$/7gXXmn.s7Ou5R1sKvWmr.4v9q1K9VbyHZOMqGAccvMgxuifGFZ7m', '02/02/2020 02:02:02'); /* password: second */

INSERT INTO users (id, name, email, password_hash, created_at)
VALUES ('0e38a4bd-87a0-447f-93fd-b904c9f7f303', 'Third User', 'third@example.com', '$2y$12$MOTaU8rec4WQrzg5p4OLVO9brmharzlkygaCaxLy.K3of4wKlRd5m', '01/01/2020 01:01:01'); /* password: third */

/*----------------------------------------------------------------*/

INSERT INTO articles (id, title, content, thumbnail_url, author)
VALUES ('82ba242e-e853-499f-8873-f271c53aca6a', 'First article', 'First content.', 'http://www.example.com/path/to/photo1', '0e38a4bd-87a0-447f-93fd-b904c9f7f303');

INSERT INTO articles (id, title, content, thumbnail_url, author, created_at)
VALUES ('c3eec2ac-0fd5-41ce-829a-6f3dd74cd102', 'Second article', 'Second content.', 'http://www.example.com/path/to/photo2', 'd1451907-e1ec-4291-ab14-a9a314b56b6a', '02/02/2021 02:02:02');

INSERT INTO articles (id, title, content, thumbnail_url, author, created_at)
VALUES ('d50f5d60-6f59-4605-96b8-a96b9e9b17ea', 'Third article', 'Third content.', 'http://www.example.com/path/to/photo3', 'd1451907-e1ec-4291-ab14-a9a314b56b6a', '01/01/2021 01:01:01');

/*----------------------------------------------------------------*/

INSERT INTO images (id, name, slug, mime, img)
VALUES ('0cf191b8-b60c-4aec-b698-ca2b64a3d0f7', 'First image', 'first-slug', 'image/svg+xml', decode('aa', 'hex'));

INSERT INTO images (id, name, slug, mime, img, created_at)
VALUES ('60b390de-cd44-4cc2-9fbb-fd9bcaa42819', 'Second image', 'second-slug', 'image/webp', decode('bb', 'hex'), '04/04/2021 04:04:04');

INSERT INTO images (id, name, slug, mime, img, created_at)
VALUES ('72e641f1-09ac-4444-b980-98be375f8efd', 'Third image', 'third-slug', 'image/webp', decode('cc', 'hex'), '03/03/2021 03:03:03');

/*----------------------------------------------------------------*/
